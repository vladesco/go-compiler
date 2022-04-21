package repl

import (
	"bufio"
	"compiler/evaluator"
	"compiler/lexer"
	"compiler/parser"
	"fmt"
	"io"
	"strings"
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
		program := parserInstance.ParseProgram()
		errors := parserInstance.GetParsingErrors()

		if len(errors) > 0 {
			fmt.Print(strings.Join(errors, "\n"), "\n")
			continue
		}

		fmt.Print(evaluator.Eval(program), "\n")
	}
}
