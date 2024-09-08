TARGET=./mini-compiler
SRCS=main.go lexer.go eval.go ast.go

all: $(TARGET)

$(TARGET): y.go $(SRCS)
	go build -o $(TARGET)

y.go: *.go.y
	goyacc $<

clean:
	rm $(TARGET) y.go y.output

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	echo "put(1*2);" | $(TARGET)
	@echo "did you see 2?"
	echo "put((1*2)+10/3);" | $(TARGET)
	@echo "did you see 5?"
	echo "put((40+1)+2/2);" | $(TARGET)
	@echo "did you see 42?"
	echo "a=40;b=2;c=a+b;put(c);" | $(TARGET)
	@echo "did you see c = 42?"
