package main

import (
	"NixxiLang/pkg/interpreter"
	"errors"
	"fmt"
	"log"
	"os"
)

func main() {
	realMain()
}

func realMain() {
	var args []string = os.Args
	var filePath string

	if len(args) > 1 {
		filePath = args[1]
	} else {
		log.Fatal(errors.New("file path is empty"))
	}

	tokens, err := interpreter.Lexer(filePath)

	if err != nil {
		log.Fatal(err)
		return
	}

	fmt.Println(len(tokens))
}
