package main

import (
	"log"
	tok "github.com/swarajrb7/json-goparser/token"
)

func parse(tokens []tok.Token) {
	if len(tokens) == 0 {
		log.Fatalf("Parser Error: empty json string")
	}

	token := tokens[0]
	if token.Id != tok.JsonSyntax {
		log.Fatalf("Parser Error: invalid json string")
	}
	switch token.Value {
	case "{":
		parseObject(tokens)
	case "[":
		parseArray(tokens)
	default:
		log.Fatalf("Parser Error: invalid json string")
	}	
} 