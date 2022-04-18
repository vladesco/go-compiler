package repl

import (
	"bufio"
	"compiler/lexer"
	"compiler/parser"
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

		parserInstance := parser.New(lexer.New(scanner.Text()))

		fmt.Print(parserInstance.ParseProgram().ToString(), "\n")
		fmt.Print("errors:", parserInstance.GetParsingErrors(), "\n")

	}

}
