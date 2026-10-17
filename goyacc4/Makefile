TARGET=./mini-compiler
SRCS=main.go lexer.go eval.go ast.go scope.go
LL=ERROR

.PHONY: all clean generate run test testcli

all: $(TARGET)

$(TARGET): grammar.go $(SRCS)
	go build -o $(TARGET)

generate: grammar.go

grammar.go: grammar.go.y
	go generate ./

clean:
	rm -f $(TARGET) y.output grammar.go

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	go test -v ./...

testcli: $(TARGET)
	@set -e; out=$$(printf 'put("ans=",1*2);\n' | $(TARGET) -logLevel $(LL)); \
	if [ ! "$$out" = "ans=2" ]; then echo 'FAIL: put("ans=",1*2); ->'" got '$$out'"; exit 1; fi; \
	echo 'ok: put("ans=",1*2); -> ans=2'
	@set -e; out=$$(printf 'put("ans=",(1*2)+10/3);\n' | $(TARGET) -logLevel $(LL)); \
	if [ ! "$$out" = "ans=5" ]; then echo 'FAIL: put("ans=",(1*2)+10/3); ->'" got '$$out'"; exit 1; fi; \
	echo 'ok: put("ans=",(1*2)+10/3); -> ans=5'
	@set -e; out=$$(printf 'put("ans=",(40+1)+2/2);\n' | $(TARGET) -logLevel $(LL)); \
	if [ ! "$$out" = "ans=42" ]; then echo 'FAIL: put("ans=",(40+1)+2/2); ->'" got '$$out'"; exit 1; fi; \
	echo 'ok: put("ans=",(40+1)+2/2); -> ans=42'
	@set -e; out=$$(printf 'a=40;b=2;c=a+b;put("c=",c);\n' | $(TARGET) -logLevel $(LL)); \
	if [ ! "$$out" = "c=42" ]; then echo 'FAIL: a=40;b=2;c=a+b;put("c=",c); ->'" got '$$out'"; exit 1; fi; \
	echo 'ok: a=40;b=2;c=a+b;put("c=",c); -> c=42'
	@set -e; out=$$(printf 'a=0;while(a<5){a=a+1;};put(a);\n' | $(TARGET) -logLevel $(LL)); \
	if [ ! "$$out" = "5" ]; then echo 'FAIL: a=0;while(a<5){a=a+1;};put(a); ->'" got '$$out'"; exit 1; fi; \
	echo 'ok: a=0;while(a<5){a=a+1;};put(a); -> 5'
	@set -e; out=$$(printf 'a=0;while(a<5){a=a+1;b=b+1;};put(a+b);\n' | $(TARGET) -logLevel $(LL)); \
	if [ ! "$$out" = "5" ]; then echo 'FAIL: a=0;while(a<5){a=a+1;b=b+1;};put(a+b); ->'" got '$$out'"; exit 1; fi; \
	echo 'ok: a=0;while(a<5){a=a+1;b=b+1;};put(a+b); -> 5 (undefined b reads as 0)'
	@if printf 'put("a"+"b");\n' | $(TARGET) -logLevel $(LL) >/dev/null 2>&1; \
	then echo 'FAIL: put("a"+"b"); should exit non-zero (strings are not supported in arithmetic)'; exit 1; \
	else echo 'ok: put("a"+"b"); exits with an evaluation error'; fi
