package interpreter

import (
	"NixxiLang/pkg/util"
	"bufio"
	"fmt"
	"log"
	"os"
)

func NewInterpreter() {
	var args []string = os.Args
	if len(args) > 1 {
		tokens, err := LoadToLexerFromFile(args[1])
		if err != nil {
			log.Fatalln(err)
		}
		util.WriteJsonFile("language_token.json", tokens)
	} else {
		interactiveInterpreter()
	}
}

func interactiveInterpreter() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Nixxi : ")
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(text)
		token, err := convertTextLineToTokan(text)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(token)
	}
}
