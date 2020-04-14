package ptypes

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type Properties map[string]interface{}

func NewTimeRecord(time time.Time, i interface{}, p Properties, q ...string) (*TimeRecord, error) {
	pbt, err := ptypes.TimestampProto(time)
	if err != nil {
		return nil, err
	}
	var kind isValue_Kind
	switch v := i.(type) {
	case int32:
		kind = &Value_Int32Value{Int32Value: v}
	case int64:
		kind = &Value_Int64Value{Int64Value: v}
	case int:
		kind = &Value_Int64Value{Int64Value: int64(v)}
	case float64:
		kind = &Value_DoubleValue{DoubleValue: v}
	case float32:
		kind = &Value_FloatValue{FloatValue: v}
	case nil:
		kind = &Value_NullValue{NullValue: NullValue_NULL_VALUE}
	case uint32:
		kind = &Value_Uint32Value{Uint32Value: v}
	case uint64:
		kind = &Value_Uint64Value{Uint64Value: v}
	case uint:
		kind = &Value_Uint64Value{Uint64Value: uint64(v)}
	case string:
		kind = &Value_StringValue{StringValue: v}
	default:
		kind = &Value_NullValue{NullValue: NullValue_NULL_VALUE}
	}
	return &TimeRecord{
		Datetime:   pbt,
		Value:      &Value{Kind: kind},
		Qualifiers: q,
		Properties: ToStruct(p),
	}, nil
}

type Record struct {
	Time       time.Time
	Qualifiers []string
	Properties Properties
	Missing    bool
}

type StringRecord struct {
	Record
	Value string
}

func NewStrRecord(r *TimeRecord) (StringRecord, error) {
	t, err := ptypes.Timestamp(r.Datetime)
	if err != nil {
		return StringRecord{}, err
	}
	if r.Value == nil {
		r.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
	}
	var valueStr string
	var isMissing bool
	switch typ := r.Value.Kind.(type) {
	case *Value_StringValue:
		valueStr = typ.StringValue
	case *Value_NullValue:
		isMissing = true
	default:
		isMissing = true
		return StringRecord{Record: Record{Missing: isMissing}}, fmt.Errorf("ptypes: type was not string or null")
	}
	p, err := GetMap(r.Properties)
	if err != nil {
		return StringRecord{}, err
	}
	return StringRecord{
		Record: Record{
			Time:       t,
			Properties: p,
			Qualifiers: r.Qualifiers,
			Missing:    isMissing,
		},
		Value: valueStr,
	}, nil
}

type Int32Record struct {
	Record
	Value int32
}

func NewInt32Record(r *TimeRecord) (Int32Record, error) {
	t, err := ptypes.Timestamp(r.Datetime)
	if err != nil {
		return Int32Record{}, err
	}
	if r.Value == nil {
		r.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
	}
	var val int32
	var isMissing bool
	switch typ := r.Value.Kind.(type) {
	case *Value_Int32Value:
		val = typ.Int32Value
	case *Value_NullValue:
		isMissing = true
	default:
		isMissing = true
		return Int32Record{Record: Record{Missing: isMissing}}, fmt.Errorf("ptypes: type was not int32 or null")
	}
	p, err := GetMap(r.Properties)
	if err != nil {
		return Int32Record{}, err
	}
	return Int32Record{
		Record: Record{
			Time:       t,
			Properties: p,
			Qualifiers: r.Qualifiers,
			Missing:    isMissing,
		},
		Value: val,
	}, nil
}

type Int64Record struct {
	Record
	Value int64
}

func NewInt64Record(r *TimeRecord) (Int64Record, error) {
	t, err := ptypes.Timestamp(r.Datetime)
	if err != nil {
		return Int64Record{}, err
	}
	if r.Value == nil {
		r.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
	}
	var val int64
	var isMissing bool
	switch typ := r.Value.Kind.(type) {
	case *Value_Int64Value:
		val = typ.Int64Value
	case *Value_NullValue:
		isMissing = true
	default:
		isMissing = true
		return Int64Record{Record: Record{Missing: isMissing}}, fmt.Errorf("ptypes: type was not int64 or null")
	}
	p, err := GetMap(r.Properties)
	if err != nil {
		return Int64Record{}, err
	}
	return Int64Record{
		Record: Record{
			Time:       t,
			Properties: p,
			Qualifiers: r.Qualifiers,
			Missing:    isMissing,
		},
		Value: val,
	}, nil
}

type Float32Record struct {
	Record
	Value float32
}

func NewFloat32Record(r *TimeRecord) (Float32Record, error) {
	t, err := ptypes.Timestamp(r.Datetime)
	if err != nil {
		return Float32Record{}, err
	}
	if r.Value == nil {
		r.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
	}
	var val float32
	var isMissing bool
	switch typ := r.Value.Kind.(type) {
	case *Value_FloatValue:
		val = typ.FloatValue
	case *Value_NullValue:
		isMissing = true
	default:
		isMissing = true
		return Float32Record{Record: Record{Missing: isMissing}}, fmt.Errorf("ptypes: type was not float32 or null")
	}
	p, err := GetMap(r.Properties)
	if err != nil {
		return Float32Record{}, err
	}
	return Float32Record{
		Record: Record{
			Time:       t,
			Properties: p,
			Qualifiers: r.Qualifiers,
			Missing:    isMissing,
		},
		Value: val,
	}, nil
}

type Float64Record struct {
	Record
	Value float64
}

func NewFloat64Record(r *TimeRecord) (Float64Record, error) {
	t, err := ptypes.Timestamp(r.Datetime)
	if err != nil {
		return Float64Record{}, err
	}
	if r.Value == nil {
		r.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
	}
	var val float64
	var isMissing bool
	switch typ := r.Value.Kind.(type) {
	case *Value_DoubleValue:
		val = typ.DoubleValue
	case *Value_NullValue:
		isMissing = true
	default:
		isMissing = true
		return Float64Record{Record: Record{Missing: isMissing}}, fmt.Errorf("ptypes: type was not float64 or null")
	}
	p, err := GetMap(r.Properties)
	if err != nil {
		return Float64Record{}, err
	}
	return Float64Record{
		Record: Record{
			Time:       t,
			Properties: p,
			Qualifiers: r.Qualifiers,
			Missing:    isMissing,
		},
		Value: val,
	}, nil
}

type Uint32Record struct {
	Record
	Value uint32
}

func NewUint32Record(r *TimeRecord) (Uint32Record, error) {
	t, err := ptypes.Timestamp(r.Datetime)
	if err != nil {
		return Uint32Record{}, err
	}
	if r.Value == nil {
		r.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
	}
	var val uint32
	var isMissing bool
	switch typ := r.Value.Kind.(type) {
	case *Value_Uint32Value:
		val = typ.Uint32Value
	case *Value_NullValue:
		isMissing = true
	default:
		isMissing = true
		return Uint32Record{Record: Record{Missing: isMissing}}, fmt.Errorf("ptypes: type was not uint32 or null")
	}
	p, err := GetMap(r.Properties)
	if err != nil {
		return Uint32Record{}, err
	}
	return Uint32Record{
		Record: Record{
			Time:       t,
			Properties: p,
			Qualifiers: r.Qualifiers,
			Missing:    isMissing,
		},
		Value: val,
	}, nil
}

type Uint64Record struct {
	Record
	Value uint64
}

func NewUint64Record(r *TimeRecord) (Uint64Record, error) {
	t, err := ptypes.Timestamp(r.Datetime)
	if err != nil {
		return Uint64Record{}, err
	}
	if r.Value == nil {
		r.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
	}
	var val uint64
	var isMissing bool
	switch typ := r.Value.Kind.(type) {
	case *Value_Uint64Value:
		val = typ.Uint64Value
	case *Value_NullValue:
		isMissing = true
	default:
		isMissing = true
		return Uint64Record{Record: Record{Missing: isMissing}}, fmt.Errorf("ptypes: type was not uint64 or null")
	}
	p, err := GetMap(r.Properties)
	if err != nil {
		return Uint64Record{}, err
	}
	return Uint64Record{
		Record: Record{
			Time:       t,
			Properties: p,
			Qualifiers: r.Qualifiers,
			Missing:    isMissing,
		},
		Value: val,
	}, nil
}
