package mmio

import (
	"fmt"
	"image/color"
)

func interp(f float64) color.RGBA {
	if f >= 1. {
		c := linearL[len(linearL)-1]
		return color.RGBA{uint8(c[1]), uint8(c[2]), uint8(c[3]), uint8(c[4])}
	} else if f < 0. {
		c := linearL[0]
		return color.RGBA{uint8(c[1]), uint8(c[2]), uint8(c[3]), uint8(c[4])}
	} else {
		for _, c := range linearL {
			if c[0]/100. > f {
				return color.RGBA{uint8(c[1]), uint8(c[2]), uint8(c[3]), uint8(c[4])}
			}
		}
	}
	return color.RGBA{0, 0, 0, 0}
}

// https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func ParseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff
	switch len(s) {
	case 7:
		_, err = fmt.Sscanf(s, "#%02x%02x%02x", &c.R, &c.G, &c.B)
	case 4:
		_, err = fmt.Sscanf(s, "#%1x%1x%1x", &c.R, &c.G, &c.B)
		// Double the hex digits:
		c.R *= 17
		c.G *= 17
		c.B *= 17
	default:
		err = fmt.Errorf("invalid length, must be 7 or 4")

	}
	return
}

// func interp(f float64) uint8 {
// 	if f > 1. {
// 		return 255
// 	} else if f < 0. {
// 		return 0
// 	} else {
// 		return uint8(f * 255)
// 	}
// }
