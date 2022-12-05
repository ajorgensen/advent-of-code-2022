package d02

import (
	"io"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End int
}

func (r Range) Contains(other Range) bool {
	if r.Start <= other.Start && r.End >= other.End {
		return true
	}

	return false
}

func (r Range) Overlap(other Range) bool {
	if r.Contains(other) {
		return true
	}

	if r.End >= other.Start && r.End <= other.End {
		return true
	}

	if r.Start >= other.Start && r.Start <= other.End {
		return true
	}

	return false
}

func PartOne(r io.Reader) (int, error) {
	pairs, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	total := 0
	for _, p := range pairs {
		if p[0].Contains(p[1]) || p[1].Contains(p[0]) {
			total += 1
		}
	}

	return total, nil
}

func PartTwo(r io.Reader) (int, error) {
	pairs, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	total := 0
	for _, p := range pairs {
		if p[0].Overlap(p[1]) || p[1].Overlap(p[0]) {
			total += 1
		}
	}

	return total, nil
}

func parseInput(r io.Reader) ([][]Range, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return [][]Range{}, err
	}

	result := make([][]Range, 0)
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ",")

		leftRange := strings.Split(parts[0], "-")
		rightRange := strings.Split(parts[1], "-")

		r1, err := strconv.Atoi(leftRange[0])
		if err != nil {
			return result, err
		}

		r2, err := strconv.Atoi(leftRange[1])
		if err != nil {
			return result, err
		}

		r3, err := strconv.Atoi(rightRange[0])
		if err != nil {
			return result, err
		}

		r4, err := strconv.Atoi(rightRange[1])
		if err != nil {
			return result, err
		}

		result = append(result, []Range{ {r1, r2}, {r3, r4}})
	}

	return result, nil
}