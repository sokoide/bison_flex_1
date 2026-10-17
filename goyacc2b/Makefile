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
	@set -e; out=$$(printf 'put(42);\n' | $(TARGET)); \
	if [ "$$out" != "42" ]; then echo "FAIL: put(42); -> got '$$out', want '42'"; exit 1; fi; \
	echo "ok: put(42); -> 42"
	@set -e; out=$$(printf 'put(1+2);put(3*4);put(55+66);\n' | $(TARGET)); \
	expected=$$(printf '3\n12\n121'); \
	if [ "$$out" != "$$expected" ]; then echo "FAIL: put(1+2);put(3*4);put(55+66); -> got '$$out', want '3\n12\n121'"; exit 1; fi; \
	echo "ok: put(1+2);put(3*4);put(55+66); -> 3,12,121"
	@set -e; out=$$(printf 'a=3;put(a+4);\n' | $(TARGET)); \
	if [ "$$out" != "7" ]; then echo "FAIL: a=3;put(a+4); -> got '$$out', want '7'"; exit 1; fi; \
	echo "ok: a=3;put(a+4); -> 7"
