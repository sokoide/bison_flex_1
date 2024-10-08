package main

import (
	"bufio"
	"flag"
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
	source := []string{}
	for s.Scan() {
		source = append(source, s.Text())
	}

	scanner.Init(strings.Join(source, "\n"))

	var prog []expression = parse(scanner)
	_, err := evaluateStmts(prog)
	if err != nil {
		panic(err)
	}
}

func parse(s *scanner) []expression {
	l := lexer{s: s}
	yyErrorVerbose = true
	if yyParse(&l) != 0 {
		panic("Parse error")
	}
	return l.program
}
