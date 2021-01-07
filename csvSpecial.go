package mmio

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

// ReadCsvDateValueFlag reads temporal csv file "date,value,flag"
func ReadCsvDateValueFlag(csvfp string) (map[time.Time]float64, error) {
	f, err := os.Open(csvfp)
	if err != nil {
		fmt.Printf("ReadCSV failed: %v\n", err)
		return nil, fmt.Errorf("ReadCSV failed: %v", err)
	}
	defer f.Close()

	recs := LoadCSV(io.Reader(f))
	o := make(map[time.Time]float64, len(recs)-1)
	for rec := range recs {
		t, err := time.Parse("2006-01-02", rec[0])
		if err != nil {
			t, err = time.Parse("2006-01-02T15:04:05", rec[0])
			if err != nil {
				return nil, fmt.Errorf("date parse error: %v", err)
			}
		}
		v, err := strconv.ParseFloat(rec[1], 64)
		if err != nil {
			return nil, fmt.Errorf("value parse error: %v", err)
		}
		o[t] = v
	}
	return o, nil
}
