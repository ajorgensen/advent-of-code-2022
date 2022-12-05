package d02

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Move struct {
	Number int
	From int
	To int
}

func (m Move) String() string {
	return fmt.Sprintf("move %d from %d to %d", m.Number, m.From, m.To)
}

type Stack struct {
	elements []string
}

func NewStack() *Stack {
	return &Stack {
		elements: make([]string, 0),
	}
}

func (s *Stack) Push(e []string) {
	s.elements = append(s.elements, e...)
}

func (s *Stack) Pop(n int) []string {
	lower := len(s.elements) - n
	if lower < 0 {
		lower = 0
	}

	e := s.elements[lower : len(s.elements)]
	s.elements = s.elements[:lower]
	return e
}

func (s *Stack) Peek() string {
	if len(s.elements) == 0 {
		return ""
	}

	return s.elements[len(s.elements) - 1]
}

func (s *Stack) String() string {
	return fmt.Sprintf("[%s]", strings.Join(s.elements, ","))
}

func PartOne(r io.Reader) (string, error) {
	stacks, moves, err := parseInput(r)
	if err != nil {
		return "", err
	}

	for _, move := range moves {
		from := stacks[move.From - 1]
		to := stacks[move.To - 1]

		for i := 0; i < move.Number; i += 1 {
			e := from.Pop(1)
			to.Push(e)
		}
	}

	result := []string{}
	for _, s := range stacks {
		result = append(result, s.Peek())
	}

	return strings.Join(result, ""), nil
}

func PartTwo(r io.Reader) (string, error) {
	stacks, moves, err := parseInput(r)
	if err != nil {
		return "", err
	}

	for _, move := range moves {
		from := stacks[move.From - 1]
		to := stacks[move.To - 1]

		e := from.Pop(move.Number)
		to.Push(e)
	}

	result := []string{}
	for _, s := range stacks {
		result = append(result, s.Peek())
	}

	return strings.Join(result, ""), nil
}

func parseInput(r io.Reader) ([]*Stack, []Move, error) {
	var stacks []*Stack
	var moves []Move

	data, err := io.ReadAll(r)
	if err != nil {
		return stacks, moves, err
	}

	p := strings.Split(string(data), "\n\n")

	stacks = parseStacks(p[0])
	moves, err = parseMoves(p[1])
	if err != nil {
		return stacks, moves, err
	}

	return stacks, moves, nil
}

func parseMoves(data string) ([]Move, error) {
	lines := strings.Split(data, "\n")

	moves := make([]Move, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")

		number, err := strconv.Atoi(parts[1])
		if err != nil {
			return moves, err
		}

		from, err := strconv.Atoi(parts[3])
		if err != nil {
			return moves, err
		}

		to, err := strconv.Atoi(parts[5])
		if err != nil {
			return moves, err
		}

		moves[i] = Move{
			Number: number,
			To: to,
			From: from,
		}
	}

	return moves, nil
}

func parseStacks(data string) []*Stack {
	lines := strings.Split(data, "\n")
	for i, j := 0, len(lines) - 1; i < j; i, j = i+1, j-1 {
		lines[i], lines[j] = lines[j], lines[i]
	}

	n := strings.Split(lines[0], "   ")

	stacks := make([]*Stack, len(n))
	for i := 0; i < len(n); i += 1 {
		stacks[i] = NewStack()
	}

	for i := 1; i <= len(lines) - 1; i += 1 {
		parts := strings.Split(lines[i], "")

		container := 0
		for j := 1; j <= len(parts); j += 4 {
			if parts[j] != " " {
				stacks[container].Push([]string{parts[j]})
			}
			container += 1
		}
	}

	return stacks
}