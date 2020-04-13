package ptypes

import "testing"

func TestIterInt32TimeSeries(t *testing.T) {
	ts := &TimeSeries{}
	var itr TimeSeriesStringIterator
	itr, err := IterStringTimeSeries(ts)
	if err != nil {
		return
	}

	for i := 0; i < itr.Len(); i++ {
		itr.GetProperties(i)
	}
}
