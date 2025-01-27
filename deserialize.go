package main

import (
	lex "github.com/swarajrb7/json-goparser/lexer"
	parser"github.com/swarajrb7/json-goparser/parser"
)

func Deserialize(data string) (any, error) {
	var json any
	var err error

	tokens, err := lex.Lexer(data)
	if err != nil {
		return nil, err
	}
	
   json,err = parser.Parse(tokens)
   if err != nil{
		return nil, err
   }
	
	return json, nil
}