package main

import (
	"fmt"
	"log"
	"os"
	"github.com/swarajrb7/json-goparser/lexer"
)

func main() {
	log.SetPrefix("Error:  ")

	if len(os.Args) < 2 {
		log.Fatalf("usage: json-parser %s", os.Args[0])
	}

	content, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	tokens := lexer.Lexer(string(content))

	for _, token := range tokens {
		fmt.Println(token)
	}

	parse(tokens)
}