package d02

import (
	"io"
	"strconv"
	"strings"
)

type Number interface {
	int | float32 | float64
}

func Min[K Number](s []K) K {
	min := s[0]

	for i := 1; i < len(s); i += 1 {
		if s[i] < min {
			min = s[i]
		}
	}

	return min
}

func Max[K Number](s []K) K {
	max := s[0]

	for i := 1; i < len(s); i += 1 {
		if s[i] > max {
			max = s[i]
		}
	}

	return max
}

func row(grid [][]int, n int) []int {
	return grid[n]
}

func col(grid [][]int, n int) []int {
	result := []int{}

	for _, row := range grid {
		for j, col := range row {
			if j == n {
				result = append(result, col)
			}
		}
	}

	return result
}

func PartOne(r io.Reader) (int, error) {
	grid, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	visible := 0
	for i, r := range grid {
		for j := range r {
			// Edges
			if i == 0 || i == len(grid)-1 || j == 0 || j == len(r)-1 {
				visible += 1
				continue
			}

			v := grid[i][j]

			upVis := true
			downVis := true
			leftVis := true
			rightVis := true

			// Check up
			for a := i - 1; a >= 0; a -= 1 {
				if grid[a][j] >= v {
					upVis = false
				}
			}

			// Check Down
			for a := i + 1; a < len(grid); a += 1 {
				if grid[a][j] >= v {
					downVis = false
				}
			}

			// Check Left
			for a := j - 1; a >= 0; a -= 1 {
				if grid[i][a] >= v {
					leftVis = false
				}
			}

			// Check Right
			for a := j + 1; a < len(grid[i]); a += 1 {
				if grid[i][a] >= v {
					rightVis = false
				}
			}

			if upVis || downVis || leftVis || rightVis {
				visible += 1
			}
		}
	}

	return visible, nil
}

func PartTwo(r io.Reader) (int, error) {
	grid, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	scores := make([]int, 0)
	for i := 1; i < len(grid)-1; i += 1 {
		for j := 1; j < len(grid[i])-1; j += 1 {
			v := grid[i][j]

			// Up
			upView := 0
			for a := i - 1; a >= 0; a -= 1 {
				upView += 1
				if grid[a][j] >= v {
					break
				}
			}

			// Down
			downView := 0
			for a := i + 1; a < len(grid); a += 1 {
				downView += 1
				if grid[a][j] >= v {
					break
				}
			}

			// Left
			leftView := 0
			for a := j - 1; a >= 0; a -= 1 {
				leftView += 1
				if grid[i][a] >= v {
					break
				}
			}

			// Right
			rightView := 0
			for a := j + 1; a < len(grid[i]); a += 1 {
				rightView += 1
				if grid[i][a] >= v {
					break
				}
			}

			scores = append(scores, upView*downView*leftView*rightView)
		}
	}

	maxScore := Max(scores)

	return maxScore, nil
}

func parseInput(r io.Reader) ([][]int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return [][]int{}, err
	}

	lines := strings.Split(string(data), "\n")

	result := make([][]int, len(lines))
	for i, line := range lines {

		cols := strings.Split(line, "")
		result[i] = make([]int, len(cols))

		for j, col := range cols {
			d, err := strconv.Atoi(col)
			if err != nil {
				return result, err
			}

			result[i][j] = d
		}
	}

	return result, nil
}
