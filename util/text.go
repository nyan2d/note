package util

import "strings"

func GetTitle(content string) string {
	firstLine := strings.TrimSpace(content)
	if strings.ContainsRune(content, '\n') {
		firstLine = strings.TrimSpace(strings.SplitN(firstLine, "\n", 2)[0])
	}

	if strings.HasPrefix(firstLine, "# ") {
		firstLine = firstLine[2:]
	}

	if len(firstLine) == 0 {
		return "Unknown title"
	}

	if len(firstLine) <= 80 {
		return firstLine
	}

	return firstLine[:80]
}
