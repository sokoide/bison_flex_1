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
	echo "20%7 + 10%3" | $(TARGET)
	@echo "did you see 7?"
	echo "(1 * 2) +6/3" | $(TARGET)
	@echo "did you see 4?"
	echo "(4 + 1) + 2/2" | $(TARGET)
	@echo "did you see 6?"
	echo "(40 + 1) + 2/2" | $(TARGET)
	@echo "did you see 42?"
	echo "1 + 10*20*30" | $(TARGET)
	@echo "did you see 6001?"
