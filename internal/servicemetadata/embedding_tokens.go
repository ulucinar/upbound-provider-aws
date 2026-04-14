// SPDX-FileCopyrightText: 2026 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package servicemetadata

import (
	"sync"

	"github.com/pkoukk/tiktoken-go"
)

// ServiceDescriptionTokenLimit is the exclusive upper bound on embedding-style token count:
// ServiceDescription must contain fewer than this many tokens (i.e. count < limit).
const ServiceDescriptionTokenLimit = 255

// DescriptionEmbeddingTokenCount returns a subword token count suitable for checking
// text length against vector-embedding limits. It uses the cl100k_base BPE encoding
// (same family as tiktoken / many OpenAI and API-compatible embedding pipelines).
//
// Elasticsearch inference and semantic_text do not ship a tokenizer in this repo; the
// exact count for a given deployment depends on the trained model's vocabulary. This
// function matches a widely used BPE and gives deterministic, offline CI validation.
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

// DescriptionEmbeddingTokenCount counts subword tokens in s using cl100k_base.
func DescriptionEmbeddingTokenCount(s string) (int, error) {
	enc, err := descriptionEncoding()
	if err != nil {
		return 0, err
	}
	return len(enc.Encode(s, nil, nil)), nil
}
