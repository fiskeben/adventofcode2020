package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var hairColorRule = regexp.MustCompile("^#[a-f0-9]{6}$")
var passportIdRule = regexp.MustCompile("^[0-9]{9}$")

func main() {
	b, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := string(b)
	lines := strings.Split(s, "\n")

	m := make(map[string]string)
	var valid int

	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			err := isValid(m)
			if err == nil {
				valid++
			}

			m = make(map[string]string)
			continue
		}

		parts := strings.Split(l, " ")
		for _, p := range parts {
			keyval := strings.Split(p, ":")
			m[keyval[0]] = keyval[1]
		}
	}

	fmt.Printf("valid passports: %d\n", valid)
}

func isValid(m map[string]string) error {
	if byr := m["byr"]; !validateYear(byr, 1920, 2002) {
		return fmt.Errorf("bad byr '%s'", byr)
	}

	if iyr := m["iyr"]; !validateYear(iyr, 2010, 2020) {
		return fmt.Errorf("bad iyr %s", iyr)
	}

	if eyr := m["eyr"]; !validateYear(eyr, 2020, 2030) {
		return fmt.Errorf("bad eyr %s", eyr)
	}

	if h := m["hgt"]; !validateHeight(h) {
		return fmt.Errorf("bad hgt '%s'", h)
	}

	if hcl := m["hcl"]; !hairColorRule.Match([]byte(hcl)) {
		return fmt.Errorf("bad hcl '%s'", hcl)
	}

	if ecl := m["ecl"]; ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		return fmt.Errorf("bad ecl '%s'", ecl)
	}

	if pid := m["pid"]; !passportIdRule.Match([]byte(pid)) {
		return fmt.Errorf("bad pid '%s'")
	}

	return nil
}

func validateYear(in string, min, max int) bool {
	yr, err := strconv.Atoi(in)
	if err != nil {
		return false
	}
	return yr >= min && yr <= max
}

func validateHeight(h string) bool {
	if i := strings.Index(h, "cm"); i > -1 {
		val, err := strconv.Atoi(h[:i])
		if err != nil {
			return false
		}
		return val >= 150 && val <= 193
	} else if i := strings.Index(h, "in"); i > -1 {
		val, err := strconv.Atoi(h[:i])
		if err != nil {
			return false
		}
		return val >= 59 && val <= 76
	}
	return false
}

func validateEyeColor(ecl string) bool {
	return ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth"
}
