package lexer

import (
	"log"
	"unicode"
	"regexp"
	tok "github.com/swarajrb7/json-goparser/token"
	"github.com/swarajrb7/json-goparser/utils"
)

var (
	jsonTrue = []rune("true")
	jsonFalse = []rune("false")
	jsonNull = []rune("null")
)

func Lexer(s string) []tok.Token {	
	tokens := []tok.Token{}
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

		var token tok.Token
		var ok bool

		token, runes, ok = lexString(runes, lineNum, colNum)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.Value) + 2
			continue
		}

		token, runes, ok = lexNumber(runes, lineNum, colNum)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.Value) 
			continue
		}

		token, runes, ok = lexBool(runes, lineNum, colNum)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.Value)
			continue
		}				
		
		token, runes, ok = lexNull(runes, lineNum, colNum)
		if ok {
			tokens = append(tokens, token)
			colNum += len(token.Value)
			continue
		}
		_, ok = tok.JsonSyntaxChars[char]
		if ok {
			tokens = append(tokens, tok.Token{tok.JsonSyntax ,string(char),lineNum, colNum})
			colNum++
			runes = runes[1:]
		} else {			
			log.Fatalf("lexer error: char %s at line %d col %d is invalid" , string(char), lineNum, colNum)
		}
	}
	return tokens

}

func lexString(runes []rune, lineNum int, colNum int) (tok.Token, []rune, bool) {
	if runes[0] != '"' {
		return tok.Token{}, runes, false
	}
	rune := runes[1:]
	for i, char := range rune {
		if char == '"' {
			return tok.Token{tok.JsonString, string(runes[:i+1]), lineNum, colNum}, rune[i+1:], true
		}
	}

	log.Fatalf("lexer error: EOF quote.")
	return tok.Token{}, runes, false
}

func lexNumber(runes []rune, lineNum int, colNum int) (tok.Token, []rune, bool) {
	if !unicode.IsDigit(runes[0]) {
		return tok.Token{}, runes, false
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
		log.Fatalf("lexer error: invalid number  %s", tokenValue)
	}
	return tok.Token{tok.JsonNumber, tokenValue, lineNum, colNum}, runes[end+1:], true

}

func lexBool(runes []rune, lineNum int, colNum int) (tok.Token, []rune, bool) {
	if utils.CompareRuneSlice(runes, jsonTrue, len(jsonTrue)) {
		return tok.Token{tok.JsonBool, string(runes[:len(jsonTrue)]), lineNum, colNum}, runes[len(jsonTrue):], true
	}

	if utils.CompareRuneSlice(runes, jsonFalse, len(jsonFalse)) {
		return tok.Token{tok.JsonBool, string(runes[:len(jsonFalse)]), lineNum, colNum}, runes[len(jsonFalse):], true
	}
	return tok.Token{}, runes, false
}

func lexNull(runes []rune, lineNum int, colNum int) (tok.Token, []rune, bool) {
	if utils.CompareRuneSlice(runes, jsonNull, len(jsonNull)) {
		return tok.Token{tok.JsonNull, string(runes[:len(jsonNull)]), lineNum, colNum}, runes[len(jsonNull):], true
	}
	return tok.Token{}, runes, false
}

































