TARGET=./hoge
SRCS=main.go lexer.go

all: $(TARGET)

$(TARGET): y.go $(SRCS)
	go build -o $(TARGET)

y.go: hoge.go.y
	goyacc $<

clean:
	rm $(TARGET) y.go

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	echo "(1*2)+6/3" | $(TARGET)
	@echo "did you see 4?"
	echo "(4+1)+2/2" | $(TARGET)
	@echo "did you see 6?"
