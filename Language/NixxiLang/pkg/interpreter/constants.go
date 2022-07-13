package interpreter

type TokenType string
type TokenValue string
type VarType int8
type OperatorValue string
type Keyword string

const (
	OPERATOR_UNKNOW            OperatorValue = "UNKNOW"
	OPERATOR_PLUS              OperatorValue = "PLUS"
	OPERATOR_MINUS             OperatorValue = "MINUS"
	OPERATOR_MULTIPLY          OperatorValue = "MULTIPLY"
	OPERATOR_DIVIDE            OperatorValue = "DIVIDE"
	OPERATOR_ASSIGN            OperatorValue = "ASSIGN"
	OPERATOR_EQUAL             OperatorValue = "EQUAL"
	OPERATOR_NOT_EQUAL         OperatorValue = "NOT_EQUAL"
	OPERATOR_MORETHAN          OperatorValue = "MORETHAN"
	OPERATOR_LESS              OperatorValue = "LESS"
	OPERATOR_MORETHAN_OR_EQUAL OperatorValue = "MORETHAN_OR_EQUAL"
	OPERATOR_LESS_OR_EQUAL     OperatorValue = "LESS_OR_EQUAL"
	OPERATOR_OPEN_BRACKET      OperatorValue = "OPEN_BRACKET"
	OPERATOR_CLOSE_BRACKET     OperatorValue = "CLOSE_BRACKET"
)

const (
	VARTYPE_NULL    VarType = 0
	VARTYPE_STRING  VarType = 1
	VARTYPE_NUMBER  VarType = 2
	VARTYPE_BOOLEAN VarType = 3
	VARTYPE_OBJECT  VarType = 4
	VARTYPE_ARRAY   VarType = 5
)

const (
	KEYWORD_UNKNOW Keyword = "unknow"
	KEYWORD_VAR    Keyword = "var"
	KEYWORD_MUT    Keyword = "mut"
	KEYWORD_IF     Keyword = "if"
	KEYWORD_ELSE   Keyword = "else"
	KEYWORD_ELSEIF Keyword = "elif"
	KEYWORD_LOOP   Keyword = "loop"
	KEYWORD_BREAK  Keyword = "break"
	KEYWORD_RETURN Keyword = "return"
	// KEYWORD_USE    Keyword = "use"
	// KEYWORD_CLASS      Keyword = "class"
	// KEYWORD_INTERFACE  Keyword = "interface"
	// KEYWORD_IMPLEMENTS Keyword = "implements"
	// KEYWORD_UNSAFE     Keyword = "unsafe"
)

const (
	TOKEN_UNKNOW   TokenType = "UNKNOW"
	TOEKN_SYMBOL   TokenType = "SYMBOL"
	TOKEN_NUMBER   TokenType = "NUMBER"
	TOKEN_STRING   TokenType = "STRING"
	TOKEN_BOOLEAN  TokenType = "BOOLEAN"
	TOKEN_OPERATOR TokenType = "OPERATOR"
	TOKEN_KEYWORD  TokenType = "KEYWORD"
)

type Token struct {
	Type  TokenType  `json:"type"`
	Value TokenValue `json:"value"`
}
