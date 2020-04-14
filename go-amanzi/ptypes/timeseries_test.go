package ptypes

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"io/ioutil"
	"testing"
	"time"
)

func TestNewInt32TimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []int32
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-int32",
			times:  []time.Time{time.Now(), time.Now()},
			values: []int32{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int32-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []int32{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int32-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int32{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int32-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int32{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int32-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int32{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewInt32TimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewUint32TimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []uint32
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-uint32",
			times:  []time.Time{time.Now(), time.Now()},
			values: []uint32{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint32-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []uint32{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint32-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint32{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint32-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint32{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint32-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint32{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewUint32TimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewInt64TimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []int64
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-int64",
			times:  []time.Time{time.Now(), time.Now()},
			values: []int64{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int64-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []int64{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int64-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int64{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int64-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int64{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int64-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int64{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewInt64TimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewUint64TimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []uint64
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-uint64",
			times:  []time.Time{time.Now(), time.Now()},
			values: []uint64{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint64-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []uint64{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint64-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint64{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint64-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint64{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint64-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint64{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewUint64TimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewIntTimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []int
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-int",
			times:  []time.Time{time.Now(), time.Now()},
			values: []int{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []int{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "int-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []int{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewIntTimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewUintTimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []uint
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-uint",
			times:  []time.Time{time.Now(), time.Now()},
			values: []uint{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []uint{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "uint-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []uint{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewUintTimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewFloat32TimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []float32
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-float32",
			times:  []time.Time{time.Now(), time.Now()},
			values: []float32{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float32-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []float32{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float32-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []float32{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float32-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []float32{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float32-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []float32{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewFloat32TimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewFloat64TimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []float64
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-float64",
			times:  []time.Time{time.Now(), time.Now()},
			values: []float64{1, 2},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float64-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []float64{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float64-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []float64{1},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float64-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []float64{1, 2},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "float64-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []float64{1, 2},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewFloat64TimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}

func TestNewStringTimeSeries(t *testing.T) {

	for _, tc := range []struct {
		name     string
		willFail bool
		times    []time.Time
		values   []string
		q        [][]string
		p        []Properties
		m        *TimeSeriesMetaInfo
	}{
		{
			name:   "basic-string",
			times:  []time.Time{time.Now(), time.Now()},
			values: []string{"1", "2"},
			q:      [][]string{{"a", "p"}, {"d", "g"}},
			p:      []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "string-arraymismatch-times",
			willFail: true,
			times:    []time.Time{time.Now()},
			values:   []string{"1", "2"},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "string-arraymismatch-values",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []string{"1"},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "string-arraymismatch-qualifiers",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []string{"1", "2"},
			q:        [][]string{{"a", "p"}},
			p:        []Properties{{"test": 5}, {"test": 6}},
		},
		{
			name:     "string-arraymismatch-properties",
			willFail: true,
			times:    []time.Time{time.Now(), time.Now()},
			values:   []string{"1", "2"},
			q:        [][]string{{"a", "p"}, {"d", "g"}},
			p:        []Properties{{"test": 5}},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			pb, err := NewStringTimeSeries(tc.m, tc.times, tc.values, tc.q, tc.p)
			if err != nil {
				if tc.willFail {
					fmt.Printf("received expected error: %s\n", err.Error())
					return
				}
				t.Fatalf(err.Error())
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			m := jsonpb.Marshaler{
				EnumsAsInts:  false,
				EmitDefaults: false,
				Indent: "	",
				OrigName:    false,
				AnyResolver: nil,
			}
			var buf bytes.Buffer
			if err := m.Marshal(&buf, pb); err != nil {
				t.Fatalf(err.Error())
			}
			if err := ioutil.WriteFile(fmt.Sprintf("./testdata/%s.json", tc.name), buf.Bytes(), 0644); err != nil {
				t.Fatalf(err.Error())
			}
		})
	}
}
