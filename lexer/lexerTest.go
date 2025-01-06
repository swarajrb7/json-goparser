package lexer

import (
	"testing"
	"github.com/stretchr/testify/assert"
	tok "github.com/swarajrb7/json-goparser/token"
)

const json = `{
	"key" : "value",
	"key-n" : 101,
	"key-0" : {
		"inner-key" : "inner-value"
	},
	"key-1" : ["list value"],
}`

func testLexer(t *testing.T) {
	expectedTokens := []tok.Token{
		{tok.JsonSyntax, "{", 1, 1},
		{tok.JsonString, `"key"`, 1, 3},
		{tok.JsonSyntax, ":", 1, 8},
		{tok.JsonString, `"value"`, 1, 10},
		{tok.JsonSyntax, ",", 1, 16},
		{tok.JsonString, `"key-n"`, 1, 18},
		{tok.JsonSyntax, ":", 1, 25},
		{tok.JsonNumber, "101", 1, 27},
		{tok.JsonSyntax, ",", 1, 31},
		{tok.JsonString, `"key-0"`, 1, 33},
		{tok.JsonSyntax, ":", 1, 40},
		{tok.JsonSyntax, "{", 1, 42},
		{tok.JsonString, `"inner-key"`, 1, 44},
		{tok.JsonSyntax, ":", 1, 54},
		{tok.JsonString, `"inner-value"`, 1, 56},
		{tok.JsonSyntax, "}", 1, 66},
		{tok.JsonSyntax, ",", 1, 68},
		{tok.JsonString, `"key-1"`, 1, 70},
		{tok.JsonSyntax, ":", 1, 77},
		{tok.JsonSyntax, "[", 1, 79},
		{tok.JsonString, `"list value"`, 1, 81},
		{tok.JsonSyntax, "]", 1, 92},
		{tok.JsonSyntax, "}", 1, 94},
	}

	tokens, err := Lexer(json)
	assert.NoError(t, err)
	assert.Equal(t, expectedTokens, tokens)
}