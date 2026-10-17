# netyacc2
> 英語版: [README.md](README.md)

> **Deprecated（非推奨）**: この例は**メンテナンスされていません**。
> 参照用として残しています。

## 概要

* 小さな言語（代入、`if`/`else`、`while`、`return`、型付き変数）を**stack machine用のbytecodeにコンパイル**し、そのstack machine（VM）自体も実装したものです。コンパイラはYaccLexTools（yacc/flexの.NET版であるgppg/gplex）で構築しています。
* コンパイルのpipeline: ソース → 中間コード（`PushN`, `Pop`, `PutI`, `PutS`, `Jump`, ...） → ラベル解決 → VMによる実行。`--verbose` で各段階を表示できます。

## ファイル

* `interp-lib/Interp/Interp.Language.grammar.y` /
  `Interp.Language.analyzer.lex` ... grammarとlexer（YaccLexToolsへの入力）。
  `Interp.Parser.Generated.cs` / `Interp.Scanner.Generated.cs` が生成結果です（コミット済み）
* `interp-lib/Interp/VM.cs`（`Cpu.cs`）... stack machine本体
* `interp-exe/` ... コマンドラインフロントエンド（`--file <script>`、`--verbose`、`demo`）
* `interp-test/` ... unit test（`dotnet test`、34テスト）
* `samples/` ... サンプルプログラム（`sample1.txt` ... `sample5.txt`、`sample_fib.txt`）

## 前提条件

* <https://github.com/ernstc/YaccLexTools>
  * `dotnet tool install dotnet-ylt --global`

## 便利なコマンド

* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet ylt add-parser -p <parserName>`

## 実行方法

これらのプロジェクトは **.NET 8** をターゲットにしています。マシンに新しいランタイムしか入っていない場合は、`DOTNET_ROLL_FORWARD=Major` を設定してください（下記のテストの注意を参照）。

* テスト

```sh
cd netyacc2
dotnet test
```

  `You must install or update .NET to run this application` で失敗する場合
  （新しいランタイムしか入っていないとき）は、次のように実行します:

  ```sh
  DOTNET_ROLL_FORWARD=Major dotnet test
  ```

* ビルド

```sh
dotnet build
```

* demoの実行

プログラムの引数は `--` の後ろに渡します:

```sh
dotnet run --project interp-exe/interp-exe.csproj -- demo
```

* または `dotnet build` の後なら（ターゲットフレームワークは `net8.0`）:

```sh
./interp-exe/bin/Debug/net8.0/interp-exe demo
```

* サンプルスクリプトの実行

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

* 出力（`--verbose` を付けると、実行前に生成コード・文字列テーブル・変数テーブルがダンプされます）:

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

* その他のサンプル（`--verbose` は省略可能）:
  * `sample2.txt` — `while` ループ
  * `sample3.txt` — `if` / `else`
  * `sample4.txt` — `return`
  * `sample5.txt` — 型付き（`int` / `string`）変数
  * `sample_fib.txt` — `fib(39)` までのフィボナッチ数

  どれでも実行できます。例:

  ```sh
  dotnet run --project interp-exe/interp-exe.csproj -- --file samples/sample_fib.txt --verbose
  ```
