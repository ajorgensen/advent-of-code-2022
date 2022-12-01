package d01

import (
	"log"
	"os"
	"testing"
)

func TestPartOne(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	if err := PartOne(file, os.Stdout); err != nil {
		log.Fatalf("could not solve: %v", err)
	}
}
func TestPartTwo(t *testing.T) {
	file, err := os.Open("testdata/input.txt")
	if err != nil {
		log.Fatalf("could not open test input: %v", err)
	}
	defer file.Close()

	if err := PartTwo(file, os.Stdout); err != nil {
		log.Fatalf("could not solve: %v", err)
	}
}