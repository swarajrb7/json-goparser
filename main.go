package main

import (
	"fmt"
	"log"
	"os"
	lex"github.com/swarajrb7/json-goparser/lexer"
	parser"github.com/swarajrb7/json-goparser/parser"
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

	tokens, err := lex.Lexer(string(content))	
	if err != nil {
		log.Fatal(err)
	}

	_, err = parser.Parse(tokens)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	for _, token := range tokens {
		fmt.Println(token)
	}
}