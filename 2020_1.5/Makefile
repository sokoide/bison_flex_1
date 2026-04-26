TARGET=./calc
CC=gcc
CFLAGS=-O2
LDFLAGS=-lm

all: $(TARGET)

$(TARGET): calc.tab.c
	$(CC) $(CFLAGS) -o $(TARGET) calc.tab.c $(LDFLAGS)

calc.tab.c: calc.y lex.yy.c
	bison -v -t calc.y

lex.yy.c: calc.l
	flex calc.l

clean:
	rm -f $(TARGET) *.tab.c lex.yy.c

run: $(TARGET)
	$(TARGET)
