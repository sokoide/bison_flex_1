# Goyacc3

> 英語版: [README.md](README.md)

* goyacc 電卓: 式の AST と Go による評価

## ファイル

* calc.go.y ... 文法ファイル。y.go（parser）と y.output（状態機械のダンプ）を生成する
* y.go, y.output ... 生成された parser とその conflict/状態レポート
* lexer.go ... 手書きの scanner。parser が次の token を得るために `Lex()` を呼び、
  構文エラーのときには `Error()` が呼ばれる
* ast.go ... AST ノードの型（number / paren / 二項演算子）
* eval.go ... ツリーをたどって評価する evaluator
* main.go ... stdin を 1 つの式として読み込み、parse してから評価する

## Build / Run / Test

```
make          # build ./calc (runs goyacc first if y.go is out of date)
make run      # read a program from stdin
make test     # exact assertions, exits non-zero on mismatch
make clean    # removes ./calc y.go y.output (make regenerates them)
```

例:

```
$ printf '(1*2)+10/3\n' | ./calc
5
$ printf '1+@\n' | ./calc
2026/10/17 12:45:37 parse error: Line 1, Column 3: "@" syntax error: unexpected ILLEGAL, expecting NUMBER or '('
(exit 1; the date/time prefix comes from Go's default logger)
```

goyacc3 は、[2020_2](../2020_2/README_ja.md) の C 言語電卓の Go 版にあたるものです。
両者を並べて読むと、同じ precedence/AST の考え方が両方の toolchain で
実現されていることがわかります。

## これは何？

1. `%union` は parser と semantic action の間で値を受け渡します。
   ここでは `expr of type expression`（AST ノード）と `tok of type token`
   を保持しています。

2. `program: expr { $$ = $1; yylex.(*lexer).program = $$ }` が、
   完成した AST を呼び出し元に返します。goyacc では reduce の結果は
   `yySymType` の union を通してしか受け渡せないため、action で
   lexer に根ノードを保存し、そこから `parse()` が `main` へ返します。

3. `%left '+' '-'` / `%left '*' '/' '%'` は、自然に書いた
   `expr op expr` 規則の曖昧さを解消します。これらが無いと文法は曖昧になり、
   goyacc は shift/reduce conflict を報告します。後ろの precedence 宣言ほど
   強く結合し、同じレベルの演算子は左結合になります。生成された状態機械と
   `0 shift/reduce, 0 reduce/reduce conflicts reported` という行は
   `y.output` で確認できます。

4. `yyErrorVerbose = true` にすると、parser が `syntax error` の代わりに
   `unexpected ILLEGAL` と表示するようになり、`yyDebug` と合わせて
   学習にとても役立ちます。

5. Lexer の契約: `Lex()` が 0（または負の値）を返すと goyacc は
   end-of-input とみなすため、未知の文字に対して token 番号 0 を返しては
   いけません。この例では未知の文字に対して `ILLEGAL` token を返すので、
   parser が残りの行を黙って捨てるのではなく、目に見えるエラーで
   失敗します。

進行の流れ: goyacc3 は 1 つの式を評価してその値を表示します。たとえば
`(1*2)+10/3` は `5` と表示され、`(40+1)+2/2` は `42`、`3%4` は `3` と
表示されます。変数と statement は次の goyacc4 で扱います。
