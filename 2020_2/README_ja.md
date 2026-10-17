# 2020_2 — flex + bisonの整数電卓
> 英語版: [README.md](README.md)

整数演算用のREPL電卓です。`+ - * /`、括弧、単項マイナスをサポートし、1行ずつ評価します。

## 学べること

- bisonでの演算子のprecedence宣言:
  `%left '+' '-'` / `%left '*' '/'` / `%right UMINUS`
- `%prec UMINUS` を使った単項マイナス: `- expr %prec UMINUS { $$ = -$2; }`
- `%union { int int_value; }` と、`calc.l` で各integerトークンに対して設定される `yylval.int_value` によるlexerからparserへの値渡し
- `YYERROR` を使ったエラー処理: 除算ルールがゼロ除算を検出すると `yyerror("division by zero")` を呼び出し、`YYERROR` で現在のparseを中断します

## 構成ファイル

- `calc.y` — grammar、precedence宣言、エラー報告（`yyerror`）
- `calc.l` — lexer: 整数、演算子、`LF` トークン。空白はスキップされ、未知の文字はstderrに報告されます
- `calc.y` は `#include "lex.yy.c"` によって `lex.yy.c` を直接取り込んでいます

## ビルドと実行

    make
    ./calc

または `make run` でも実行できます。その他のターゲット: `make debug`（`-g -DDEBUG` で再ビルド）と `make help`（ターゲット一覧を表示）。

## 実行例

    $ printf '1 + 2\n(40+1)+2/2\n1/0\n' | ./calc
    Simple Calculator
    Enter expressions (Ctrl+D to exit):
    3
    42
    Error: division by zero

    Goodbye!

（`Error: division by zero` の行はstderrに出力されるため、stdoutがバッファリングされるとtranscript上の位置が変わることがあります。）

注意: `1/0` は現在の行をスキップするだけではありません。grammarに `error` リカバリルールがないため、`YYERROR` が `yyparse()` 全体を中断し、プログラムが終了します:

    $ printf '1/0\n2+2\n' | ./calc
    Simple Calculator
    Enter expressions (Ctrl+D to exit):
    Error: division by zero

    Goodbye!
    (exit code 1; the "2+2" line is never evaluated)

`make clean` で `calc`、`calc.tab.c`、`calc.tab.h`、`lex.yy.c`、`calc.output` を削除します。
