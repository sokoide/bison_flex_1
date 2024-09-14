# bison_flex_1

## About

* very basic bison and flex examples
  * 2020_1: simple lex example (C)
  * 2020_1.5: simple lex + yacc example (C)
  * 2020_2: simple lex + yacc calculator (C)
  * goyacc1: simple goyacc example (Go)
  * goyacc2: simple goyacc example2 (Go)
  * goyacc3: another goyacc example with parser and expression (Go)
  * goyacc4: tiny language interpreter (Go)
  * TBD: goyacc5: Prolog like language interpreter (Go)
  * netyacc1: .net version of goyacc1 (C#)
  * netyacc2: simple compiler for a stack machine (C#)

## Prereqs

* C
  * Binson & Flex
* Go
  * go install golang.org/x/tools/cmd/goyacc@latest
* .NET <https://github.com/ernstc/YaccLexTools>
  * dotnet tool install dotnet-ylt --global

## Useful Commands

* .NET <https://github.com/ernstc/YaccLexTools>
  * dotnet ylt add-parser -p <parserName>

## How to run

* C/Go

```bash
make test
make run
```

* .NET
  * Please refer to [README.md in netyacc2](netyacc2/README.md)
