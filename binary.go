package mmio

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
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
	b, err := os.ReadFile(filepath)
	if err != nil {
		log.Fatalf("Fatal error: binary.OpenBinary failed: %v\n", err)
	}
	return bytes.NewReader(b)
}

// ReadBinary general binary reader
func ReadBinary(filepath string, data ...interface{}) error {
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ReadBinary failed: %v\n", err)
		return fmt.Errorf("os.ReadFile failed: %v", err)
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

// ReadString reads and returns string from binary file
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

// ReadFloat32 reads next float32 from buffer
func ReadFloat32(b *bytes.Reader) float32 {
	var f float32
	if err := binary.Read(b, binary.LittleEndian, &f); err != nil {
		log.Fatalf("ReadFloat32 failed: %v", err)
	}
	return f
}

// ReadFloat64 reads next float64 from buffer
func ReadFloat64(b *bytes.Reader) float64 {
	var f float64
	if err := binary.Read(b, binary.LittleEndian, &f); err != nil {
		log.Fatalf("ReadFloat64 failed: %v", err)
	}
	return f
}

// ReadBytes reads next n-byte array from buffer
func ReadBytes(b *bytes.Reader, n int) []byte {
	i := make([]byte, n)
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadUInt8 failed:", err)
	}
	return i
}

// ReadUInt8 reads next uint8 from buffer
func ReadUInt8(b *bytes.Reader) uint8 {
	var i uint8
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadUInt8 failed:", err)
	}
	return i
}

// ReadInt8 reads next int8 (signed byte) from buffer
func ReadInt8(b *bytes.Reader) int8 {
	var i int8
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadInt8 failed:", err)
	}
	return i
}

// ReadUInt16 reads next uint16 from buffer
func ReadUInt16(b *bytes.Reader) uint16 {
	var i uint16
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadUInt16 failed:", err)
	}
	return i
}

// ReadUInt32 reads next uint32 from buffer
func ReadUInt32(b *bytes.Reader) uint32 {
	var i uint32
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadUInt32 failed:", err)
	}
	return i
}

// ReadUInt64 reads next uint64 from buffer
func ReadUInt64(b *bytes.Reader) uint64 {
	var i uint64
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadUInt64 failed:", err)
	}
	return i
}

// ReadInt32 reads next int32 from buffer
func ReadInt32(b *bytes.Reader) int32 {
	var i int32
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadInt32 failed:", err)
	}
	return i
}

// ReadInt32check reads next int32 from buffer
func ReadInt32check(b *bytes.Reader) (int32, bool) {
	var i int32
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		if err == io.EOF {
			return 0, false
		}
		log.Fatalf("ReadInt32check failed: %v", err)
	}
	return i, true
}

// ReadInt64 reads next int64 from buffer
func ReadInt64(b *bytes.Reader) int64 {
	var i int64
	if err := binary.Read(b, binary.LittleEndian, &i); err != nil {
		fmt.Println("ReadInt64 failed:", err)
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

// ReadBinaryFloats reads an entire file and returns a slice of floats
func ReadBinaryFloats(filepath string) ([]float64, error) {
	var err error
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ReadBinaryFloats failed: %v\n", err)
		return nil, fmt.Errorf("os.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	n := len(b) / 8
	a := make([]float64, n)
	err = binary.Read(buf, binary.LittleEndian, a)
	if err != nil {
		fmt.Printf("ReadBinaryFloats failed: %v\n", err)
		return nil, fmt.Errorf("binary.Read failed: %v", err)
	}
	return a, nil
}

// ReadBinaryFloats reads an entire file and returns a slice of d dimensions
func ReadBinaryFloat64s(filepath string, d int) ([][]float64, int, error) {
	var err error
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ReadBinaryFloats failed: %v\n", err)
		return nil, 0, fmt.Errorf("os.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	n := len(b) / 8 / d
	a := make([][]float64, d)
	for i := 0; i < d; i++ {
		v := make([]float64, n)
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Printf("ReadBinaryFloats failed: %v\n", err)
			return nil, 0, fmt.Errorf("binary.Read failed: %v", err)
		}
		a[i] = v
	}
	return a, n, nil
}

// ReadBinaryFloat32s reads an entire file and returns a slice of d dimensions
func ReadBinaryFloat32s(filepath string, d int) ([][]float32, int, error) {
	var err error
	b, err := os.ReadFile(filepath)
	if err != nil {
		fmt.Printf("ReadBinaryFloats failed: %v\n", err)
		return nil, 0, fmt.Errorf("os.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	n := len(b) / 4 / d
	a := make([][]float32, d)
	for i := 0; i < d; i++ {
		v := make([]float32, n)
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			fmt.Printf("ReadBinaryFloats failed: %v\n", err)
			return nil, 0, fmt.Errorf("binary.Read failed: %v", err)
		}
		a[i] = v
	}
	return a, n, nil
}

// ReadBinaryInts reads an entire file and returns a slice of d dimensions
func ReadBinaryInts(filepath string, d int) ([][]int32, int, error) {
	var err error
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, 0, fmt.Errorf("ReadBinaryInts: os.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	n := len(b) / 4 / d
	a := make([][]int32, d)
	for i := 0; i < d; i++ {
		v := make([]int32, n)
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			return nil, 0, fmt.Errorf("ReadBinaryInts: binary.Read failed: %v", err)
		}
		a[i] = v
	}
	return a, n, nil
}

// ReadBinaryShorts reads an entire file and returns a slice of d dimensions
func ReadBinaryShorts(filepath string, d int) ([][]int16, int, error) {
	var err error
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, 0, fmt.Errorf("ReadBinaryShorts: os.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	n := len(b) / 2 / d
	a := make([][]int16, d)
	for i := 0; i < d; i++ {
		v := make([]int16, n)
		err := binary.Read(buf, binary.LittleEndian, v)
		if err != nil {
			return nil, 0, fmt.Errorf("ReadBinaryShorts: binary.Read failed: %v", err)
		}
		a[i] = v
	}
	return a, n, nil
}

// ReadBinaryIMAP reads a map[int]int for an entire file
func ReadBinaryIMAP(filepath string) (map[int]int, error) {
	var err error
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("ReadBinaryIMAP: os.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	n := len(b) / 8
	m := make(map[int]int, n)
	v := make([]int32, 2*n)
	if err := binary.Read(buf, binary.LittleEndian, v); err != nil {
		return nil, fmt.Errorf("ReadBinaryIMAP: binary.Read failed: %v", err)
	}
	for i := 0; i < n; i++ {
		m[int(v[2*i])] = int(v[2*i+1])
	}
	return m, nil
}

// ReadBinaryRMAP reads a map[int]float64 for an entire file
func ReadBinaryRMAP(filepath string) (map[int]float64, error) {
	var err error
	b, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("ReadBinaryRMAP: os.ReadFile failed: %v", err)
	}
	buf := bytes.NewReader(b)
	type Dat struct {
		I int32
		F float64
	}
	n := len(b) / 12
	v := make([]Dat, n)
	if err := binary.Read(buf, binary.LittleEndian, v); err != nil {
		return nil, fmt.Errorf("ReadBinaryRMAP: %v", err)
	}
	m := make(map[int]float64, len(v))
	for _, d := range v {
		m[int(d.I)] = d.F
	}
	return m, nil
}

// WriteBinary general binary writer
func WriteBinary(filepath string, data ...interface{}) error {
	buf := new(bytes.Buffer)
	for _, v := range data {
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			return fmt.Errorf("mmio.WriteBinary failed: %v", err)
		}
	}
	if err := os.WriteFile(filepath, buf.Bytes(), 0644); err != nil { // see: https://en.wikipedia.org/wiki/File_system_permissions
		return fmt.Errorf("mmio.WriteBinary failed: %v", err)
	}
	return nil
}

// WriteIMAP general map writer
func WriteIMAP(filepath string, data map[int]int) error {
	buf := new(bytes.Buffer)
	for k, v := range data {
		if err := binary.Write(buf, binary.LittleEndian, int32(k)); err != nil {
			log.Fatalln("WriteBinary failed:", err)
		}
		if err := binary.Write(buf, binary.LittleEndian, int32(v)); err != nil {
			log.Fatalln("WriteBinary failed:", err)
		}
	}
	if err := os.WriteFile(filepath, buf.Bytes(), 0644); err != nil { // see: https://en.wikipedia.org/wiki/File_system_permissions
		return fmt.Errorf(" os.WriteIMAP failed: %v", err)
	}
	return nil
}

// WriteRMAP general map writer
func WriteRMAP(filepath string, data map[int]float64, append bool) error {
	buf := new(bytes.Buffer)
	for k, v := range data {
		if err := binary.Write(buf, binary.LittleEndian, int32(k)); err != nil {
			log.Fatalln("WriteBinary failed:", err)
		}
		if err := binary.Write(buf, binary.LittleEndian, v); err != nil {
			log.Fatalln("WriteBinary failed:", err)
		}
	}

	if append {
		// If the file doesn't exist, create it, or append to the file
		f, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		if _, err := f.Write(buf.Bytes()); err != nil {
			return err
		}
		if err := f.Close(); err != nil {
			return err
		}
	} else {
		if err := os.WriteFile(filepath, buf.Bytes(), 0644); err != nil { // see: https://en.wikipedia.org/wiki/File_system_permissions
			return fmt.Errorf(" os.WriteRMAP failed: %v", err)
		}
	}
	return nil
}
