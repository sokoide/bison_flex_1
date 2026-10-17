package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	var source string
	for s.Scan() {
		source += s.Text() + "\n"
	}
	lexer := &lexer{src: source, index: 0}
	if yyParse(lexer) != 0 {
		fmt.Fprintln(os.Stderr, "parse failed")
		os.Exit(1)
	}
}
