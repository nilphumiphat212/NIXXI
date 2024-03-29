package interpreter

import (
	"os"
	"testing"
)

func TestIsSymbol(t *testing.T) {
	symbol := isSymbol("nil")
	if !symbol {
		t.Error("Can not detect symbol from text")
	}

	notSymbol := isSymbol("\"nil\"")
	if notSymbol {
		t.Error("Can not defect difference string or symbol")
	}
}

func TestIsNumberShouldReturnTrue(t *testing.T) {
	res := isNumber("123456")

	if !res {
		t.Error("True result is expected")
	}
}

func TestIsNumberShouldReturnFalse(t *testing.T) {
	res := isNumber("ABCdEFghi")

	if res {
		t.Error("False result is expected")
	}
}

func TestIsString(t *testing.T) {
	isAString := isString("\"test\"")

	if !isAString {
		t.Error("True result is expected")
	}

	isNotString := isString("2022")

	if isNotString {
		t.Error("False result is expected")
	}
}

func TestCanFilterOperator(t *testing.T) {
	isPlus, plus := isOperator("+")
	if !isPlus || plus != OPERATOR_PLUS {
		t.Error("Can not filter 'Plus' operator")
	}

	isMinus, minus := isOperator("-")
	if !isMinus || minus != OPERATOR_MINUS {
		t.Error("Can not filter 'Minus' operator")
	}

	isMultiply, multiply := isOperator("*")
	if !isMultiply || multiply != OPERATOR_MULTIPLY {
		t.Error("Can not filter 'Multiply' operator")
	}

	isDivide, divide := isOperator("/")
	if !isDivide || divide != OPERATOR_DIVIDE {
		t.Error("Can not filter 'Divide' operator")
	}

	isAssign, assign := isOperator("=")
	if !isAssign || assign != OPERATOR_ASSIGN {
		t.Error("Can not filter 'Assign' operator")
	}

	isEqual, equal := isOperator("==")
	if !isEqual || equal != OPERATOR_EQUAL {
		t.Error("Can not filter 'Equal' operator")
	}

	isNotEqual, notequal := isOperator("!=")
	if !isNotEqual || notequal != OPERATOR_NOT_EQUAL {
		t.Error("Can not filter 'Not Equal' operator")
	}

	isMorethan, morethan := isOperator(">")
	if !isMorethan || morethan != OPERATOR_MORETHAN {
		t.Error("Can not filter 'Morethan' operator")
	}

	isLess, less := isOperator("<")
	if !isLess || less != OPERATOR_LESS {
		t.Error("Can not filter 'Less' operator")
	}

	isMorethanOrEqual, morethanOrEqual := isOperator(">=")
	if !isMorethanOrEqual || morethanOrEqual != OPERATOR_MORETHAN_OR_EQUAL {
		t.Error("Can not filter 'Morehan or Equal' operator")
	}

	isLessnOrEqual, lessOrEqual := isOperator("<=")
	if !isLessnOrEqual || lessOrEqual != OPERATOR_LESS_OR_EQUAL {
		t.Error("Can not filter 'Less or Equal' operator")
	}

	isOpenBracket, openBracket := isOperator("(")
	if !isOpenBracket || openBracket != OPERATOR_OPEN_BRACKET {
		t.Error("Can not filter 'Open Bracket or Equal' operator")
	}

	isCloseBracket, closeBracket := isOperator(")")
	if !isCloseBracket || closeBracket != OPERATOR_CLOSE_BRACKET {
		t.Error("Can not filter 'Close Bracket or Equal' operator")
	}
}

func TestIsKeywordCanfilterWithString(t *testing.T) {
	isNotKeyword, unknowKeyword := isKeyword("test")
	if isNotKeyword && unknowKeyword != KEYWORD_UNKNOW {
		t.Error("Can not filter wrong string")
	}

	isVar, varKeyword := isKeyword("var")
	if !isVar || varKeyword != KEYWORD_VAR {
		t.Error("Can not filter 'var' keyword")
	}

	isMut, mutKeyword := isKeyword("mut")
	if !isMut || mutKeyword != KEYWORD_MUT {
		t.Error("Can not filter 'mut' keyword")
	}

	isIf, ifKeyword := isKeyword("if")
	if !isIf || ifKeyword != KEYWORD_IF {
		t.Error("Can not filter 'if' keyword")
	}

	isElse, elseKeyword := isKeyword("else")
	if !isElse || elseKeyword != KEYWORD_ELSE {
		t.Error("Can not filter 'else' keyword")
	}

	isElseIf, elseIfKeyword := isKeyword("elif")
	if !isElseIf || elseIfKeyword != KEYWORD_ELSEIF {
		t.Error("Can not filter 'elif' keyword")
	}

	isLoop, loopKeyword := isKeyword("loop")
	if !isLoop || loopKeyword != KEYWORD_LOOP {
		t.Error("Can not filter 'loop' keyword")
	}

	isBreak, breakKeyword := isKeyword("break")
	if !isBreak || breakKeyword != KEYWORD_BREAK {
		t.Error("Can not filter 'break' keyword")
	}

	isReturn, returnKeyword := isKeyword("return")
	if !isReturn || returnKeyword != KEYWORD_RETURN {
		t.Error("Can not filter 'return' keyword")
	}
}

func TestIsBoolean(t *testing.T) {
	trueToken := isBoolean("true")
	if !trueToken {
		t.Error("Can not filter boolean (true)")
	}

	falseToken := isBoolean("true")
	if !falseToken {
		t.Error("Can not filter boolean (false)")
	}

	notBoolean := isBoolean("test")
	if notBoolean {
		t.Error("Can not filter wrong string")
	}
}

func TestCanGetAllWordingOfTextLine(t *testing.T) {
	textLine := "This is a    wording"
	result := getWordingsOfTextLine(textLine)

	if len(result) != 4 || result[0] != "This" || result[1] != "is" || result[2] != "a" || result[3] != "wording" {
		t.Error("Wordings from function is not collect")
	}
}

func TestCanFilterToken(t *testing.T) {
	numberToken, err := toToken("2022")
	if err != nil || numberToken.Type != TOKEN_NUMBER {
		t.Error("Can not filter number token")
	}

	stringToken, err := toToken("\"Create by nilphumiphat\"")
	if err != nil || stringToken.Type != TOKEN_STRING {
		t.Error("Can not filter string token")
	}

	booleanToken, err := toToken("true")
	if err != nil || booleanToken.Type != TOKEN_BOOLEAN {
		t.Error("Can not filter boolean token")
	}

	keywordToken, err := toToken("var")
	if err != nil || keywordToken.Type != TOKEN_KEYWORD {
		t.Error("Can not filter keyword token")
	}

	unexpectedToken, err := toToken("`^@$")
	if err == nil || unexpectedToken.Type != TOKEN_UNKNOW {
		t.Error("Can not detect unexpet token")
	}
}

func TestLoadToLexerFromFile(t *testing.T) {
	tokens, err := LoadToLexerFromFile("")
	if err == nil || tokens != nil {
		t.Error("Can not detect empty source path")
	}

	createMockFileErr := os.WriteFile("test.nixxi", []byte("var year = 2022"), 0644)
	if createMockFileErr != nil {
		t.Error("Mock file not create")
	}
	tokens2, err2 := LoadToLexerFromFile("test.nixxi")
	if err2 != nil || tokens2 == nil {
		t.Error("Can not pass source file with convert to token array")
	}
	os.Remove("test.nixxi")
}
