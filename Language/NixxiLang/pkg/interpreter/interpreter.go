package interpreter

import (
	"NixxiLang/pkg/util"
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
		fmt.Println("Interactive runtime comming soon")
	}
}
