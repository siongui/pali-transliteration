// Package palitrans provides Pāli transliteration to/from
// Thai/Devanagari/Sinhalese/Burmese/... script.
package palitrans

import (
	"strings"
)

// RomanToThai converts Romanized Pāli to Thai script Pāli.
func RomanToThai(rm string) string {

	input := strings.Replace(strings.ToLower(rm), "ṁ", "ṃ", -1)

	vowel := make(map[string]string)
	vowel["a"] = "1"
	vowel["ā"] = "1"
	vowel["i"] = "1"
	vowel["ī"] = "1"
	vowel["iṃ"] = "1"
	vowel["u"] = "1"
	vowel["ū"] = "1"
	vowel["e"] = "2"
	vowel["o"] = "2"

	thair := make(map[string]string)
	thair["a"] = "อ"
	thair["ā"] = "า"
	thair["i"] = "\u0e34"
	thair["ī"] = "\u0e35"
	thair["iṃ"] = "\u0e36"
	thair["u"] = "\u0e38"
	thair["ū"] = "\u0e39"
	thair["e"] = "เ"
	thair["o"] = "โ"
	thair["ṃ"] = "\u0e4d"
	thair["k"] = "ก"
	thair["kh"] = "ข"
	thair["g"] = "ค"
	thair["gh"] = "ฆ"
	thair["ṅ"] = "ง"
	thair["c"] = "จ"
	thair["ch"] = "ฉ"
	thair["j"] = "ช"
	thair["jh"] = "ฌ"
	thair["ñ"] = "ญ"
	thair["ṭ"] = "ฏ"
	thair["ṭh"] = "ฐ"
	thair["ḍ"] = "ฑ"
	thair["ḍh"] = "ฒ"
	thair["ṇ"] = "ณ"
	thair["t"] = "ต"
	thair["th"] = "ถ"
	thair["d"] = "ท"
	thair["dh"] = "ธ"
	thair["n"] = "น"
	thair["p"] = "ป"
	thair["ph"] = "ผ"
	thair["b"] = "พ"
	thair["bh"] = "ภ"
	thair["m"] = "ม"
	thair["y"] = "ย"
	thair["r"] = "ร"
	thair["l"] = "ล"
	thair["ḷ"] = "ล"
	thair["v"] = "ว"
	thair["s"] = "ส"
	thair["h"] = "ห"

	cons := make(map[string]string)
	cons["k"] = "1"
	cons["g"] = "1"
	cons["ṅ"] = "1"
	cons["c"] = "1"
	cons["j"] = "1"
	cons["ñ"] = "1"
	cons["ṭ"] = "1"
	cons["ḍ"] = "1"
	cons["ṇ"] = "1"
	cons["t"] = "1"
	cons["d"] = "1"
	cons["n"] = "1"
	cons["p"] = "1"
	cons["b"] = "1"
	cons["m"] = "1"
	cons["y"] = "1"
	cons["r"] = "1"
	cons["l"] = "1"
	cons["ḷ"] = "1"
	cons["v"] = "1"
	cons["s"] = "1"
	cons["h"] = "1"

	var i0 = ""
	var i1 = ""
	var i2 = ""
	var i3 = ""
	var output = ""
	var i = 0

	input = strings.Replace(input, "&quot;", "`", -1)

	for i < length(input) {
		im := CharAt(input, i-2)
		i0 = CharAt(input, i-1)
		i1 = CharAt(input, i)
		i2 = CharAt(input, i+1)
		i3 = CharAt(input, i+2)

		if _, ok := vowel[i1]; ok {
			if i1 == "o" || i1 == "e" {
				output += thair[i1] + thair["a"]
				i++
			} else {
				if i == 0 {
					output += thair["a"]
				}
				if i1 == "i" && i2 == "ṃ" { // special i.m character
					output += thair[i1+i2]
					i++
				} else if i1 != "a" {
					output += thair[i1]
				}
				i++
			}
		} else if _, ok2 := thair[i1+i2]; ok2 && i2 == "h" { // two character match
			if i3 == "o" || i3 == "e" {
				output += thair[i3]
				i++
			}
			output += thair[i1+i2]
			if _, ok3 := cons[i3]; ok3 {
				output += "\u0e3a"
			}
			i = i + 2
		} else if _, ok4 := thair[i1]; ok4 && i1 != "a" { // one character match except a
			if i2 == "o" || i2 == "e" {
				output += thair[i2]
				i++
			}
			output += thair[i1]
			if _, ok5 := cons[i2]; ok5 && i1 != "ṃ" {
				output += "\u0e3a"
			}
			i++
		} else if _, ok6 := thair[i1]; !ok6 {
			output += i1
			_, ok7 := cons[i0]

			if _, ok71 := cons[im]; ok7 || (i0 == "h" && ok71) {
				output += "\u0e3a"
			}

			i++
			if i2 == "o" || i2 == "e" { // long vowel first
				output += thair[i2]
				i++
			}
			if _, ok8 := vowel[i2]; ok8 { // word-beginning vowel marker
				output += thair["a"]
			}
		} else { // a
			i++
		}
	}

	// Not sure if the following block is correct or not
	/*
		if _, ok9 := cons[i1]; ok9 {
			output += "\u0e3a"
		}
	*/

	//js: output = output.replace(/\`+/g, '"');
	return output
}

// CharAt is Go implementation of JavaScript charAt function.
// See https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/String/charAt
func CharAt(str string, index int) string {
	if index < 0 {
		return ""
	}

	i := 0
	for _, runeValue := range str {
		if i == index {
			return string(runeValue)
		}
		i++
	}

	return ""
}

func length(str string) int {
	i := 0
	for range str {
		i++
	}
	return i
}
