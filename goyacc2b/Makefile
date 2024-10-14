TARGET=./hoge
SRCS=main.go lexer.go

all: $(TARGET)

$(TARGET): y.go $(SRCS)
	go build -o $(TARGET)

y.go: hoge.go.y
	goyacc $<

clean:
	rm $(TARGET) y.go y.output

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	echo "1+2;3*4;55+66;" | $(TARGET)
	@echo "did you see 3,12,121?"
	echo "put(42);" | $(TARGET)
	@echo "did you see 42?"
	echo "a=3;put(a+4);" | $(TARGET)
	@echo "did you see 7?"
