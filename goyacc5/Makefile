TARGET=./prolog-interpreter
TESTSRC=./tmp.pro
LL := $(or $(shell echo $$LOGLEVEL), "DEBUG")

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
	echo "* Testing $$file"; \
	$(TARGET) -logLevel $(LL) $$file; \
	done

testquery: $(TARGET)
	num=`ls query/*.pro |wc -l`; \
	for i in `seq $$num`; do \
	echo "* Testing query$$i.pro"; \
	$(TARGET) -logLevel $(LL) ./test/test$$i.pro ./query/query$$i.pro; \
	done