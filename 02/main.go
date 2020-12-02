package main

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type password struct {
	min  int
	max  int
	char rune
	pass string
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	var downTheStreetValid int
	var officialValid int

	s := string(b)
	lines := strings.Split(s, "\n")

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}

		p, err := parsePassword(l)
		if err != nil {
			log.Fatal(err)
		}

		if downTheStreetValidation(p) {
			downTheStreetValid++
		}
		if officialValidation(p) {
			officialValid++
		}
	}

	log.Printf("%d valid passwords with old rules, %d with official rules", downTheStreetValid, officialValid)
}

// 1-3 a: abcde
func parsePassword(s string) (password, error) {
	var p password

	parts := strings.Split(s, " ")
	minmax := strings.Split(parts[0], "-")
	min, err := strconv.Atoi(minmax[0])
	if err != nil {
		return p, err
	}
	p.min = min

	max, err := strconv.Atoi(minmax[1])
	if err != nil {
		return p, err
	}
	p.max = max

	p.char = rune(parts[1][:len(parts[1])-1][0])
	p.pass = parts[2]

	return p, nil
}

func downTheStreetValidation(p password) bool {
	var chars int
	for _, r := range p.pass {
		if r == p.char {
			chars++
		}
		if chars > p.max {
			return false
		}
	}
	return chars >= p.min
}

func officialValidation(p password) bool {
	if len(p.pass) < p.max {
		return false
	}
	a := rune(p.pass[p.min-1])
	b := rune(p.pass[p.max-1])

	return (a == p.char && b != p.char) || (a != p.char && b == p.char)
}
