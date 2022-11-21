package reggenerator

import (
	"fmt"
	"math/rand"
	"strings"

	"golang.org/x/sync/errgroup"
)

// RandFn is a pointer to a function that returns random generated number.
var RandFn func() int = rand.Int

// Generate generates N random strings based on supplied regex.
func Generate(regex string, count int) ([]string, error) {
	tokens, err := tokenize(regex)
	if err != nil {
		return nil, fmt.Errorf("diagnosing: %w", err)
	}

	parsedTokens, err := parse(tokens)
	if err != nil {
		return nil, fmt.Errorf("parsing tokens: %w", err)
	}

	g := errgroup.Group{}
	ch := make(chan string, count)
	for i := 0; i < count; i++ {
		g.Go(func() error {
			s := strings.Builder{}
			for _, t := range parsedTokens {
				_, err := s.Write(generatePart(t.possibleCharacters, t.quantifier))
				if err != nil {
					return fmt.Errorf("writing to string builder: %w", err)
				}
			}

			ch <- s.String()
			return nil
		})
	}

	go func() {
		err = g.Wait()
		close(ch)
	}()

	generated := make([]string, 0, count)
	for {
		s, ok := <-ch
		if !ok {
			break
		}
		generated = append(generated, s)
	}

	return generated, err
}

func generatePart(charList []byte, quantifier quantifier) []byte {
	if quantifier.randomPossibility && RandFn()%2 == 0 {
		return nil
	}

	count := quantifier.max // Fixed number of generations

	if quantifier.max != quantifier.min {
		// Generation range {1,3}
		count = (RandFn() % (quantifier.max - quantifier.min + 1)) + quantifier.min
	}

	bytes := make([]byte, 0, count)
	charListLen := len(charList)
	for i := 0; i < count; i++ {
		bytes = append(bytes, charList[RandFn()%charListLen])
	}

	return bytes
}
