TARGET=./hoge
SRCS=$(wildcard *.go)

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
	@set -e; out=$$(printf 'if(1){put(1);}\n' | $(TARGET)); \
	if [ "$$out" != "1" ]; then echo "FAIL: if(1){put(1);} -> got '$$out', want '1'"; exit 1; fi; \
	echo "ok: if(1){put(1);} -> 1"
	@set -e; out=$$(printf 'if(0){put(2);}\n' | $(TARGET)); \
	if [ ! -z "$$out" ]; then echo "FAIL: if(0){put(2);} -> got '$$out', want empty"; exit 1; fi; \
	echo "ok: if(0){put(2);} -> empty"
	@set -e; out=$$(printf 'a=2*5;if(a){put(a);}\n' | $(TARGET)); \
	if [ "$$out" != "10" ]; then echo "FAIL: a=2*5;if(a){put(a);} -> got '$$out', want '10'"; exit 1; fi; \
	echo "ok: a=2*5;if(a){put(a);} -> 10"
	@set -e; out=$$(printf 'a=0;if(a){put(4);}\n' | $(TARGET)); \
	if [ ! -z "$$out" ]; then echo "FAIL: a=0;if(a){put(4);} -> got '$$out', want empty"; exit 1; fi; \
	echo "ok: a=0;if(a){put(4);} -> empty"
	@set -e; out=$$(printf 'put(1+2*3);put(12/3-2);put(10%%3);\n' | $(TARGET)); \
	expected=$$(printf '7\n2\n1'); \
	if [ "$$out" != "$$expected" ]; then echo "FAIL: put(1+2*3);put(12/3-2);put(10%3); -> got '$$out', want '7\n2\n1'"; exit 1; fi; \
	echo "ok: put(1+2*3);put(12/3-2);put(10%3); -> 7,2,1"
