package mmio

import (
	"fmt"
	"image/color"
	"math"
	"sort"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

// Histo creates a generic histogram
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
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// HistoGT0 creates a generic histogram of all values >0.
func HistoGT0(fp string, x []float64, nbins int) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	n0 := 0
	for _, d := range x {
		if d <= 0. {
			n0++
		}
	}

	p.Title.Text = fmt.Sprintf("%s (n= %d; n0=%d)", fp, len(x), n0)

	v, i := make(plotter.Values, len(x)-n0), 0
	for _, d := range x {
		if d > 0. {
			v[i] = d
			i++
		}
	}

	h, err := plotter.NewHist(v, nbins)
	if err != nil {
		panic(err)
	}

	p.Add(h)

	// Save the plot to a PNG file.
	if err := p.Save(4*vg.Inch, 4*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// ObsSim is used to create simple observed vs. simulated hydrographs
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
	p.Legend.Top = true
	// p.X.Tick.Marker = plot.TimeTicks{Format: "Jan"}

	// Save the plot to a PNG file.
	if err := p.Save(24*vg.Inch, 8*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// ObsSimFDC is used to create simple observed vs. simulated flow-duration curves
func ObsSimFDC(fp string, o, s []float64) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// create copies
	ocopy, scopy := make([]float64, len(o)), make([]float64, len(s))
	copy(ocopy, o)
	copy(scopy, s)

	// p.Title.Text = fp
	p.X.Label.Text = ""
	p.Y.Label.Text = "discharge"

	ps, err := plotter.NewLine(cumulativeDistributionLine(scopy))
	if err != nil {
		panic(err)
	}
	ps.Color = color.RGBA{R: 255, A: 255}

	po, err := plotter.NewLine(cumulativeDistributionLine(ocopy))
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

// Scatter creates a generic scatter plot
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

// Line creates a generic line plot
func Line(fp string, x []float64, ys map[string][]float64, width float64) {
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
	p.Legend.Top = true

	// Save the plot to a PNG file.
	if err := p.Save(vg.Length(width)*vg.Inch, 8*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// LineCol creates a generic line plot with specified colour scheme
func LineCol(fp string, x []float64, ys map[string][]float64, colours map[string]color.RGBA) {
	p, err := plot.New()
	if err != nil {
		panic(err)
	}

	// p.Title.Text = fp
	// p.X.Label.Text = ""
	// p.Y.Label.Text = ""

	for l, y := range ys {
		ps, err := plotter.NewLine(points(x, y))
		if err != nil {
			panic(err)
		}
		if _, ok := colours[l]; !ok {
			panic("colour not found")
		}
		ps.Color = colours[l]
		p.Add(ps)
		p.Legend.Add(l, ps)
	}
	p.Legend.Top = true

	// Save the plot to a PNG file.
	if err := p.Save(16*vg.Inch, 8*vg.Inch, fp); err != nil {
		panic(err)
	}
}

// LinePoints1 creates a generic plot of one xy set of data only
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

// LinePoints2 creates a generic plot of lines from 2 sets of xy data
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
	v = OnlyPositive(v)
	sort.Float64s(v)
	RevF(v)
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
