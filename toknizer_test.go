package reggenerator

import "testing"

func TestTokenize(t *testing.T) {
	isEqualtokens := func(t *testing.T, expected, actual []*token) {
		if len(expected) != len(actual) {
			t.Fatalf("token len not equal: expected %d, got %d", len(expected), len(actual))
		}

		for i := range expected {
			if len(expected[i].charRange) != len(actual[i].charRange) {
				t.Fatalf("char range is not equal, expected %d, got %d", len(expected[i].charRange), len(actual[i].charRange))
			}
			for j := range expected[i].charRange {
				if expected[i].charRange[j] != actual[i].charRange[j] {
					t.Fatalf("char should be '%c' got '%c'", expected[i].charRange[j], actual[i].charRange[j])
				}
			}

			if expected[i].quantifier != actual[i].quantifier {
				t.Fatalf("quantifier should be '%s', got '%s'", expected[i].quantifier, actual[i].quantifier)
			}

			if expected[i].anyCharacter != actual[i].anyCharacter {
				t.Fatalf("any characher doesn't match")
			}
		}
	}

	tests := []struct {
		name      string
		str       string
		expected  []*token
		shouldErr bool
	}{
		{
			name:      "empty string",
			str:       "",
			shouldErr: true,
		},
		{
			name:      "missing slash before",
			str:       "fds/",
			shouldErr: true,
		},
		{
			name:      "missing slash after",
			str:       "fds/",
			shouldErr: true,
		},
		{
			name:      "invalid regex",
			str:       "/[fds/",
			shouldErr: true,
		},
		{
			name:      "invalid regex 2",
			str:       "/[fds]{/",
			shouldErr: true,
		},
		{
			name:      "invalid regex 3",
			str:       "/[f-]{/",
			shouldErr: true,
		},
		{
			name: "single char",
			str:  "/f/",
			expected: []*token{
				{
					charRange: "f",
				},
			},
		},
		{
			name: "only characters",
			str:  "/f[.]/",
			expected: []*token{
				{
					charRange: "f",
				},
				{
					charRange: ".",
				},
			},
		},
		{
			name: "only char special chars",
			str:  "/-[?.],!/",
			expected: []*token{
				{
					charRange: "-",
				},
				{
					charRange: "?.",
				},
				{
					charRange: ",",
				},
				{
					charRange: "!",
				},
			},
		},
		{
			name: "char range",
			str:  "/[a-b]/",
			expected: []*token{
				{
					charRange: "a-b",
				},
			},
		},
		{
			name: "char range 2",
			str:  "/[a-bA-Z]/",
			expected: []*token{
				{
					charRange: "a-bA-Z",
				},
			},
		},
		{
			name: "char range with quantifier",
			str:  "/[a-bA-Z]{1,3}/",
			expected: []*token{
				{
					charRange:  "a-bA-Z",
					quantifier: "1,3",
				},
			},
		},
		{
			name: "char range with optional",
			str:  "/[ab]?/",
			expected: []*token{
				{
					charRange:  "ab",
					quantifier: "?",
				},
			},
		},
		{
			name: "negated chars",
			str:  "/[^ab]?/",
			expected: []*token{
				{
					charRange:  "^ab",
					quantifier: "?",
				},
			},
		},
		{
			name: "dot",
			str:  "/[.]/",
			expected: []*token{
				{
					charRange: ".",
				},
			},
		},
		{
			name: "any char",
			str:  "/./",
			expected: []*token{
				{
					anyCharacter: true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tokens, err := tokenize(test.str)
			if err != nil {
				if test.shouldErr {
					return
				}

				t.Errorf("unexpected error: %s", err)
				return
			}

			isEqualtokens(t, test.expected, tokens)
		})
	}
}
