package ptypes

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
	"reflect"
)

func getNativeValue(v *structpb.Value) (val interface{}) {
	switch v.Kind.(type) {
	case *structpb.Value_NullValue:
		return nil
	case *structpb.Value_NumberValue:
		return v.GetNumberValue()
	case *structpb.Value_StringValue:
		return v.GetStringValue()
	case *structpb.Value_BoolValue:
		return v.GetBoolValue()
	case *structpb.Value_StructValue:
		// recurse
		m, _ := GetMap(v.GetStructValue())
		return m
	case *structpb.Value_ListValue:
		var list []interface{}
		for _, lv := range v.GetListValue().Values {
			list = append(list, getNativeValue(lv))
		}
		return list
	}
	return false
}

func GetMap(s *structpb.Struct) (map[string]interface{}, error) {
	m := make(map[string]interface{})
	if s != nil {
		for k, v := range s.Fields {
			m[k] = getNativeValue(v)
		}
	}
	return m, nil
}

// ToStruct converts a map[string]interface{} to a ptypes.Struct
func ToStruct(v map[string]interface{}) *structpb.Struct {
	size := len(v)
	if size == 0 {
		return nil
	}
	fields := make(map[string]*structpb.Value, size)
	for k, v := range v {
		fields[k] = toValue(v)
	}
	return &structpb.Struct{
		Fields: fields,
	}
}

// ToValue converts an interface{} to a ptypes.Value
func toValue(v interface{}) *structpb.Value {
	switch v := v.(type) {
	case nil:
		return nil
	case bool:
		return &structpb.Value{
			Kind: &structpb.Value_BoolValue{
				BoolValue: v,
			},
		}
	case int:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int8:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int32:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint8:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint32:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float32:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: v,
			},
		}
	case string:
		return &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: v,
			},
		}
	case error:
		return &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: v.Error(),
			},
		}
	default:
		// Fallback to reflection for other types
		return toValue(reflect.ValueOf(v))
	}
}
