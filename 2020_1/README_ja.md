# 2020_1 — flexのみのトークンプリンタ
> 英語版: [README.md](README.md)

このリポジトリの中で最もシンプルな例です。**純粋なflex**だけを使ったlexer（bison/yaccは登場しません）。入力トークンを認識し、マッチした内容を出力します。

## 学べること

- flexでの名前付きパターンの定義（`WHITESPACE`, `INTEGER`, `SYMBOL`, `LF`）
- 現在のマッチのテキストである `yytext` を出力するルールアクション
- すべてのルールに明示的なアクションが必要な理由: アクションを持たないルールにはflexの**デフォルトアクション（ECHO）**が適用され、マッチした内容がstdoutに出力されてしまいます。この例では、空白・改行・キャッチオールの `.` ルールすべてに明示的な空のアクションを与えており、トークンだけが出力されるようになっています。

## 構成ファイル

- `calc.l` — flexのlexer仕様（`main()` を含む）
- `Makefile` — `flex calc.l` を実行し、`lex.yy.c` を `$(CC) $(CFLAGS) ... $(LDFLAGS)` でコンパイルします（デフォルト: `clang`, `-O2`, `-lm`）

## ビルドと実行

    make
    ./calc

または `make run` でも実行できます。

## 実行例

    $ printf '12 + 3\n' | ./calc
    INTEGER: 12
    SYMBOL: +
    INTEGER: 3

空白、改行、未知の文字は黙ってスキップされます。
`make clean` で `calc` と `lex.yy.c` を削除します。
