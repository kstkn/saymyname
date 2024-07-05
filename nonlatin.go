package saymyname

import (
	"strings"
)

var NonLatin = map[rune]string{
	// German Characters
	'Ä': "Ae", 'ä': "ae",
	'Ö': "Oe", 'ö': "oe",
	'Ü': "Ue", 'ü': "ue",
	'ß': "ss",

	// Nordic Characters
	'Å': "Aa", 'å': "aa",
	'Æ': "Ae", 'æ': "ae",
	'Ø': "Oe", 'ø': "oe",

	// French Characters
	'É': "E", 'é': "e",
	'È': "E", 'è': "e",
	'Ê': "E", 'ê': "e",
	'À': "A", 'à': "a",
	'Ç': "C", 'ç': "c",

	// Spanish Characters
	'Ñ': "N", 'ñ': "n",
	'Á': "A", 'á': "a",
	'Í': "I", 'í': "i",
	'Ó': "O", 'ó': "o",
	'Ú': "U", 'ú': "u",
	// 'Ü': "U", 'ü': "u",  // Already covered under German Characters

	// Turkish Characters
	// 'Ç': "C", 'ç': "c",  // Already covered under French Characters
	'Ğ': "G", 'ğ': "g",
	'İ': "I", 'i': "i",
	// 'Ö': "O", 'ö': "o",  // Already covered under German Characters
	'Ş': "S", 'ş': "s",
	// 'Ü': "U", 'ü': "u",  // Already covered under German Characters

	// Polish Characters
	'Ą': "A", 'ą': "a",
	'Ć': "C", 'ć': "c",
	'Ę': "E", 'ę': "e",
	'Ł': "L", 'ł': "l",
	'Ń': "N", 'ń': "n",
	'Ś': "S", 'ś': "s",
	'Ź': "Z", 'ź': "z",
	'Ż': "Z", 'ż': "z",

	// Czech Characters
	'Č': "C", 'č': "c",
	'Ď': "D", 'ď': "d",
	'Ě': "E", 'ě': "e",
	'Ň': "N", 'ň': "n",
	'Ř': "R", 'ř': "r",
	'Š': "S", 'š': "s",
	'Ť': "T", 'ť': "t",
	'Ů': "U", 'ů': "u",
	'Ž': "Z", 'ž': "z",
	'Ý': "Y", 'ý': "y",

	// Hungarian Characters
	// 'Á': "A", 'á': "a",  // Already covered under Spanish Characters
	// 'É': "E", 'é': "e",  // Already covered under French Characters
	// 'Í': "I", 'í': "i",  // Already covered under Spanish Characters
	// 'Ó': "O", 'ó': "o",  // Already covered under Spanish Characters
	// 'Ö': "O", 'ö': "o",  // Already covered under German Characters
	'Ő': "O", 'ő': "o",
	// 'Ú': "U", 'ú': "u",  // Already covered under Spanish Characters
	// 'Ü': "U", 'ü': "u",  // Already covered under German Characters
	'Ű': "U", 'ű': "u",

	// Pinyin Characters
	'Ā': "A", 'ā': "a", 'Ǎ': "A", 'ǎ': "a",
	'Ē': "E", 'ē': "e",
	'Ī': "I", 'ī': "i", 'Ǐ': "I", 'ǐ': "i", 'Ì': "I", 'ì': "i",
	'Ō': "O", 'ō': "o", 'Ǒ': "O", 'ǒ': "o", 'Ò': "O", 'ò': "o",
	'Ū': "U", 'ū': "u", 'Ǔ': "U", 'ǔ': "u", 'Ù': "U", 'ù': "u",
	'Ǖ': "U", 'ǖ': "u", 'Ǘ': "U", 'ǘ': "u", 'Ǚ': "U", 'ǚ': "u", 'Ǜ': "U", 'ǜ': "u",
	// 'á': "a", 'à': "a", 'é': "e", 'ě': "e", 'è': "e", 'í': "i", 'ó': "o", 'ú': "u", // Already covered under Spanish Characters
	// 'Á': "A", 'À': "A", 'É': "E", 'Ě': "E", 'È': "E", 'Í': "I", 'Ó': "O", 'Ú': "U", // Already covered under Spanish Characters
}

func Latin(in string) string {
	var out strings.Builder

	for _, char := range in {
		if substitute, found := NonLatin[char]; found {
			out.WriteString(substitute)
		} else {
			out.WriteRune(char)
		}
	}

	return out.String()
}
