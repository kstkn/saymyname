package saymyname

var Affixes = map[string]struct{}{
	"A":    {},            // (Romanian) "son of"
	"Ab":   {},            // (Welsh, Cornish, Breton) "son of"
	"Af":   {},            // (Danish, Swedish) "of"
	"Av":   {},            // (Norwegian) "of"
	"Ap":   {},            // (Welsh) "son of"
	"Abu":  {},            // (Arabic) "father of";
	"Aït":  {},            // (Berber) "of"
	"Al":   {},            // (Arabic) "the"
	"Ālam": {},            // (Persian) "world"
	"At":   {}, "Ath": {}, // (Berber) "(son) of"
	"Aust":   {}, // (Norwegian) "east", "estern"
	"Austre": {},
	"Bar":    {},            // (Aramaic) "son of"
	"Bath":   {}, "Bat": {}, // (Hebrew) "daughter of"
	"Ben": {}, "Bin": {}, "Ibn": {}, // (Arabic and Hebrew) "son of"
	"Bet":   {},              // (Arabic from "Beyt") "house of"
	"Bint":  {},              // (Arabic) "daughter of"
	"Binti": {}, "Binte": {}, // (Malaysian) "daughter of"
	"D":     {},
	"Da":    {},              // (Italian) "from", "of"; (Portuguese) "from the" (before a feminine singular noun)
	"Das":   {},              // (Portuguese) "from the", "of the", preceding a feminine plural noun
	"De":    {},              // (Italian, French, Spanish, Portuguese, Filipino) "of"; indicates region of origin, often a sign of nobility; in Spanish-speaking countries a married woman will sometimes append her name with "de XXXX" where "XXXX" is her husband's last name; (Dutch) "the"
	"Degli": {},              // (Italian) "of the", preceding a masculine plural noun starting with either sp, sc, ps, z, gn, or st.
	"Del":   {},              // (Italian, Spanish) "of the", preceding a masculine singular noun
	"Dele":  {},              // Southern French, Spanish, Filipino, and Occitan, equivalent of Du
	"Della": {},              // (Italian) "of the", preceding a feminine singular noun
	"Der":   {},              // (Western Armenian) "son/daughter of a priest"; (German) "the" (masculine nominative), "of the" (feminine genitive)
	"Di":    {},              // (Italian) "son of"
	"Dos":   {},              // (Portuguese) "from the, of the", preceding a masculine plural noun
	"Du":    {},              // (French) "of the", preceding a masculine singular noun
	"E":     {},              // (Portuguese) "and", between surnames (Maria Eduarda de Canto e Mello)
	"El":    {},              // (Arabic and Spanish) "the"
	"Fetch": {}, "Vetch": {}, // (Welsh) "daughter of"
	"Fitz": {},            // (Irish, from Norman French) "son of", from Latin "filius" meaning "son" (mistakenly thought to mean illegitimate son, because of its use for certain illegitimate sons of English kings)
	"i":    {},            // (Catalan) "and", always in lowercase, used to identify both surnames (e.g. Antoni Gaudí i Cornet)
	"ka":   {},            // (Zulu) "(son/daughter) of", always in lower case and preceding the name of the father.
	"Kil":  {}, "Gil": {}, // (English, Irish, Scottish) "son of", "servant of", "devotee of", originating from the Irish "Mac Giolla", typically followed by a Saint's name (e.g. Mac Giolla Bhríde).
	"La":    {},                                                        // (Italian, French, Spanish) "the", feminine singular
	"Le":    {},                                                        // (Northern French) "the", masculine singular
	"Lille": {},                                                        // (Norwegian) "small", "little"
	"Lu":    {},                                                        // (Latin and Roman) "Master"
	"M'":    {}, "Mac": {}, "Mc": {}, "Mck": {}, "Mhic": {}, "Mic": {}, // (Irish, Scottish, and Manx Gaelic) "son". Both Mac and Mc are sometimes written Mac and Mc (with superscript ac or c). In some names, Mc is pronounced Mac.
	"Mala":   {},               // (Kurdish) "House of"
	"Mellom": {}, "Myljom": {}, // (Norwegian) "between"
	"Na":  {},              // (Thai) "at"
	"Ned": {}, "Nedre": {}, // (Norwegian) "low", "lower"
	"Neder": {},           // (Swedish) "lower", "under"
	"Ngā":   {},           // (Te Reo Māori) "the (plural)"
	"Nic":   {}, "Ní": {}, // (Irish, Scottish) "daughter of", from Irish "iníon" meaning "daughter"
	"Nin":  {},             // (Serbian)
	"Nord": {}, "Norr": {}, // (German, Swedish, Danish)
	"Nordre": {},                              // (Norwegian) "north", "northern"
	"Ny":     {},                              // (Swedish, Danish, Norwegian) "new"
	"O":      {}, "Ó": {}, "Ua": {}, "Uí": {}, // (Irish, Scottish, and Manx Gaelic) "son of", "grandson of", "descendant of"
	"Opp": {}, "Upp": {}, // (Norwegian) "up"
	"Öfver": {},              // (Swedish) "upper", "over" (archaic spelling)
	"Ost":   {},              // (German) "east", "eastern"
	"Öst":   {}, "Öster": {}, // (Swedish) "east", "eastern"
	"Øst":   {},                         // (Danish)  "east", "eastern"
	"Østre": {},                         // (Norwegian)  "east", "eastern"
	"Över":  {},                         // (Swedish) "upper", "over"
	"Øvste": {}, "Øvre": {}, "Øver": {}, // (Norwegian) "upper"
	"Öz":   {}, // (Turkish) "pure"
	"Pour": {}, // (Persian) "son of"
}
