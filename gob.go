package mmio

import (
	"encoding/gob"
	"os"
)

// SaveGOB saves map[int]int
func SaveGOB(fp string, d map[int]int) error {
	f, err := os.Create(fp)
	defer f.Close()
	if err != nil {
		return err
	}
	enc := gob.NewEncoder(f)
	err = enc.Encode(d)
	if err != nil {
		return err
	}
	return nil
}

// LoadGOB saves map[int]int
func LoadGOB(fp string) (map[int]int, error) {
	var d map[int]int
	f, err := os.Open(fp)
	defer f.Close()
	if err != nil {
		return nil, err
	}
	enc := gob.NewDecoder(f)
	err = enc.Decode(&d)
	if err != nil {
		return nil, err
	}
	return d, nil
}
