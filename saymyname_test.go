package saymyname

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAffix(t *testing.T) {
	cases := []struct {
		in  string
		out Name
	}{
		{
			"Chris",
			Name{First: "Chris"},
		},
		{
			"Clark Kent",
			Name{"Clark", "Kent"},
		},
		{
			"Clark O Connor",
			Name{"Clark", "OConnor"},
		},
		{
			"Clark D Agostino",
			Name{"Clark", "DAgostino"},
		},
		{
			"Robert Downey Jr.",
			Name{"Robert", "Downey Jr."},
		},
		{
			"Downey Jr. Robert",
			Name{"Downey Jr.", "Robert"},
		},
	}
	for i, c := range cases {
		c := c
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := ParseFullname(c.in)
			assert.Equal(t, c.out, actual)
		})
	}
}
