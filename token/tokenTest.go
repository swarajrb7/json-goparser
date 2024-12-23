package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func testGetTokenKind(t *testing.T) {
	assert.Equal(t, "Bool", getTokenKind(JsonBool))
	assert.Equal(t, "Null", getTokenKind(JsonNull))
	assert.Equal(t, "Number", getTokenKind(JsonNumber))
	assert.Equal(t, "String", getTokenKind(JsonString))	
	assert.Equal(t, "json Syntax", getTokenKind(JsonSyntax))
}