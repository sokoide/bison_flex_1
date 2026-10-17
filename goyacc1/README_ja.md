# Goyacc1

> 英語版: [README.md](README.md)

* 最もシンプルなGoyaccのサンプル: 整数式の評価器

## Build / Run / Test

```
make          # runs goyacc on hoge.go.y -> y.go, then builds ./hoge
make test     # exact assertions, exits non-zero on mismatch
make run      # read an expression from stdin
make clean    # removes ./hoge y.go y.output (make regenerates them)
```

実行例:

```
$ printf '4+1*4/2\n' | ./hoge
Result: 6
$ printf '1+\n' | ./hoge
error: syntax error, at token: +
parse failed     (exit 1)
```

## ファイル

* `hoge.go.y` ... 文法ファイル（goyaccへの入力）
* `y.go` ... `hoge.go.y` から `goyacc` が生成したもの（gitignore対象。
  `make` が再生成します）
* `y.output` ... 生成されたparserの状態表/コンフリクト報告
  （gitignore対象）
* `lexer.go` ... 手書きのlexer
  * `Lex()` はparserが次のトークンを取得するために呼び出します。
    0を返す（または任意の負数、ここでは-1）と、goyaccにとってEOFの合図になります。
  * `Error()` はparserによる構文（SYNTAX）エラーの報告のためにgoyaccから
    呼び出されます（たとえば `1+`）。字句上の問題に対しては呼ばれないことに
    注意してください。不正な文字が現れたときに何をするかはlexerの作者が
    決めるもので、このlexerは単純にpanicします。
  * `yyErrorVerbose = true`（生成コードのグローバル変数）を設定すると、
    parserがエラー文字列に予期しないトークン/期待されたトークンの名前を
    含めるようになります。`yyDebug = 1..4` はparserの各ステップをstderrに
    トレース出力します。どちらも実行時に設定できます。
    このサンプルでは `yyErrorVerbose` をデフォルト（`false`）のままにして
    あります。そのため上のメッセージは単なる `syntax error` の形になり
    ます。[goyacc2](../goyacc2/README_ja.md) はこれを有効にしているので、
    両者を比較できます。
* `main.go` ... stdinを読み、lexerに渡し、`yyParse` を実行します

## 文法を一言で言うと

`program: expr` — 1回の実行につき1つの式。

## `%token` と `%union`

* `%union{ val int }` はyaccの `YYSTYPE` を置き換えるGoの型を定義します。
* `%token<val> NUMBER` はトークンを宣言し、その意味値（semantic value）が
  `lval.val`（`%union` のメンバー `val`）に入ることを示します。
* `%left '+' '-'` と `%left '*' '/'` により `'*'` `/` のグループに高い
  precedenceが与えられ、`4+1*4/2` は `4+((1*4)/2)` としてパースされます。

## `hoge.go.y` を読み歩く

1. パッケージとimport:
   `.y` ファイル先頭の `%{ ... %}` ブロックは生成される `y.go` にそのまま
   貼り込まれるため、生成されたparserは必要なimport（以下のアクションで使う
   `fmt`）とともに `package main` に置かれます。

2. union宣言:
   `%union` は文法シンボルに結び付けられる型を定義します。
   ここでは `int` 型の `val` だけがあります。

3. トークンと型の宣言:
   `%token <val> NUMBER` は数値のためのトークンを宣言し、
   `%type <val> expr` は非終端記号 `expr` が整数値を持つことを示します。

4. 演算子のprecedence:
   `%left '+' '-'` と `%left '*' '/'` — 後に宣言したものほど強く結合する
   （bind tighter）ため、`4+1*4/2` が6として計算されます。

5. 文法ルール:
   ルールはC版（`2020_2` を参照）とほぼ同じですが、アクションはGoです:
   Cのyaccと同様に `$$` への代入と `$1`、`$3`、... の参照を行います。
   アクションはパースの最中に実行されます（ASTを使う代替案については
   [goyacc2c](../goyacc2c/README_ja.md) のREADMEを参照）。

6. lexer実装:
   `lexer` 構造体とそのメソッドがトークン化を担当します — `Lex()` が入力を
   読んでトークンを返し、`Error()` が構文エラーを報告します。

7. エントリポイント（`main.go`）: stdinをすべてlexerに読み込ませ、
   `yyParse` を実行し、パースに失敗したら `parse failed` で終了コード1で
   終わります。`main` は値を一切見ないことに注意してください —
   `program: expr` のアクションが、パースの最中に自前で `Result: %d` を
   出力します。[goyacc2c](../goyacc2c/README_ja.md) と比較してください。
   あちらは構築した木を `main` に返し、`yyParse` が返った後に評価します。
