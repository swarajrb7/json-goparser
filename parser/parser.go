package main

import (
	"log"
)

func parse(tokens []Token) {
	if len(tokens) == 0 {
		log.Fatalf("Parser Error: empty json string")
	}

	token := tokens[0]
	if token.id != typeJsonSyntax {
		log.Fatalf("Parser Error: invalid json string")
	}
	switch token.value {
	case "{":
		parseObject(tokens)
	case "[":
		parseArray(tokens)
	default:
		log.Fatalf("Parser Error: invalid json string")
	}	
}