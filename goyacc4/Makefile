TARGET=./mini-compiler
SRCS=main.go lexer.go eval.go ast.go
LL=DEBUG

.PHONY: all clean generate run test testcli

all: $(TARGET)

$(TARGET): grammar.go $(SRCS)
	go build -o $(TARGET)

generate: grammar.go

grammar.go: grammar.go.y
	go generate ./

clean:
	rm $(TARGET) y.output grammar.go

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	go test -v ./...

testcli: $(TARGET)
	echo 'put("ans=",1*2);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see ans=2?'
	echo 'put("ans=",(1*2)+10/3);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see ans=5?'
	echo 'put("ans=",(40+1)+2/2);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see ans=42?'
	echo 'a=40;b=2;c=a+b;put("c=",c);' | $(TARGET) -logLevel $(LL)
	@echo 'did you see c=42?'
	echo 'a=0;while(a<5){a=a+1;};put(a);' | $(TARGET) -logLevel $(LL)
	@echo 'did you 5?'
