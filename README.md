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
  * netyacc2: simple interpreter

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

* C/Go

```bash
make test
make run
```

* .NET

```bash
cd netyacc2
# to run unit tests in interp-test
dotnet test -v
# to build interp-exe
dotnet build
# to run a demo
dotnet run --project interp-exe/interp-exe.csproj demo
# or,
./interp-exe/bin/Debug/net7.0/interp-exe demo
# to run your app
cat samples/sample1.txt | dotnet run --project interp-exe/interp-exe.csproj
```
