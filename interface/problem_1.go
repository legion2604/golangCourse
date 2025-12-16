package main

import (
	"bytes"
	"io"
	"os"
)

type WriteWrapper struct {
	write io.Writer
}

func (w *WriteWrapper) Write(p []byte) (n int, err error) {
	return w.write.Write(p)
}

func Print(p io.Writer) {
	p.Write([]byte("Hello, World!"))
}

func main() {
	file, _ := os.Create("./problem.txt")
	defer file.Close()
	fileWriter := &WriteWrapper{write: file}
	Print(fileWriter) // Записывает "Hello, World!" в файл problem.txt

	consoleWriter := &WriteWrapper{write: os.Stdout}
	Print(consoleWriter) // Выводит "Hello, World!" в консоль

	var bufferWriter bytes.Buffer
	bufferWrapper := &WriteWrapper{write: &bufferWriter}
	Print(bufferWrapper) // Записывает "Hello, World!" в буфер
}

/*
ну реально же сложно пздец, там чёто с интерфейсами, он типо должен реализовать интерфейс io.writer
крч понял немного как это работает, но не до конца
а тут просто запысиаваю в файл без логера, там уже логер изучить придется
*/
