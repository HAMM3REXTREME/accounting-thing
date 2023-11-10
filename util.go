package main

import (
	"regexp"
)

func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func vStr(s string) string {
	// Filters \033 color codes used around here. Useful for finding lengths for padding
	// \033 codes really mess up formatting otherwise
	// Regular expression to match ANSI escape codes
	re := regexp.MustCompile(`\x1b\[[0-9;]*[a-zA-Z]`)

	// Remove ANSI escape codes from the string
	plainText := re.ReplaceAllString(s, "")

	return plainText
}

func padded(item string, padding int) string {
	// Takes a string and returns a padded string. The padding is at the front for negative
	// Pads with " " until it is a certain length
	var padThing string // Spaces needed to have a consistent cell size for each cell in a row.
	padLength := intAbs(padding) - len(vStr(item))
	if padLength < 0 {
		padLength = 0 // Ensure padding is non-negative
	}
	for s := 0; s < padLength; s++ {
		padThing += " " // deltaPadding = padLength * " "
	}

	if padding < 0 { // negative padding aligns to the right
		return padThing + item
	}
	return item + padThing

}
