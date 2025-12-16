package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CountLines(r io.Reader) (int, error) {
	reader := bufio.NewReader(r)
	count := 0

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) > 0 { // для последней строки без \n
				count++
			}
			return count, nil
		}
		if err != nil {
			fmt.Println(line)
			return count, err
		}
		count++
	}
}

func main() {
	a, _ := os.Open("./problem.txt")
	FileLines, _ := CountLines(a)
	fmt.Println("Файл:", FileLines)

	StringLines, _ := CountLines(strings.NewReader("one\ntwo\nthree"))
	fmt.Println("Строка:", StringLines)
}

/*
считывает количество строк из любого источника, реализующего интерфейс io.Reader
*/
