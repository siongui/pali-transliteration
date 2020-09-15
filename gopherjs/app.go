package main

import (
	. "github.com/siongui/godom"
	"github.com/siongui/gopalilib/libfrontend/velthuis"
	palitrans "github.com/siongui/pali-transliteration"
)

func main() {
	input := Document.QuerySelector("#input")
	output := Document.QuerySelector("#output")

	// add pali input method to input text element
	// call velthuis before keyup event (order of keyevent handlers matters)
	velthuis.BindPaliInputMethodToInputTextElementById("input")

	input.AddEventListener("keyup", func(e Event) {
		r := input.Value()
		//println(r)
		t := palitrans.RomanToThai(r)
		//println(t)
		output.SetInnerHTML(t)
	})

	input.Focus()
}
