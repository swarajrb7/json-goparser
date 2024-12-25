package parser

import (
	"os"
	"os/exec"
	"testing"
	"github.com/stretchr/testify/assert"
	lexer "github.com/swarajrb7/json-goparser/lexer"
)

const  ValidJson = `{
	"key" : "value",
	"key-n" : 101,
	"key-0" : {
		"inner-key" : "inner-value"
	},
	"key-1" : ["list value"],
}`

const InvalidJson = `{
	"key" : "value",	
	"key-n" : 101,
	"key-0" : {
		"inner-key" : "inner-value"
	},
	"key-1" : ['list value'], 
}`

func testParse(t *testing.T) {
	Parse(lexer.Lexer(ValidJson))
}

func testParseInvalid(t *testing.T) {
	if os.Getenv("FLAG") != "1" { 
		Parse(lexer.Lexer(InvalidJson))
		return
	} 
	
	cmd := exec.Command("go", "test")
	cmd.Env = append(os.Environ(), "FLAG=1")
	
	err := cmd.Run()
	
	assert.NotEqual(t, nil, err)
	if err != nil {
		assert.Equal(t, "exit status 1", err.Error())
	}
}