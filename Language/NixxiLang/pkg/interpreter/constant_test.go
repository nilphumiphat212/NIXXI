package interpreter

import "testing"

func TestOperatorConstant(t *testing.T) {
	unknow := string(OPERATOR_UNKNOW)
	if unknow != "UNKNOW" {
		t.Error("'UNKNOW' is expected")
	}

	plus := string(OPERATOR_PLUS)
	if plus != "PLUS" {
		t.Error("'PLUS' is expected")
	}

	minus := string(OPERATOR_MINUS)
	if minus != "MINUS" {
		t.Error("'MINUS' is expected")
	}

	multiply := string(OPERATOR_MULTIPLY)
	if multiply != "MULTIPLY" {
		t.Error("'MULTIPLY' is expected")
	}

	divide := string(OPERATOR_DIVIDE)
	if divide != "DIVIDE" {
		t.Error("'DIVIDE' is expected")
	}

	assign := string(OPERATOR_ASSIGN)
	if assign != "ASSIGN" {
		t.Error("'ASSIGN' is expected")
	}
}

func TestVarTypeConstant(t *testing.T) {
	null := VARTYPE_NULL
	if int(null) != 0 {
		t.Error("0 is expect of null type")
	}

	stringType := VARTYPE_STRING
	if int(stringType) != 1 {
		t.Error("1 is expect of string type")
	}

	numberType := VARTYPE_NUMBER
	if int(numberType) != 2 {
		t.Error("2 is expect of number type")
	}

	booleanType := VARTYPE_BOOLEAN
	if int(booleanType) != 3 {
		t.Error("3 is expect of boolean type")
	}

	objectType := VARTYPE_OBJECT
	if int(objectType) != 4 {
		t.Error("4 is expect of object type")
	}

	arrayType := VARTYPE_ARRAY
	if int(arrayType) != 5 {
		t.Error("5 is expect of object type")
	}
}

func TestKeywordConstant(t *testing.T) {
	varKeyword := KEYWORD_VAR
	if string(varKeyword) != "var" {
		t.Error("'var' is expected of var keyword")
	}

	mutKeyword := KEYWORD_MUT
	if string(mutKeyword) != "mut" {
		t.Error("'mut' is expected of mut keyword")
	}

	ifKeyword := KEYWORD_IF
	if string(ifKeyword) != "if" {
		t.Error("'if' is expected of if keyword")
	}

	elseKeyword := KEYWORD_ELSE
	if string(elseKeyword) != "else" {
		t.Error("'else' is expected of else keyword")
	}

	elseifKeyword := KEYWORD_ELSEIF
	if string(elseifKeyword) != "elif" {
		t.Error("'elif' is expected of else if keyword")
	}

	loopKeyword := KEYWORD_LOOP
	if string(loopKeyword) != "loop" {
		t.Error("'loop' is expected of loop keyword")
	}
}

func TestTokenConstant(t *testing.T) {
	unknow := TOKEN_UNKNOW
	if string(unknow) != "UNKNOW" {
		t.Error("'UNKNOW' is expected")
	}

	symbol := TOEKN_SYMBOL
	if string(symbol) != "SYMBOL" {
		t.Error("'SYMBOL' is expected")
	}

	number := TOKEN_NUMBER
	if string(number) != "NUMBER" {
		t.Error("'NUMBER' is expected")
	}

	stringTk := TOKEN_STRING
	if string(stringTk) != "STRING" {
		t.Error("'STRING' is expected")
	}

	boolean := TOKEN_BOOLEAN
	if string(boolean) != "BOOLEAN" {
		t.Error("'BOOLEAN' is expected")
	}

	operator := TOKEN_OPERATOR
	if string(operator) != "OPERATOR" {
		t.Error("'OPERATOR' is expected")
	}

	keyword := TOKEN_KEYWORD
	if string(keyword) != "KEYWORD" {
		t.Error("'KEYWORD' is expected")
	}
}

func TestTokenStruct(t *testing.T) {
	var test Token = Token{TOKEN_UNKNOW, TokenValue("test")}
	if test.Type != TOKEN_UNKNOW || test.Value != TokenValue("test") {
		t.Error("Can not set type or value to 'Token' struct")
	}
}
