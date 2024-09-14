package prolog

func YyParse(lexer *Lexer) int {
	yyErrorVerbose = true
	return yyNewParser().Parse(lexer)
}
