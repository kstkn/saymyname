package saymyname

import "strings"

type Name struct {
	First string
	Last  string
}

func ParseFullname(name string) Name {
	parts := strings.Split(strings.TrimSpace(name), " ")

	if a := affixIndex(parts); a > -1 {
		parts[a] += parts[a+1]
		parts = append(parts[:a+1], parts[a+2:]...)
	}

	if len(parts) == 1 {
		return Name{First: name}
	}

	if len(parts) == 2 {
		return Name{
			First: parts[0],
			Last:  parts[1],
		}
	}

	return Name{}
}

func affixIndex(parts []string) int {
	for i, p := range parts {
		if _, ok := Affixes[p]; ok {
			return i
		}
	}
	return -1
}
