package mmio

import (
	"fmt"
	"os"
)

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
