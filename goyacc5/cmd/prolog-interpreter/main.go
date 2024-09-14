package main

import (
	"flag"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"sokoide.com/bison_flex_1/goyacc5/pkg/prolog"
)

type options struct {
	logLevel string
	ll       log.Level
}

var o options

func parseFlags() {
	var err error
	var ll log.Level
	flag.StringVar(&o.logLevel, "logLevel", "debug", "TRACE | DEBUG | INFO | WARN | ERROR")
	flag.Parse()

	ll, err = log.ParseLevel(o.logLevel)
	if err == nil {
		o.ll = ll
	} else {
		o.ll = log.InfoLevel
		log.Warnf("logLevel: %s is not supported. falling back to INFO", o.logLevel)
	}
}

// func yyError(msg string) {
// 	fmt.Fprintf(os.Stderr, "Error: %s\n", msg)
// }

func main() {
	parseFlags()

	if flag.NArg() < 1 {
		fmt.Println("Usage: prolog <input>")
		os.Exit(1)
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	//yyErrorVerbose = true
	lexer := &prolog.Lexer{Reader: file}
	log.Println("lexer: %+v", lexer)
	prolog.YyParse(lexer)
}
