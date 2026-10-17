TARGET=./prolog-interpreter
TESTSRC=./tmp.pro
# LL := $(or $(shell echo $$LOGLEVEL), "DEBUG")
LL := $(or $(shell echo $$LOGLEVEL), "INFO")

SRCS=./cmd/prolog-interpreter/*.go ./pkg/prolog/*.go
.PHONY: all clean generate run test testparser testquery

all: $(TARGET)

$(TARGET): pkg/prolog/grammar.go $(SRCS)
	go build -o $(TARGET) ./cmd/prolog-interpreter

generate: pkg/prolog/grammar.go

pkg/prolog/grammar.go: pkg/prolog/grammar.go.y
	go generate ./pkg/prolog/...

clean:
	rm -f $(TARGET) pkg/prolog/y.output pkg/prolog/grammar.go

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	go test -v ./...

testparser: $(TARGET)
	set -e; for file in ./test/*.pro; do \
	echo "* Testing $$file"; \
	$(TARGET) -logLevel $(LL) $$file; \
	done

testquery: $(TARGET)
	set -e; for query in ./query/*.pro; do \
	name=$${query##*/}; number=$${name#query}; number=$${number%.pro}; \
	echo "* Testing $$name"; \
	$(TARGET) -logLevel $(LL) ./test/test$$number.pro $$query; \
	done

debug: $(TARGET)
	$(TARGET) -logLevel $(LL) ./test/test5.pro ./query/query5.pro
