# goyacc5

> 英語版: [README.md](README.md)

## 概要

* goyacc の上に載った Prolog 風言語: fact、rule、query に
  unification と深さ優先の backtracking を備えます。
* 以前のステップと同じ構成 — goyacc 文法 → 生成された parser →
  手書きの lexer — ですが、今回は本物の（小さな）言語になっています。

## Build / Run / Test

```sh
make           # goyacc on pkg/prolog/grammar.go.y, build ./prolog-interpreter
make test      # go test -v ./... (unifier + lexer unit tests)
make testparser  # parse/evaluate each test/*.pro program on its own
make testquery   # for each query/queryN.pro: run it against test/testN.pro
make clean     # removes the binary, pkg/prolog/y.output and grammar.go
               # (make regenerates grammar.go automatically)
```

テストスクリプトは `LOGLEVEL` を尊重します（デフォルトは `INFO`）:

```sh
make testparser            # INFO logging
LOGLEVEL=DEBUG make testparser   # verbose lexer/parser traces
make testquery
LOGLEVEL=DEBUG make testquery
```

`make testquery` はファイルを**番号でペアにします**: `query/query5.pro` は
`test/test5.pro` に対して実行されます。そのためファイルが欠けていたり
名前が間違っていたりするとターゲットが失敗します
（`set -e` により最初の失敗で中断します）。

### 手で実行する — このディレクトリの中からのみ

```sh
./prolog-interpreter -logLevel INFO test/test1.pro query/query1.pro
```

interpreter は `resource/builtin.pro` を**現在のディレクトリからの相対
パス**で読み込むため、`goyacc5/` の中で起動する必要があります
（すべての `make` ターゲットはあなたの代わりにそこで実行します）。
リポジトリのルートから実行すると次のようになります:

```
level=fatal msg="... failed to open resource/builtin.pro: no such file or directory"
```

`make run` は引数なしでバイナリを起動しますが、これは
`Usage: prolog <source-file> <query-file>` を表示するだけなので、
上のようにバイナリを直接呼び出してください
（引数 0 = 読み込むプログラム、引数 1 = 実行する query）。

> **Transcript format**: 以下の例は対話的なターミナルでの出力を示します。
> そこでは logrus が `INFO[0000] ...` のように表示します。出力を
> リダイレクトしたりパイプすると、同じログ行が
> `time="..." level=info msg="..."` の形で出力されます — 内容は同一です。

## ファイル

* `pkg/prolog/grammar.go.y` ... 文法（goyacc への入力）。`grammar.go`
  （コミットされているので、生成された parser を読めます）と
  `y.output`（状態機械 / conflict レポート）を生成する
* `pkg/prolog/lexer.go` ... tokenizer。lexer エラーは panic する代わりに
  保存され、`Lexer.Err()` 経由で返される
* `pkg/prolog/term.go` ... term の表現（atom、number、string、
  variable、複合項、list）
* `pkg/prolog/clause.go` ... fact と rule の clause
* `pkg/prolog/unifier.go` ... `unify()` と rule 本体の backtracking
* `pkg/prolog/eval.go` ... `Load()`（builtin とプログラムを parse）、
  `Query()`、解の表示
* `pkg/prolog/lexer_test.go`, `pkg/prolog/unifier_test.go` ... unit tests
* `cmd/prolog-interpreter/main.go` ... flags、2 つのファイルの読み込み、
  query の実行
* `resource/builtin.pro` ... 組み込み predicate（`write/1`、`writeln/1`）
* `test/testN.pro` ... サンプルプログラム。`query/queryN.pro` ... 同じ
  番号のプログラムに対して実行する query

## 仕組み

1. `Load()` は `resource/builtin.pro` を goyacc の parser で parse し、
   それからプログラムファイルを parse します。action
   `input: clause_list { yylex.(*lexer).program = $1 }` が clause の
   リストを lexer 経由で返します — goyacc2c/3/4 と同じトリックです。
2. query ファイルも同じ方法で parse されます（query も clause の一種です:
   `food(orange).` はドットで終わります）。
3. `Query()` は各 query term に対して clause を上から順にたどります。
   `unify()` が goal を clause の head と照合し、substitution
   （variable への束縛）を生成します。rule 本体は再帰的に解かれ、
   すべての選択点が alternatives を保持するため、後の goal が失敗すると
   そこへ backtracking します（深さ優先）。
4. 解は束縛を query に適用して表示されます（`X = 5`）。また rule 本体の
   中の `writeln/1` は出力を書くという副作用を持ちます。研究に値する
   クセ: `test5.pro` の `print_list([5,6,7,8,9])` はすべての要素を表示
   したうえで、最終的に `false` になります — `[H|T]` がリストを使い
   果たしたあと、基本ケース `print_list([]).` は空の tail と unify
   しなくなるからです（空リストの unification はリテラル対リテラルでしか
   働きません。後述の "List is partially supported" を参照）。副作用は
   取り消されません。

## 文法

* [Grammar file](./pkg/prolog/grammar.go.y)、要約（action は省略）:

  ```
  input       : /* empty */ | clause_list
  clause_list : clause | clause_list clause
  clause      : fact_clause | rule_clause
  fact_clause : term '.'
  rule_clause : term COLON_DASH term_list '.'        -- "head :- body."
  term_list   : term | term_list ',' term
  term        : IDENT | VAR | '_' | NUMBER_LITERAL | STRING_LITERAL
              | IDENT '(' term_list ')'              -- compound: f(A,B)
              | '[' term_list ']'                    -- [5,6,7]
              | '[' term_list '|' term ']'           -- [H|T]
              | '[' ']'                              -- []
              | BUILTIN_WRITE '(' term_list ')' | BUILTIN_NL
  ```

  内部ではリストが 3 通りの方法で表現されていることに注意してください
  （`Args`、`Head`/`Tail`、`IsEmpty`）— この近道こそが、list 対応を
  「部分的」にしている理由です（How it works の `print_list` のクセを
  参照）。

* サポートされる組み込み predicate
  * write/1
  * writeln/1
* List は部分的にしかサポートされていません。`test5.pro` と
  `query5.pro` を参照してください。

## TODO

* 次を含む多くの機能が未実装です:
  * comparison (> >= < =< =:= =\= or etc) operators
  * arithmetic (+ - * / ** // mod) operators
  * cut (!) operator
  * その他の組み込み predicate
* lexer は `+ - * / = ; { }` の tokenize には既に対応していますが、
  これらを受け付ける文法規則はまだありません。対応する規則が文法に
  追加されるまで、.pro ファイルでこれらを使うと構文エラーになります。

## 例

### 単純な fact / rule の評価

```sh
$ cat test/test1.pro
food(orange).
food(apple).

meal(X) :- food(X).

$ cat query/query1.pro
food(orange).
meal(orange).
food(car).

$ ./prolog-interpreter -logLevel INFO  test/test1.pro  query/query1.pro
INFO[0000] querying...
INFO[0000] food(orange). -> true
INFO[0000] meal(orange). -> true
INFO[0000] food(car). -> false
INFO[0000] query completed.
```

### 再帰を伴う単純な fact / rule の評価

```sh
$ cat test/test2.pro
% facts
mammal(dog).
mammal(cat).
mammal(human).
fish(cod).
fish(swordfish).
fish(mackerel).

% rules
animal(X) :- mammal(X).
animal(X) :- fish(X).

print_if_mammal(X) :-
    mammal(X),
    write(X),
    writeln(" is mammal").

% end of file

$ cat query/query2.pro
mammal(dog).
fish(dog).
animal(dog).

mammal(cod).
fish(cod).
animal(cod).

print_if_mammal(dog).
print_if_mammal(cod).

$ ./prolog-interpreter -logLevel INFO  test/test2.pro  query/query2.pro
INFO[0000] querying...
INFO[0000] mammal(dog). -> true
INFO[0000] fish(dog). -> false
INFO[0000] animal(dog). -> true
INFO[0000] mammal(cod). -> false
INFO[0000] fish(cod). -> true
INFO[0000] animal(cod). -> true
dog is mammal
INFO[0000] print_if_mammal(dog). -> true
INFO[0000] print_if_mammal(cod). -> false
INFO[0000] query completed.
```

### 複合した rule（grandparent）

```sh
$ cat test/test3.pro
father(ichiro, jiro).
father(jiro, saburo).
mother(alice, jiro).
mother(becky, saburo).

parent(X,Y) :- father(X,Y).
parent(X,Y) :- mother(X,Y).
grandparent(A, B) :- parent(A, C), parent(C, B).

% end of file

$ cat query/query3.pro
parent(ichiro, jiro).
parent(alice, jiro).
parent(becky, jiro).
grandparent(ichiro, saburo).
grandparent(alice, saburo).
grandparent(alice, jiro).

$ ./prolog-interpreter -logLevel INFO  test/test3.pro  query/query3.pro
INFO[0000] querying...
INFO[0000] parent(ichiro, jiro). -> true
INFO[0000] parent(alice, jiro). -> true
INFO[0000] parent(becky, jiro). -> false
INFO[0000] grandparent(ichiro, saburo). -> true
INFO[0000] grandparent(alice, saburo). -> true
INFO[0000] grandparent(alice, jiro). -> false
INFO[0000] query completed.
```

### List

```sh
$ cat test/test5.pro
first([X|_], X).
second([_,X|_], X).

print_list([]).

print_list([H|T]) :-
    writeln(H),
    print_list(T).

$ cat query/query5.pro
first([5,6,7], X).
second([5,6,7], X).
first([hello,world], Y).
print_list([5,6,7,8,9]).

$ ./prolog-interpreter -logLevel INFO  test/test5.pro query/query5.pro
INFO[0000] querying...
X = 5
INFO[0000] first([5, 6, 7], X). -> true
X = 6
INFO[0000] second([5, 6, 7], X). -> true
Y = hello
INFO[0000] first([hello, world], Y). -> true
5
6
7
8
9
INFO[0000] print_list([5, 6, 7, 8, 9]). -> false
INFO[0000] query completed.
```
