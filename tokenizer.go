package reggenerator

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
)

type token struct {
	charRange    string
	quantifier   string
	anyCharacter bool
}

func tokenize(str string) ([]*token, error) {
	if len(str) == 0 ||
		str[0] != '/' ||
		str[len(str)-1] != '/' {
		return nil, fmt.Errorf("invalid regex")
	}

	reader := bufio.NewReader(bytes.NewReader([]byte(str[1 : len(str)-1])))
	tokens := make([]*token, 0, 10)

	for {
		char, err := reader.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("reading error: %w", err)
		}

		token := new(token)
		switch char {
		case '.':
			token.anyCharacter = true
		case '[':
			token.charRange, err = readCharRange(reader)
			if err != nil {
				return nil, fmt.Errorf("reading char range: %w", err)
			}
		default:
			token.charRange = string(char)
		}

		quantifier, err := readQuantifier(reader)
		if err != nil {
			return nil, fmt.Errorf("reading quantifier: %w", err)
		}
		token.quantifier = quantifier

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func readCharRange(reader *bufio.Reader) (string, error) {
	bytes, err := reader.ReadBytes(']')
	if err != nil {
		return "", fmt.Errorf("reading bytes: %w", err)
	}

	return string(bytes[:len(bytes)-1]), nil // Cut last char ']'
}

func readQuantifier(reader *bufio.Reader) (string, error) {
	bytes, err := reader.Peek(1)
	if err != nil {
		if err == io.EOF { // End of string, means no quantifier
			return "", nil
		}

		return "", fmt.Errorf("quantifier peeking: %w", err)
	}
	switch bytes[0] {
	case '+', '*':
		// '+' and '*' are always considered as quantifiers unless they are first char
		return "", fmt.Errorf("not supporting '+' quantifier, use range instead {1,4}")
	case '?':
		_, err := reader.Discard(1)
		if err != nil {
			return "", fmt.Errorf("discarding byte: %w", err)
		}

		return "?", nil
	case '{':
		bytes, err := reader.ReadBytes('}')
		if err != nil {
			return "", fmt.Errorf("reading bytes: %w", err)
		}

		return string(bytes[1 : len(bytes)-1]), nil // Cut first ('{') and last ('}') char
	}

	return "", nil
}
