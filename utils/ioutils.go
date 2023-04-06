package utils

import (
	"bufio"
	"io"
	"os"
	"strings"
)

type Reader interface {
	ReadString() string
	ReadRune() rune
}

type StdinReader struct {
	r io.Reader
}

func NewStdinReader() *StdinReader {
	return &StdinReader{
		r: os.Stdin,
	}
}

func (s *StdinReader) ReadString() string {
	reader := bufio.NewReader(s.r)
	text, _ := reader.ReadString('\n')
	text = strings.Replace(text, "\r\n", "", -1)
	text = strings.Replace(text, "\n", "", -1)
	return text
}

func (s *StdinReader) ReadRune() rune {
	reader := bufio.NewReader(s.r)
	char, _, _ := reader.ReadRune()
	return char
}
