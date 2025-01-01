package token

import (
	"fmt"
	"strconv"
)


type tokenID int32
type Token struct {
	Id tokenID
	Value string
	LineNum int	
	ColNum int
}

const (
	JsonBool       tokenID = iota
	JsonNull       
	JsonNumber
	JsonString     
	JsonSyntax     
)

var JsonSyntaxChars = map[rune]struct{} {	
	'{' : {},
	'}' : {},
	'[' : {},
	']' : {},
	':' : {},
	',' : {},
}

func GetTokenKind(kind tokenID) string {
	switch kind {
	case JsonBool:
		return "tokenBool"
	case JsonNull:
		return "jsonNull"
	case JsonNumber:
		return "jsonNumber"
	case JsonString:
		return "jsonString"
	case JsonSyntax:
		return "jsonSyntax"
	default:
		return "unknown"
	}
	
}

func ConvertTokenToType(token Token) (any, error) {
	
	var value any
	var err error

	switch token.Id {
	case JsonBool:
		value = token.Value == "true"
	case JsonNull:
		value = nil
	case JsonNumber:
		value, err = strconv.ParseFloat(token.Value,64)
	case JsonString:
		value = token.Value
	default:
		err = tokenError(token)
	}

	if err != nil {
		return nil, err
	}

	return value, nil
}

func tokenError(token Token)  error  {
	return fmt.Errorf("Parser Error: unexpected %s token '%s' at line %d col %d", GetTokenKind(token.Id), token.Value, token.LineNum, token.ColNum)
}
