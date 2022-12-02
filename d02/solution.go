package d02

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

const (
	Rock int = 1
	Paper int = 2
	Scissors int = 3
)

func PartOne(r io.Reader, w io.Writer) error {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	lines := strings.Split(string(data), "\n")

	total := 0
	for _, l := range lines {
		strat := strings.Split(l, " ")

		total += scoreRound(strat[0], strat[1])
	}

	fmt.Fprintf(w, "Part One: %d\n", total)

	return nil
}

func PartTwo(r io.Reader, w io.Writer) error {
	rounds, err := parseInput(r)
	if err != nil {
		return err
	}

	total := 0
	for _, round := range rounds {
		theirs, result := round[0], round[1]

		// lose
		if result == "X" {
			// Rock
			if theirs == "A" {
				total += 3
			}

			// Paper
			if theirs == "B" {
				total += 1
			}

			// Scissors
			if theirs == "C" {
				total += 2
			}
		}

		//Draw
		if result == "Y" {
			// Rock
			if theirs == "A" {
				total += 1 + 3 
			}

			// Paper
			if theirs == "B" {
				total += 2 + 3
			}

			// Scissors
			if theirs == "C" {
				total += 3 + 3
			}
		}

		//Win
		if result == "Z" {
			// Rock
			if theirs == "A" {
				total += 2 + 6
			}

			// Paper
			if theirs == "B" {
				total += 3 + 6
			}

			// Scissors
			if theirs == "C" {
				total += 1 + 6
			}
		}
	}

	fmt.Fprintf(w, "PartTwo: %d\n", total)

	return nil
}

func parseInput(r io.Reader) ([][]string, error) {
	result := make([][]string, 0)

	data, err := ioutil.ReadAll(r)
	if err != nil {
		return result, err
	}

	lines := strings.Split(string(data), "\n")

	for _, l := range lines {
		strat := strings.Split(l, " ")
		result = append(result, []string{strat[0], strat[1]})
	}

	return result, err
}

// A/X: Rock
// B/Y: Paper
// C/Z: Scissors
func scoreRound(theirs, ours string) int {
	if theirs == "A" {
		if ours == "X" {
			return 3 + 1
		}

		if ours == "Y" {
			return 6 + 2
		}

		if ours == "Z" {
			return 3
		}
	}

	if theirs == "B" {
		if ours == "X" {
			return 1
		}

		if ours == "Y" {
			return 3 + 2
		}

		if ours == "Z" {
			return 6 + 3
		}
	}

	if theirs == "C" {
		if ours == "X" {
			return 6 + 1
		}

		if ours == "Y" {
			return 2
		}

		if ours == "Z" {
			return 3 + 3
		}
	}

	return 0
}