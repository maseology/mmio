package mmio

import (
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
