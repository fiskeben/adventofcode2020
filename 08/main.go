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

	acc := run(lines)
	fmt.Printf("accumulator is %d\n", acc)

	acc = runFixed(lines)
	fmt.Printf("patched accumulator is %d\n", acc)
}

func run(lines []string) int {
	var acc int
	index := make([]bool, len(lines))
	var pointer int

	for {
		if index[pointer] {
			return acc
		}
		index[pointer] = true

		line := lines[pointer]
		args := strings.Split(line, " ")
		op := args[0]
		val, _ := strconv.Atoi(args[1])

		switch op {
		case "acc":
			acc += val
			pointer++
		case "jmp":
			pointer += val
		case "nop":
			pointer++
		}
	}
}

func runFixed(lines []string) int {
	var acc, pointer int
	index := make([]bool, len(lines))
	patches := make([]bool, len(lines))
	var patched bool

	for {
		if pointer == len(lines) {
			return acc
		}
		if index[pointer] {
			fmt.Println("infinite loop detected, restarting")
			acc = 0
			pointer = 0
			index = make([]bool, len(lines))
			patched = false
			continue
		}
		index[pointer] = true

		line := lines[pointer]
		args := strings.Split(line, " ")
		op := args[0]
		val, _ := strconv.Atoi(args[1])

		if !patched {
			switch op {
			case "jmp":
				if !patches[pointer] {
					op = "nop"
					patches[pointer] = true
					patched = true
				}
			case "nop":
				if !patches[pointer] {
					op = "jmp"
					patches[pointer] = true
					patched = true
				}
			}
		}

		switch op {
		case "acc":
			acc += val
			pointer++
		case "jmp":
			pointer += val
		case "nop":
			pointer++
		}
	}
}
