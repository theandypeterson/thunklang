package testprograms

import (
	"fmt"
	"interpreter/object"
	"interpreter/repl"
	"os"
	"testing"
)

func TestAddTwo(t *testing.T) {
	testProgram(t, "add-two.andy", int64(7))
}

func TestFib(t *testing.T) {
	testProgram(t, "fib.andy", int64(5702887))
}

func testProgram(t *testing.T, filename string, expectedValue int64) {
	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filename, err.Error())
	}

	evaluated, errors := repl.Run(string(dat))
	fmt.Println(string(dat))

	if len(errors) > 0 {
		t.Fatalf("program had errors: %s", errors)
	}

	result := evaluated.(*object.Integer).Value
	if result != expectedValue {
		t.Fatalf("incorrect value, expected=%d, got=%d", expectedValue, result)
	}
}
