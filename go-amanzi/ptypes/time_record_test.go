package ptypes

import (
	"bytes"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
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
		willFail bool
	}{
		{
			name:     "int32 test",
			t:        time.Now(),
			val:      int32(3),
			q:        []string{},
			willFail: false,
		},
		{
			name:     "int test",
			t:        time.Now(),
			val:      int32(3),
			q:        []string{},
			willFail: false,
		},
		{
			name:     "int64 test",
			t:        time.Now(),
			val:      int64(3),
			q:        []string{},
			willFail: false,
		},
		{
			name:     "uint32 test",
			t:        time.Now(),
			val:      uint32(3),
			q:        []string{},
			willFail: false,
		},
		{
			name:     "int test",
			t:        time.Now(),
			val:      uint32(3),
			q:        []string{},
			willFail: false,
		},
		{
			name:     "uint64 test",
			t:        time.Now(),
			val:      uint64(3),
			q:        []string{},
			willFail: false,
		},
		{
			name:     "null test",
			t:        time.Now(),
			val:      nil,
			q:        []string{},
			willFail: false,
		},
		{
			name:     "double test",
			t:        time.Now(),
			val:      3.5,
			q:        []string{},
			willFail: false,
		},
		{
			name:     "float test",
			t:        time.Now(),
			val:      float32(3.5),
			q:        []string{},
			willFail: false,
		},
		{
			name:     "string test",
			t:        time.Now(),
			val:      "string",
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
			q:        []string{},
			willFail: true,
		},
		{
			name: "map test",
			t:    time.Now(),
			val: map[string]interface{}{
				"test": 5,
			},
			q:        []string{},
			willFail: true,
		},
		{
			name:     "slice test",
			t:        time.Now(),
			val:      []interface{}{5.3, "biscuit", 4},
			willFail: true,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			record, err := NewTimeRecord(tc.t, tc.val, tc.q...)
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
		OrigName:    true,
		AnyResolver: nil,
	}
	if err := m.Marshal(&buf, v); err != nil {
		t.Fatalf(err.Error())
	}
	if err := ioutil.WriteFile("test.json", buf.Bytes(), 0644); err != nil {
		t.Fatalf(err.Error())
	}
}

func TestNewTimeRecord_NoValueSet(t *testing.T) {
	b, err := ioutil.ReadFile("test.json")
	if err != nil {
		t.Fail()
	}

	var record TimeRecord
	if err := proto.Unmarshal(b, &record); err != nil {
		t.Fail()
	}
	if record.Value == nil {

	}
	switch v := record.Value.Kind.(type) {
	case *Value_Uint64Value:
	case *Value_Uint32Value:
	case *Value_Int32Value:
	case *Value_Int64Value:
	case *Value_DoubleValue:
	case *Value_FloatValue:
	case *Value_NullValue:
	case *Value_StringValue:
	default:
		t.Fatalf("failed : %v", v)
	}

}
