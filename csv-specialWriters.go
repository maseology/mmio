package mmio

import "time"

func WriteCsvFloats(csvfp, header string, d ...[]float64) error {
	csv := NewCSVwriter(csvfp)
	defer csv.Close()
	if err := csv.WriteHead(header); err != nil {
		return err
	}
	nc := len(d)
	nr := len(d[0])
	for i := 0; i < nr; i++ {
		iv := make([]interface{}, nc)
		for j := 0; j < nc; j++ {
			iv[j] = d[j][i]
		}
		if err := csv.WriteLine(iv...); err != nil {
			return err
		}
	}
	return nil
}

func WriteCsvDateFloats(csvfp, header string, t []time.Time, d ...[]float64) error {
	csv := NewCSVwriter(csvfp)
	defer csv.Close()
	if err := csv.WriteHead("date," + header); err != nil {
		return err
	}
	nc := len(d)
	nr := len(d[0])
	for i := 0; i < nr; i++ {
		iv := make([]interface{}, nc+1)
		iv[0] = t[i].Format("2006-01-02 15:04:05")
		for j := 0; j < nc; j++ {
			iv[j+1] = d[j][i]
		}
		if err := csv.WriteLine(iv...); err != nil {
			return err
		}
	}
	return nil
}

func WriteCsvIntInts(csvfp, header string, ii map[int]int) error {
	csv := NewCSVwriter(csvfp)
	defer csv.Close()
	if err := csv.WriteHead(header); err != nil {
		return err
	}
	for k, v := range ii {
		iv := make([]interface{}, 2)
		iv[0] = k
		iv[1] = v
		if err := csv.WriteLine(iv...); err != nil {
			return err
		}
	}
	return nil
}
