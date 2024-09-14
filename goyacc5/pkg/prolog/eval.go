package prolog

import log "github.com/sirupsen/logrus"

func YyParse(lexer *Lexer) int {
	yyErrorVerbose = true
	return yyNewParser().Parse(lexer)
}

func Evaluate(lexer *Lexer) error {
	for _, c := range lexer.program {
		log.Infof("%s", c)
	}
	return nil
}
