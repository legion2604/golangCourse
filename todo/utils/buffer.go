package utils

import (
	"bufio"
	"os"
)

func Print() string {
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	text = text[:len(text)-1]

	return text
}
