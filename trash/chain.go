package main

import (
	"bytes"
	"fmt"
	"runtime"
)

type StringBuffer struct {
	buf  bytes.Buffer
	size int
	err  error
}

func (s *StringBuffer) Append(text string) *StringBuffer {
	if s.err != nil {
		return s // Nothing to do.
	}
	size, err := s.buf.WriteString(text)
	s.size += size
	s.err = err
	return s
}

func (s *StringBuffer) String() string {
	return s.buf.String()
}

func NewBuffer() StringBuffer {
	return StringBuffer{}
}

const LF = "\n"

var mem runtime.MemStats

func PrintMemory() {
	runtime.ReadMemStats(&mem)
	fmt.Println(mem.Alloc, mem.TotalAlloc, mem.HeapAlloc, mem.HeapSys)
}

func main() {
	PrintMemory()

	buf := NewBuffer()

	for i := 0; i < 25000; i++ {
		buf.Append("Hello").Append(", ").Append("world.")
		buf.Append(LF).Append("I have a ").Append("pen.")
	}
	// fmt.Println(buf.String(), buf.size, buf.err)
	fmt.Println(buf.size, buf.err)

	PrintMemory()
}
