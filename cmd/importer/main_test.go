package main

import (
	"reflect"
	"testing"
)

// TestRemoveNullValues checks the function against various map configurations
func TestRemoveNullValues(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]any
		expected map[string]any
	}{
		{
			name: "single level no nil",
			input: map[string]any{
				"language": "Go",
				"type":     "compiled",
			},
			expected: map[string]any{
				"language": "Go",
				"type":     "compiled",
			},
		},
		{
			name: "single level with nil",
			input: map[string]any{
				"language": "Go",
				"version":  nil,
			},
			expected: map[string]any{
				"language": "Go",
			},
		},
		{
			name: "nested maps with nil",
			input: map[string]any{
				"details": map[string]any{
					"type": "language",
					"year": nil,
				},
			},
			expected: map[string]any{
				"details": map[string]any{
					"type": "language",
				},
			},
		},
		{
			name: "deeply nested maps with arrays",
			input: map[string]any{
				"details": map[string]any{
					"type": "language",
					"specs": []interface{}{
						map[string]any{"speed": "fast", "safety": nil},
					},
				},
				"version": nil,
			},
			expected: map[string]any{
				"details": map[string]any{
					"type": "language",
					"specs": []interface{}{
						map[string]any{"speed": "fast"},
					},
				},
			},
		},
		{
			name: "nested array of maps and nils",
			input: map[string]any{
				"features": []interface{}{
					nil,
					map[string]any{"reliable": true, "secure": nil},
					[]interface{}{nil, map[string]any{"async": true}, nil},
				},
			},
			expected: map[string]any{
				"features": []interface{}{
					map[string]any{"reliable": true},
					[]interface{}{map[string]any{"async": true}},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			removeNullValues(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("removeNullValues() = %v, want %v", test.input, test.expected)
			}
		})
	}
}
