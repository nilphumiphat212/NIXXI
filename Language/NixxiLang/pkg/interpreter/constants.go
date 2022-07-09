package interpreter

type TokenType string
type TokenValue string
type VarType int
type Keyword string

const (
	OPARATOR_PLUS     = "PLUS"
	OPARATOR_MINUS    = "MINUS"
	OPARATOR_MULTIPLY = "MULTIPLY"
	OPARATOR_DIVIDE   = "DIVIDE"
)

const (
	VARTYPE_NULL    VarType = 0
	VARTYPE_STRING  VarType = 1
	VARTYPE_BOOLEAN VarType = 2
	VARTYPE_OBJECT  VarType = 3
	VARTYPE_ARRAY   VarType = 4
)

const (
	KEYWORD_UNKNOW Keyword = "unknow"
	KEYWORD_VAR    Keyword = "var"
	KEYWORD_MUT    Keyword = "mut"
	KEYWORD_IF     Keyword = "if"
	KEYWORD_ELSE   Keyword = "else"
	KEYWORD_ELSEIF Keyword = "elif"
	KEYWORD_LOOP   Keyword = "loop"
	// KEYWORD_CLASS      Keyword = "class"
	// KEYWORD_INTERFACE  Keyword = "interface"
	// KEYWORD_IMPLEMENTS Keyword = "implements"
	// KEYWORD_UNSAFE     Keyword = "unsafe"
)

const (
	TOKEN_KEYWORD  TokenType = "KEYWORD"
	TOKEN_OPARATOR TokenType = "OPARATOR"
	TOKEN_STRING   TokenType = "STRING"
	TOKEN_NUMBER   TokenType = "NUMBER"
	TOKEN_BOOLEAN  TokenType = "BOOLEAN"
)

type Token struct {
	Type  TokenType
	Value TokenValue
}
