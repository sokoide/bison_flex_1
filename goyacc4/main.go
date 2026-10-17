package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
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
	log.SetFormatter(&log.TextFormatter{})
	parseFlags()
	log.SetLevel(o.ll)

	s := bufio.NewScanner(os.Stdin)
	scanner := new(scanner)
	var source []string
	for s.Scan() {
		source = append(source, s.Text())
	}

	scanner.Init(strings.Join(source, "\n"))

	prog, err := parse(scanner)
	if err != nil {
		log.Errorf("Parse error: %v", err)
		os.Exit(1)
	}

	_, err = evaluateStmts(prog)
	if err != nil {
		log.Errorf("Evaluation error: %v", err)
		os.Exit(1)
	}
}

func parse(s *scanner) ([]expression, error) {
	l := lexer{s: s}
	yyErrorVerbose = true
	if yyParse(&l) != 0 {
		err := l.err
		if err == nil {
			err = fmt.Errorf("parse error")
		}
		return nil, err
	}
	return l.program, nil
}
