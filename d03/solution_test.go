package d02

import (
	"os"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	file, err := os.Open("testdata/example.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()


	if err := PartOne(file, os.Stdout); err != nil {
		t.Fatalf("could not solve: %v", err)
	}
}

func TestPartOne(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()


	if err := PartOne(file, os.Stdout); err != nil {
		t.Fatalf("could not solve: %v", err)
	}
}

func TestPartTwoExample(t *testing.T) {
	file, err := os.Open("testdata/example.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()


	if err := PartTwo(file, os.Stdout); err != nil {
		t.Fatalf("could not solve: %v", err)
	}
}
func TestPartTwo(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		t.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()


	if err := PartTwo(file, os.Stdout); err != nil {
		t.Fatalf("could not solve: %v", err)
	}
}