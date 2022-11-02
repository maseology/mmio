package mmio

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// ReadCsvDateFloat reads temporal csv file "date,value,flag,..."
func ReadCsvDateFloat(csvfp string) (map[int64]float64, error) {
	f, err := os.Open(csvfp)
	if err != nil {
		fmt.Printf("ReadCsvDateFloat failed: %v\n", err)
		return nil, fmt.Errorf("ReadCsvDateFloat failed: %v", err)
	}
	defer f.Close()

	recs := LoadCSV(io.Reader(f))
	o := make(map[int64]float64, len(recs)-1)
	for rec := range recs {
		t, err := time.Parse("2006-01-02", rec[0])
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05", rec[0])
			if err != nil {
				t, err = time.Parse("2006-01-02 15:04:05 +0000 UTC", rec[0])
				if err != nil {
					return nil, fmt.Errorf("date parse error in %s: %v", csvfp, err)
				}
			}
		}
		v, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			return nil, fmt.Errorf("value parse error: %v", err)
		}
		o[t.Unix()] = v
	}
	return o, nil
}

// ReadCsvDateFloat reads temporal csv file "date,value,flag,..."
func ReadCsvDateFloats(csvfp string) (map[time.Time][]float64, error) {
	f, err := os.Open(csvfp)
	if err != nil {
		fmt.Printf("ReadCsvDateFloats failed: %v\n", err)
		return nil, fmt.Errorf("ReadCsvDateFloats failed: %v", err)
	}
	defer f.Close()

	ncol := ncolsCSV(io.Reader(f))
	recs := LoadCSV(io.Reader(f))
	o := make(map[time.Time][]float64, len(recs)-1)
	for rec := range recs {
		t, err := time.Parse("2006-01-02", rec[0])
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05", rec[0])
			if err != nil {
				t, err = time.Parse("2006-01-02 15:04:05 +0000 UTC", rec[0])
				if err != nil {
					return nil, fmt.Errorf("date parse error in %s: %v", csvfp, err)
				}
			}
		}
		vs := make([]float64, ncol)
		for i := 1; i < ncol; i++ {
			vs[i], err = strconv.ParseFloat(rec[i], 64)
			if err != nil {
				return nil, fmt.Errorf("value parse error: %v", err)
			}
		}
		o[t] = vs
	}
	return o, nil
}

// ReadCsvStringInt reads temporal csv file ith column type "<str>,<int>"
func ReadCsvStringInt(csvfp string) (map[string]int, error) {
	f, err := os.Open(csvfp)
	if err != nil {
		fmt.Printf("ReadCSV failed: %v\n", err)
		return nil, fmt.Errorf("ReadCSV failed: %v", err)
	}
	defer f.Close()

	recs := LoadCSV(io.Reader(f))
	o := make(map[string]int, len(recs)-1)
	for rec := range recs {
		v, err := strconv.Atoi(rec[1])
		if err != nil {
			return nil, fmt.Errorf("value parse error: %v", err)
		}
		o[rec[0]] = v
	}
	return o, nil
}

// ReadCsvStringFloat reads temporal csv file ith column type "<str>,<float>"
func ReadCsvStringFloat(csvfp string) (map[string]float64, error) {
	f, err := os.Open(csvfp)
	if err != nil {
		fmt.Printf("ReadCSV failed: %v\n", err)
		return nil, fmt.Errorf("ReadCSV failed: %v", err)
	}
	defer f.Close()

	recs := LoadCSV(io.Reader(f))
	o := make(map[string]float64, len(recs)-1)
	for rec := range recs {
		v, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			return nil, fmt.Errorf("value parse error: %v", err)
		}
		o[rec[0]] = v
	}
	return o, nil
}
