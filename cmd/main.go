// HO.

package main

import (
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/tools/go/packages"

	"github.com/pkg/errors"
)

func main() {
	// File path to your Go source file
	pattern := "./apis/..."

	pkgs, err := packages.Load(&packages.Config{
		Mode: packages.NeedName | packages.NeedFiles | packages.NeedImports | packages.NeedDeps | packages.NeedTypes | packages.NeedSyntax,
	}, pattern)
	if err != nil {
		panic(err)
	}
	packages.PrintErrors(pkgs)
	for _, p := range pkgs {
		if len(p.Errors) != 0 {
			continue
		}
		for i, f := range p.GoFiles {
			if filepath.Base(f) != "zz_generated.resolvers.go" {
				continue
			}
			processFile(p.Fset, p.Syntax[i], f)
		}
	}

	/*file := "/Users/alper/github/upbound/provider-aws/apis/workspaces/v1beta1/zz_generated.resolvers.go"
	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, file, nil, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	processFile(fset, f, file)*/
}

type importUsage struct {
	path string
	used bool
}

func processFile(fset *token.FileSet, node *ast.File, filePath string) {
	importMap, err := addMRVariableDeclarations(node)
	if err != nil {
		panic(err)
	}

	// Map to track imports used in reference.To structs
	importsUsed := make(map[string]importUsage)
	var block *ast.BlockStmt
	var assign *ast.AssignStmt
	var group, version, kind, listKind string

	// Traverse the AST and modify the nodes
	ast.Inspect(node, func(n ast.Node) bool {
		switch x := n.(type) {
		case *ast.ImportSpec:
			// Initially, mark all imports as not used
			if x.Name != nil {
				importsUsed[x.Name.Name] = importUsage{
					path: strings.Trim(x.Path.Value, `"`),
					used: true,
				}
			} else if x.Path != nil {
				importsUsed[x.Path.Value] = importUsage{
					path: strings.Trim(x.Path.Value, `"`),
					used: true,
				}
			}

		case *ast.FuncDecl:
			block = x.Body

		case *ast.AssignStmt:
			assign = x

		case *ast.KeyValueExpr:
			// Check if the key is "To" and the value is a CompositeLit
			if key, ok := x.Key.(*ast.Ident); ok && key.Name == "To" {
				if cl, ok := x.Value.(*ast.CompositeLit); ok {
					//printer.Fprint(os.Stdout, fset, n)
					//fmt.Println()
					//fmt.Println()

					// Check if there are any package qualifiers in the CompositeLit
					for _, elt := range cl.Elts {
						if kv, ok := elt.(*ast.KeyValueExpr); ok {
							if uexpr, ok := kv.Value.(*ast.UnaryExpr); ok {
								if cl, ok := uexpr.X.(*ast.CompositeLit); ok {
									if sexpr, ok := cl.Type.(*ast.SelectorExpr); ok {
										if ident, ok := sexpr.X.(*ast.Ident); ok {
											path := importsUsed[ident.Name].path
											importsUsed[ident.Name] = importUsage{
												path: path,
												used: false,
											}
											tokens := strings.Split(path, "/")
											version = tokens[len(tokens)-1]
											group = tokens[len(tokens)-2] + ".aws.upbound.io"
											if sexpr.Sel != nil {
												if strings.HasSuffix(sexpr.Sel.Name, "List") {
													listKind = sexpr.Sel.Name
												} else {
													kind = sexpr.Sel.Name
												}
											}
										}
									} else {
										tokens := strings.Split(filePath, "/")
										version = tokens[len(tokens)-2]
										group = tokens[len(tokens)-3] + ".aws.upbound.io"
										if ident, ok := cl.Type.(*ast.Ident); ok {
											if strings.HasSuffix(ident.Name, "List") {
												listKind = ident.Name
											} else {
												kind = ident.Name
											}
										}
										// fmt.Println(group, version, kind, listKind)
									}
								}
							}
						}
					}

					// Replace the value with a CompositeLit of type reference.To
					x.Value = &ast.CompositeLit{
						Type: &ast.SelectorExpr{
							X:   ast.NewIdent("reference"),
							Sel: ast.NewIdent("To"),
						},
						Elts: []ast.Expr{
							&ast.KeyValueExpr{
								Key:   ast.NewIdent("List"),
								Value: ast.NewIdent("l"),
							},
							&ast.KeyValueExpr{
								Key:   ast.NewIdent("Managed"),
								Value: ast.NewIdent("m"),
							},
						},
					}

					mrImports, stmts, err := getManagedResourceStatements(group, version, kind, listKind)
					if err != nil {
						panic(err)
					}
					if !insertStatements(stmts, block, assign) {
						panic("MR statements not injected")
					}
					for k, v := range mrImports {
						importMap[k] = v
					}
				}
			}
		}
		return true
	})

	// Remove the unused imports
	for _, decl := range node.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok && gd.Tok == token.IMPORT {
			newSpecs := []ast.Spec{}
			for _, spec := range gd.Specs {
				if imp, ok := spec.(*ast.ImportSpec); ok {
					var name string
					if imp.Name != nil {
						name = imp.Name.Name
					} else {
						name = imp.Path.Value
					}
					if usage, exists := importsUsed[name]; !exists || usage.used {
						newSpecs = append(newSpecs, spec)
					}
				}
			}
			gd.Specs = newSpecs

			for path, name := range importMap {
				gd.Specs = append(gd.Specs, &ast.ImportSpec{
					Name: &ast.Ident{
						Name: name,
					},
					Path: &ast.BasicLit{
						Kind:  token.STRING,
						Value: path,
					},
				})
			}
		}
	}

	// Open the file for writing
	outFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Write the modified AST back to the file
	if err := format.Node(outFile, fset, node); err != nil {
		panic(err)
	}
}

func insertStatements(stmts []ast.Stmt, block *ast.BlockStmt, assign *ast.AssignStmt) bool {
	if block == nil {
		return false
	}

	for i, s := range block.List {
		if forStmt, ok := s.(*ast.ForStmt); ok {
			if insertStatements(stmts, forStmt.Body, assign) {
				return true
			}
			continue
		}

		if s == assign {
			block.List = append(block.List[:i], append(stmts, block.List[i:]...)...)
			return true
		}
	}
	return false
}

func addMRVariableDeclarations(f *ast.File) (map[string]string, error) {
	varSrc := `package main

import (
	xpresource "github.com/crossplane/crossplane-runtime/pkg/resource"
)

// reference resolver source objects
var m xpresource.Managed
var l xpresource.ManagedList
`
	fset := token.NewFileSet()
	varFile, err := parser.ParseFile(fset, "", varSrc, parser.ParseComments)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse the managed resource variables file")
	}
	var varDecls []ast.Stmt
	importMap := make(map[string]string, 0)
	for _, decl := range varFile.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			switch genDecl.Tok {
			case token.VAR:
				varDecls = append(varDecls, &ast.DeclStmt{Decl: genDecl})

			case token.IMPORT:
				for _, spec := range genDecl.Specs {
					if importSpec, ok := spec.(*ast.ImportSpec); ok {
						name := ""
						if importSpec.Name != nil {
							name = importSpec.Name.Name
						}
						importMap[importSpec.Path.Value] = name
					}
				}
			}
		}
	}

	ast.Inspect(f, func(n ast.Node) bool {
		fn, ok := n.(*ast.FuncDecl)
		if !ok {
			return true
		}

		if fn.Name.Name == "ResolveReferences" && len(fn.Recv.List) > 0 {
			fn.Body.List = append(varDecls, fn.Body.List...)
		}

		return true
	})

	return importMap, nil
}

func getManagedResourceStatements(group, version, kind, listKind string) (map[string]string, []ast.Stmt, error) {
	stmtSrc := `package main

import (
	apisresolver "github.com/upbound/provider-aws/internal/apis"
)

func f() {
	m, l, err = apisresolver.GetManagedResource("%s", "%s", "%s", "%s")
	if err != nil {
		return errors.Wrap(err, "failed to get the reference target managed resource and its list for reference resolution")
	}
}
`
	stmtSrc = fmt.Sprintf(stmtSrc, group, version, kind, listKind)

	fset := token.NewFileSet()
	stmtFile, err := parser.ParseFile(fset, "", stmtSrc, parser.ParseComments)
	if err != nil {
		panic(err)
	}
	importMap := make(map[string]string, 0)
	var stmts []ast.Stmt
	for _, decl := range stmtFile.Decls {
		switch x := decl.(type) {
		case *ast.GenDecl:
			if x.Tok == token.IMPORT {
				for _, spec := range x.Specs {
					if importSpec, ok := spec.(*ast.ImportSpec); ok {
						name := ""
						if importSpec.Name != nil {
							name = importSpec.Name.Name
						}
						importMap[importSpec.Path.Value] = name
					}
				}
			}

		case *ast.FuncDecl:
			stmts = x.Body.List
		}

	}
	return importMap, stmts, nil
}
