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

These projects target **.NET 8**. If your machine has only a newer runtime
installed, set `DOTNET_ROLL_FORWARD=Major` (see the test note below).

* Test

```sh
cd netyacc2
dotnet test
```

  If this fails with `You must install or update .NET to run this application`
  (only a newer runtime is installed), run:

  ```sh
  DOTNET_ROLL_FORWARD=Major dotnet test
  ```

* Build

```sh
dotnet build
```

* to run a demo

Program arguments are passed to the app after `--`:

```sh
dotnet run --project interp-exe/interp-exe.csproj -- demo
```

* or, after `dotnet build` (the target framework is `net8.0`):

```sh
./interp-exe/bin/Debug/net8.0/interp-exe demo
```

* to run sample scripts

```sh
dotnet run --project interp-exe/interp-exe.csproj -- --file samples/sample1.txt --verbose
```

* `samples/sample1.txt`:

  ```text
  a=1;
  b=2;
  c=42;
  put(c);
  put(b);
  put(a);
  ```

* output (`--verbose` dumps the generated code, the string table and the
  variable table before running):

  ```text
  * Source
  a=1;
  b=2;
  c=42;
  put(c);
  put(b);
  put(a);

  * Original. Jump/JumpF's operands mean Label name
  [0000] PushN 1
  [0001] Pop 1
  [0002] PushN 2
  [0003] Pop 2
  [0004] PushN 42
  [0005] Pop 3
  [0006] PutI 3
  [0007] PutS 1
  [0008] PutI 2
  [0009] PutS 1
  [0010] PutI 1
  [0011] PutS 1
  * Label Resolved. Jump/JumpF's operands mean PC
  [0000] PushN 1
  [0001] Pop 1
  [0002] PushN 2
  [0003] Pop 2
  [0004] PushN 42
  [0005] Pop 3
  [0006] PutI 3
  [0007] PutS 1
  [0008] PutI 2
  [0009] PutS 1
  [0010] PutI 1
  [0011] PutS 1
  * String table
  [0001] \n
  * Variable table
  Name: a, Type: INT, Index: 1
  Name: b, Type: INT, Index: 2
  Name: c, Type: INT, Index: 3

  42
  2
  1
  ```

* other samples (`--verbose` is optional):
  * `sample2.txt` — `while` loop
  * `sample3.txt` — `if` / `else`
  * `sample4.txt` — `return`
  * `sample5.txt` — typed (`int` / `string`) variables
  * `sample_fib.txt` — fibonacci numbers up to `fib(39)`

  Run any of them, for example:

  ```sh
  dotnet run --project interp-exe/interp-exe.csproj -- --file samples/sample_fib.txt --verbose
  ```
