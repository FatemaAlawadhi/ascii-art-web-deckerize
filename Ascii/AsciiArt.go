package Ascii

import (
	"ascii-art-web-dockerize/pkg"
	"strings"
)

func AsciiArt(Text string, Font string) (string, error) {
	Text = strings.ReplaceAll(Text, "\r\n", "\\n")
	Text = strings.ReplaceAll(Text, "\r", "\\n")
	TextArray := strings.Split(Text, "\\n")
	Banner := pkg.Banner(Font)
	FinalText := pkg.PrintString(Banner, TextArray)
	TextString := strings.Join(FinalText, "")

	return TextString, nil
}
