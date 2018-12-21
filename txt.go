package mmio

import (
	"bufio"
	"log"
	"os"
)

// TXTwriter : general text writer
type TXTwriter struct {
	file   *os.File
	writer *bufio.Writer
}

// NewTXTwriter : TXTwriter constructor
func NewTXTwriter(fp string) *TXTwriter {
	file, err := os.Create(fp)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	nt := &TXTwriter{
		file:   file,
		writer: bufio.NewWriter(file),
	}
	return nt
}

// Close : closes TXTwriter
func (w *TXTwriter) Close() {
	w.writer.Flush()
	w.file.Close()
}

// WriteLine : general textfile line writer method for TXTwriter
func (w *TXTwriter) WriteLine(line string) error {
	_, err := w.writer.WriteString(line + "\n")
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
	return nil
}

// ReadTextLines reads and returns string lines from binary file
func ReadTextLines(fp string) []string {
	file, err := os.Open(fp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner, a := bufio.NewScanner(file), make([]string, 0)
	for scanner.Scan() {
		a = append(a, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return a
}
