package saymyname

import "strings"

type Name struct {
	First string
	Last  string
}

func ParseFullname(name string) Name {
	parts := strings.Split(strings.TrimSpace(name), " ")

	if a := affixIndex(parts); a > -1 && a < len(parts)-1 {
		parts[a] += parts[a+1]
		parts = append(parts[:a+1], parts[a+2:]...)
	}

	if a := suffixIndex(parts); a > -1 && a > 0 {
		parts[a-1] += " " + parts[a]
		parts = append(parts[:a], parts[a+1:]...)
	}

	if len(parts) == 1 {
		return Name{First: name}
	}

	if len(parts) >= 2 {
		return Name{
			First: parts[0],
			Last:  strings.Join(parts[1:], " "),
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

func suffixIndex(parts []string) int {
	for i, p := range parts {
		if _, ok := Suffixes[p]; ok {
			return i
		}
	}
	return -1
}
