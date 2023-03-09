package saymyname

var Titles = map[string]struct{}{
	// English
	"Mr":   {}, // for men, regardless of marital status, who do not have another professional or academic title
	"Miss": {}, // for girls, unmarried women and (in the UK) married women who continue to use their maiden name
	"Mrs":  {}, // for married women who do not have another professional or academic title, an abbreviation of Mistress
	"Ms":   {}, // for women, regardless of marital status or when marital status is unknown
	"Mx":   {}, // a gender-neutral honorific for those who do not wish to specify their gender or who do not identify with Mr/Master or Ms/Mrs/Miss

	// German
	"Herr": {}, "Hr": {}, // for men (broadly equivalent to Mr., Lord and Sir in English)
	"Frau": {}, "Fr": {}, //
}
