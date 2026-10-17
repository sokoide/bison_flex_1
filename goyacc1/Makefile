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
	@set -e; out=$$(printf '1*2+6/3\n' | $(TARGET)); \
	if [ "$$out" != "Result: 4" ]; then echo "FAIL: 1*2+6/3 -> got '$$out', want 'Result: 4'"; exit 1; fi; \
	echo "ok: 1*2+6/3 -> 4"
	@set -e; out=$$(printf '4+1*4/2\n' | $(TARGET)); \
	if [ "$$out" != "Result: 6" ]; then echo "FAIL: 4+1*4/2 -> got '$$out', want 'Result: 6'"; exit 1; fi; \
	echo "ok: 4+1*4/2 -> 6"
	@set -e; out=$$(printf '(4+1)*4/2\n' | $(TARGET)); \
	if [ "$$out" != "Result: 10" ]; then echo "FAIL: (4+1)*4/2 -> got '$$out', want 'Result: 10'"; exit 1; fi; \
	echo "ok: (4+1)*4/2 -> 10"
	@set -e; out=$$(printf '1+10*(10+10)*30\n' | $(TARGET)); \
	if [ "$$out" != "Result: 6001" ]; then echo "FAIL: 1+10*(10+10)*30 -> got '$$out', want 'Result: 6001'"; exit 1; fi; \
	echo "ok: 1+10*(10+10)*30 -> 6001"
