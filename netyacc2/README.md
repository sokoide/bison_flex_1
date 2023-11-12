# netyacc2

## About

* Simple compiler for a stack machine and the stack machine

## Prereqs

* <https://github.com/ernstc/YaccLexTools>
  * `dotnet tool install dotnet-ylt --global`

## Useful Commands

* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet ylt add-parser -p <parserName>`

## How to run

* Test

```sh
cd netyacc2
# to run unit tests in interp-test
dotnet test -v
```

* Build

```sh
dotnet build
```

* to run a demo

``` sh
dotnet run --project interp-exe/interp-exe.csproj demo
```

* or

```sh
./interp-exe/bin/Debug/net7.0/interp-exe demo
```

* to run sample scripts
  * basic

  ```sh
  $ dotnet run  --project interp-exe/interp-exe.csproj --file samples/sample1.txt --verbose
  * Source
  a=42;
  put(a);

  * Original. Jump/JumpF's operands mean Label name
  [0000] PushN 42
  [0001] Pop 0
  [0002] PutI 0
  [0003] PutS 1
  * Label Resolved. Jump/JumpF's operands mean PC
  [0000] PushN 42
  [0001] Pop 0
  [0002] PutI 0
  [0003] PutS 1
  * String table
  [0001] \n

  42
  ```

  * fibonacci

  ```sh
  $ dotnet run --project interp-exe/interp-exe.csproj  --file samples/sample_fib.txt --verbose
  * Source
  put("fib(0)=0");
  a=0;
  put("fib(1)=1");
  b=1;
  i=2;
  while(i<40){
      f=a+b;
      put("fib(",i,")=",f);
      a=b;
      b=f;
      i=i+1;
  }
  * Original. Jump/JumpF's operands mean Label name
  [0000] PutS 1
  [0001] PutS 2
  [0002] PushN 0
  [0003] Pop 0
  [0004] PutS 3
  [0005] PutS 2
  [0006] PushN 1
  [0007] Pop 1
  [0008] PushN 2
  [0009] Pop 8
  [0010] Label 1001
  [0011] PushI 8
  [0012] PushN 40
  [0013] Calc LTOP
  [0014] JumpF 1002
  [0015] PushI 0
  [0016] PushI 1
  [0017] Calc ADD
  [0018] Pop 5
  [0019] PutS 4
  [0020] PutI 8
  [0021] PutS 5
  [0022] PutI 5
  [0023] PutS 2
  [0024] PushI 1
  [0025] Pop 0
  [0026] PushI 5
  [0027] Pop 1
  [0028] PushI 8
  [0029] PushN 1
  [0030] Calc ADD
  [0031] Pop 8
  [0032] Jump 1001
  [0033] Label 1002
  * Label Resolved. Jump/JumpF's operands mean PC
  [0000] PutS 1
  [0001] PutS 2
  [0002] PushN 0
  [0003] Pop 0
  [0004] PutS 3
  [0005] PutS 2
  [0006] PushN 1
  [0007] Pop 1
  [0008] PushN 2
  [0009] Pop 8
  [0010] Label 1001
  [0011] PushI 8
  [0012] PushN 40
  [0013] Calc LTOP
  [0014] JumpF 33
  [0015] PushI 0
  [0016] PushI 1
  [0017] Calc ADD
  [0018] Pop 5
  [0019] PutS 4
  [0020] PutI 8
  [0021] PutS 5
  [0022] PutI 5
  [0023] PutS 2
  [0024] PushI 1
  [0025] Pop 0
  [0026] PushI 5
  [0027] Pop 1
  [0028] PushI 8
  [0029] PushN 1
  [0030] Calc ADD
  [0031] Pop 8
  [0032] Jump 10
  [0033] Label 1002
  * String table
  [0001] fib(0)=0
  [0002] \n
  [0003] fib(1)=1
  [0004] fib(
  [0005] )=

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
  ``
