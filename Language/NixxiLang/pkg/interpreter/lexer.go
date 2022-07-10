package interpreter

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
)

func isSymbol(text string) bool {
	var re *regexp.Regexp = regexp.MustCompile("[a-zA-Z]+$")
	if len(re.FindAllString(text, -1)) > 0 {
		return true && !isString(text)
	} else {
		return false
	}
}

func isNumber(text string) bool {
	var re *regexp.Regexp = regexp.MustCompile("[0-9]+")
	if len(re.FindAllString(text, -1)) > 0 {
		return true
	} else {
		return false
	}
}

func isString(text string) bool {
	var runes []rune = []rune(text)
	return string(runes[0]) == "\"" && string(runes[len(runes)-1]) == "\""
}

func isBoolean(text string) bool {
	return text == "true" || text == "false"
}

func isOperator(text string) (bool, OperatorValue) {
	if text == "+" {
		return true, OPERATOR_PLUS
	} else if text == "-" {
		return true, OPERATOR_MINUS
	} else if text == "*" {
		return true, OPERATOR_MULTIPLY
	} else if text == "/" {
		return true, OPERATOR_DIVIDE
	} else if text == "=" {
		return true, OPERATOR_ASSIGN
	}

	return false, OPERATOR_UNKNOW
}

func isKeyword(text string) (bool, Keyword) {
	var allKeywords []string = []string{
		string(KEYWORD_VAR),
		string(KEYWORD_MUT),
		string(KEYWORD_IF),
		string(KEYWORD_ELSE),
		string(KEYWORD_ELSEIF),
		string(KEYWORD_LOOP),
	}

	for _, keyword := range allKeywords {
		if text == keyword {
			return true, Keyword(keyword)
		}
	}

	return false, KEYWORD_UNKNOW
}

func getWordingsOfTextLine(textLine string) []string {
	var wordings []string
	var wording string = ""
	for _, char := range textLine + " " {
		var ch string = string(char)
		if ch != " " && ch != "\t" {
			wording = wording + ch
		} else {
			if wording != "" {
				wordings = append(wordings, wording)
			}
			wording = ""
		}
	}
	return wordings
}

func toToken(text string) (Token, error) {
	if isSymbol(text) {
		return Token{TOEKN_SYMBOL, TokenValue(text)}, nil
	}

	if isNumber(text) {
		return Token{TOKEN_NUMBER, TokenValue(text)}, nil
	}

	if isString(text) {
		return Token{TOKEN_STRING, TokenValue(text)}, nil
	}

	if isBoolean(text) {
		return Token{TOKEN_BOOLEAN, TokenValue(text)}, nil
	}

	isOperator, operator := isOperator(text)
	if isOperator {
		return Token{TOKEN_OPERATOR, TokenValue(operator)}, nil
	}

	isKeyword, keyword := isKeyword(text)
	if isKeyword {
		return Token{TOKEN_KEYWORD, TokenValue(keyword)}, nil
	}

	return Token{TOKEN_UNKNOW, TokenValue(text)}, errors.New("unexpected token")
}

func Lexer(filePath string) ([]Token, error) {
	var tokens []Token

	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("can not read file")
	}

	defer file.Close()

	var scanner *bufio.Scanner = bufio.NewScanner(file)

	for scanner.Scan() {
		for _, wording := range getWordingsOfTextLine(scanner.Text()) {
			token, err := toToken(wording)
			if err != nil {
				log.Fatal(err)
			}
			tokens = append(tokens, token)
		}
	}

	return tokens, nil
}
