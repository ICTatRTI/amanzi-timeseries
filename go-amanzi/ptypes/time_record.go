package ptypes

import (
	"github.com/golang/protobuf/ptypes"
	"time"
)

func NewTimeRecord(time time.Time, i interface{}, q ...string) (*TimeRecord, error) {
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
	}, nil
}
