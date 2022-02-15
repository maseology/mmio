package mmio

import (
	"fmt"
	"os"
	"strconv"
)

// ReadInts is a simple routine that reads an integer slice to an ascii file
func ReadInts(fp string) ([]int, error) {
	sa, err := ReadTextLines(fp)
	if err != nil {
		return nil, err
	}
	da := make([]int, len(sa))
	for i, ln := range sa {
		d, err := strconv.Atoi(ln)
		if err != nil {
			return nil, err
		}
		da[i] = d
	}
	return da, nil
}

// WriteInts is a simple routine that writes an integer slice to an ascii file
func WriteInts(fp string, d []int) error {
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	// f, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // append
	if err != nil {
		return err
	}
	for _, v := range d {
		if _, err := f.Write([]byte(fmt.Sprintf("%d\n", v))); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// ReadFloats is a simple routine that reads an float slice to an ascii file
func ReadFloats(fp string) ([]float64, error) {
	sa, err := ReadTextLines(fp)
	if err != nil {
		return nil, err
	}
	da := make([]float64, len(sa))
	for i, ln := range sa {
		d, err := strconv.ParseFloat(ln, 64)
		if err != nil {
			return nil, err
		}
		da[i] = d
	}
	return da, nil
}

// WriteFloats is a simple routine that writes an float slice to an ascii file
func WriteFloats(fp string, d []float64) error {
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	// f, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // append
	if err != nil {
		return err
	}
	for _, v := range d {
		if _, err := f.Write([]byte(fmt.Sprintf("%f\n", v))); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

// WriteStrings is a simple routine that writes a slice of strings to an ascii file
func WriteStrings(fp string, s []string) error {
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	// f, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // append
	if err != nil {
		return err
	}
	for _, v := range s {
		if _, err := f.Write([]byte(v + "\n")); err != nil {
			return err
		}
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}

func WriteString(fp, content string) error {
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_WRONLY, 0644)
	// f, err := os.OpenFile(fp, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644) // append
	if err != nil {
		return err
	}
	if _, err := f.Write([]byte(content)); err != nil {
		return err
	}
	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
