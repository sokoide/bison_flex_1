# netyacc1
> 英語版: [README.md](README.md)

> **Deprecated（非推奨）**: この例は**メンテナンスされていません**。
> 参照用として残しています。

## 概要

* [goyacc1](../goyacc1/README_ja.md) の.NET (C#) 版: YaccLexToolsが生成したparserとscannerで動く、小さな算術式 evaluator。
* `goyacc1/hoge.go.y` と同一構成です: `%left '+' '-'` / `%left '*' '/'` のprecedence、`NUMBER` トークン、`( expr )`、結果を出力する `program : expr` ルール。
* 他の例と違い、これはREPLです: 1行に1つの式を読み込みます。

## ファイル

* `Calc/Calc.Language.grammar.y` ... grammar（YaccLexToolsへの入力）
* `Calc/Calc.Language.analyzer.lex` ... lexer仕様
* `Calc/Calc.Parser.Generated.cs` / `Calc/Calc.Scanner.Generated.cs`
  ... ビルド時にYaccLexToolsが生成（リポジトリにコミット済み）
* `Calc/Calc.Parser.cs` / `Calc/Calc.Scanner.cs` ... 手書きのpartial class
* `Program.cs` ... REPLのエントリポイント

## 前提条件

* .NET SDK
* <https://github.com/ernstc/YaccLexTools>
  * `dotnet tool install dotnet-ylt --global`
  * `.y` / `.lex` ファイルからparser/scannerを再生成するときにのみ必要です
    （生成済みの `.cs` ファイルがコミットされているため、普通のビルドには不要です）

## 実行方法

```sh
cd netyacc1
dotnet run
```

1行に1つの式を入力します（`Ctrl+D` または空行で終了）:

```text
> 1+2*3
7

> (4+1)*4/2
10
```

## メモ

* このプロジェクトはサポートが終了した **.NET 7** をターゲットにしているため、ビルド時に警告 `NETSDK1138` が出力されます。それでも例は動作します。マシンに新しいランタイムしか入っていない場合は、`DOTNET_ROLL_FORWARD=Major dotnet run` で実行してください。
* ここにはunit testプロジェクトがありません。テスト付きのGo版は `goyacc1` を参照してください。
