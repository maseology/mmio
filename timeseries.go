package mmio

import "time"

// TimeSeries is a collection of temporal data
type TimeSeries map[time.Time]float64

// TimeSeriesMonthly returns a monthly timeseries
type TimeSeriesMonthly map[Yr]map[Mo]float64

// MonthlySumCount converts a timeseries to a sum TimeSeriesMonthly and a count TimeSeriesMonthly
func MonthlySumCount(ts TimeSeries) (_, _ TimeSeriesMonthly) {
	tb, te := MinMaxTimeseries(ts)
	yrb, yre := Yr(tb.Year()), Yr(te.Year())
	sum, cnt := make(TimeSeriesMonthly, yre-yrb+1), make(TimeSeriesMonthly, yre-yrb+1)
	for yr := yrb; yr <= yre; yr++ { // initialize
		sum[yr] = make(map[Mo]float64, 12)
		cnt[yr] = make(map[Mo]float64, 12)
	}
	for dt, v := range ts {
		mo, yr := Mo(dt.Month()), Yr(dt.Year())
		sum[yr][mo] += v
		cnt[yr][mo]++
	}
	return sum, cnt
}

// MinMaxTimeseries returns the range of a Timeseries
func MinMaxTimeseries(ts TimeSeries) (_, _ time.Time) {
	tx, tn := MinMaxTime()
	for dt := range ts {
		if dt.Before(tn) {
			tn = dt
		}
		if dt.After(tx) {
			tx = dt
		}
	}
	return tn, tx
}
