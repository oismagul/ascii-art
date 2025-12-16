package functions

import (
	"os"
	"strings"
)
// LoadBanner loads the ASCII art banner from the specified file path.
func LoadBanner(path string) ([][]string, error) {
	var blocks [][]string
    var block []string

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.SplitSeq(string(data), "\n")
	for line := range lines {
		if line != "" {
			block = append(block, line)
		} else {
			blocks = append(blocks, block)
			block = []string{}
		}
	}

	if len(block) > 0 {
		blocks = append(blocks, block)
	}

	return blocks, nil
}
