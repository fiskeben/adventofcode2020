package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")

	one := traverse(lines, 1, 1)
	two := traverse(lines, 3, 1)
	three := traverse(lines, 5, 1)
	four := traverse(lines, 7, 1)
	five := traverse(lines, 1, 2)

	trees := one * two * three * four * five

	fmt.Printf("trees: %d\n", trees)
}

func traverse(lines []string, right, down int) int {
	var trees, progress int
	first := true

	for i := 0; i < len(lines); i += down {
		l := lines[i]

		if strings.TrimSpace(l) == "" {
			continue
		}

		if first {
			first = false
			continue
		}

		progress += right

		if progress >= len(l) {
			progress = progress % len(l)
		}

		if l[progress] == byte('#') {
			trees++
		}
	}
	return trees
}
