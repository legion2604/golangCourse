package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type LineCounter struct {
	r     io.Reader
	lines int
}

func (l *LineCounter) Read(p []byte) (int, error) {
	n, err := l.r.Read(p)

	for _, b := range p[:n] {
		if b == '\n' {
			l.lines++
		}
	}

	// Если это конец файла и последняя строка не заканчивается '\n'
	if err == io.EOF && n > 0 {
		if p[n-1] != '\n' {
			l.lines++
		}
	}

	return n, err
}

func PrintLineCount(r io.Reader) {
	l := &LineCounter{r: r}
	_, _ = io.Copy(io.Discard, l)
	fmt.Println("Количество строк:", l.lines)
}

func main() {
	stringLine := strings.NewReader("dog\ncat\ncow\nchicken\n")
	stringLineCounter := &LineCounter{r: stringLine}
	PrintLineCount(stringLineCounter)

	fileLine, _ := os.Open("./problem.txt")
	fileLineCounter := &LineCounter{r: fileLine}
	PrintLineCount(fileLineCounter)
}

/*
	считывает количество строк из любого источника, реализующего интерфейс io.Reader
*/
