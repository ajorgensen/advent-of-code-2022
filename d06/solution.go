package d02

import (
	"errors"
	"io"

	"github.com/ajorgensen/advent-of-code-2022/set"
)

func PartOne(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return -1, err
	}

	for i := 0; i < len(data) - 4; i += 1 {
		e := []byte{data[i], data[i+1], data[i+2], data[i+3]}

		if len(e) == set.Of(data[i], data[i+1], data[i+2], data[i+3]).Len() {
			return i+4, nil
		}
	}

	return 0, nil
}

func PartTwo(r io.Reader) (int, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return -1, err
	}

	for i := 0; i < len(data) - 14; i += 1 {
		e := make([]byte, 14)

		for j := 0; j < 14; j += 1 {
			e[j] = data[i+j]
		}

		if len(e) == set.Of(e...).Len() {
			return i+14, nil
		}
	}

	return 0, errors.New("todo")
}