TARGET=./hoge
SRCS=main.go

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
	echo "(1 * 2) +6/3;" | $(TARGET)
	@echo "did you see 4?"
	echo "(4 + 1) + 2/2;" | $(TARGET)
	@echo "did you see 6?"
	echo "(40 + 1) + 2/2;" | $(TARGET)
	@echo "did you see 42?"
	echo "1 + 10*20*30;" | $(TARGET)
	@echo "did you see 6001?"
	echo "1+2;3*4;55+66;" | $(TARGET)
	@echo "did you see 3,12,121?"
	echo "34%7+4;" | $(TARGET)
	@echo "did you see 10?"
	echo "3*(1+2);" | $(TARGET)
	@echo "did you see 9?"
