package mmio

import (
	"io/ioutil"
	"log"

	"github.com/maseology/mmaths"
	geojson "github.com/paulmach/go.geojson"
)

func LineSegmentsToGeojson(lns map[int]mmaths.LineSegment, outfp string) {
	DeleteFile(outfp)
	fc := geojson.NewFeatureCollection()
	for i, l := range lns {
		ft := geojson.NewLineStringFeature([][]float64{{l.P0.X, l.P0.Y}, {l.P1.X, l.P1.Y}})
		ft.Properties["fid"] = i
		// ft.Properties["length"] = area / 1000 / 1000
		fc.AddFeature(ft)
	}

	rawJSON, err := fc.MarshalJSON()
	if err != nil {
		panic(err)
	}

	WriteString(outfp, string(rawJSON)+"\n")
}

func ReadGeojsonLines(fp string) [][][]float64 {
	flines, err := ioutil.ReadFile(fp)
	if err != nil {
		panic(err)
	}
	glines, err := geojson.UnmarshalFeatureCollection(flines)
	if err != nil {
		panic(err)
	}

	var o [][][]float64
	for _, f := range glines.Features {
		// fmt.Printf("  feature properties: %v\n",f.Properties)
		switch f.Geometry.Type {
		case "LineString":
			o = append(o, f.Geometry.LineString)
		case "MultiLineString":
			o = append(o, f.Geometry.MultiLineString...)
		default:
			log.Fatalf("Routing.LoadNetwork: unsupported type, given %v\n", f.Geometry.Type)
		}
	}
	return o
}
