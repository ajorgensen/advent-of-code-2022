package d02

import (
	"errors"
	"io"
	"sort"
	"strconv"
	"strings"
)

type Stack[K comparable] struct {
	elements []K
}

func NewStack[K comparable]() Stack[K] {
	return Stack[K]{
		elements: make([]K, 0),
	}
}

func (s *Stack[K]) Push(e K) {
	s.elements = append(s.elements, e)
}

func (s *Stack[K]) Pop() K {
	if len(s.elements) == 0 {
		var zero K
		return zero
	}

	e := s.elements[len(s.elements)-1]
	s.elements = s.elements[0 : len(s.elements)-1]

	return e
}

func PartOne(r io.Reader) (int, error) {
	lines, err := parseInput(r)
	if err != nil {
		return -1, err
	}
	pwd := NewStack[string]()
	dirSize := map[string]int{}

	// This assumes that ls is only called once
	for _, line := range lines {
		parts := strings.Split(line, " ")

		if parts[1] == "cd" {
			if parts[2] == ".." {
				pwd.Pop()
			} else {
				pwd.Push(parts[2])
			}
		}

		if parts[0] != "dir" && parts[0] != "$" {
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				return -1, err
			}
			// add the size to all the directors in the current path
			for i := range pwd.elements {
				d := strings.Join(pwd.elements[0:i+1], "/")
				dirSize[d] += size
			}
		}
	}

	totalSize := 0
	for _, v := range dirSize {
		if v <= 100000 {
			totalSize += v
		}
	}

	return totalSize, nil
}

func PartTwo(r io.Reader) (int, error) {
	lines, err := parseInput(r)
	if err != nil {
		return -1, err
	}
	pwd := NewStack[string]()
	dirSize := map[string]int{}

	// This assumes that ls is only called once
	for _, line := range lines {
		parts := strings.Split(line, " ")

		if parts[1] == "cd" {
			if parts[2] == ".." {
				pwd.Pop()
			} else {
				pwd.Push(parts[2])
			}
		}

		if parts[0] != "dir" && parts[0] != "$" {
			size, err := strconv.Atoi(parts[0])
			if err != nil {
				return -1, err
			}
			// add the size to all the directors in the current path
			for i := range pwd.elements {
				d := strings.Join(pwd.elements[0:i+1], "/")
				dirSize[d] += size
			}
		}
	}

	current := dirSize["/"]
	sizes := make([]int, 0, len(dirSize))

	for _, v := range dirSize {
		sizes = append(sizes, v)
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] < sizes[j]
	})

	for _, v := range sizes {
		available := (70000000 - current) + v
		if available >= 30000000 {
			return v, nil
		}
	}

	return -1, errors.New("something went wrong")
}

func parseInput(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(data), "\n"), nil
}
