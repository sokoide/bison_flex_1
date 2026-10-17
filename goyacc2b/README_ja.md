# Goyacc2b

> 英語版: [README.md](README.md)

* 第3歩: 変数、代入、`put(...)` ステートメント、そして文法に `%`

[goyacc2](../goyacc2/README_ja.md) の上に構築しています。同じ三ファイル構成
（main.go、lexer.go、y.go）です。

## goyacc2からの追加点

* `stmt: IDENT '=' expr ';'` — 代入ステートメント。アクションはエピローグ
  （最後の `%%` の後のセクション）で宣言されたグローバルな `vars` マップに
  格納します: `var vars = map[string]int{}`。
* `stmt: PUT '(' expr ')' ';'` — 出力ステートメント。`put` は
  `getString()` 内で特別扱いされます: lexerがその語を検索し、`IDENT` の
  代わりに `PUT` トークンを返します（手書きのキーワード検索）。
* `expr: IDENT { $$ = vars[$1] }` — 変数の読み出し。既知の癖として:
  未定義の変数は診断なしに 0 として読めます。
* `expr '%' expr` がついに文法に実装されました（goyacc2でも `%` を切り出して
  いました。ここでは代入と組み合わせられます: `a=20%7`）。
* 布石として置いたトークン `'{'  '}'` はここでも未使用です — goyacc2cの
  `if` が消費します。

## Build / Run / Test

```
make          # goyacc hoge.go.y -> y.go, then build ./hoge
make test     # exact assertions, exits non-zero on mismatch
make run      # read statements from stdin
make clean    # removes ./hoge y.go y.output
```

## 教えどころ: 手書きlexerでのキーワード処理

lexerは自分で書くコードなので、「キーワード」に追加の仕組みは要りません:
`getString()` が語 `put` を作ったら `PUT` を返し、そうでなければ `IDENT`
を返します。実際の言語処理系ではキーワードは通常マップで管理します、と
比べてみてください。

意味アクションはgoyacc1/2とまったく同じくパースの最中に実行されるため、
ゼロ除算は **`yyParse` 自身の内側でpanicします**。文法アクションの中で:

```
$ printf 'put(1/0);\n' | ./hoge
panic: runtime error: integer divide by zero
    ... hoge.go.y:46 (the expr '/' expr action) inside (*yyParserImpl).Parse ...
```

（ここでは *裸の* `1/0;` はすでに構文エラーになる点に注意: すべての
ステートメントは代入か `put(...)` でなければなりません。）goyacc2cは
ASTによる代替案を示します。

## ファイル

* hoge.go.y ... 文法。`%%` の後のエピローグにある `var vars` マップに注意
* y.go/y.output ... `goyacc` が生成
* lexer.go ... 手書きのlexer（`PUT` と `IDENT` の振り分け）
* main.go ... stdinを読み、lexerに渡し、`yyParse` を実行します
