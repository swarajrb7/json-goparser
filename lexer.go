package main

import (
	"log"
	"unicode"
	"regexp"
)

type tokenID int32

const (
	tokenBool      tokenID = iota
	tokenNull      tokenID = iota
	tokenNumber    tokenID = iota
	tokenString    tokenID = iota
	typeJsonSyntax  tokenID = iota
)				

				
type Token struct {
	id tokenID
	value string
}

var (
	jsonTrue = []rune("true")
	jsonFalse = []rune("false")
	jsonNull = []rune("null")
)

var jsonSyntaxToken = map[rune]struct{} {
	'{' : {},
	'}' : {},
	'[' : {},
	']' : {},
	':' : {},
	',' : {},
}

func lexer(s string) []Token {	
	var tokens []Token	
	lineNum := 1
	colNum := 1


	//rune is used to ensure that the string is a treated as a sequence of full Unicode characters, not just raw bytes.	

	runes := []rune(s)

	for len(runes) > 0 {
		char := runes[0]
		if unicode.IsSpace(char) {
			colNum++
			if char == '\n' {				
				lineNum++
				colNum = 1
			}
			runes = runes[1:]
			continue
		}

		var token Token
		var ok bool

		token, runes, ok = lexString(runes)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.value) + 2
			continue
		}

		token, runes, ok = lexNumber(runes)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.value) 
			continue
		}

		token, runes, ok = lexBool(runes)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.value)
			continue
		}				
		
		token, runes, ok = lexNull(runes)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.value)
			continue
		}
		_, ok = JsonSyntaxTokens[char]
		if ok {
			token = append(tokens, Token{typeJsonSyntax, string(char)})
			colNum++
			runes = runes[1:]
		} else {			
			log.Fatalf("lexer error: char %c at line %d col %d is invalid" , string(char), lineNum, colNum)
		}
	}
	return tokens
	
}

func lexString(runes []rune) (Token, []rune, bool) {
	if runes[0] != '"' {
		return Token{}, runes, false
	}
	rune := runes[1:]
	for i, char := range rune {
		if char == '"' {
			return Token{tokenString, string(runes[:i+1])}, rune[i+1:], true
		}
	}

	log.Fatalf("lexer error: EOF quote.")
	return Token{}, runes, false
}

func lexNumber(runes []rune) (Token, []rune, bool) {
	if !unicode.IsDigit(runes[0]) {
		return Token{}, runes, false
	}

	var end int = len(runes)
	for i, char := range runes {
		if !unicode.IsDigit(char) && char != '.' && char != 'e' && char != 'E' {
			end = i -1
			break
		}
	}

	tokenValue := string(runes[:end+1])
	if !regexp.MustCompile(`^\d+(?:\.\d+)?(?:e\d+)?$)`).MatchString(tokenValue) {
		log.Fatal("lexer error: invalid number  %s", tokenValue)
	}
	return Token{tokenNumber, tokenValue}, runes[end+1:], true

}

func lexBool(runes []rune) (Token, []rune, bool) {
	if CompareRuneSlice(runes, jsonTrue, len(jsonTrue)) {
		return Token{tokenBoolean, string(runes[:len(JsonTrue)])}, runes[len(JsonTrue):], true
	}

	if CompareRuneSlice(runes, JsonFalse, len(JsonFalse)) {
		return Token{tokenBoolean, string(runes[:len(JsonFalse)])}, runes[len(JsonFalse):], true
	}
	return Token{}, runes, false
}

func lexNull(runes []rune) (Token, []rune, bool) {
	if CompareRuneSlice(runes, jsonNull, len(jsonNull)) {
		return Token{tokenNull, string(runes[:len(jsonNull)])}, runes[len(jsonNull):], true
	}
	return Token{}, runes, false
}

































