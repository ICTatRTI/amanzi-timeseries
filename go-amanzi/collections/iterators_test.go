package collections

import (
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"os"
	"testing"
	"time"
)

func TestLazyIterator(t *testing.T) {
	itr := NewLazyIterator(&ptypes.TimeSeries{})
	var sum float64
	for itr.Next() {
		rec := itr.Float64Record()
		// if our record isn't missing
		if !rec.Missing {
			sum += rec.Value
		}

		// do something with value
	}
	if err := itr.Err(); err != nil {
		// handle error
	}
}

func TestEagerIterator_Float32Record_Compatability(t *testing.T) {
	for _, tc := range []struct {
		name     string
		pbPath   string
		expected float32
	}{
		{
			name:     "Basic",
			pbPath:   "./testdata/test_float32.json",
			expected: 8.6,
		},
		{
			name:     "mixed fields",
			pbPath:   "./testdata/mixed_types.json",
			expected: 142.610001,
		},
		{
			name:     "json ts",
			pbPath:   "./testdata/compat_json_ts.json",
			expected: 5659.088867,
		},
		{
			name:     "missing values",
			pbPath:   "./testdata/missing_values.json",
			expected: 142.610001,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f, err := os.Open(tc.pbPath)
			if err != nil {
				t.Fatal(err.Error())
			}
			var pb ptypes.TimeSeries
			if err := jsonpb.Unmarshal(f, &pb); err != nil {
				t.Fatalf(err.Error())
			}
			itr, err := NewEagerIterator(&pb)
			if err != nil {
				t.Fatal(err)
			}
			var sum float32
			for itr.Next() {
				rec := itr.Float32Record()
				if !rec.Missing {
					sum += rec.Value
				}
				itr.Len()
			}
			if sum != tc.expected {
				t.Fatalf("expected %f got %f", tc.expected, sum)
			}
		})
	}
}

func TestEagerIterator_Float32Record(t *testing.T) {
	times := []time.Time{time.Now(), time.Now()}
	values := []float32{1.3, 4.5}
	q := [][]string{{"a", "p"}, {"d", "g"}}
	p := []ptypes.Properties{{"test": 5}, {"test": 6}}
	meta := &ptypes.TimeSeriesMetaInfo{}
	tsPb, err := ptypes.NewFloat32TimeSeries(meta, times, values, q, p)
	itr, err := NewEagerIterator(tsPb)
	if err != nil {
		t.Fatal(err)
	}
	var sum float32
	for itr.Next() {
		rec := itr.Float32Record()
		if !rec.Missing {
			sum += rec.Value
		}
	}
	if sum != 5.8 {
		t.Fatalf("expected %f got %f", 5.8, sum)
	}
}
