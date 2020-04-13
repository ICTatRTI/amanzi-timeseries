package ptypes

import (
	structpb "github.com/golang/protobuf/ptypes/struct"
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
	for k, v := range s.Fields {
		m[k] = getNativeValue(v)
	}
	return m, nil
}
