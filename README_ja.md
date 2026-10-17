# bison_flex_1

> 英語版: [README.md](README.md)

lexer/parserジェネレータを学ぶための、段階的な最小サンプル集です:
**flex + bison**（C）、**goyacc**（Go）、**YaccLexTools**（C#）。

各サンプルは意図的に小さく作られ、前のサンプルの上に少しずつ積み上げていきます。
順番に読み進めることで、lexer/parserのペアが小さなインタプリタへと育っていく
過程を見ることができます。

## 例（Examples）

| ディレクトリ | ツール           | 内容                                                                 |
| ---------- | ---------------- | -------------------------------------------------------------------- |
| `2020_1`   | flex             | flexのみ: 整数/演算子をトークン化して表示する                          |
| `2020_1.5` | flex + bison     | 最初の文法: 小さな "SUBJECT VERB OBJECT" 言語                           |
| `2020_2`   | flex + bison     | 整数電卓: precedence、`%prec UMINUS`、`YYERROR`、REPL                  |
| `goyacc1`  | goyacc           | 最小のgoyacc parser: 算術式の評価器                                    |
| `goyacc2`  | goyacc           | 複数ステートメント、空白の扱い、`%`                                    |
| `goyacc2b` | goyacc           | 変数と `put(...)` ステートメント                                       |
| `goyacc2c` | goyacc           | ASTを構築し `if` を追加（評価はパース後に行う）                        |
| `goyacc3`  | goyacc           | 式のASTを `%union` で持ち回る（interface型）                          |
| `goyacc4`  | goyacc           | ステートメント、`while` ループ、スコープ付き変数、ユニットテスト        |
| `goyacc5`  | goyacc           | Prolog風インタプリタ（単一化 + バックトラック）                        |
| `netyacc1` | YaccLexTools     | `goyacc1` の.NET移植（非推奨）                                         |
| `netyacc2` | YaccLexTools     | 小さな言語向けの.NETコンパイラ/インタプリタ（非推奨）                   |

各サンプルのディレクトリ（C、Go、.NET）には、ビルド/実行の手順を記した
固有の `README.md` があります。

## 学習パス

各トラックは順番に読むことを想定しています:

* **flex/bison (C)**: `2020_1`（lexerのみ。`yytext` とデフォルトのECHOが
  何を意味するか）→ `2020_1.5`（最初の文法、`yylval`/`%union`）→ `2020_2`
  （precedence、単項マイナス、エラー処理、REPL）。
* **goyacc (Go)**: `goyacc1`（文法とlexerの契約）→ `goyacc2`
  （複数ステートメント）→ `goyacc2b`（変数）→ `goyacc2c`（AST +
  遅延評価）→ `goyacc3`（`%union` による型付きAST）→ `goyacc4`
  （スコープ、`while`、テスト）→ `goyacc5`（goyaccの上に作る「本物の」言語）。
* **.NET**（非推奨・メンテナンス停止）: `netyacc1`（`goyacc1` の移植）→ `netyacc2`（コンパイラ + VM）。

## 必要なもの（Prereqs）

* C
  * bison & flex（macOSではXcodeのコマンドラインツールに含まれています）
  * Cコンパイラ（`clang` または `gcc`）
* Go
  * `go install golang.org/x/tools/cmd/goyacc@latest`
* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet tool install dotnet-ylt --global`
    （`.y`/`.lex` の入力を再生成する場合にのみ必要。生成済みのソースは
    コミットされています）

## 実行方法

* Go（各ディレクトリが独立したモジュールです）
  * `cd goyaccN && make` — parserの生成（goyacc）とビルド
  * `make test` — 検証可能なアサーション（一致しなければ非ゼロで終了）;
    `goyacc4`/`goyacc5` はさらに `go test` も実行します
  * `make run` — stdinをバイナリに渡す（goyacc1–4）
  * goyacc5はファイル引数が必要で、`goyacc5/` の中で実行する必要があります:
    `cd goyacc5 && make && ./prolog-interpreter -logLevel INFO test/test1.pro query/query1.pro`
* C
  * `cd 2020_2 && make && ./calc`（または `make run`; `2020_1` と `2020_1.5`
    にも `make run` があります）
* .NET（2つの例はどちらも非推奨・メンテナンス停止です）
  * [netyacc1/README.md](netyacc1/README.md) と
    [netyacc2/README.md](netyacc2/README.md) を参照してください

## 学習メモ（goyacc）

* **`Lex()` の契約**: トークン番号を返し、入力の終端では `0`（または負の値）
  を返します。goyaccは `0`/負の値をEOFトークン（`yyTok1[0]`）として扱うため、
  実際のトークンに対して *絶対に* `0` を返してはいけません — 「未知」の文字には
  宣言済みのトークンを与える必要があります（`goyacc3`/`goyacc4` の `ILLEGAL` を
  参照）。そうしないとparserはそこで黙って停止してしまいます。
* **`Error()`** は *構文* エラーに対してgoyaccから呼ばれます。*字句* エラーで
  何が起こるかはlexer次第です（panic、スキップ、またはエラートークンの生成）。
* **`yyErrorVerbose = true`** を設定すると、エラーメッセージに予期しない
  トークンと期待されたトークンの名前が表示されます。例:
  `syntax error: unexpected $end, expecting NUMBER or '('`。
  学習中に最も有用なスイッチです。
* **`yyDebug = 1..4`** でparserのshift/reduceトレースが出力されます。
* **`y.output`** は生成されたparserの状態表です。
  `N shift/reduce, M reduce/reduce conflicts reported` という行を見てください:
  `expr '/' expr` のような曖昧な文法は、`%left`/`%right` のprecedence宣言で
  コンフリクトなしにできます。
* 生成されたparserファイルは `make` が `.y` ファイルから生成します。その一部
  （`goyacc4/grammar.go`、`goyacc5/pkg/prolog/grammar.go`）は生成コードを読める
  ようにコミットされています。それ以外（`y.go`、`*.tab.c`、`y.output`）は
  gitignoreされており、必要に応じて再生成されます。`make clean` で削除できます。

## 学習メモ（flex/bison）

* **デフォルトのアクションはECHO**: `{...}` のないルールはマッチした文字列を
  stdoutに出力します。すべてのルールに明示的なアクションを与えましょう
  （`2020_1` を参照）。
* **`yylval` はlexer→parserの値の通り道**: lexerは `yylval.<member>` に値を
  格納し、文法はそれを `$1`/`$2`... として読みます
  （`2020_1.5` の `%union { int int_value; }` を参照）。
* **`%prec UMINUS`** は単項マイナスに専用のprecedenceを与えます: `'*'` の後に
  `%right UMINUS` を宣言すると、`-2 + 3` は `-(2+3) = -5` ではなく
  `(-2)+3 = 1` と評価され、自然な `expr: '-' expr` ルールがコンフリクトなしに
  なります（`2020_2`）。
* **`YYERROR` は現在のパースを中止します**: `error` リカバリルールがないと、
  ゼロ除算1回で `yyparse()` 全体が終了してしまい、REPLもそれとともに終わります
  （`2020_2` はこの挙動を実演しています）。
* **`bison -v` は `*.output` を書き出します**: goyaccの `y.output` に相当する
  C版で、同じshift/reduceとreduce/reduceのコンフリクト報告が含まれます。

## 便利なコマンド

* parserを手動で再生成する: `goyacc hoge.go.y`（Makefileが代わりにやってくれます）、
  `bison -v -t calc.y`、`flex calc.l`
* .NET <https://github.com/ernstc/YaccLexTools>
  * `dotnet ylt add-parser -p <parserName>`
