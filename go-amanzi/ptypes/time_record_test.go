package ptypes

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/ptypes"
	"io/ioutil"
	"testing"
	"time"
)

func TestNewTimeRecord(t *testing.T) {
	for _, tc := range []struct {
		name     string
		t        time.Time
		val      interface{}
		q        []string
		p        Properties
		willFail bool
	}{
		{
			name:     "int32 test",
			t:        time.Now(),
			val:      int32(3),
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "int test",
			t:        time.Now(),
			val:      int32(3),
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "int64 test",
			t:        time.Now(),
			val:      int64(3),
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "uint32 test",
			t:        time.Now(),
			val:      uint32(3),
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "int test",
			t:        time.Now(),
			val:      uint32(3),
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "uint64 test",
			t:        time.Now(),
			val:      uint64(3),
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "null test",
			t:        time.Now(),
			val:      nil,
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "double test",
			t:        time.Now(),
			val:      3.5,
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "float test",
			t:        time.Now(),
			val:      float32(3.5),
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "string test",
			t:        time.Now(),
			val:      "string",
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name: "struct test",
			t:    time.Now(),
			val: struct {
				name string
			}{
				name: "a bad type",
			},
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name: "map test",
			t:    time.Now(),
			val: map[string]interface{}{
				"test": 5,
			},
			p:        Properties{"test": 5},
			q:        []string{},
			willFail: false,
		},
		{
			name:     "slice test",
			t:        time.Now(),
			val:      []interface{}{5.3, "biscuit", 4},
			p:        Properties{"test": 5},
			willFail: false,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			record, err := NewTimeRecord(tc.t, tc.val, tc.p, tc.q...)
			if err != nil {

				if tc.willFail {
					fmt.Printf("received expected error: %s", err.Error())
					return
				}
				t.Fatalf("error: %s", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			got := record.Value.Kind
			switch typ := got.(type) {
			case *Value_Uint64Value:
			case *Value_Uint32Value:
			case *Value_Int32Value:
			case *Value_Int64Value:
			case *Value_DoubleValue:
			case *Value_FloatValue:
			case *Value_NullValue:
			case *Value_StringValue:
			default:
				t.Fatalf("error: invalid type %v", typ)
			}
		})
	}
}

func TestNewInt32Record(t *testing.T) {
	for _, tc := range []struct {
		name      string
		willFail  bool
		isMissing bool
		pb        *TimeRecord
	}{
		{
			name:     "basic int32",
			willFail: false,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int32Value{Int32Value: 4}},
			},
		},
		{
			name:     "int32 with int64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int64Value{Int64Value: 4}},
			},
		},
		{
			name:     "int32 with uint32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint32Value{Uint32Value: 4}},
			},
		},
		{
			name:     "int32 with uint64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint64Value{Uint64Value: 4}},
			},
		},
		{
			name:     "int32 with float64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_DoubleValue{DoubleValue: 4.3}},
			},
		},
		{
			name:     "int32 with float32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_FloatValue{FloatValue: 4.3}},
			},
		},
		{
			name:     "int32 with string",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_StringValue{StringValue: "4.3"}},
			},
		},
		{
			name:      "int32 with null",
			willFail:  false,
			isMissing: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec, err := NewInt32Record(tc.pb)
			if err != nil {
				if tc.willFail {
					if tc.willFail != rec.Missing {
						t.Fatalf("expected missing to be %t but got %t", tc.willFail, rec.Missing)
					}
					fmt.Printf("received expected error: %s\n", err.Error())
					return

				}
				t.Fatalf("error: %s\n", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			if rec.Missing != tc.isMissing {
				t.Fatalf("missing was expected to be %t but got %t", tc.isMissing, rec.Missing)
			}
		})
	}
}

func TestNewInt64Record(t *testing.T) {
	for _, tc := range []struct {
		name      string
		willFail  bool
		isMissing bool
		pb        *TimeRecord
	}{
		{
			name:     "basic int64",
			willFail: false,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int64Value{Int64Value: 4}},
			},
		},
		{
			name:     "int64 with int32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int32Value{Int32Value: 4}},
			},
		},
		{
			name:     "int64 with uint32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint32Value{Uint32Value: 4}},
			},
		},
		{
			name:     "int64 with uint64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint64Value{Uint64Value: 4}},
			},
		},
		{
			name:     "int64 with float64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_DoubleValue{DoubleValue: 4.3}},
			},
		},
		{
			name:     "int64 with float32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_FloatValue{FloatValue: 4.3}},
			},
		},
		{
			name:     "int64 with string",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_StringValue{StringValue: "4.3"}},
			},
		},
		{
			name:      "int64 with null",
			willFail:  false,
			isMissing: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec, err := NewInt64Record(tc.pb)
			if err != nil {
				if tc.willFail {
					if tc.willFail != rec.Missing {
						t.Fatalf("expected missing to be %t but got %t", tc.willFail, rec.Missing)
					}
					fmt.Printf("received expected error: %s\n", err.Error())
					return

				}
				t.Fatalf("error: %s\n", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			if rec.Missing != tc.isMissing {
				t.Fatalf("missing was expected to be %t but got %t", tc.isMissing, rec.Missing)
			}
		})
	}
}

func TestNewUInt64Record(t *testing.T) {
	for _, tc := range []struct {
		name      string
		willFail  bool
		isMissing bool
		pb        *TimeRecord
	}{
		{
			name:     "basic uint64",
			willFail: false,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint64Value{Uint64Value: 4}},
			},
		},
		{
			name:     "uint64 with  int64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int64Value{Int64Value: 4}},
			},
		},
		{
			name:     "uint64 with int32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int32Value{Int32Value: 4}},
			},
		},
		{
			name:     "uint64 with uint32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint32Value{Uint32Value: 4}},
			},
		},
		{
			name:     "uint64 with float64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_DoubleValue{DoubleValue: 4.3}},
			},
		},
		{
			name:     "uint64 with float32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_FloatValue{FloatValue: 4.3}},
			},
		},
		{
			name:     "uint64 with string",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_StringValue{StringValue: "4.3"}},
			},
		},
		{
			name:      "uint64 with null",
			willFail:  false,
			isMissing: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec, err := NewUint64Record(tc.pb)
			if err != nil {
				if tc.willFail {
					if tc.willFail != rec.Missing {
						t.Fatalf("expected missing to be %t but got %t", tc.willFail, rec.Missing)
					}
					fmt.Printf("received expected error: %s\n", err.Error())
					return

				}
				t.Fatalf("error: %s\n", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			if rec.Missing != tc.isMissing {
				t.Fatalf("missing was expected to be %t but got %t", tc.isMissing, rec.Missing)
			}
		})
	}
}

func TestNewUInt32Record(t *testing.T) {
	for _, tc := range []struct {
		name      string
		willFail  bool
		isMissing bool
		pb        *TimeRecord
	}{
		{
			name:     "basic uint32",
			willFail: false,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint32Value{Uint32Value: 4}},
			},
		},
		{
			name:     "uint32 with uint64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint64Value{Uint64Value: 4}},
			},
		},
		{
			name:     "uint32 with int64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int64Value{Int64Value: 4}},
			},
		},
		{
			name:     "uint32 with int32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int32Value{Int32Value: 4}},
			},
		},
		{
			name:     "uint32 with float64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_DoubleValue{DoubleValue: 4.3}},
			},
		},
		{
			name:     "uint32 with float32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_FloatValue{FloatValue: 4.3}},
			},
		},
		{
			name:     "uint32 with string",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_StringValue{StringValue: "4.3"}},
			},
		},
		{
			name:      "uint32 with null",
			willFail:  false,
			isMissing: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec, err := NewUint32Record(tc.pb)
			if err != nil {
				if tc.willFail {
					if tc.willFail != rec.Missing {
						t.Fatalf("expected missing to be %t but got %t", tc.willFail, rec.Missing)
					}
					fmt.Printf("received expected error: %s\n", err.Error())
					return

				}
				t.Fatalf("error: %s\n", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			if rec.Missing != tc.isMissing {
				t.Fatalf("missing was expected to be %t but got %t", tc.isMissing, rec.Missing)
			}
		})
	}
}

func TestNewFloat32Record(t *testing.T) {
	for _, tc := range []struct {
		name      string
		willFail  bool
		isMissing bool
		pb        *TimeRecord
	}{
		{
			name:     "basic float32",
			willFail: false,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_FloatValue{FloatValue: 4.3}},
			},
		},
		{
			name:     "float32 with uint32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint32Value{Uint32Value: 4}},
			},
		},
		{
			name:     "float32 with uint64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint64Value{Uint64Value: 4}},
			},
		},
		{
			name:     "float32 with  int64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int64Value{Int64Value: 4}},
			},
		},
		{
			name:     "float32 with int32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int32Value{Int32Value: 4}},
			},
		},
		{
			name:     "float32 with float64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_DoubleValue{DoubleValue: 4.3}},
			},
		},
		{
			name:     "float32 with string",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_StringValue{StringValue: "4.3"}},
			},
		},
		{
			name:      "float32 with null",
			willFail:  false,
			isMissing: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec, err := NewFloat32Record(tc.pb)
			if err != nil {
				if tc.willFail {
					if tc.willFail != rec.Missing {
						t.Fatalf("expected missing to be %t but got %t", tc.willFail, rec.Missing)
					}
					fmt.Printf("received expected error: %s\n", err.Error())
					return

				}
				t.Fatalf("error: %s\n", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			if rec.Missing != tc.isMissing {
				t.Fatalf("missing was expected to be %t but got %t", tc.isMissing, rec.Missing)
			}
		})
	}
}

func TestNewFloat64Record(t *testing.T) {
	for _, tc := range []struct {
		name      string
		willFail  bool
		isMissing bool
		pb        *TimeRecord
	}{
		{
			name:     "basic float64",
			willFail: false,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_DoubleValue{DoubleValue: 4.3}},
			},
		},
		{
			name:     "float64 with float32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_FloatValue{FloatValue: 4.3}},
			},
		},
		{
			name:     "float64 with uint32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint32Value{Uint32Value: 4}},
			},
		},
		{
			name:     "float64 with uint64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint64Value{Uint64Value: 4}},
			},
		},
		{
			name:     "float64 with  int64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int64Value{Int64Value: 4}},
			},
		},
		{
			name:     "float64 with int32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int32Value{Int32Value: 4}},
			},
		},
		{
			name:     "float64 with string",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_StringValue{StringValue: "4.3"}},
			},
		},
		{
			name:      "float64 with null",
			willFail:  false,
			isMissing: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec, err := NewFloat64Record(tc.pb)
			if err != nil {
				if tc.willFail {
					if tc.willFail != rec.Missing {
						t.Fatalf("expected missing to be %t but got %t", tc.willFail, rec.Missing)
					}
					fmt.Printf("received expected error: %s\n", err.Error())
					return

				}
				t.Fatalf("error: %s\n", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			if rec.Missing != tc.isMissing {
				t.Fatalf("missing was expected to be %t but got %t", tc.isMissing, rec.Missing)
			}
		})
	}
}

func TestNewStringRecord(t *testing.T) {
	for _, tc := range []struct {
		name      string
		willFail  bool
		isMissing bool
		pb        *TimeRecord
	}{
		{
			name:     "string with float64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_DoubleValue{DoubleValue: 4.3}},
			},
		},
		{
			name:     "string with float32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_FloatValue{FloatValue: 4.3}},
			},
		},
		{
			name:     "string with uint32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint32Value{Uint32Value: 4}},
			},
		},
		{
			name:     "string with uint64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Uint64Value{Uint64Value: 4}},
			},
		},
		{
			name:     "string with  int64",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int64Value{Int64Value: 4}},
			},
		},
		{
			name:     "string with int32",
			willFail: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_Int32Value{Int32Value: 4}},
			},
		},
		{
			name:     "basic string",
			willFail: false,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_StringValue{StringValue: "4.3"}},
			},
		},
		{
			name:      "string with null",
			willFail:  false,
			isMissing: true,
			pb: &TimeRecord{
				Datetime: ptypes.TimestampNow(),
				Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
			},
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			rec, err := NewStrRecord(tc.pb)
			if err != nil {
				if tc.willFail {
					if tc.willFail != rec.Missing {
						t.Fatalf("expected missing to be %t but got %t", tc.willFail, rec.Missing)
					}
					fmt.Printf("received expected error: %s\n", err.Error())
					return

				}
				t.Fatalf("error: %s\n", err)
			}
			if tc.willFail {
				t.Fatalf("expected error but was nil")
			}
			if rec.Missing != tc.isMissing {
				t.Fatalf("missing was expected to be %t but got %t", tc.isMissing, rec.Missing)
			}
		})
	}
}

func TestNewTimeRecord2(t *testing.T) {
	v := &TimeRecord{
		Datetime: ptypes.TimestampNow(),
		Value:    &Value{Kind: &Value_NullValue{NullValue: NullValue_NULL_VALUE}},
	}
	var buf bytes.Buffer
	m := jsonpb.Marshaler{
		EnumsAsInts:  false,
		EmitDefaults: false,
		Indent: "	",
		OrigName:    false,
		AnyResolver: nil,
	}
	if err := m.Marshal(&buf, v); err != nil {
		t.Fatalf(err.Error())
	}
	if err := ioutil.WriteFile("test.json", buf.Bytes(), 0644); err != nil {
		t.Fatalf(err.Error())
	}
}
