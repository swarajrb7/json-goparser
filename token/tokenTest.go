package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetTokenKind(t *testing.T) {
	assert.Equal(t, "Bool", GetTokenKind(JsonBool))
	assert.Equal(t, "Null", GetTokenKind(JsonNull))
	assert.Equal(t, "Number", GetTokenKind(JsonNumber))
	assert.Equal(t, "String", GetTokenKind(JsonString))	
	assert.Equal(t, "json Syntax", GetTokenKind(JsonSyntax))
}