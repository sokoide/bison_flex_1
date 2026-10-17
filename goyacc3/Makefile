TARGET=./calc
SRCS=main.go lexer.go eval.go ast.go

all: $(TARGET)

$(TARGET): y.go $(SRCS)
	go build -o $(TARGET)

y.go: calc.go.y
	goyacc $<

clean:
	rm -f $(TARGET) y.go y.output

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	@set -e; out=$$(printf '(1*2)+10/3\n' | $(TARGET)); \
	if [ ! "$$out" = "5" ]; then echo "FAIL: (1*2)+10/3 -> got '$$out'"; exit 1; fi; \
	echo "ok: (1*2)+10/3 -> 5"
	@set -e; out=$$(printf '(40+1)+2/2\n' | $(TARGET)); \
	if [ ! "$$out" = "42" ]; then echo "FAIL: (40+1)+2/2 -> got '$$out'"; exit 1; fi; \
	echo "ok: (40+1)+2/2 -> 42"
	@set -e; out=$$(printf '3%%4\n' | $(TARGET)); \
	if [ ! "$$out" = "3" ]; then echo "FAIL: 3%4 -> got '$$out'"; exit 1; fi; \
	echo "ok: 3%4 -> 3"
	@if printf '10/0\n' | $(TARGET) 2>/dev/null; \
	then echo "FAIL: 10/0 should exit non-zero"; exit 1; \
	else echo "ok: 10/0 exits with an error"; fi
	@if printf '1+@\n' | $(TARGET) >/dev/null 2>&1; \
	then echo "FAIL: unknown character should exit non-zero"; exit 1; \
	else echo "ok: unknown character (ILLEGAL) exits with an error"; fi
