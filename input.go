package main

import (
	"os"
	"syscall"
)

type Coord struct {
	filename string
	ppline   int
	line     int
	col      int
}

const EOF = 255

type Input struct {
	filename string
	base     []byte
	cursor   int
	lineHead int
	line     int
	file     *os.File
	size     int
}

var input Input

func (i *Input) Cursor() int {
	return i.cursor
}

func (i *Input) Move() *Input {
	i.cursor++
	return i
}

func (i *Input) Peek() byte {
	return i.base[i.cursor+1]
}

func (i *Input) Peek2() byte {
	return i.base[i.cursor+2]
}

func (i *Input) Char() byte {
	return i.base[i.cursor]
}

func (i *Input) Token(start int, len int) []byte {
	return i.base[start:start+len]
}

func (i*Input) TokenTo(start int,stop int)[]byte  {
	return i.base[start:stop]
}

func (*Input) ReadSourceFile(filename string) {
	srcFile, err := os.OpenFile(filename, os.O_RDONLY, os.ModePerm)
	if err != nil {
		Fatal("Can't open file: %s:%s", filename, err)
	}

	fileInfo, err := os.Stat(filename)
	if err != nil {
		Fatal("can't stat file: %s,error:%s", filename, err)
	}

	data, err := syscall.Mmap(int(srcFile.Fd()), 0, int(fileInfo.Size())+1, syscall.PROT_WRITE, syscall.MAP_PRIVATE)
	if err != nil {
		Fatal("%s mmap error:%s", filename, err)
	}
	input.filename = filename
	input.file = srcFile
	input.base = data
	input.cursor = 0
	input.lineHead = 0
	input.line = 1
	input.size = int(fileInfo.Size())
	input.base[input.size] = EndOfFile

}

func (i *Input) CloseSourceFile() {
	syscall.Munmap(i.base)
	i.file.Close()
}
