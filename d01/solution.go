package d01

import (
	_ "embed"
	"fmt"
	"io"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func PartOne(r io.Reader, w io.Writer) error {
	data, err := parseInput(r)
	if err != nil {
		return err
	}
	
	max := 0
	for _, elf := range data {
		t := 0
		for _, cal := range elf {
			t += cal
		}

		if t > max {
			max = t
		}
	}

	fmt.Fprintf(w, "PartOne: %d\n", max)

	return nil
}

func PartTwo(r io.Reader, w io.Writer) error {
	data, err := parseInput(r)
	if err != nil {
		return err
	}

	totals := make([]int, len(data))
	for _, elf := range data  {
		t := 0
		for _, cal := range elf {
			t += cal
		}
		totals = append(totals, t)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totals)))

	top3 := totals[0] + totals[1] + totals[2]

	fmt.Fprintf(w,  "PartTwo: %d\n", top3)

	return nil
}

func parseInput(r io.Reader) ([][]int, error) {
	input, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	data := make([][]int, 0)
	for _, elf := range strings.Split(string(input), "\n\n") {
		row := make([]int, 0)
		for _, cal := range strings.Split(elf, "\n") {
			if cal == "" {
				continue
			}

			i, err := strconv.Atoi(cal)
			if err != nil {
				return data, err
			}

			row = append(row, i)
		}

		data = append(data, row)
	}

	return data, nil
}

// func main() {
// 	lines := strings.Split(input, "\n")

// 	e := make([]Elf, 0)

// 	elfNum := 0
// 	curCal := 0
// 	for _, l := range(lines) {
// 		// new elf
// 		if l == "" {
// 			elfNum += 1
// 			e = append(e, Elf{Calories: curCal, Number: elfNum})
// 			curCal = 0
// 			continue
// 		}

// 		c, err := strconv.Atoi(l)
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		curCal += c
// 	}

// 	sort.Slice(e, func(i, j int) bool {
// 		return e[i].Calories > e[j].Calories
// 	})

// 	fmt.Printf("Part 1: %d\n", e[0].Calories)

// 	total := e[0].Calories + e[1].Calories + e[2].Calories
// 	fmt.Printf("Part 2: %d\n", total)
// }