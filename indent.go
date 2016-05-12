package html2rst

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func StringToLines(s string) []string {
	var lines []string

	scanner := bufio.NewScanner(strings.NewReader(s))
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	return lines
}

func indentEachLine(s string) string {
	lines := StringToLines(s)
	var indentedLines []string
	for _, line := range lines {
		indentedLines = append(indentedLines, "  "+line)
	}
	return strings.Join(indentedLines, "\n")
}
