// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package servicemetadata

import (
	"os"
	"path/filepath"
	"slices"
	"sort"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

// TestServiceMetadataCompleteness ensures every cmd/provider/<name> subdirectory has a
// matching top-level key in package/service-metadata.yaml (and vice versa), with
// ServiceDescription and at least one ServiceCategories entry.
func TestServiceMetadataCompleteness(t *testing.T) {
	t.Parallel()

	root, err := moduleRoot()
	if err != nil {
		t.Fatalf("module root: %v", err)
	}

	provDir := filepath.Join(root, "cmd", "provider")
	entries, err := os.ReadDir(provDir)
	if err != nil {
		t.Fatalf("read cmd/provider: %v", err)
	}

	var want []string
	for _, e := range entries {
		if !e.IsDir() || strings.HasPrefix(e.Name(), ".") {
			continue
		}
		want = append(want, e.Name())
	}
	sort.Strings(want)

	metaPath := filepath.Join(root, "package", "service-metadata.yaml")
	raw, err := os.ReadFile(metaPath)
	if err != nil {
		t.Fatalf("read %s: %v", metaPath, err)
	}

	var meta map[string]map[string]any
	if err := yaml.Unmarshal(raw, &meta); err != nil {
		t.Fatalf("parse service-metadata.yaml: %v", err)
	}

	got := make([]string, 0, len(meta))
	for k := range meta {
		got = append(got, k)
	}
	sort.Strings(got)

	if !slices.Equal(got, want) {
		t.Errorf("cmd/provider subdirectories must match package/service-metadata.yaml keys exactly.\n"+
			"want (%d): %q\n"+
			"got (%d): %q\n"+
			"add or remove metadata entries when adding or removing a subprovider under cmd/provider.",
			len(want), want, len(got), got)
	}

	for _, name := range want {
		entry, ok := meta[name]
		if !ok {
			continue // already reported above
		}
		desc, _ := entry["ServiceDescription"].(string)
		if strings.TrimSpace(desc) == "" {
			t.Errorf("service-metadata.yaml %q: ServiceDescription must be non-empty", name)
			continue
		}
		tokens, err := DescriptionEmbeddingTokenCount(desc)
		if err != nil {
			t.Errorf("service-metadata.yaml %q: token count: %v", name, err)
		} else if tokens >= ServiceDescriptionTokenLimit {
			t.Errorf("service-metadata.yaml %q: ServiceDescription must be fewer than %d embedding-style tokens (cl100k_base), got %d",
				name, ServiceDescriptionTokenLimit, tokens)
		}
		cats, ok := entry["ServiceCategories"].([]any)
		if !ok || len(cats) == 0 {
			t.Errorf("service-metadata.yaml %q: ServiceCategories must be a non-empty list", name)
		}
	}
}

func moduleRoot() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", os.ErrNotExist
		}
		dir = parent
	}
}
