package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")
	var largestID, alsoLargest int
	seats := make([]bool, 100)

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}

		row, col := locateSeat(l)
		id := row*8 + col
		if id > largestID {
			largestID = id
		}

		binID := binaryLocate(l)
		if binID > alsoLargest {
			alsoLargest = binID
		}

		if len(seats) < id {
			t := make([]bool, id*2)
			copy(t, seats)
			seats = t
		}
		seats[id] = true
	}

	fmt.Printf("largest ID is %d or %d\n", largestID, alsoLargest)

	beginning := true
	for id, taken := range seats {
		if id >= len(seats)-2 {
			continue
		}

		if beginning && taken {
			beginning = false
		}

		if candidate := id + 1; !beginning && !seats[candidate] && seats[candidate+1] {
			fmt.Printf("my seat is %d\n", candidate)
		}
	}
}

func locateSeat(e string) (int, int) {
	row := divide(e[:7], space{0, 127}, 'F', 'B')
	col := divide(e[7:], space{0, 7}, 'L', 'R')
	return row, col
}

type space struct {
	from int
	to   int
}

func divide(e string, s space, front, back rune) int {
	for _, c := range e {
		diff := (s.to - s.from) / 2
		switch c {
		case front:
			s = space{s.from, s.from + diff}
		case back:
			s = space{s.from + diff + 1, s.to}
		}
	}
	return s.from
}

/*
Second solution
*/

func binaryLocate(s string) int {
	row := parseAsBinary(s[:7], "F", "B")
	col := parseAsBinary(s[7:], "L", "R")
	return row*8 + col
}

func parseAsBinary(a, low, high string) int {
	a = strings.ReplaceAll(a, low, "0")
	a = strings.ReplaceAll(a, high, "1")
	n, err := strconv.ParseInt(a, 2, 64)
	if err != nil {
		panic(err)
	}
	return int(n)
}
