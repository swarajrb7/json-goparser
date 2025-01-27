package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
	lexer "github.com/swarajrb7/json-goparser/lexer"
)

const InvalidJson = `{
	"key" : "value",	
	"key-n" : 101,
	"key-0" : {
		"inner-key" : "inner-value"
	},
	"key-1" : ['list value'], 
}`


func testParseInvalid(t *testing.T) {
	const  ValidJson = `{
		"key" : "value",
		"key-n" : 101,
		"key-0" : {
			"inner-key" : "inner-value"
		},
		"key-1" : ["list value"],
	}`

	expectedJson := map[string]any{
		"key": "value",
		"key-n": 101,
		"key-0": map[string]any{
			"inner-key": "inner-value",
		},
		"key-1": []any{"list value"},
	}

	tokens, err := lexer.Lexer(ValidJson)
	assert.Equal(t, nil, err)

	json, err := Parse(tokens)
	assert.Equal(t, nil, err)
	assert.Equal(t, expectedJson, json)
}
