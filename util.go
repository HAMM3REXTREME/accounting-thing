package main

func intAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
func padded(item string, padding int) string {
	// Takes a string and returns a padded string. The padding is at the front for negative
	// Pads with " " until it is a certain length
	var padThing string // Spaces needed to have a consistent cell size for each cell in a row.
	padLength := intAbs(padding) - len(item)
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
