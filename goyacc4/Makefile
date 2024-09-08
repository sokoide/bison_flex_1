TARGET=./mini-compiler
SRCS=main.go lexer.go eval.go ast.go
LL=DEBUG

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
	echo 'put("ans=",1*2);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see ans=2?'
	echo 'put("ans=",(1*2)+10/3);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see ans=5?'
	echo 'put("ans=",(40+1)+2/2);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see ans=42?'
	echo 'a=40;b=2;c=a+b;put("c=",c);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see c=42?'
