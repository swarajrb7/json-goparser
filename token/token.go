package token

type Token struct {
	Id tokenID
	Value string
	LineNum int	
	ColNum int
}

type tokenID int32

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