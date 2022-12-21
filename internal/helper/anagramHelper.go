package helper

func StringToInt(w string) rune {
	var chucky rune = 0
	for _, letter := range w {
		chucky += letter
	}
	return chucky
}
