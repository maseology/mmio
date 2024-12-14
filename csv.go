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

// ReadCSV general CSV reader (must be completely numeric)
func ReadCSV(filepath string, nHeaderLines int) ([][]float64, error) {
	f, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("ReadCSV failed: %v\n", err)
		return nil, fmt.Errorf("ReadCSV failed: %v", err)
	}
	defer f.Close()
	var fout [][]float64
	recs := LoadCSV(io.Reader(f), nHeaderLines)
	// for rec := range LoadCSV(io.Reader(f), nHeaderLines) {
	for rec := range recs {
		f1 := make([]float64, 0, len(rec))
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

func ncolsCSV(rc io.Reader) int {
	r := csv.NewReader(rc)
	head, err := r.Read()
	if err != nil { //read header
		log.Fatal(err)
	}
	return len(head)
}

// LoadCSV  use: for rec := range LoadCSV(io.Reader(f)) {
func LoadCSV(rc io.Reader, nHeaderLines int) (ch chan []string) {
	ch = make(chan []string)

	go func() {
		r := csv.NewReader(rc)
		for l := 0; l < nHeaderLines; l++ {
			if _, err := r.Read(); err != nil { //read header(s)
				log.Fatalf("LoadCSV error: %v", err)
			}
		}
		defer close(ch)
		for {
			rec, err := r.Read()
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatalf("LoadCSV error: %v", err)
			}
			ch <- rec
		}
	}()
	return
}

func LoadCsvArray(fp string, nHeaderLines int) [][]string {
	a := make([][]string, 0)

	f, err := os.Open(fp)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	r := csv.NewReader(io.Reader(f))

	for l := 0; l < nHeaderLines; l++ {
		if _, err := r.Read(); err != nil { //read header(s)
			log.Fatalf("LoadCSV error: %v", err)
		}
		// if rec, err := r.Read(); err != nil { //read header(s)
		// 	log.Fatalf("LoadCSV error: %v", err)
		// } else {
		// 	fmt.Println(rec)
		// }
	}

	for {
		rec, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("LoadCSV error: %v", err)
		}
		a = append(a, rec)
	}
	return a
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

	n := len(data) // get line length
	for _, v := range data {
		switch v.(type) {
		case []float64:
			for range v.([]float64) {
				n++
			}
			n--
		case []string, []int, []float32:
			return fmt.Errorf("CSVwriter.WriteLine TODO: lists of type %T", v)
		default:
		}
	}

	a, ii := make([]string, n), 0
	for _, v := range data {
		switch v.(type) {
		case []float64:
			for _, vv := range v.([]float64) {
				a[ii] = fmt.Sprint(vv)
				ii++
			}
		default:
			a[ii] = fmt.Sprint(v)
			ii++
		}
	}

	if err := w.writer.Write(a); err != nil {
		return fmt.Errorf("CSVwriter.WriteLine error: %v", err)
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

// WriteCSV2d writes csv from a complete dataset dat[row][col]
func WriteCSV2d(fp, h string, dat [][]interface{}) {
	csv := NewCSVwriter(fp)
	defer csv.Close()
	csv.WriteHead(h)
	for _, ln := range dat {
		iv := make([]interface{}, len(ln))
		copy(iv, ln)
		// for i, v := range ln {
		// 	iv[i] = v
		// }
		csv.WriteLine(iv...)
	}
}

// WriteCSV writes csv from a complete dataset
func WriteCSV(fp, h string, d ...[]interface{}) {
	csv := NewCSVwriter(fp)
	defer csv.Close()
	csv.WriteHead(h)
	nc := len(d)
	nr := len(d[0])
	for i := 0; i < nr; i++ {
		iv := make([]interface{}, nc)
		notnil := false
		for j := 0; j < nc; j++ {
			if d[j][i] != nil {
				notnil = true
			}
			iv[j] = d[j][i]
		}
		if notnil {
			csv.WriteLine(iv...)
		}
	}
}
