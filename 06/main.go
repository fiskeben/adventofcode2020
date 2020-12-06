package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type group []rune

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")

	anyone := anyoneAnsweredYes(lines)
	everyone := everyoneAnsweredYes(lines)

	fmt.Printf("the sum of anyone answering yes is %d and everyone answering yes is %d\n", anyone, everyone)
}

func anyoneAnsweredYes(lines []string) int {
	var total int
	group := make(map[rune]struct{})

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			total += len(group)
			group = make(map[rune]struct{})
			continue
		}

		for _, r := range l {
			group[r] = struct{}{}
		}
	}
	return total
}

func everyoneAnsweredYes(lines []string) int {
	var total, groupLength int
	group := make(map[rune]int)

	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			for _, n := range group {
				if n == groupLength {
					total++
				}
			}

			group = make(map[rune]int)
			groupLength = 0
			continue
		}

		for _, r := range l {
			group[r]++
		}

		groupLength++
	}
	return total
}
