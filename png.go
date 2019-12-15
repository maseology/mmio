package mmio

import (
	"fmt"
	"image"
	"image/png"
	"math"
	"os"
)

// ArrayToPNG prints a 2D array (as row-major 1D array) to a png
func ArrayToPNG(fp string, v []float64, nr, nc int) {
	img := image.NewRGBA(image.Rect(0, 0, nc, nr))
	min, max := math.MaxFloat64, -math.MaxFloat64
	for _, vv := range v {
		if vv < min {
			min = vv
		}
		if vv > max {
			max = vv
		}
	}
	for j := 0; j < nc; j++ {
		for i := 0; i < nr; i++ {
			img.Set(j, i, interp((v[i*nc+j]-min)/(max-min)))
		}
	}

	f, err := os.Create(fp)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	err = png.Encode(f, img)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// fmt.Println(img)
	// matrix := [5][5]bool{}
	// // fill in matrix here
	// m := image.NewRGBA(image.Rect(0, 0, 5, 5))
	// fg := color.RGBA{255, 0, 255, 255}
	// bg := color.RGBA{255, 255, 255, 255}
	// draw.Draw(m, m.Bounds(), &image.Uniform{bg}, image.ZP, draw.Src)
	// for x, row := range matrix {
	// 	for y, cell := range row {
	// 		if cell {
	// 			m.Set(x, y, fg)
	// 		}
	// 	}
	// }
	// // buffer := new(bytes.Buffer)
	// // png.Encode(buffer, *img)
	// // w.Write(buffer.Bytes())

}
