package parser

import (
	"log"
 tok "github.com/swarajrb7/json-goparser/token"
)

func Parse(tokens []tok.Token) {
	if len(tokens) == 0 {
		log.Fatalf("Parser Error: empty json string")
	}

	token := tokens[0]
	if token.Id != tok.JsonSyntax {
		log.Fatalf("Unexpected %s token '%s' at line %d col %d", tok.GetTokenKind(token.Id), token.Value, token.LineNum, token.ColNum)
	}
	switch token.Value {
	case "{":
		tokens = parseObject(tokens[1:])
	case "[":
		tokens = parseArray(tokens[1:])
	default:
		log.Fatalf("Unexpected %s token '%s' at line %d col %d", tok.GetTokenKind(token.Id), token.Value, token.LineNum, token.ColNum)
	}	

	if len(tokens) > 0 {
		log.Fatalf("Parser Error: unexpected token '%s' at line %d col %d", token.Value, token.LineNum, token.ColNum)
	}
} 

func parseObject(tokens []tok.Token) []tok.Token {	

	if len(tokens) == 0 {
		log.Fatalf("Parser Error: unexpected End-of-Object brace '}' ")
	}

	json := map[string]any{}

	token := tokens[0] 
	if token.Id != tok.JsonSyntax  && token.Value != "}" {
		return tokens[1:]
	}

	keys := map[string]struct{}{}

	const (
		checkKey   = iota
		checkColon  = iota
		checkValue = iota
		checkEnd   = iota 
	)

	var check = checkKey

	var currentKey string

	for len(tokens) > 0 {
	 
		token = tokens[0]
		
		switch check {
		case checkKey:
			if token.Id !=  tok.JsonString {
				tokenError(token)
			}
			_, ok := keys[token.Value]
			if ok {
				log.Fatal("Parser Error: duplicate key '%s' ", token.Value)
			}
			keys[token.Value] = struct{}{}
			currentKey = token.Value
			tokens = tokens[1:]
			check = checkColon

		case checkColon:
			if token.Id != tok.JsonSyntax || (token.Id == tok.JsonSyntax && token.Value != ":") {
				tokenError(token)
			}
			tokens = tokens[1:]
			check = checkValue

		case checkValue:
			var value any
			if token.Id == tok.JsonSyntax {
				switch token.Value {
				case "{":
					tokens = parseObject(tokens[1:])
					json[currentKey] = value
				case "[":
					tokens = parseArray(tokens[1:])
					json[currentKey] = value
				default:
					tokenError(token)
				}
			} else {
				value, err := tok.ConvertTokenToType(token)
				if err != nil {
					return []tok.Token{}
				}
				json[currentKey] = value
				tokens = tokens[1:]
			}
			check = checkEnd
		case checkEnd:
			if token.Id != tok.JsonSyntax {
				tokenError(token)
			}

			switch token.Value {
			case ",":
				tokens = tokens[1:]
			case "}":
				return tokens[1:]			
			default:				
				tokenError(token)			
		}
			check = checkKey
		}
		
	}

	switch check {
	case checkKey:
		log.Fatal("Parser Error: Expected a key string '}' ")
	case checkColon:
		log.Fatal("Parser Error: Expected a colon ':' ")
	case checkValue:
		log.Fatal("Parser Error: Expected value ")
	default:
		log.Fatal("Parser Error: unexpected End-of-Object brace '}' ")
	}

	return []tok.Token{}
}

func parseArray(tokens []tok.Token) []tok.Token {

	if len(tokens) == 0 {
		log.Fatalf("Parser Error: unexpected End-of-Array bracket ']'")
	}

	token := tokens[0]
	if token.Id == tok.JsonSyntax && token.Value != "]" {
		return tokens[1:]
	}

	prevElement := false
	for len(tokens) > 0 {		
		token = tokens[0]

		if token.Id == tok.JsonSyntax {
			switch {
			case token.Value == "[" && !prevElement:
				tokens = parseArray(tokens[1:])
				prevElement = true
			case token.Value == "{" && !prevElement: 
				tokens = parseObject(tokens[1:])
				prevElement = true
			case token.Value == "]" && prevElement:
				return tokens[1:]
			case token.Value == "," && prevElement: 
				prevElement = false
				tokens = tokens[1:]
			default:
				tokenError(token)
			}
		}else if prevElement {
			tokenError(token)
		}else {
			prevElement = true
			tokens = tokens[1:]
		}
	}
	log.Fatalf("Parser Error: unexpected End-of-Array bracket ']'")
	return []tok.Token{}
}

func tokenError(token tok.Token) {
	log.Fatalf("Parser Error: unexpected %s token '%s' at line %d col %d", tok.GetTokenKind(token.Id), token.Value, token.LineNum, token.ColNum)
}