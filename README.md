# bison_flex_1

## About

* very basic bison and flex examples
  * 2020_1: simple lex example (C)
  * 2020_1.5: simple lex + yacc example (C)
  * 2020_2: simple lex + yacc calculator (C)
  * goyacc1: simple goyacc example (Go)
  * goyacc2: simple goyacc example2 (Go)
  * goyacc3: another goyacc example with parser and expression (Go)
  * netyacc1: .net version of goyacc1 (C#)
  * netyacc2: simple interpreter (C#)

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
dotnet run --project ./interp-exe/interp-exe.csproj  demo

* Source
put("*** Demo ***");
put("counting down...");
e = 3;
while (e > 0)
{
    put("e=", e);
    e = e - 1;
    }

* Original. Jump/JumpF's operands mean Label name
[0000] PutS 1
[0001] PutS 2
[0002] PutS 3
[0003] PutS 2
[0004] PushN 3
[0005] Pop 4
[0006] Label 1001
[0007] PushI 4
[0008] PushN 0
[0009] Calc GTOP
[0010] JumpF 1002
[0011] PutS 4
[0012] PutI 4
[0013] PutS 2
[0014] PushI 4
[0015] PushN 1
[0016] Calc SUB
[0017] Pop 4
[0018] Jump 1001
[0019] Label 1002
* Label Resolved. Jump/JumpF's operands mean PC
[0000] PutS 1
[0001] PutS 2
[0002] PutS 3
[0003] PutS 2
[0004] PushN 3
[0005] Pop 4
[0006] Label 1001
[0007] PushI 4
[0008] PushN 0
[0009] Calc GTOP
[0010] JumpF 19
[0011] PutS 4
[0012] PutI 4
[0013] PutS 2
[0014] PushI 4
[0015] PushN 1
[0016] Calc SUB
[0017] Pop 4
[0018] Jump 6
[0019] Label 1002
* String table
[0001] *** Demo ***
[0002] \n
[0003] counting down...
[0004] e=
* Executing...
*** Demo ***
counting down...
e=3
e=2
e=1

cat samples/sample_fib.txt | dotnet run --project interp-exe/interp-exe.csproj

fib(0)=0
fib(1)=1
fib(2)=1
fib(3)=2
fib(4)=3
fib(5)=5
fib(6)=8
fib(7)=13
fib(8)=21
fib(9)=34
fib(10)=55
fib(11)=89
fib(12)=144
fib(13)=233
fib(14)=377
fib(15)=610
fib(16)=987
fib(17)=1597
fib(18)=2584
fib(19)=4181
fib(20)=6765
fib(21)=10946
fib(22)=17711
fib(23)=28657
fib(24)=46368
fib(25)=75025
fib(26)=121393
fib(27)=196418
fib(28)=317811
fib(29)=514229
fib(30)=832040
fib(31)=1346269
fib(32)=2178309
fib(33)=3524578
fib(34)=5702887
fib(35)=9227465
fib(36)=14930352
fib(37)=24157817
fib(38)=39088169
fib(39)=63245986
```
