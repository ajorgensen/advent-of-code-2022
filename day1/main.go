package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Elf struct {
	Number int
	Calories int
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(string(data), "\n")

	e := make([]Elf, 0)

	elfNum := 0
	currCal := 0
	for _, l := range(lines) {
		// new elf
		if l == "" {
			elfNum += 1
			e = append(e, Elf{Calories: currCal, Number: elfNum})
			currCal = 0
			continue
		}

		c, err := strconv.Atoi(l)
		if err != nil {
			log.Fatal(err)
		}

		currCal += c
	}

	sort.Slice(e, func(i, j int) bool {
		return e[i].Calories > e[j].Calories
	})

	fmt.Printf("Part 1: %d\n", e[0].Calories)

	total := e[0].Calories + e[1].Calories + e[2].Calories
	fmt.Printf("Part 2: %d\n", total)
}