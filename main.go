package main

import (
	"fmt"
	"os"
	"strings"

	"art/functions"
)

func main() {
	args := os.Args[1:]

	// argument count handling
	if len(args) == 0 {
		fmt.Println("Аргументов нет\nNo arguments")
		return
	}

	if len(args) != 1 {
		fmt.Println("Слишком много аргументов не души\nToo many arguments")
		return
	}

	arg := strings.ReplaceAll(args[0], `\n`, "\n")
	if arg == "" {
		return
	}

	if arg == "\n" {
		fmt.Print("\n")
		return
	}

	blocks, err := functions.LoadBanner("standard.txt")
	// path error handling
	if err != nil {
		fmt.Println("Error loading banner file:", err)
		return
	}

	lines := strings.SplitSeq(arg, "\n")
	for line := range lines {
		if line == "" {
			fmt.Print("\n")
			continue
		} else {
			output, ok := functions.ArtConvert(line, blocks)
			// invalid character handling
			if !ok {
				fmt.Println("Error: invalid character in input string")
				return
			}
			fmt.Print(output)
		}
	}
}
