package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/toncek345/reggenerator"
)

func main() {
	regex := flag.String("regex", "", "regex for random string")
	count := flag.Int("count", 1, "number of random string")
	flag.Parse()

	strings, err := reggenerator.Generate(*regex, *count)
	if err != nil {
		log.Fatalf("error happened: %s", err)
	}

	for _, s := range strings {
		fmt.Printf("%s\n", s)
	}
}
