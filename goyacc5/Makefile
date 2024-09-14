TARGET=./prolog-interpreter
TESTSRC=./tmp.pro
LL=DEBUG
SRCS=./cmd/prolog-interpreter/*.go ./pkg/prolog/*.go
.PHONY: all clean generate run test testcli

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

testcli: $(TARGET)
	@echo 'authorize(taro).' > $(TESTSRC)
	$(TARGET) -logLevel $(LL) $(TESTSRC)
	@echo 'did you see true?'
	@rm $(TESTSRC)