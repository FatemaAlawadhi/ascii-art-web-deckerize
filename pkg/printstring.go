package pkg

func PrintString(Banner []string, TextArray []string) []string {
	var word []rune
	var LineBanner int
	var Line string
	var Text []string
	NewLine := 0
	ChrCount := 0
	for str := 0; str < len(TextArray); str++ {
		word = []rune(TextArray[str])
		ChrCount = len(word) + ChrCount + 1
		if ChrCount > 11 {
			Text = append(Text, "\n")
			ChrCount = 0
		}
		if len(word) > 0 {
			for Linenum := 1; Linenum <= 8; Linenum++ {
				for chr := 0; chr < len(word); chr++ {
					LineBanner = ((int(word[chr]) - 32) * 9) + (1 * Linenum)
					if LineBanner < 0 || LineBanner > 856 {
						Text = []string{"Wrong ", "Format"}
						return Text
					}
					Line = Line + Banner[LineBanner]
				}
				Text = append(Text, Line)
				Text = append(Text, "\n")
				Line = ""
			}

		} else if len(word) == 0 {
			NewLine++
			if str == (len(TextArray)-1) && NewLine == len(TextArray) {
				return Text
			}
			Text = append(Text, "\n")
		}
	}
	return Text
}
