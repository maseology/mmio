package mmio

// InterfaceToFloat converts a slice of interface, and converts to float (assuming possible)
func InterfaceToFloat(d []interface{}) []float64 {
	out := make([]float64, len(d))
	for i := range d {
		out[i] = d[i].(float64)
	}
	return out
}
