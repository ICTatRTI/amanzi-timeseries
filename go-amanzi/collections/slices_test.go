package collections

import (
	"fmt"
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"os"
	"testing"
)

func TestFloat32Slices(t *testing.T) {
	for _, tc := range []struct {
		name        string
		pbPath      string
		expected    float32
		willFail    bool
		missingSize int
	}{
		{
			name:        "Basic",
			pbPath:      "./testdata/test_float32.json",
			expected:    8.6,
			missingSize: 0,
		},
		{
			name:        "mixed fields",
			willFail:    true,
			pbPath:      "./testdata/mixed_types.json",
			expected:    142.610001,
			missingSize: 0,
		},
		{
			name:        "json ts",
			pbPath:      "./testdata/compat_json_ts.json",
			expected:    5659.088867,
			missingSize: 0,
		},
		{
			name:        "missing values",
			pbPath:      "./testdata/missing_values.json",
			expected:    142.610001,
			missingSize: 1,
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
			data := Float32TsData{
				Values:  make([]float32, 0, len(pb.Data)),
				Missing: make(map[int]struct{}),
			}
			if err := Float32Slices(&pb, &data); err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			var sum float32
			for i := 0; i < len(pb.Data); i++ {
				sum += data.Values[i]
			}
			if sum != tc.expected {
				t.Fatalf("expected %f but got %f", tc.expected, sum)
			}
			if len(data.Missing) != tc.missingSize {
				t.Fatalf("expected missing size to be %d but got %d", tc.missingSize, len(data.Missing))
			}
		})
	}
}
