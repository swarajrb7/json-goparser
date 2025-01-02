package parser

import (
	"log"
 tok "github.com/swarajrb7/json-goparser/token"
    "errors"
	"fmt"
)

func Parse(tokens []tok.Token) error {
	if len(tokens) == 0 {
		return errors.New("Parser Error: empty json string")
	}

	token := tokens[0]
	if token.Id != tok.JsonSyntax {
		return tokenError(token)
	}
	var err error
	switch token.Value {
	case "{":
		tokens, err = parseObject(tokens[1:])
		if err != nil {
			return err
		}
	case "[":
		tokens, err = parseArray(tokens[1:])
		if err != nil {
			return err
		}
	default:
		return tokenError(token)
	}	

	if len(tokens) > 0 {
		return tokenError(tokens[0])
	}
	return nil
} 

func parseObject(tokens []tok.Token) ([]tok.Token, error) {	  

	if len(tokens) == 0 {
		return []tok.Token{}, errors.New("Parser Error: unexpected End-of-Object brace '}' ")
	}

	json := map[string]any{}

	token := tokens[0] 
	if token.Id != tok.JsonSyntax  && token.Value != "}" {
		return tokens[1:], nil
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

	var err error

	for len(tokens) > 0 {
	 
		token = tokens[0]
		
		switch check {
		case checkKey:
			if token.Id !=  tok.JsonString {
				return []tok.Token{}, tokenError(token)
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
				return []tok.Token{}, tokenError(token)
			}
			tokens = tokens[1:]
			check = checkValue

		case checkValue:
			var value any
			if token.Id == tok.JsonSyntax {
				switch token.Value {
				case "{":
					tokens, err = parseObject(tokens[1:])
					if err != nil {
						return []tok.Token{}, err
					}
					json[currentKey] = value
				case "[":
					tokens, err = parseArray(tokens[1:])
					if err != nil {
						return []tok.Token{}, err
					}
					json[currentKey] = value
				default:
					return []tok.Token{}, tokenError(token)
				}
			} else {
				value, err := tok.ConvertTokenToType(token)
				if err != nil {
					return []tok.Token{}, tokenError(token)
				}
				json[currentKey] = value
				tokens = tokens[1:]
			}
			check = checkEnd
		case checkEnd:
			if token.Id != tok.JsonSyntax {
				return []tok.Token{}, tokenError(token)
			}

			switch token.Value {
			case ",":
				tokens = tokens[1:]
			case "}":
				return tokens[1:], nil			
			default:				
				return []tok.Token{}, tokenError(token)			
		}
			check = checkKey
		}
		
	}

	switch check {
	case checkKey:
		err = errors.New("Parser Error: Expected a key string '}' ")
	case checkColon:
		err = errors.New("Parser Error: Expected a colon ':' ")
	case checkValue:
		err = errors.New("Parser Error: Expected value ")
	default:
		err = errors.New("Parser Error: unexpected End-of-Object brace '}' ")
	}

	return []tok.Token{}, err
}

func parseArray(tokens []tok.Token) ([]tok.Token, error) {

	if len(tokens) == 0 {
		return []tok.Token{}, errors.New("Parser Error: unexpected End-of-Array bracket ']'")
	}

	token := tokens[0]
	if token.Id == tok.JsonSyntax && token.Value != "]" {
		return tokens[1:], nil
	}

	prevElement := false
	var err error

	for len(tokens) > 0 {		
		token = tokens[0]

		if token.Id == tok.JsonSyntax {
			switch {
			case token.Value == "[" && !prevElement:
				tokens,err  = parseArray(tokens[1:])
				if err != nil {
					return []tok.Token{}, err
				}
				prevElement = true
			case token.Value == "{" && !prevElement: 
				tokens, err = parseObject(tokens[1:])
				if err != nil {
					return []tok.Token{}, err
				}
				prevElement = true
			case token.Value == "]" && prevElement:
				return tokens[1:], nil
			case token.Value == "," && prevElement: 
				prevElement = false
				tokens = tokens[1:]
			default:
				return []tok.Token{}, tokenError(token)
			}
		}else if prevElement {
			return []tok.Token{}, tokenError(token)
		}else {
			prevElement = true
			tokens = tokens[1:]
		}
	}
	log.Fatalf("Parser Error: unexpected End-of-Array bracket ']'")
	return []tok.Token{}, errors.New("Parser Error: unexpected End-of-Array bracket ']'")
}

func tokenError(token tok.Token) error {
    return fmt.Errorf("Parser Error: unexpected %s token '%s' at line %d col %d", 
        tok.GetTokenKind(token.Id), 
        token.Value, 
        token.LineNum, 
        token.ColNum)
}