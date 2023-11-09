# bison_flex_1

## About

* very basic bison and flex examples
  * 2020_1: simple lex example (C)
  * 2020_1.5: simple lex + yacc example (C)
  * 2020_2: simple lex + yacc calculator (C)
  * goyacc1: simple goyacc example (Go)
  * goyacc2: simple goyacc example2 (Go)
  * goyacc3: another goyacc example with parser and expression (Go)
  * netyacc1: .net version of goyacc1 (C#)_

## Prereqs

* C
  * Binson & Flex
* Go
  * go get -u golang.org/x/tools/cmd/goyacc
* .NET <https://github.com/ernstc/YaccLexTools>
  * dotnet tool install dotnet-ylt --global

## Useful Commands

* .NET <https://github.com/ernstc/YaccLexTools>
  * dotnet ylt add-parser -p <parserName>

## How to run

```bash
make test
make run
```
