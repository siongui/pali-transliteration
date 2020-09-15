package palitrans

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
