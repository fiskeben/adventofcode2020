package main

import "testing"

func Test_IsValid(t *testing.T) {
	tests := []struct {
		name string
		arg  map[string]string
		want bool
	}{
		{
			name: "valid passport 1",
			arg: map[string]string{
				"pid": "087499704",
				"hgt": "74in",
				"ecl": "grn",
				"iyr": "2012",
				"eyr": "2030",
				"byr": "1980",
				"hcl": "#623a2f",
			},
			want: true,
		},
		{
			name: "valid passport 2",
			arg: map[string]string{
				"eyr": "2029",
				"ecl": "blu",
				"cid": "129",
				"byr": "1989",
				"iyr": "2014", "pid": "896056539", "hcl": "#a97842", "hgt": "165cm",
			},
			want: true,
		},
		{
			name: "valid passport 3",
			arg: map[string]string{
				"hcl": "#888785",
				"hgt": "164cm", "byr": "2001", "iyr": "2015", "cid": "88",
				"pid": "545766238", "ecl": "hzl",
				"eyr": "2022",
			},
			want: true,
		},
		{
			name: "invalid passport 1",
			arg: map[string]string{
				"eyr": "1972", "cid": "100",
				"hcl": "#18171d", "ecl": "amb", "hgt": "170", "pid": "186cm", "iyr": "2018", "byr": "1926",
			},
		},
		{
			name: "invalid passport 2",
			arg: map[string]string{
				"iyr": "2019",
				"hcl": "#602927", "eyr": "1967", "hgt": "170cm",
				"ecl": "grn", "pid": "012533040", "byr": "1946",
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := isValid(tt.arg); res != tt.want {
				t.Errorf("expected %t got %t", tt.want, res)
			}
		})
	}
}

func Test_ValidateHeight(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "valid height in inches",
			arg:  "60in",
			want: true,
		},
		{
			name: "valid height in cm",
			arg:  "160cm",
			want: true,
		},
		{
			name: "invalid height in inches",
			arg:  "40in",
		},
		{
			name: "invalid height in cm",
			arg:  "200cm",
		},
		{
			name: "invalid height",
			arg:  "40",
		},
		{
			name: "invalid height in text",
			arg:  "tall",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := validateHeight(tt.arg); res != tt.want {
				t.Errorf("expected %t got %t", tt.want, res)
			}
		})
	}
}

func Test_PassportIdRule(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "valid ID",
			arg:  "123456789",
			want: true,
		},
		{
			name: "invalid ID with letters",
			arg:  "1234512asf",
		},
		{
			name: "too long",
			arg:  "12345128435348584343875483",
		},
		{
			name: "too short",
			arg:  "123",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := passportIdRule.Match([]byte(tt.arg)); res != tt.want {
				t.Errorf("expected %t want %t", tt.want, res)
			}
		})
	}

}

func Test_hairColorRule(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "valid",
			arg:  "#cfa07d",
			want: true,
		},
		{
			name: "missing hash",
			arg:  "cfa07d",
		},
		{
			name: "too long",
			arg:  "#cfa07de",
		},
		{
			name: "too short",
			arg:  "#cfa07",
		},
		{
			name: "hash in wrong place",
			arg:  "cfa#07",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := hairColorRule.Match([]byte(tt.arg)); res != tt.want {
				t.Errorf("expected %t want %t", tt.want, res)
			}
		})
	}

}

func Test_validateEyeColor(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "amber",
			arg:  "amb",
			want: true,
		},
		{
			name: "blue",
			arg:  "blu",
			want: true,
		},
		{
			name: "brown",
			arg:  "brn",
			want: true,
		},
		{
			name: "gray",
			arg:  "gry",
			want: true,
		},
		{
			name: "green",
			arg:  "grn",
			want: true,
		},
		{
			name: "hazel",
			arg:  "hzl",
			want: true,
		},
		{
			name: "other",
			arg:  "oth",
			want: true,
		},
		{
			name: "empty",
			arg:  "",
		},
		{
			name: "whatever",
			arg:  "whatevs",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if res := validateEyeColor(tt.arg); res != tt.want {
				t.Errorf("expected %t got %t", tt.want, res)
			}
		})
	}
}
