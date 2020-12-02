package main

import (
	"io/ioutil"
	"log"
	"os"
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

	for n, a := range lines {
		if a == "" {
			continue
		}
		i, err := strconv.Atoi(a)
		if err != nil {
			log.Fatal(err)
		}

		for m, b := range lines[n:] {
			if b == "" {
				continue
			}
			j, err := strconv.Atoi(b)
			if err != nil {
				log.Fatal(err)
			}

			for _, c := range lines[m:] {
				if c == "" {
					continue
				}
				k, err := strconv.Atoi(c)
				if err != nil {
					log.Fatal(err)
				}

				if i+j+k == 2020 {
					log.Printf("Result: %d * %d * %d = %d\n", i, j, k, i*j*k)
					os.Exit(0)
				}
			}
		}
	}
}
