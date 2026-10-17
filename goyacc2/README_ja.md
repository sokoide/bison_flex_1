# Goyacc2

> 英語版: [README.md](README.md)

* 第2歩: 1回の実行で複数ステートメント（`stmts`/`stmt`）、空白の扱い、そして文法が使わないトークンまでlexerが切り出す例

[goyacc1](../goyacc1/README_ja.md) の上に構築しています。同じ単一ファイルの
パイプライン（main.go → lexer.go → y.go）に、以下を追加します:

## goyacc1からの追加点

* `program: stmts` と `stmts: stmts stmt | /* empty */` — 入力が複数の
  ステートメントを持てるようになり、それぞれが自前の `Result:` を出力します。
  空の `stmts` ルールは、EOF上でのreduceの最初の例でもあります。
* `isWhiteSpace()` — lexerは `Lex()` ループの先頭で `' '`、`'\t'`、`'\n'`
  をスキップするようになったため、入力は自由に整形できます（goyacc1のlexerも
  同じことをするようになりましたが、ステートメントが行をまたぐ最初のステップは
  goyacc2です）。
* 新しく切り出されるトークン: `'='`、`';'`、`'{'`、`'}'` はlexerが返しますが、
  この段階ではどのルールでも使われません。これは意図的です: それらに対する
  parserの出力は `syntax error` となり、lexerと文法は文法が使うものについてのみ
  合致していればよいことを示しています。そしてこれらのトークンは
  goyacc2b/2cへの布石です。
* `%` は切り出され、ルールが与えられます: `expr '%' expr { $$ = $1 % $3 }`。
* `IDENT` ([a-z]+) は切り出されますが、文法はまだ受け付けないため、
  `a=3;` は `syntax error: unexpected IDENT` でパースに失敗します。
  変数はgoyacc2bで登場します。

## Build / Run / Test

```
make          # goyacc hoge.go.y -> y.go, then build ./hoge
make test     # exact assertions, exits non-zero on mismatch
make run      # read statements from stdin
make clean    # removes ./hoge y.go y.output
```

## 教えどころ

* `yyErrorVerbose = true` は `lexer.go` の `init()` で設定されています。
  `printf '1+\n' | ./hoge` を実行すると、単なる `syntax error` の代わりに
  詳細な形式
  （`syntax error: unexpected $end, expecting NUMBER or '('`）
  が見られます。文法に何かおかしな点があるときは、まずこのメッセージを
  調べるのがおすすめです。
* テストレシピ内の `%%` はmakeのエスケープではなく **printf** の
  エスケープです: `printf '20%%7 + 10%%3\n'` — 単一の `%7` はprintfに
  書式指定として消費されてしまいます。
* `1;` は `syntax error: unexpected $unk, at token: ;` で失敗します —
  文字を切り出すことと、文法がそれを受け付けることは別物です。`'='`、
  `';'`、`'{'`、`'}'` はこの段階ではすべて「切り出されるが未使用」です。
* goyacc1と同じく、不正な文字（たとえば `@`）の扱いは *lexer* の判断です:
  このlexerは `Error()`（構文エラー専用）を呼ばずにpanicします。
* main.goは `yyParse()` の戻り値をチェックし、失敗時には非ゼロで終了します。

## ファイル

* hoge.go.y ... 文法ファイル
* y.go/y.output ... `goyacc` が生成
* lexer.go ... 手書きのlexer
* main.go ... stdinを読み、lexerに渡し、`yyParse` を実行します
