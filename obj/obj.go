package obj

import (
	"fmt"
	"strings"
)

type Parser struct {
	LinesIgnored int
}

func Parse(s string) *Parser {
	linesIgnored := 0
	for _, line := range strings.Split(s, "\n") {
		fmt.Println(line)
		line = strings.Trim(line, "\n\t\r ")
		if line == "" {
			continue
		}

		switch {
		default:
			linesIgnored++
		}
	}

	return &Parser{
		LinesIgnored: linesIgnored,
	}
}
