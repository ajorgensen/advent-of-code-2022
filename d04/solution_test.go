package d02

import (
	"fmt"
	"log"
	"os"
	"testing"
)

func TestPartOneExample(t *testing.T) {
	file, err := os.Open("testdata/example.txt")
	if err != nil {
		log.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()


	result, err := PartOne(file)
	if err != nil {
		log.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartOneExample: %d\n", result)
	expected := 2
	if result != expected {
		log.Fatalf("expected %d got %d", expected, result)
	}
}

func TestPartOne(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	result, err := PartOne(file)
	if err != nil {
		log.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartOne: %d\n", result)
	expected := 509
	if result != expected {
		log.Fatalf("expected %d got %d", expected, result)
	}
}

func TestPartTwoExample(t *testing.T) {
	file, err := os.Open("testdata/example.txt")
	if err != nil {
		log.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	result, err := PartTwo(file)
	if err != nil {
		log.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartTwoExample: %d\n", result)
	expected := 4
	if result != expected {
		log.Fatalf("expected %d got %d", expected, result)
	}
}
func TestPartTwo(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	result, err := PartTwo(file)
	if err != nil {
		log.Fatalf("could not solve: %v", err)
	}

	fmt.Printf("PartTwo: %d\n", result)
	expected := 870
	if result != expected {
		log.Fatalf("expected %d got %d", expected, result)
	}
}