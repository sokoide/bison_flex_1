TARGET=./calc
SRCS=main.go lexer.go eval.go ast.go

all: $(TARGET)

$(TARGET): y.go $(SRCS)
	go build -o $(TARGET)

y.go: calc.go.y
	goyacc $<

clean:
	rm $(TARGET) y.go y.output

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	echo "(1*2)+10/3" | $(TARGET)
	@echo "did you see 5?"
	echo "(40+1)+2/2" | $(TARGET)
	@echo "did you see 42?"
