package saymyname

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

type Name struct {
	First      string
	Last       string
	Salutation string
}

func ParseFullname(name string) Name {
	name = strings.TrimSpace(name)
	if name == "" {
		return Name{}
	}

	parts := strings.Split(name, " ")

	if a := affixIndex(parts); a > -1 && a < len(parts)-1 {
		parts[a] += parts[a+1]
		parts = append(parts[:a+1], parts[a+2:]...)
	}

	if a := suffixIndex(parts); a > -1 && a > 0 {
		parts[a-1] += " " + parts[a]
		parts = append(parts[:a], parts[a+1:]...)
	}

	out := Name{}
	if a := titleIndex(parts); a > -1 && a > 0 {
		out.Salutation = parts[a]
		parts = append(parts[:a], parts[a+1:]...)
	}

	if len(parts) == 1 {
		out.First = name
		return out
	}

	if len(parts) > 1 {
		out.First = parts[0]
		out.Last = strings.Join(parts[1:], " ")
	}

	return out
}

func affixIndex(parts []string) int {
	for i, p := range parts {
		if _, ok := Affixes[strings.Title(p)]; ok {
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

func titleIndex(parts []string) int {
	for i, p := range parts {
		if _, ok := Titles[cases.Title(language.Und).String(p)]; ok {
			return i
		}
	}
	return -1
}
