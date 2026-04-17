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
	"sync"
	"testing"

	"github.com/pkoukk/tiktoken-go"
	"gopkg.in/yaml.v3"
)

// serviceDescriptionTokenLimit is the exclusive upper bound on embedding-style token count:
// ServiceDescription must contain fewer than this many tokens (i.e. count < limit).
const serviceDescriptionTokenLimit = 255

var (
	embeddingEncoding     *tiktoken.Tiktoken
	embeddingEncodingErr  error
	embeddingEncodingOnce sync.Once
)

func descriptionEncoding() (*tiktoken.Tiktoken, error) {
	embeddingEncodingOnce.Do(func() {
		embeddingEncoding, embeddingEncodingErr = tiktoken.GetEncoding("cl100k_base")
	})
	return embeddingEncoding, embeddingEncodingErr
}

// descriptionEmbeddingTokenCount counts subword tokens in s using cl100k_base (test-only;
// used to validate service-metadata.yaml offline in CI).
func descriptionEmbeddingTokenCount(s string) (int, error) {
	enc, err := descriptionEncoding()
	if err != nil {
		return 0, err
	}
	return len(enc.Encode(s, nil, nil)), nil
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

func TestDescriptionEmbeddingTokenCount(t *testing.T) {
	t.Parallel()
	n, err := descriptionEmbeddingTokenCount("hello world")
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("descriptionEmbeddingTokenCount(%q): got %d tokens, want 2", "hello world", n)
	}
}

func TestDescriptionEmbeddingTokenCount_EdgeCases(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		input     string
		wantExact bool
		exact     int
		minTokens int
	}{
		{
			name:      "empty",
			input:     "",
			wantExact: true,
			exact:     0,
		},
		{
			name:      "unicode",
			input:     "日本語テスト",
			minTokens: 1,
		},
		// Golden value: strings.Repeat("word ", 300) is exactly 301 cl100k_base tokens (tiktoken).
		{
			name:      "long_near_token_limit",
			input:     strings.Repeat("word ", 300),
			wantExact: true,
			exact:     301,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			n, err := descriptionEmbeddingTokenCount(tt.input)
			if err != nil {
				t.Fatalf("descriptionEmbeddingTokenCount: %v", err)
			}
			if tt.wantExact {
				if n != tt.exact {
					t.Errorf("got %d tokens, want exactly %d", n, tt.exact)
				}
				return
			}
			if tt.minTokens > 0 && n < tt.minTokens {
				t.Errorf("got %d tokens, want >= %d", n, tt.minTokens)
			}
		})
	}
}

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
		tokens, err := descriptionEmbeddingTokenCount(desc)
		if err != nil {
			t.Errorf("service-metadata.yaml %q: token count: %v", name, err)
		} else if tokens >= serviceDescriptionTokenLimit {
			t.Errorf("service-metadata.yaml %q: ServiceDescription must be fewer than %d embedding-style tokens (cl100k_base), got %d",
				name, serviceDescriptionTokenLimit, tokens)
		}
		cats, ok := entry["ServiceCategories"].([]any)
		if !ok || len(cats) == 0 {
			t.Errorf("service-metadata.yaml %q: ServiceCategories must be a non-empty list", name)
		}
	}
}
