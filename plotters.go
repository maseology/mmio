package mmio

import (
	"image/color"
	"math"
	"sort"

	"github.com/maseology/mmaths"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Histo : create generic histogram
func Histo(fp string, x []float64, nbins int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fp

	v := make(plotter.Values, len(x))
	for i, d := range x {
		v[i] = d
	}

	h, err := plotter.NewHist(v, nbins)
	if err != nil {
		panic(err)
	}

	// Normalize the area under the histogram to
	// sum to one.
	h.Normalize(1)
	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "hist.png"); err != nil {
		panic(err)
	}
}

// ObsSim used to create simple observed vs. simulated hydrographs
func ObsSim(fp string, o, s []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// p.Title.Text = fp
	p.X.Label.Text = ""
	p.Y.Label.Text = "discharge"

	ps, err := plotter.NewLine(sequentialLine(s))
	if err != nil {
		panic(err)
	}
	ps.Color = color.RGBA{R: 255, A: 255}

	po, err := plotter.NewLine(sequentialLine(o))
	if err != nil {
		panic(err)
	}
	po.Color = color.RGBA{B: 255, A: 255}

	// Add the functions and their legend entries.
	p.Add(ps, po)
	p.Legend.Add("obs", po)
	p.Legend.Add("sim", ps)

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// ObsSimFDC used to create simple observed vs. simulated flow-duration curves
func ObsSimFDC(fp string, o, s []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// p.Title.Text = fp
	p.X.Label.Text = ""
	p.Y.Label.Text = "discharge"

	ps, err := plotter.NewLine(cumulativeDistributionLine(s))
	if err != nil {
		panic(err)
	}
	ps.Color = color.RGBA{R: 255, A: 255}

	po, err := plotter.NewLine(cumulativeDistributionLine(o))
	if err != nil {
		panic(err)
	}
	po.Color = color.RGBA{B: 255, A: 255}

	// Add the functions and their legend entries.
	p.Add(ps, po)
	p.Legend.Add("obs", po)
	p.Legend.Add("sim", ps)
	p.Y.Scale = plot.LogScale{}
	p.Y.Tick.Marker = plot.LogTicks{}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// Wbal used to review waterbalance
func Wbal(fp string, f, a, q, g, s []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	err = plotutil.AddLines(p,
		"pre", sequentialLine(f),
		"aet", sequentialLine(a),
		"ro", sequentialLine(q),
		"rch", sequentialLine(g),
		"sto", sequentialLine(s))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// Scatter : create generic scatter plot
func Scatter(fp string, x, y []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	p.Title.Text = fp
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	err = plotutil.AddScatters(p, points(x, y))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// Line : generic line plot
func Line(fp string, x []float64, ys map[string][]float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	lines := make([]interface{}, 0)
	for l, y := range ys {
		lines = append(lines, l)
		lines = append(lines, points(x, y))
	}
	err = plotutil.AddLinePoints(p, lines...)
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// LinePoints1 : generic plot of 1 line
func LinePoints1(fp string, x, y []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	err = plotutil.AddLinePoints(p, "v1", points(x, y))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// LinePoints2 : generic plot of 2 lines
func LinePoints2(fp string, x, y1, y2 []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	err = plotutil.AddLinePoints(p,
		"v1", points(x, y1),
		"v2", points(x, y2))
	if err != nil {
		panic(err)
	}

	// Save the plot to a PNG file.
	if err := p.Save(12*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

func points(x, y []float64) plotter.XYs {
	if len(x) != len(y) {
		panic("mmplt.scatter error: unequal points array sizes")
	}
	pts := make(plotter.XYs, len(x))
	for i := range pts {
		pts[i].X = x[i]
		pts[i].Y = y[i]
	}
	return pts
}

func sequentialLine(v []float64) plotter.XYs {
	pts, c := make(plotter.XYs, len(v)), 0
	for i := range pts {
		if math.IsNaN(v[i]) {
			continue
		}
		pts[c].X = float64(i)
		pts[c].Y = v[i]
		c++
	}
	return pts[:c]
}

func cumulativeDistributionLine(v []float64) plotter.XYs {
	v = mmaths.OnlyPositive(v)
	sort.Float64s(v)
	mmaths.RevF(v)
	pts, c, x := make(plotter.XYs, len(v)), 0, float64(len(v))/100.
	for i := range pts {
		if math.IsNaN(v[i]) {
			continue
		}
		pts[c].X = float64(i) / x
		pts[c].Y = v[i]
		c++
	}
	return pts[:c]
}
