TARGET=./calc

all: $(TARGET)

$(TARGET): y.go
	go build -o $(TARGET)

y.go: calc.go.y
	goyacc calc.go.y

clean:
	rm $(TARGET) y.go

run: $(TARGET)
	$(TARGET)

test: $(TARGET)
	@echo "(1*2)+10/3" | $(TARGET)
	@echo "did you see 5?"
	@echo "(40+1)+2/2" | $(TARGET)
	@echo "did you see 42?"