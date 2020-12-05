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
	var largestID int
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

		if len(seats) < id {
			t := make([]bool, id*2)
			copy(t, seats)
			seats = t
		}
		seats[id] = true
	}

	fmt.Printf("largest ID is %d\n", largestID)

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
