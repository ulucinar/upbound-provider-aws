// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package servicemetadata

import (
	"strings"
	"testing"
)

func TestDescriptionEmbeddingTokenCount(t *testing.T) {
	t.Parallel()
	n, err := DescriptionEmbeddingTokenCount("hello world")
	if err != nil {
		t.Fatal(err)
	}
	if n != 2 {
		t.Fatalf("DescriptionEmbeddingTokenCount(%q): got %d tokens, want 2", "hello world", n)
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
		// Golden value: strings.Repeat("word ", 300) is exactly 301 cl100k_base tokens (DescriptionEmbeddingTokenCount / tiktoken).
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
			n, err := DescriptionEmbeddingTokenCount(tt.input)
			if err != nil {
				t.Fatalf("DescriptionEmbeddingTokenCount: %v", err)
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
