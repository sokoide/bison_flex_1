TARGET=calc

all: $(TARGET)

$(TARGET): lex.yy.c
	gcc -O2 -o $(TARGET) lex.yy.c -lm -ly

lex.yy.c: calc.l
	flex calc.l

clean:
	rm $(TARGET) lex.yy.c
