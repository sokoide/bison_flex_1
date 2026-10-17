# Goyacc2c

> 英語版: [README.md](README.md)

* 第4歩: パース中にASTを構築し、`yyParse` が返った後に評価する。そして `if` を追加

[goyacc2b](../goyacc2b/README_ja.md) の上に構築しています。同じファイル構成に
`ast.go` を追加し、評価を文法アクションの外へ移します。

## goyacc2bからの追加点

* `ast.go` — 小さな汎用の `Node{Kind, S, I, N1, N2, Ns}` 木と
  `evaluate()`/`evaluateNode()`。文字列のKindがそのままノード型になるため、
  ルールごとに個別のノード構造体は不要です
  （`newNode("+", "", 0, $1, $3, nil)`）。
* 文法アクションはもう何も計算しません: ノードを構築するだけです
  （`%type<val>` の代わりに `%type<node>`/`%type<nodes>`）。そして `if`
  ステートメントがルールになります:
  `stmt: IF '(' expr ')' '{' stmts '}'`。これがgoyacc2でlexerが `'{'`
  `'}'` トークンを布石として切り出しておいた理由です。
* 評価は後回し（deferred）です: `main()` は `yyParse(lexer)` が返った
  *後* に `evaluate(lexer)` を実行します。

## なぜ評価を後回しにするのか？ `1/0` の教えどころ

goyacc1/2/2bではアクションはパースの最中に実行されるため、ゼロ除算は
パースの途中、`expr '/' expr` アクションの内側 — parserがまだスタック上に
ある間に — panicします。`yyParse` からエラーの戻り値を得ることはできません
（panicが bypassするため）。

goyacc2cではパースは木を構築するだけなので、`yyParse` は正常に返り、panic
（もしあれば）はその後、評価器の中で起こります:

```
$ printf 'put(1/0);\n' | ./hoge
panic: runtime error: integer divide by zero
    ... ast.go:53 (evaluateNode), called from evaluateNodes ...
```

同じ入力で、2つの異なるクラッシュ地点 — この違いこそが学ぶべき点です:
パースと実行がきれいに分離されました。これは後のインタプリタ最適化、
バイトコード、バックエンドのためにまさに必要となる構造です。

## Build / Run / Test

```
make          # goyacc hoge.go.y -> y.go, then build ./hoge
make test     # exact assertions, exits non-zero on mismatch
make run      # read statements from stdin
make clean    # removes ./hoge y.go y.output
```

## 構築した木を呼び出し元へ渡す

program非終端記号の結果を `main()` まで届かせる必要があります。トリックは
これです: `program: stmts { $$ = $1; yylex.(*lexer).program = $$ }` が
`yylex` インターフェイス引数を具象型の `*lexer` ポインタにキャストし、
main() がその `program []*Node` フィールドを読みます。（このキャストが
機能するのは、goyaccが呼び出し元の型に `yylex` インターフェイスを実装する
ことを要求しているためです。）

これまでどおり: 代入/putはノードを格納し、評価時に出力します。未定義の変数は
0 として読めます。単項マイナスなし。識別子は小文字の `[a-z]+` のみです。

## ファイル

* hoge.go.y ... 文法（状態機械のみ。アクションに副作用なし）
* y.go/y.output ... `goyacc` が生成
* lexer.go ... 手書きのlexer（`PUT`、`IF` キーワードの振り分け）
* ast.go ... Node型とパース後の評価器
* main.go ... stdinを読み、パースし、その後評価します
