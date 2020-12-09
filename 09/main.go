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

	var data []int
	for _, l := range lines {
		l = strings.TrimSpace(l)
		if l == "" {
			continue
		}

		n, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		data = append(data, n)
	}

	val := findNumber(data)
	fmt.Printf("the number is %d\n", val)

	hack := findContiguous(data, val)
	fmt.Printf("the hack is %d\n", hack)
}

func findNumber(data []int) int {
	var pointer int

	for {
		index := pointer + 25
		if pointer >= len(data) {
			panic("not found")
		}

		window := data[pointer:index]
		if v := data[index]; !valid(window, v) {
			return v
		}
		pointer++
	}
}

func valid(window []int, n int) bool {
	for x, i := range window {
		for _, j := range window[x+1:] {
			if i+j == n {
				return true
			}
		}
	}
	return false
}

func findContiguous(data []int, n int) int {
	start := 0
	end := 1
	sum := data[0]

	for {
		sum += data[end]
		if sum == n {
			var smallest, largest int
			for _, j := range data[start : end+1] {
				if j < smallest || smallest == 0 {
					smallest = j
				}
				if j > largest {
					largest = j
				}
			}
			return smallest + largest
		}
		if sum > n {
			// reset
			start++
			end = start + 1
			sum = data[start]
			continue
		}
		end++
	}
}
