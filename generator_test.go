package reggenerator_test

import (
	"testing"

	"github.com/toncek345/reggenerator"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		name     string
		randomFn func() func() int
		regex    string
		expected []string
		count    int
	}{
		{
			name: "generate",
			randomFn: func() func() int {
				return func() int {
					return 1
				}
			},
			regex:    "/g{4}/",
			expected: []string{"gggg"},
			count:    1,
		},
		{
			name: "generate random possibility",
			randomFn: func() func() int {
				return func() int {
					return 2
				}
			},
			regex:    "/g?/",
			expected: []string{""},
			count:    1,
		},
		{
			name: "generate random possibility2",
			randomFn: func() func() int {
				return func() int {
					return 1
				}
			},
			regex:    "/g?/",
			expected: []string{"g"},
			count:    1,
		},
		{
			name: "generate with char range",
			randomFn: func() func() int {
				i := 1
				return func() int {
					i++
					return i
				}
			},
			regex:    "/[ghj]{4}/",
			expected: []string{"jghj"},
			count:    1,
		},
		{
			name: "generate range",
			randomFn: func() func() int {
				i := 0
				return func() int {
					i++
					return i / 2
				}
			},
			regex:    "/g{2,4}/",
			expected: []string{"gg", "gggg", "ggg"},
			count:    3,
		},
	}

	for _, v := range tests {
		t.Run(v.name, func(t *testing.T) {
			reggenerator.RandFn = v.randomFn()

			got, err := reggenerator.Generate(v.regex, v.count)
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}

			if len(got) != len(v.expected) {
				t.Fatalf("expected length '%d', got '%d'", len(v.expected), len(got))
			}

			for i := range v.expected {
				if got[i] != v.expected[i] {
					t.Logf("%#v; %#v\n", got, v.expected)
					t.Fatalf("expected result '%s', got '%s'", v.expected[i], got[i])
				}
			}
		})
	}
}
