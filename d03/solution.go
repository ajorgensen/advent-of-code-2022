package d02

import (
	"fmt"
	"io"
	"strings"
)

var scores = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func PartOne(r io.Reader, w io.Writer) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	rucksacks := parseInput(string(data))

	total := 0
	for _, r := range rucksacks {
		left, right := r[0:len(r)/2], r[len(r)/2:]

		c := common(left, right)
		s := score(c[0])

		total += s
	}

	fmt.Fprintf(w, "PartOne: %d\n", total)

	return nil
}

func PartTwo(r io.Reader, w io.Writer) error {
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}

	rucksacks := parseInput(string(data))

	total := 0
	for i := 0; i < len(rucksacks); i+=3 {
		c := common(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		s := score(c[0])

		total += s
	}


	fmt.Fprintf(w, "PartTwo: %d\n", total)

	return nil
}

func score(a string) int {
	return strings.Index(scores, a) + 1
}

func common(slices ...[]string) []string {
	if len(slices) == 0 {
		return []string{}
	}

	common := make(map[string]bool)
	for _, l := range slices[0] {
		common[l] = true
	}

	for _, s := range slices[1:] {
		c := make(map[string]bool)
		for _, l := range s {
			if common[l] {
				c[l] = true
			}
		}

		common = c
	}

	keys := make([]string, 0, len(common))
	for k := range common {
		keys = append(keys, k)
	}

	return keys
}

func parseInput(data string) [][]string {
	d := strings.Split(data, "\n")

	r := make([][]string, len(d))
	for i, ruck := range d {
		items := strings.Split(ruck, "")
		r[i] = items
	}

	return r
}