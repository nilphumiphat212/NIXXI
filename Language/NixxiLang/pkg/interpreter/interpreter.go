package interpreter

import (
	"NixxiLang/pkg/util"
	"bufio"
	"fmt"
	"os"
)

func NewInterpreter() error {
	var args []string = os.Args
	if len(args) > 1 {
		tokens, err := LoadToLexerFromFile(args[1])
		if err != nil {
			return err
		}
		util.WriteJsonFile("language_token.json", tokens)
	} else {
		err := interactiveInterpreter()
		if err != nil {
			return err
		}
	}
	return nil
}

func interactiveInterpreter() error {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	for {
		var text string
		fmt.Print("Nixxi : ")
		scanner.Scan()
		text = scanner.Text()
		if len(text) > 0 {
			if text == "exit" {
				break
			}
			token, err := convertTextLineToTokan(text)
			if err != nil {
				return err
			}
			fmt.Println(token)
		}
	}
	return nil
}
