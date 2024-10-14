TARGET=./hoge
SRCS=$(wildcard *.go)

all: $(TARGET)

$(TARGET): y.go $(SRCS)
	go build -o $(TARGET)

y.go: hoge.go.y
	goyacc $<

clean:
	rm $(TARGET) y.go y.output

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	echo "if(1){put(1);}" | $(TARGET)
	@echo "did you see 1?"
	echo "if(0){put(2);}" | $(TARGET)
	@echo "did you see empty?"
	echo "a=2*5;if(a){put(a);}" | $(TARGET)
	@echo "did you see 10?"
	echo "a=0;if(a){put(4);}" | $(TARGET)
	@echo "did you see empty?"
	echo "put(1+2*3);put(12/3-2);put(10%3);" | $(TARGET)
	@echo "did you see 7, 2, 1?"
