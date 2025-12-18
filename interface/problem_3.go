package main

import (
	"bytes"
	"os"
)

func PrintWithOutInterfaceForFile(p *os.File) {
	p.Write([]byte("Hello, World!"))
}
func PrintWithOutInterfaceForBuffer(p bytes.Buffer) {
	p.Write([]byte("Hello, World!"))
}

func main() {
	file, _ := os.Create("./problem.txt")
	defer file.Close()
	PrintWithOutInterfaceForFile(file)

	var bufferWriter bytes.Buffer
	PrintWithOutInterfaceForBuffer(bufferWriter)
}
