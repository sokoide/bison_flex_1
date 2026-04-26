TARGET=calc
CC=clang
CFLAGS=-O2
LDFLAGS=-lm

all: $(TARGET)

$(TARGET): lex.yy.c
	$(CC) $(CFLAGS) -o $(TARGET) lex.yy.c $(LDFLAGS)

lex.yy.c: calc.l
	flex calc.l

run: $(TARGET)
	./$(TARGET)

clean:
	rm $(TARGET) lex.yy.c
