package d02

import (
	"fmt"
	"io"
	"strconv"
	"strings"
)

type Knot struct {
	X int
	Y int
}

type Rope struct {
	knots []Knot
}

func (r *Rope) Tail() Knot {
	return r.knots[len(r.knots)-1]
}

func NewRope(length int) Rope {
	knots := make([]Knot, length)

	for i := 0; i < length; i += 1 {
		knots[i] = Knot{0, 0}
	}

	return Rope{
		knots: knots,
	}
}

func (r *Rope) isValid(head, tail Knot) bool {
	if abs(head.X-tail.X) > 1 || abs(head.Y-tail.Y) > 1 {
		return false
	}

	return true
}

func (r *Rope) Move(direction string) {
	// Move the head knot
	head := r.knots[0]
	switch direction {
	case "U":
		head.Y += 1
	case "D":
		head.Y -= 1
	case "L":
		head.X -= 1
	case "R":
		head.X += 1
	default:
	}
	r.knots[0] = head

	// Move the rest of the knots if needed
	for i := 1; i < len(r.knots); i += 1 {
		h, t := r.knots[i-1], r.knots[i]
		if r.isValid(h, t) {
			continue
		}

		if h.X-t.X > 0 {
			t.X += 1
		}

		if h.X-t.X < 0 {
			t.X -= 1
		}

		if h.Y-t.Y > 0 {
			t.Y += 1
		}

		if h.Y-t.Y < 0 {
			t.Y -= 1
		}

		r.knots[i] = t
	}
}

func (r *Rope) Draw(size int) string {
	center := size / 2

	grid := make([][]string, size)
	for i := 0; i < size; i += 1 {
		grid[i] = make([]string, size)
		for j := 0; j < size; j += 1 {
			grid[i][j] = "."
		}
	}

	for i := len(r.knots) - 1; i >= 0; i -= 1 {
		knot := r.knots[i]

		if i == 0 {
			grid[knot.Y+center][knot.X+center] = "H"
		} else {
			grid[knot.Y+center][knot.X+center] = fmt.Sprint(i)
		}
	}

	result := ""
	for i := len(grid) - 1; i >= 0; i -= 1 {
		for j := 0; j < len(grid[i]); j += 1 {
			result = fmt.Sprintf("%s%s", result, grid[i][j])
		}
		result = fmt.Sprintf("%s%s", result, "\n")
	}

	return result
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func PartOne(r io.Reader) (int, error) {
	lines, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	commands, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	locations := map[Knot]bool{}
	rope := NewRope(2)
	for _, cmd := range commands {
		for i := 0; i < cmd.Amount; i += 1 {
			rope.Move(cmd.Direction)
			locations[rope.Tail()] = true
		}
	}

	return len(locations), nil
}

func PartTwo(r io.Reader) (int, error) {
	lines, err := parseInput(r)
	if err != nil {
		return -1, err
	}

	cmds, err := parseLines(lines)
	if err != nil {
		return -1, err
	}

	locations := map[Knot]bool{}
	rope := NewRope(10)
	for _, cmd := range cmds {
		for i := 0; i < cmd.Amount; i += 1 {
			rope.Move(cmd.Direction)
			locations[rope.Tail()] = true
		}
	}

	return len(locations), nil
}

type cmd struct {
	Direction string
	Amount    int
}

func parseLines(lines []string) ([]cmd, error) {
	commands := make([]cmd, len(lines))

	for i, line := range lines {
		parts := strings.Split(line, " ")
		direction := parts[0]
		amount, err := strconv.Atoi(parts[1])
		if err != nil {
			return commands, err
		}

		commands[i] = cmd{direction, amount}
	}

	return commands, nil
}

func parseInput(r io.Reader) ([]string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return []string{}, err
	}

	return strings.Split(string(data), "\n"), nil
}
