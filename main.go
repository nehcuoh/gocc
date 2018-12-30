package main

import (
	"flag"
	"fmt"
)

func main() {

	sourceFile := flag.String("s", "", "source file name")
	flag.Parse()

	input.ReadSourceFile(*sourceFile)

	setupScanner()

	token := getNextToken()
	for ; token != TK_END; {
		fmt.Printf("(%v,%v) %s\n", tokenCoord.line, tokenCoord.col, token)
		token = getNextToken()
		//fmt.Printf("%s ", currentToken)
	}

	input.CloseSourceFile()
}
