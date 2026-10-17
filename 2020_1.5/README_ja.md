# 2020_1.5 — flex + bisonの"I LIKE <目的語>"デモ
> 英語版: [README.md](README.md)

初めてのflex + bisonの組み合わせの例です。`SUBJECT VERB OBJECT`（主語・動詞・目的語）という形の文を受け付け、目的語の値を出力する小さなgrammarです。

## 学べること

- lexerがbisonに渡すtokenコード（`SUBJECT`, `VERB`, `OBJECT`, `INTEGER`, `LF`）を返す仕組み
- `yylval` を介してlexerからparserに意味値（semantic value）を渡す方法（`%union { int int_value; }`）: `OBJECT` は常に値42を持ち、`INTEGER` はマッチした数値を持ちます
- 任意の行数を処理する開始記号を持つ、2ルールのgrammar

## grammar

    line:  LF                 { exit(0); }        -- a blank line ends the demo
        | SUBJECT VERB obj LF { printf("result: %d\n", $3); }
    obj  : OBJECT   {$$ = $1;}
         | INTEGER  {$$ = $1;}

lexerは単語をハードコードしています: `I` = SUBJECT、`LIKE` = VERB、`SOCCER` = OBJECT。空白はスキップされ、その他の文字は `Unknown character: '<c>'` としてstderrに報告されます。

## ビルドと実行

    make
    ./calc

または `make run` でも実行できます。

## 入力と出力の例

    $ printf 'I LIKE SOCCER\n' | ./calc
    result: 42
    $ printf 'I LIKE 3\n' | ./calc
    result: 3

## 不正な入力の場合

grammarにマッチしない入力は、stderr上でparse errorを引き起こします。例:

    $ printf 'dog barks 3\n' | ./calc
    Unknown character: 'd'
    Unknown character: 'o'
    Unknown character: 'g'
    Unknown character: 'b'
    Unknown character: 'a'
    Unknown character: 'r'
    Unknown character: 'k'
    Unknown character: 's'
    parse error near 3

`make clean` で `calc`、`calc.tab.c`、`lex.yy.c` を削除します。
