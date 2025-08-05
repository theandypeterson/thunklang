package testprograms

import (
	"fmt"
	"interpreter/object"
	"interpreter/repl"
	"os"
	"testing"
)

func TestAddTwo(t *testing.T) {
	testProgramInt(t, "add-two.andy", int64(7))
}

func TestFib(t *testing.T) {
	testProgramInt(t, "fib.andy", int64(5702887))
}

func TestStrings(t *testing.T) {
	testProgramString(t, "strings.andy", "this is a string!")
}

func TestCaching(t *testing.T) {
	testProgramInt(t, "caching.andy", int64(2))
}

func testProgram(t *testing.T, filename string) object.Object {
	dat, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error reading file %s: %s\n", filename, err.Error())
	}

	evaluated, errors := repl.Run(string(dat))
	fmt.Println(string(dat))

	if len(errors) > 0 {
		t.Fatalf("program had errors: %s", errors)
	}

	return evaluated
}

func testProgramInt(t *testing.T, filename string, expectedValue int64) {
	evaluated := testProgram(t, filename)
	result := evaluated.(*object.Integer).Value
	if result != expectedValue {
		t.Fatalf("incorrect value, expected=%d, got=%d", expectedValue, result)
	}
}

func testProgramString(t *testing.T, filename string, expectedValue string) {
	evaluated := testProgram(t, filename)
	result := evaluated.(*object.String).Value
	if result != expectedValue {
		t.Fatalf("incorrect value, expected=%s, got=%s", expectedValue, result)
	}
}
