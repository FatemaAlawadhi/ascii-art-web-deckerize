package pkg

import (
	"bufio"
	"os"
)

func Banner(FileType string) []string {
	var Banner []string
	if FileType == "standard" {
		FileType = "Banner/standard.txt"
	} else if FileType == "shadow" {
		FileType = "Banner/shadow.txt"
	} else if FileType == "thinkertoy" {
		FileType = "Banner/thinkertoy.txt"
	}
	
	file, err := os.Open(FileType)
		if err != nil {
			return nil
		}
	defer file.Close()
	// Create a scanner for the file
	scanner := bufio.NewScanner(file)
	// Read each line of the file
	for scanner.Scan() {
		Banner = append(Banner, scanner.Text())
	}

	return Banner
}