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
	dump     bool
}

var o options

func parseFlags() {
	var err error
	var ll log.Level
	flag.StringVar(&o.logLevel, "logLevel", "INFO", "TRACE | DEBUG | INFO | WARN | ERROR")
	flag.BoolVar(&o.dump, "dump", false, "dump the parser")
	flag.Parse()

	ll, err = log.ParseLevel(o.logLevel)
	if err == nil {
		o.ll = ll
	} else {
		o.ll = log.InfoLevel
		log.Warnf("logLevel: %s is not supported. falling back to INFO", o.logLevel)
	}
}

func main() {
	parseFlags()
	log.SetLevel(o.ll)
	log.SetFormatter(&log.TextFormatter{})

	if flag.NArg() < 1 {
		fmt.Println("Usage: prolog <source-file> <query-file>")
		os.Exit(1)
	}

	lexer, err := prolog.NewLexer(flag.Arg(0))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer lexer.Close()

	log.Debugf("lexer: %+v", lexer)
	program, err := prolog.Load(lexer)
	if err != nil {
		log.Fatalf("Error loading program: %v", err)
	}
	if o.dump {
		prolog.Dump(program)
	}

	if flag.NArg() >= 2 {
		queryLexer, err := prolog.NewLexer(flag.Arg(1))
		if err != nil {
			log.Fatalf("Error opening file: %v", err)
		}
		defer queryLexer.Close()
		queryProgram, err := prolog.Load(queryLexer)
		if err != nil {
			log.Fatalf("Error loading query program: %v", err)
		}

		if o.dump {
			prolog.Dump(queryProgram)
		}

		prolog.Query(program, queryProgram)
	}
}
