package palitrans

import (
	"testing"
)

func TestRomanToThai(t *testing.T) {
	var mapping = map[string]string{
		"ปา":        "pā",
		"ลิ":        "li",
		"ลี":        "lī",
		"ปาลิ":      "pāli",
		"สาธุ":      "sādhu",
		"ก":         "ka",
		"กิ":        "ki",
		"กุ":        "ku",
		"กา":        "kā",
		"กี":        "kī",
		"กู":        "kū",
		"เก":        "ke",
		"โก":        "ko",
		"กจ":        "kaca",
		"อิติ":      "iti",
		"เอว":       "eva",
		"อปิ":       "api",
		"อาชิ":      "āji",
		"อุป":       "upa",
		"กํ":        "kaṃ",
		"นรํ":       "naraṃ",
		"ภิกฺขุ":    "bhikkhu",
		"ทุกฺข":     "dukkha",
		"พุทฺเธน":   "buddhena",
		"อหํ":       "ahaṃ",
		"พุทฺธํ":    "buddhaṃ",
		"สรณํ":      "saraṇaṃ",
		"คจฺฉามิ":   "gacchāmi",
		"นโม":       "namo",
		"ทสฺส":      "dassa",
		"ตสฺส":      "tassa",
		"ภควโต":     "bhagavato",
		"ธมฺมานิ":   "dhammāni",
		"คนฺตฺวา":   "gantvā",
		"มฺหิ":      "mhi",
		"ยฺห":       "yha",
		"มยฺหํ":     "mayhaṃ",
		"อคฺคิมฺหิ": "aggimhi",
		"กึ":        "kiṃ",
		"พุทฺธ":     "buddha",
		"พุทฺธฺ":    "buddh",
		"":          "",
	}

	for k, v := range mapping {
		if k != RomanToThai(v) {
			t.Error(k, v)
		}
	}
}

func TestCharAt(t *testing.T) {
	s := CharAt("abc", -1)
	if s != "" {
		t.Error(s)
		return
	}

	s = CharAt("I am 字串", 5)
	if s != "字" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 0)
	if s != "ā" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 1)
	if s != "d" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 2)
	if s != "ā" {
		t.Error(s)
		return
	}

	s = CharAt("ādā", 3)
	if s != "" {
		t.Error(s)
		return
	}
}

func TestLength(t *testing.T) {
	if length("I am 字串") != 7 {
		t.Error("I am 字串")
		return
	}

	if length("ādā") != 3 {
		t.Error("ādā")
		return
	}

	if length("abc") != 3 {
		t.Error("abc")
		return
	}
}
