package mmio

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// TXTwriter general text writer
type TXTwriter struct {
	file   *os.File
	writer *bufio.Writer
}

// NewTXTwriter constructor
func NewTXTwriter(fp string) (*TXTwriter, error) {
	file, err := os.Create(fp)
	if err != nil {
		return nil, fmt.Errorf("Cannot create file: %v", err)
	}
	nt := &TXTwriter{
		file:   file,
		writer: bufio.NewWriter(file),
	}
	return nt, nil
}

// Close closes TXTwriter
func (w *TXTwriter) Close() {
	w.writer.Flush()
	w.file.Close()
}

// Write is a general textfile writer method for TXTwriter
func (w *TXTwriter) Write(s string) error {
	_, err := w.writer.WriteString(s)
	if err != nil {
		return fmt.Errorf("Cannot write to file: %v", err)
	}
	return nil
}

// WriteLine is a general textfile line writer method for TXTwriter
func (w *TXTwriter) WriteLine(line string) error {
	_, err := w.writer.WriteString(line + "\n")
	if err != nil {
		return fmt.Errorf("Cannot write line to file: %v", err)
	}
	return nil
}

// WriteBytes general textfile line writer method for TXTwriter
func (w *TXTwriter) WriteBytes(b []byte) error {
	_, err := w.writer.Write(b)
	if err != nil {
		return fmt.Errorf("Cannot write to file: %v", err)
	}
	return nil
}

// ReadTextLines reads and returns string lines from binary file
func ReadTextLines(fp string) ([]string, error) {
	file, err := os.Open(fp)
	if err != nil {
		return nil, fmt.Errorf("ReadTextLines: %v", err)
	}
	defer file.Close()

	reader, a := bufio.NewReader(file), make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		// line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, fmt.Errorf("ReadTextLines: %v", err)
		}
		// a = append(a, string(line))
		a = append(a, line[:len(line)-1])
	}

	// scanner, a := bufio.NewScanner(file), make([]string, 0)
	// for scanner.Scan() {
	// 	a = append(a, scanner.Text())
	// }
	// if err := scanner.Err(); err != nil {
	// 	return nil, fmt.Errorf("ReadTextLines: %v", err)
	// }

	return a, nil
}

// RemoveBOM detects byte order mark (BOM) and moves ahead if exists
func RemoveBOM(r *bufio.Reader) error {
	rn, _, err := r.ReadRune()
	if err != nil {
		return err
	}
	if rn != '\uFEFF' {
		r.UnreadRune() // Not a BOM -- put the rune back
	}
	return nil
}
