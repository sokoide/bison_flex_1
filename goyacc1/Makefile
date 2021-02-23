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