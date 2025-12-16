package functions
// isValidArtChar checks if the character code is within the valid ASCII art range.
func isValidArtChar(n int) bool {
	return n >= 32 && n <= 126
}
