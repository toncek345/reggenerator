package reggenerator

import (
	"testing"
)

func TestParse(t *testing.T) {
	parsedOutputEqual := func(t *testing.T, expected, actual []*parsedToken) {
		if len(expected) != len(actual) {
			t.Fatalf("len of expected and actual should be the same")
		}

		for i := range expected {
			if len(expected[i].possibleCharacters) != len(actual[i].possibleCharacters) {
				t.Fatalf("len of possible characters is different")
			}

			for j := range expected[i].possibleCharacters {
				if expected[i].possibleCharacters[j] != actual[i].possibleCharacters[j] {
					t.Fatalf("character should be '%c' got '%c'", expected[i].possibleCharacters[i], actual[i].possibleCharacters[i])
				}
			}

			// TODO: check quantifier
		}
	}

	tests := []struct {
		name        string
		tokens      []*token
		expectedOut []*parsedToken
		shouldErr   bool
	}{
		{
			name:      "no char rage",
			shouldErr: true,
			tokens: []*token{
				{
					charRange: "",
				},
			},
		},
		{
			name:      "invalid range",
			shouldErr: true,
			tokens: []*token{
				{
					charRange: "a-",
				},
			},
		},
		{
			name: "single letter; single token",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: []byte{'a'},
				},
			},
			tokens: []*token{
				{
					charRange: "a",
				},
			},
		},
		{
			name: "single letter; multiple tokens",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: []byte{'a'},
				},
				{
					possibleCharacters: []byte{'b'},
				},
			},
			tokens: []*token{
				{
					charRange: "a",
				},
				{
					charRange: "b",
				},
			},
		},
		{
			name: "single token, range",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: []byte{'a', 'b', 'c'},
				},
			},
			tokens: []*token{
				{
					charRange: "a-c",
				},
			},
		},
		{
			name: "single token, range & single letter that is already in the range",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: []byte{'a', 'b', 'c'},
				},
			},
			tokens: []*token{
				{
					charRange: "aa-c",
				},
			},
		},
		{
			name: "single token, range & single letter",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: []byte{'t', 'a', 'b', 'c'},
				},
			},
			tokens: []*token{
				{
					charRange: "ta-c",
				},
			},
		},
		{
			name: "multiple token, range",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: []byte{'a', 'b', 'c'},
				},
				{
					possibleCharacters: []byte{'e', 'f', 'g'},
				},
			},
			tokens: []*token{
				{
					charRange: "a-c",
				},
				{
					charRange: "e-g",
				},
			},
		},
		{
			name: "negation",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: possibleCharArray[:len(possibleCharArray)-1],
				},
			},
			tokens: []*token{
				{
					charRange: "^~",
				},
			},
		},
		{
			name: "any char",
			expectedOut: []*parsedToken{
				{
					possibleCharacters: possibleCharArray,
				},
			},
			tokens: []*token{
				{
					anyCharacter: true,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			out, err := parse(test.tokens)
			if err != nil {
				if test.shouldErr {
					return
				}
				t.Errorf("unexpected error: %s", err)
			}

			parsedOutputEqual(t, test.expectedOut, out)
		})
	}
}
