package collections

import (
	apt "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type StringTsData struct {
	Times      []time.Time
	Values     []string
	Properties []apt.Properties
	Qualifiers [][]string
	Missing    map[int]struct{}
}

type Float32Slice []float32

type Float32TsData struct {
	Float32Slice
	Times      []time.Time
	Values     []float32
	Properties []apt.Properties
	Qualifiers [][]string
	Missing    map[int]struct{}
}

type Float64TsData struct {
	Times      []time.Time
	Values     []float64
	Properties []apt.Properties
	Qualifiers [][]string
	Missing    map[int]struct{}
}

type Int32TsData struct {
	Times      []time.Time
	Values     []int32
	Properties []apt.Properties
	Qualifiers [][]string
	Missing    map[int]struct{}
}

type Int64TsData struct {
	Times      []time.Time
	Values     []int64
	Properties []apt.Properties
	Qualifiers [][]string
	Missing    map[int]struct{}
}

type Uint32TSData struct {
	Times      []time.Time
	Values     []uint32
	Properties []apt.Properties
	Qualifiers [][]string
	Missing    map[int]struct{}
}
type Uint64TSData struct {
	Times      []time.Time
	Values     []uint64
	Properties []apt.Properties
	Qualifiers [][]string
	Missing    map[int]struct{}
}

func StrSlices(series *apt.TimeSeries, data *StringTsData) error {
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err
		}
		if data.Times != nil {
			data.Times = append(data.Times, t)
		}
		p, err := apt.GetMap(record.Properties)
		if err != nil {
			return err
		}
		if data.Properties != nil {
			data.Properties = append(data.Properties, p)
		}

		if data.Qualifiers != nil {
			data.Qualifiers = append(data.Qualifiers, record.Qualifiers)
		}

		if record.Value == nil {
			record.Value = &apt.Value{Kind: &apt.Value_NullValue{NullValue: apt.NullValue_NULL_VALUE}}
		}
		var val string
		switch typ := record.Value.Kind.(type) {
		case *apt.Value_StringValue:
			val = typ.StringValue
		case *apt.Value_NullValue:
			data.Missing[idx] = struct{}{}
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		if data.Values != nil {
			data.Values = append(data.Values, val)
		}
	}
	return nil
}

func Int32Slices(series *apt.TimeSeries, data *Int32TsData) error {
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err
		}
		if data.Times != nil {
			data.Times = append(data.Times, t)
		}
		p, err := apt.GetMap(record.Properties)
		if err != nil {
			return err
		}
		if data.Properties != nil {
			data.Properties = append(data.Properties, p)
		}

		if data.Qualifiers != nil {
			data.Qualifiers = append(data.Qualifiers, record.Qualifiers)
		}

		if record.Value == nil {
			record.Value = &apt.Value{Kind: &apt.Value_NullValue{NullValue: apt.NullValue_NULL_VALUE}}
		}
		var val int32
		switch typ := record.Value.Kind.(type) {
		case *apt.Value_Int32Value:
			val = typ.Int32Value
		case *apt.Value_NullValue:
			data.Missing[idx] = struct{}{}
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		if data.Values != nil {
			data.Values = append(data.Values, val)
		}
	}
	return nil
}

func Int64Slices(series *apt.TimeSeries, data *Int64TsData) error {
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err
		}
		if data.Times != nil {
			data.Times = append(data.Times, t)
		}
		p, err := apt.GetMap(record.Properties)
		if err != nil {
			return err
		}
		if data.Properties != nil {
			data.Properties = append(data.Properties, p)
		}

		if data.Qualifiers != nil {
			data.Qualifiers = append(data.Qualifiers, record.Qualifiers)
		}

		if record.Value == nil {
			record.Value = &apt.Value{Kind: &apt.Value_NullValue{NullValue: apt.NullValue_NULL_VALUE}}
		}
		var val int64
		switch typ := record.Value.Kind.(type) {
		case *apt.Value_Int64Value:
			val = typ.Int64Value
		case *apt.Value_NullValue:
			data.Missing[idx] = struct{}{}
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		if data.Values != nil {
			data.Values = append(data.Values, val)
		}
	}
	return nil
}

func Uint32Slices(series *apt.TimeSeries, data *Uint32TSData) error {
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err
		}
		if data.Times != nil {
			data.Times = append(data.Times, t)
		}
		p, err := apt.GetMap(record.Properties)
		if err != nil {
			return err
		}
		if data.Properties != nil {
			data.Properties = append(data.Properties, p)
		}

		if data.Qualifiers != nil {
			data.Qualifiers = append(data.Qualifiers, record.Qualifiers)
		}

		if record.Value == nil {
			record.Value = &apt.Value{Kind: &apt.Value_NullValue{NullValue: apt.NullValue_NULL_VALUE}}
		}
		var val uint32
		switch typ := record.Value.Kind.(type) {
		case *apt.Value_Uint32Value:
			val = typ.Uint32Value
		case *apt.Value_NullValue:
			data.Missing[idx] = struct{}{}
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		if data.Values != nil {
			data.Values = append(data.Values, val)
		}
	}
	return nil
}

func Uint64Slices(series *apt.TimeSeries, data *Uint64TSData) error {
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err
		}
		if data.Times != nil {
			data.Times = append(data.Times, t)
		}
		p, err := apt.GetMap(record.Properties)
		if err != nil {
			return err
		}
		if data.Properties != nil {
			data.Properties = append(data.Properties, p)
		}

		if data.Qualifiers != nil {
			data.Qualifiers = append(data.Qualifiers, record.Qualifiers)
		}

		if record.Value == nil {
			record.Value = &apt.Value{Kind: &apt.Value_NullValue{NullValue: apt.NullValue_NULL_VALUE}}
		}
		var val uint64
		switch typ := record.Value.Kind.(type) {
		case *apt.Value_Uint64Value:
			val = typ.Uint64Value
		case *apt.Value_NullValue:
			data.Missing[idx] = struct{}{}
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		if data.Values != nil {
			data.Values = append(data.Values, val)
		}
	}
	return nil
}

func Float32Slices(series *apt.TimeSeries, data *Float32TsData) error {
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err
		}
		if data.Times != nil {
			data.Times = append(data.Times, t)
		}
		p, err := apt.GetMap(record.Properties)
		if err != nil {
			return err
		}
		if data.Properties != nil {
			data.Properties = append(data.Properties, p)
		}

		if data.Qualifiers != nil {
			data.Qualifiers = append(data.Qualifiers, record.Qualifiers)
		}

		if record.Value == nil {
			record.Value = &apt.Value{Kind: &apt.Value_NullValue{NullValue: apt.NullValue_NULL_VALUE}}
		}
		var val float32
		switch typ := record.Value.Kind.(type) {
		case *apt.Value_FloatValue:
			val = typ.FloatValue
		case *apt.Value_NullValue:
			data.Missing[idx] = struct{}{}
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		if data.Values != nil {
			data.Values = append(data.Values, val)
		}
	}
	return nil
}

func Float64Slices(series *apt.TimeSeries, data *Float64TsData) error {
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err
		}
		if data.Times != nil {
			data.Times = append(data.Times, t)
		}
		p, err := apt.GetMap(record.Properties)
		if err != nil {
			return err
		}
		if data.Properties != nil {
			data.Properties = append(data.Properties, p)
		}

		if data.Qualifiers != nil {
			data.Qualifiers = append(data.Qualifiers, record.Qualifiers)
		}

		if record.Value == nil {
			record.Value = &apt.Value{Kind: &apt.Value_NullValue{NullValue: apt.NullValue_NULL_VALUE}}
		}
		var val float64
		switch typ := record.Value.Kind.(type) {
		case *apt.Value_DoubleValue:
			val = typ.DoubleValue
		case *apt.Value_NullValue:
			data.Missing[idx] = struct{}{}
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		if data.Values != nil {
			data.Values = append(data.Values, val)
		}
	}
	return nil
}
