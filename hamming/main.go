package hamming

import (
	"errors"
	"strings"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Hamming Length must be equal")
	} else {
		var h1 []string
		var h2 []string
		aReady := strings.ReplaceAll(a, " ", "")
		bReady := strings.ReplaceAll(b, " ", "")

		for _, char := range aReady {
			h1 = append(h1, string(char))
		}
		for _, char := range bReady {
			h2 = append(h2, string(char))
		}
		// GAGCCTACTAACGGGAT
		// CATCGTAATGACGGCCT
		// h1[0] = G h2[0] = C
		var diff int
		for i := 0; i < len(h1); i++ {
			if h1[i] != h2[i] {
				diff += 1
			}
		}
		return diff, nil

	}
}
