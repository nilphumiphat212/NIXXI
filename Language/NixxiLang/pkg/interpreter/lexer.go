package interpreter

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
)

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
	return string(runes[0:1]) == "\"" && string(runes[len(runes)-1:1]) == "\""
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
	if isNumber(text) {
		return Token{TOKEN_NUMBER, TokenValue(text)}, nil
	}

	if isString(text) {
		return Token{TOKEN_STRING, TokenValue(text)}, nil
	}

	isKeyword, keyword := isKeyword(text)
	if isKeyword {
		return Token{TOKEN_KEYWORD, TokenValue(keyword)}, nil
	}

	return Token{TOKEN_UNKNOW, TokenValue(text)}, nil
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
