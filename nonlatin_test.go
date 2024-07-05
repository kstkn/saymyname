package saymyname

import (
	"testing"
)

func TestLatin(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{"X Æ A-Xii", "X Ae A-Xii"},
		{"Bećirović", "Becirovic"},
		{"Ångström", "Aangstroem"},
		{"Ægir", "Aegir"},
		{"Dæhlie", "Daehlie"},
		{"Bjørn", "Bjoern"},
		{"Schröder", "Schroeder"},
		{"Yüksel", "Yueksel"},
		{"Antonín Dvořák", "Antonin Dvorak"},
		{"Průša", "Prusa"},
		{"Ōu Chángkūn", "Ou Changkun"},
		{"Shàng Zōng", "Shang Zong"},
		{"Liú Kāng", "Liu Kang"},
		{"Steinstraßer", "Steinstrasser"},
		{"Élise", "Elise"},
		{"François", "Francois"},
		{"Muñoz", "Munoz"},
		{"García", "Garcia"},
		{"İstanbul", "Istanbul"},
		{"Çağlar", "Caglar"},
		{"Łódź", "Lodz"},
		{"Wiedźmin", "Wiedzmin"},
		{"Michał", "Michal"},
		{"Český", "Cesky"},
		{"Žižka", "Zizka"},
		{"Őrség", "Orseg"},
		{"Østerbye", "Oesterbye"},
		{"Şahin", "Sahin"},
	}

	for _, test := range tests {
		result := Latin(test.input)
		if result != test.expected {
			t.Errorf("Latin(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
