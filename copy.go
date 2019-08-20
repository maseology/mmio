package mmio

// CopyMapif deep copies map[int]float64
func CopyMapif(originalMap map[int]float64) (newMap map[int]float64) {
	newMap = make(map[int]float64, len(originalMap))
	for k, v := range originalMap {
		newMap[k] = v
	}
	return
}
