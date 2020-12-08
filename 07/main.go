package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Bag struct {
	Name  string
	Rules []Rule
}

type Rule struct {
	Weight   int
	Contents Bag
}

type candidate struct {
	name  string
	index int
}

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")
	bagsContainingMyBag := findBags(lines)

	fmt.Printf("%d bags can eventually contain my bag\n", bagsContainingMyBag)

	bagsInMyBag := countBags(lines)

	fmt.Printf("my bag contains %d other bags\n", bagsInMyBag)
}

func findBags(lines []string) int {
	kinds := []string{"shiny gold"}
	var find string
	index := make([]bool, len(lines))

	for {
		if len(kinds) == 0 {
			break
		}

		find, kinds = kinds[0], kinds[1:]
		canContain := contains(lines, find)
		for _, c := range canContain {
			index[c.index] = true
			kinds = append(kinds, c.name)
		}
	}

	var counter int
	for _, b := range index {
		if b {
			counter++
		}
	}

	return counter
}

func countBags(lines []string) int {
	b, _ := findNamedBag(lines, "shiny gold")
	return countBagsRecursive(lines, b)
}

func countBagsRecursive(lines []string, b Bag) int {
	var total int
	for _, r := range b.Rules {
		c, ok := findNamedBag(lines, r.Contents.Name)
		if !ok {
			total += r.Weight
			continue
		}
		subtotal := countBagsRecursive(lines, c)
		total += r.Weight*subtotal + r.Weight
	}

	return total
}

func findNamedBag(lines []string, kind string) (Bag, bool) {
	var b Bag
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		b = parseRule(l)
		if b.Name == kind {
			return b, true
		}
	}
	return b, false
}

func contains(lines []string, kind string) []candidate {
	var res []candidate

	for i, l := range lines {
		l = strings.TrimSpace(l)

		if l == "" {
			continue
		}

		b := parseRule(l)

		for _, x := range b.Rules {
			if x.Contents.Name == kind {
				res = append(res, candidate{name: b.Name, index: i})
			}
		}
	}

	return res
}

func parseRule(rule string) Bag {
	key := "contains"
	if strings.Index(rule, key) == -1 {
		key = "contain"
	}

	parts := strings.Split(rule, key)

	name := extractName(parts[0])
	b := Bag{Name: name}

	rules := strings.Split(parts[1], ",")
	b.Rules = make([]Rule, len(rules))

	for i, r := range rules {
		r = strings.TrimSpace(r)
		if strings.Index(r, "no other bags") > -1 {
			continue
		}

		rest := strings.Split(r, " ")
		weight, err := strconv.Atoi(rest[0])
		if err != nil {
			panic(err)
		}
		containsName := strings.Join(rest[1:3], " ")

		rr := Rule{Weight: weight, Contents: Bag{Name: containsName}}
		b.Rules[i] = rr
	}

	return b
}

func extractName(s string) string {
	parts := strings.Split(s, " ")
	return strings.Join(parts[:2], " ")
}
