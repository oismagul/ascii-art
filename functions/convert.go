package functions

import (
	"strings"
)

const asciiOffset = 31
// ArtConvert converts the input string into ASCII art using the provided blocks.
func ArtConvert(s string, blocks [][]string) (string, bool) {
	var sb strings.Builder

	for i := range 8 {
		for _, v := range s {
			index := int(v) - asciiOffset
			if !isValidArtChar(int(v)) {
				return "", false
			}
			if index < 0 || index >= len(blocks) {
				continue
			}
			line := blocks[index][i]
			sb.WriteString(line)
		}
		sb.WriteString("\n")
	}

	return sb.String(), true
}
