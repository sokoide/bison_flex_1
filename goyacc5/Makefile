TARGET=./prolog-interpreter
TESTSRC=./tmp.pro
LL=INFO
SRCS=./cmd/prolog-interpreter/*.go ./pkg/prolog/*.go
.PHONY: all clean generate run test testparser testquery

all: $(TARGET)

$(TARGET): pkg/prolog/grammar.go $(SRCS)
	go build -o $(TARGET) ./cmd/prolog-interpreter

generate: pkg/prolog/grammar.go

pkg/prolog/grammar.go: pkg/prolog/grammar.go.y
	go generate ./pkg/prolog/...

clean:
	rm $(TARGET) pkg/prolog/y.output pkg/prolog/grammar.go

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	go test -v ./...

testparser: $(TARGET)
	for file in ./test/*.pro; do \
	$(TARGET) -logLevel $(LL) $$file; \
	done

testquery: $(TARGET)
	$(TARGET) -logLevel $(LL) ./test/test1.pro ./query/query1.pro
	$(TARGET) -logLevel $(LL) ./test/test2.pro ./query/query2.pro
	$(TARGET) -logLevel $(LL) ./test/test3.pro ./query/query3.pro