TARGET=./hoge
SRCS=main.go lexer.go

all: $(TARGET)

$(TARGET): y.go $(SRCS)
	go build -o $(TARGET)

y.go: hoge.go.y
	goyacc $<

clean:
	rm -f $(TARGET) y.go y.output

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	@set -e; out=$$(printf '20%%7 + 10%%3\n' | $(TARGET)); \
	if [ "$$out" != "Result: 7" ]; then echo "FAIL: 20%7+10%3 -> got '$$out', want 'Result: 7'"; exit 1; fi; \
	echo "ok: 20%7 + 10%3 -> 7"
	@set -e; out=$$(printf '(1 * 2) +6/3\n' | $(TARGET)); \
	if [ "$$out" != "Result: 4" ]; then echo "FAIL: (1*2)+6/3 -> got '$$out', want 'Result: 4'"; exit 1; fi; \
	echo "ok: (1 * 2) +6/3 -> 4"
	@set -e; out=$$(printf '(4 + 1) + 2/2\n' | $(TARGET)); \
	if [ "$$out" != "Result: 6" ]; then echo "FAIL: (4+1)+2/2 -> got '$$out', want 'Result: 6'"; exit 1; fi; \
	echo "ok: (4 + 1) + 2/2 -> 6"
	@set -e; out=$$(printf '(40 + 1) + 2/2\n' | $(TARGET)); \
	if [ "$$out" != "Result: 42" ]; then echo "FAIL: (40+1)+2/2 -> got '$$out', want 'Result: 42'"; exit 1; fi; \
	echo "ok: (40 + 1) + 2/2 -> 42"
	@set -e; out=$$(printf '1 + 10*20*30\n' | $(TARGET)); \
	if [ "$$out" != "Result: 6001" ]; then echo "FAIL: 1+10*20*30 -> got '$$out', want 'Result: 6001'"; exit 1; fi; \
	echo "ok: 1 + 10*20*30 -> 6001"
