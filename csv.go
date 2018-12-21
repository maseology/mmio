package mmio

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// ReadCSV general CSV reader
func ReadCSV(filepath string) ([][]float64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("ReadCSV failed: %v\n", err)
		return nil, fmt.Errorf("ReadCSV failed: %v", err)
	}
	defer f.Close()
	var fout [][]float64
	for rec := range processCSV(io.Reader(f)) {
		var f1 []float64
		for i, c := range rec {
			f2, err := strconv.ParseFloat(c, 64)
			if err != nil {
				fmt.Printf("ReadCSV failed: rec[%v]: %v; error: %v\n", i, rec, err)
				return nil, fmt.Errorf("ReadCSV failed: rec[%v]: %v; error: %v", i, rec, err)
			}
			f1 = append(f1, f2)
		}
		fout = append(fout, f1)
	}
	return fout, err
}

func processCSV(rc io.Reader) (ch chan []string) {
	ch = make(chan []string, 10)
	go func() {
		r := csv.NewReader(rc)
		if _, err := r.Read(); err != nil { //read header
			log.Fatal(err)
		}
		defer close(ch)
		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			ch <- rec
		}
	}()
	return
}

// CSVwriter general CSV writer
type CSVwriter struct {
	file   *os.File
	writer *csv.Writer
}

// NewCSVwriter CSVwriter constructor
func NewCSVwriter(fp string) *CSVwriter {
	file, err := os.Create(fp)
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	nc := &CSVwriter{
		file:   file,
		writer: csv.NewWriter(file),
	}
	return nc
}

// Close closes CSVwriter
func (w *CSVwriter) Close() {
	w.writer.Flush()
	w.file.Close()
}

// WriteLine general CSV line writer method for CSVwriter
func (w *CSVwriter) WriteLine(data ...interface{}) error {
	var err error
	a := make([]string, len(data))
	for i, v := range data {
		a[i] = fmt.Sprint(v)
	}
	err = w.writer.Write(a)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
	return nil
}

// WriteHead add header row to CSVwriter
func (w *CSVwriter) WriteHead(h string) error {
	var err error
	a := strings.Split(h, ",")
	err = w.writer.Write(a)
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
	return nil
}

// WriteCSV writes csv from a complete dataset
func WriteCSV(fp, h string, dat [][]float64) {
	csv := NewCSVwriter(fp)
	defer csv.Close()
	csv.WriteHead(h)
	for _, ln := range dat {
		iv := make([]interface{}, len(ln))
		for i, v := range ln {
			iv[i] = v
		}
		csv.WriteLine(iv...)
	}
}
