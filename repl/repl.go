package repl

import (
	"bufio"
	"compiler/lexer"
	"compiler/token"
	"fmt"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner((in))

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()

		if !scanned {
			return
		}

		lexerInstance := lexer.New(scanner.Text())

		for nextToken := lexerInstance.ReadNextToken(); nextToken.Type != token.EOF; nextToken = lexerInstance.ReadNextToken() {
			fmt.Printf("%+v\n", nextToken)
		}
	}

}
