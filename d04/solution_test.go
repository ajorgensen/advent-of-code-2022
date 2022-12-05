package d02

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

func TestStack(t *testing.T) {
	s := NewStack()

	s.Push([]string{"a", "b", "c"})
	e := s.Pop(1)

	if len(e) != 1 {
		t.Fatalf("not the right number of elements [%+v]. Expected 1", strings.Join(e, ","))
	}

	if e[0] != "c" {
		t.Fatalf("wrong element. got %s expected %s", e[0], "c")
	}

	e = s.Pop(2)
	if len(e) != 2 {
		t.Fatalf("not the right number of elements [%+v]. Expected 1", strings.Join(e, ","))
	}

	if e[0] != "a" {
		t.Fatalf("wrong element. got %s expected %s", e[0], "a")
	}

	if e[1] != "b" {
		t.Fatalf("wrong element. got %s expected %s", e[1], "b")
	}
}

func TestPartOneExample(t *testing.T) {
	file, err := os.Open("testdata/example.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()


	result, err := PartOne(file)
	if err != nil {
		t.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartOneExample: %v\n", result)
	expected := "CMZ"
	if result != expected {
		t.Fatalf("expected %v got %v", expected, result)
	}
}

func TestPartOne(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	result, err := PartOne(file)
	if err != nil {
		t.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartOne: %v\n", result)
	expected := "ZWHVFWQWW"
	if result != expected {
		t.Fatalf("expected %v got %v", expected, result)
	}
}

func TestPartTwoExample(t *testing.T) {
	file, err := os.Open("testdata/example.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	result, err := PartTwo(file)
	if err != nil {
		t.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartTwoExample: %v\n", result)
	expected := "MCD"
	if result != expected {
		t.Fatalf("expected %v got %v", expected, result)
	}
}
func TestPartTwo(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	result, err := PartTwo(file)
	if err != nil {
		t.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartTwo: %v\n", result)
	expected := "HZFZCCWWV"
	if result != expected {
		t.Fatalf("expected %v got %v", expected, result)
	}
}