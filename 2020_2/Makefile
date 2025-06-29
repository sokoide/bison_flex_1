# Calculator Makefile
TARGET = ./calc
CC = gcc
CFLAGS = -Wall -Wextra -O2
LIBS = -ly -lm

# Default target
all: $(TARGET)

# Build the calculator
$(TARGET): calc.tab.c calc.tab.h lex.yy.c
	$(CC) $(CFLAGS) -o $(TARGET) calc.tab.c $(LIBS)

# Generate parser from grammar
calc.tab.c calc.tab.h: calc.y
	bison -d -v -t calc.y

# Generate lexer from specification
lex.yy.c: calc.l calc.tab.h
	flex calc.l

# Clean generated files
clean:
	@echo "Cleaning generated files..."
	rm -f $(TARGET) *.tab.c *.tab.h lex.yy.c *.output

# Run the calculator
run: $(TARGET)
	@echo "Starting calculator..."
	$(TARGET)

# Debug build
debug: CFLAGS += -g -DDEBUG
debug: $(TARGET)

# Show help
help:
	@echo "Available targets:"
	@echo "  all     - Build the calculator (default)"
	@echo "  clean   - Remove generated files"
	@echo "  run     - Build and run the calculator"
	@echo "  debug   - Build with debug symbols"
	@echo "  help    - Show this help"

.PHONY: all clean run debug help