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
			"",
			Name{},
		},
		{
			" ",
			Name{},
		},
		{
			"Jr",
			Name{First: "Jr"},
		},
		{
			"Dr",
			Name{First: "Dr"},
		},
		{
			"Mr",
			Name{First: "Mr"},
		},
		{
			"D",
			Name{First: "D"},
		},
		{
			"Chris",
			Name{First: "Chris"},
		},
		{
			"Clark Kent",
			Name{"Clark", "Kent", ""},
		},
		{
			"Clark O Connor",
			Name{"Clark", "OConnor", ""},
		},
		{
			"Clark D Agostino",
			Name{"Clark", "DAgostino", ""},
		},
		{
			"Robert Downey Jr.",
			Name{"Robert", "Downey Jr.", ""},
		},
		{
			"Downey Jr. Robert",
			Name{"Downey Jr.", "Robert", ""},
		},
		{
			"CASSISA FABIEN MR",
			Name{"CASSISA", "FABIEN", "MR"},
		},
		{
			"MORENO BUSTOS RAMON",
			Name{"MORENO", "BUSTOS RAMON", ""},
		},
		{
			"BACH CONNY MRS",
			Name{"BACH", "CONNY", "MRS"},
		},
		{
			"LANUZA/LIONEL P",
			Name{"LANUZA/LIONEL", "P", ""},
		},
		{
			"JACHOWICZ DELAUNEY JOANNA",
			Name{"JACHOWICZ", "DELAUNEY JOANNA", ""},
		},
		{
			"D OVIDIO FABRIZIO MR",
			Name{"DOVIDIO", "FABRIZIO", "MR"},
		},
		{
			"HERRERO PONS JOSEP MR",
			Name{"HERRERO", "PONS JOSEP", "MR"},
		},
		{
			"PATTON ELISE MRS",
			Name{"PATTON", "ELISE", "MRS"},
		},
		{},
	}
	for i, c := range cases {
		c := c
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := ParseFullname(c.in)
			assert.Equal(t, c.out, actual)
		})
	}
}
