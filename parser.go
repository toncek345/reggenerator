package reggenerator

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"unicode"
)

var possibleCharArray []byte

func init() {
	// TODO: possible char array is bigger, including [],? etc...
	possibleCharArray = make([]byte, 0, ('z'-'a')*2)
	for i := 'a'; i <= 'z'; i++ {
		possibleCharArray = append(possibleCharArray, byte(i))
	}
	for i := 'A'; i <= 'Z'; i++ {
		possibleCharArray = append(possibleCharArray, byte(i))
	}
}

type parsedToken struct {
	possibleCharacters []byte
}

func parse(tokens []*token) ([]*parsedToken, error) {
	parsedTokens := make([]*parsedToken, 0, len(tokens))
	for _, t := range tokens {
		// TODO: if there is no character limit, add any character as a first case

		switch len(t.charRange) {
		case 0:
			return nil, fmt.Errorf("char range is 0, something is terribly wrong")
		case 1:
			t := &parsedToken{possibleCharacters: []byte{t.charRange[0]}}
			parsedTokens = append(parsedTokens, t)
			continue
		}

		negate := false
		charRange := t.charRange
		if charRange[0] == '^' {
			negate = true
			charRange = charRange[1:]
		}

		reader := bufio.NewReader(bytes.NewReader([]byte(charRange)))
		parsed := new(parsedToken)
		for {
			byte, err := reader.ReadByte()
			if err != nil {
				if err == io.EOF {
					break
				}
				return nil, fmt.Errorf("error reading byte: %w", err)
			}

			nextByte, err := reader.Peek(1)
			if err != nil && err != io.EOF {
				return nil, fmt.Errorf("error peeking byte: %w", err)
			}

			if len(nextByte) == 0 || !isRange(nextByte[0]) {
				// In this step of iteration, we have a single character, not a range
				parsed.possibleCharacters = append(parsed.possibleCharacters, byte)
				continue
			}

			secondChar, err := readSecondRangePart(reader)
			if err != nil {
				return nil, fmt.Errorf("reading second part of the range: %w", err)
			}

			bytes, err := byteRange(byte, secondChar)
			if err != nil {
				return nil, fmt.Errorf("getting bytes from range: %w", err)
			}

			parsed.possibleCharacters = append(parsed.possibleCharacters, bytes...)
		}

		// TODO: do not forget about quantifier
		if negate {
			parsed.possibleCharacters = negateCharacters(parsed.possibleCharacters)
		}
		parsed.possibleCharacters = uniqueCharacters(parsed.possibleCharacters)

		parsedTokens = append(parsedTokens, parsed)
	}

	return parsedTokens, nil
}

func isRange(char byte) bool {
	return char == '-'
}

func readSecondRangePart(reader *bufio.Reader) (byte, error) {
	bytes := make([]byte, 2)
	n, err := reader.Read(bytes)
	if err != nil {
		if err == io.EOF {
			return 0, fmt.Errorf("invalid regex")
		}
		return 0, fmt.Errorf("unexpected error reading range: %w", err)
	}
	if n != 2 {
		return 0, fmt.Errorf("wrong number of bytes read: %d", n)
	}

	return bytes[1], nil
}

// byteRange return the char range from first to second character.
func byteRange(first, second byte) ([]byte, error) {
	if unicode.IsLower(rune(first)) && unicode.IsLower(rune(second)) ||
		unicode.IsUpper(rune(first)) && unicode.IsUpper(rune(second)) {

		byteRange := make([]byte, 0, second-first+1)
		for i := first; i <= second; i++ {
			byteRange = append(byteRange, i)
		}

		return byteRange, nil
	}

	return nil, fmt.Errorf("characters do not range lower/upper case")
}

// uniqueCharacters makes all characters in the byte array unique
func uniqueCharacters(chars []byte) []byte {
	new := make([]byte, 0, len(chars))
	charSet := make(map[byte]struct{})

	for _, v := range chars {
		if _, ok := charSet[v]; ok {
			continue
		}

		charSet[v] = struct{}{}
		new = append(new, v)
	}

	return new
}

func negateCharacters(chars []byte) []byte {
	charMap := make(map[byte]struct{})
	for _, c := range chars {
		charMap[c] = struct{}{}
	}

	bytes := make([]byte, 0, len(charMap))
	for _, v := range possibleCharArray {
		if _, ok := charMap[v]; ok {
			continue
		}

		bytes = append(bytes, v)
	}

	return bytes
}
