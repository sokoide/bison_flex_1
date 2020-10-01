TARGET=calc

all: $(TARGET)

$(TARGET): calc.tab.c
	gcc -O2 -o $(TARGET) calc.tab.c -ly -lm

calc.tab.c: calc.y lex.yy.c
	bison calc.y

lex.yy.c: calc.l
	flex calc.l

clean:
	rm $(TARGET) *.tab.c lex.yy.c
