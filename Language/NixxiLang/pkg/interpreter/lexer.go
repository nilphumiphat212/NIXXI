package interpreter

import (
	"bufio"
	"errors"
	"log"
	"os"
)

const _NUMBERDIGIT = "0123456789"

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

func toToken(textLine string) (Token, error) {
	// mock := "a = 1"
	return Token{Type: TOKEN_KEYWORD, Value: "null"}, nil
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
		token, err := toToken(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		tokens = append(tokens, token)
	}

	return tokens, nil
}
