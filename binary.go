package mmio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// // example
// var routi int32 = 42
// routf := [5]float32{1., 2., 3., 4., 5.2}
// WriteBinary("out1.bin", routi, routf)

// var rini int32
// rinf := [5]float32{}
// erri := ReadBinary("out1.bin", &rini, &rinf)
// if erri != nil {
// 	fmt.Println(erri)
// }
// fmt.Println(rini, rinf)

// OpenBinary creates reader from filepath
func OpenBinary(filepath string) *bytes.Reader {
	var err error
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Fatal error: OpenBinary failed: %v\n", err)
	}
	return bytes.NewReader(b)
}

// ReadBinary general binary reader
func ReadBinary(filepath string, data ...interface{}) error {
	var err error
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ReadBinary failed: %v\n", err)
		return fmt.Errorf("ioutil.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	for _, v := range data {
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Printf("ReadBinary failed: %v\n", err)
			return fmt.Errorf("binary.Read failed: %v", err)
		}
	}
	return nil
}

// ReachedEOF tests to see if all reader data has been read
func ReachedEOF(b *bytes.Reader) bool {
	t := make([]byte, 1)
	if v, _ := b.Read(t); v != 0 {
		return false
	}
	return true
}

// ReadString : reads and returns string from binary file
func ReadString(b *bytes.Reader) string {
	var blen byte
	err := binary.Read(b, binary.LittleEndian, &blen)
	if err != nil {
		fmt.Println("ReadString failed:", err)
	}
	len := int(blen)
	if len > 127 {
		var blen2 byte
		err = binary.Read(b, binary.LittleEndian, &blen2)
		if err != nil {
			fmt.Println("ReadString failed:", err)
		}
		len += (int(blen2) - 1) * 128
	}
	str := make([]byte, len)
	err = binary.Read(b, binary.LittleEndian, &str)
	if err != nil {
		fmt.Println("ReadString failed:", err)
	}
	return string(str)
}

// ReadFloat64 reads next float64 from buffer
func ReadFloat64(b *bytes.Reader) float64 {
	var f float64
	if err := binary.Read(b, binary.LittleEndian, &f); err != nil {
		fmt.Println("ReadFloat64 failed:", err)
	}
	return f
}

// ReadInt32 reads next int32 from buffer
func ReadInt32(b *bytes.Reader) int32 {
	var i int32
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadInt32 failed:", err)
	}
	return i
}

// ReadLines reads and returns string lines from binary file
func ReadLines(b *bytes.Reader) []string {
	return strings.FieldsFunc(ReadString(b), lineParser)
}
func lineParser(r rune) bool {
	return r == '\r' || r == '\n'
}

// WriteBinary : general binary writer
func WriteBinary(filepath string, data ...interface{}) error {
	buf := new(bytes.Buffer)
	for _, v := range data {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Println("WriteBinary failed:", err)
		}
	}
	if err := ioutil.WriteFile(filepath, buf.Bytes(), 0644); err != nil { // see: https://en.wikipedia.org/wiki/File_system_permissions
		return fmt.Errorf("ioutil.WriteFile failed: %v", err)
	}
	return nil
}

// ReadBinaryFloats : reads an entire file and returns a slice of d dimensions
func ReadBinaryFloats(filepath string, d int) ([][]float64, int, error) {
	var err error
	b, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ReadBinaryFloats failed: %v\n", err)
		return nil, 0, fmt.Errorf("ioutil.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	n := len(b) / 8 / d
	a := make([][]float64, n, n)
	for i := 0; i < n; i++ {
		v := make([]float64, d, d)
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Printf("ReadBinaryFloats failed: %v\n", err)
			return nil, 0, fmt.Errorf("binary.Read failed: %v", err)
		}
		a[i] = v
	}
	return a, n, nil
}
