package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompareRuneSlice(t *testing.T) {
	assert.True(t, CompareRuneSlice([]rune("123"), []rune("123"), 3))
	assert.True(t, CompareRuneSlice([]rune("123"), []rune("123"), 2))
	assert.False(t, CompareRuneSlice([]rune("123"), []rune("1234"), 3))
	assert.False(t, CompareRuneSlice([]rune("123"), []rune("1234"), 4))
}