package main

import (
	"bufio"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var source string
	for s.Scan() {
		source += s.Text()
	}
	lexer := &lexer{src: source, index: 0}
	yyParse(lexer)
	evaluate(lexer)
}
