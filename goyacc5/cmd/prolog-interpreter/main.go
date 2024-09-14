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

func main() {
	parseFlags()
	log.SetLevel(o.ll)
	log.SetFormatter(&log.TextFormatter{})

	if flag.NArg() < 1 {
		fmt.Println("Usage: prolog <input>")
		os.Exit(1)
	}

	lexer, err := prolog.NewLexer(flag.Arg(0))
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer lexer.Close()

	log.Printf("lexer: %+v", lexer)
	prolog.YyParse(lexer)
}
