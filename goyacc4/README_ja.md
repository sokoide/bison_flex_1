# Goyacc4

> 英語版: [README.md](README.md)

* goyacc ミニコンパイラ: statement の AST、while ループ、scope のスタック

## ファイル

* grammar.go.y ... 文法ファイル。`go generate` で grammar.go と y.output を生成する
* grammar.go, y.output ... 生成された parser とその状態機械ダンプ
* lexer.go ... 手書きの scanner。parser が次の token を得るために `Lex()` を呼び、
  構文エラーのときには `Error()` が呼ばれる
* ast.go ... 式と statement 両方の AST ノード型
* eval.go ... statement と式をツリー步行で評価する interpreter
* scope.go ... while ループの本体を評価するときに使う scope のスタック
* main.go ... stdin からプログラムを読み込み、parse して評価する

## Build / Run / Test

```
make generate  # run goyacc on grammar.go.y -> grammar.go
make           # build ./mini-compiler (grammar.go must exist before go build)
make test      # go test -v ./...
make testcli   # exact assertions against the compiled binary
make clean     # removes ./mini-compiler grammar.go y.output
```

バイナリは `-logLevel (TRACE|DEBUG|INFO|WARN|ERROR)` を受け付けます。
`put()` の出力は stdout に、ログ行はすべて stderr に出るため、
`-logLevel ERROR` を指定すれば実行記録をきれいに保てます。

例（scope の規則に注意: `b` はループ内で作られ、ループの後には消えるため、
`a+b` は `5+0` になります）:

```
$ printf 'a=0;while(a<5){a=a+1;b=b+1;};put(a+b);\n' | ./mini-compiler -logLevel ERROR
5
$ printf 'put(1/0);\n' | ./mini-compiler -logLevel ERROR
time="..." level=error msg="[stmt 0] evaluation error: division by zero"
time="..." level=error msg="Evaluation error: division by zero"
(exit 1)
```

## これは何？

goyacc3 と似ていますが、各文法シンボルが 1 つの式ではなく
statement 全体を保持します:

1. `%union` は `exprs []expression`（リスト）、`expr expression`（1 ノード）、
   `tok token` を保持します。`%type<exprs> program stmts put_list` が
   non-terminal に型を付けます。

2. `program: stmts { $$ = $1; yylex.(*lexer).program = $$ }` が、
   完成した statement リストを lexer 経由で呼び出し元に返し、そこから
   `parse() <- ([]expression, error)` が `main` へ返します。

3. `%left '+' '-'` / `%left '*' '/' '%'` により、曖昧な `expr op expr`
   文法が conflict 無しになります（y.output で
   `0 shift/reduce, 0 reduce/reduce conflicts reported` と報告されます）。

4. 言語の仕様: 代入 `a=1;`、出力 `put("x=",expr,...);`、新しい scope を
   導入する `{...}` を伴う `while(cond){...}`、空 statement としての `;`、
   そして `while(...)` の中（`cond`）にしか現れない比較演算
   `== != >= > <= <`。

5. `scope.go` は scope のスタックを実装しています。`while` の本体は
   scope を push し、ループが終わると pop するため、ループ内で作られた
   変数はその後消えます（`while(a<5){a=a+1;b=b+1;}` の `b`）。

6. エラー処理: `yyErrorVerbose = true` により `unexpected ILLEGAL`
   のようなメッセージが得られます。「未知の文字が parse を終了させては
   いけない」ことは goyacc の古典的な落とし穴です。`Lex()` が 0 または
   負の値を返すと end-of-input を意味するため、scanner は代わりに
   `ILLEGAL` token を返します。lexer と parser のエラーは Go の error と
   して返され（lexer は `os.Exit` を呼ぶ代わりにメッセージを保存する）、
   評価時のエラー（0 除算、`put()` の外で使われる文字列）は
   プログラムを非ゼロの exit status で異常終了させます。

7. 既知の制限: 単項マイナスはありません（`put(-1);` は構文エラー）。
   文字列は `put()` の直接のリテラルとしてのみ有効です（そのため
   `"a"+"b"`、`a="x";`、`"a"=="b"` は評価エラーになります）。未定義の
   変数を読むと 0 になります（現在の scope に自動生成されます）。

`yyDebug` のレベル（生成された parser 内で `yyDebug = N` を設定）を指定すると
parser の状態遷移が表示され、`y.output` と合わせてもう一つの
有用な学習ツールになります。
