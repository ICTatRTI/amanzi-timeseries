package ptypes

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"time"
)

type arraysNotMatchError struct {
	timeLength      int
	valuesLength    int
	qualifierLength int
}

type valuesNotAllSameTypeError struct {
	val   interface{}
	tsId  string
	index int
}

func (v valuesNotAllSameTypeError) Error() string {
	if v.tsId == "" {
		v.tsId = "unknown"
	}
	return fmt.Sprintf("ptypes: timeseries %s has inconsistent types at index %d value %v", v.tsId, v.index, v.val)
}
func (a arraysNotMatchError) Error() string {
	return fmt.Sprintf("ptypes: array lengths do not match\n times %d\n values %d\n qualifiers %d", a.timeLength, a.valuesLength, a.qualifierLength)
}

//
func NewTimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []interface{}, qualifiers [][]string) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)

	if timesL != valuesL || timesL != qualL {
		return nil, arraysNotMatchError{
			timeLength:      timesL,
			valuesLength:    valuesL,
			qualifierLength: qualL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], qualifiers[i]...)
		if err != nil {
			return nil, err
		}
		data = append(data, record)
	}
	return &TimeSeries{
		MetaInfo: meta,
		Data:     data,
	}, nil
}

func IterInt32TimeSeries(series *TimeSeries) (error, int, []time.Time, []int32, [][]string) {
	l := len(series.Data)
	times := make([]time.Time, 0, l)
	values := make([]int32, 0, l)
	qualifiers := make([][]string, 0, l)
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err, l, nil, nil, nil
		}
		if record.Value == nil {
			record.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
		}

		switch typ := record.Value.Kind.(type) {
		case *Value_Int32Value:
			values = append(values, typ.Int32Value)
		case *Value_NullValue:
			var nilVal int32
			values = append(values, nilVal)
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}, l, nil, nil, nil
		}
		times = append(times, t)
		qualifiers = append(qualifiers, record.Qualifiers)
	}

	return nil, l, times, values, qualifiers
}

func IterInt64TimeSeries(series *TimeSeries) (error, int, []time.Time, []int64, [][]string) {
	l := len(series.Data)
	times := make([]time.Time, 0, l)
	values := make([]int64, 0, l)
	qualifiers := make([][]string, 0, l)
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err, l, nil, nil, nil
		}
		switch typ := record.Value.Kind.(type) {
		case *Value_Int64Value:
			values = append(values, typ.Int64Value)
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}, l, nil, nil, nil
		}
		times = append(times, t)
		qualifiers = append(qualifiers, record.Qualifiers)
	}

	return nil, l, times, values, qualifiers
}

func IterUint32TimeSeries(series *TimeSeries) (error, int, []time.Time, []uint32, [][]string) {
	l := len(series.Data)
	times := make([]time.Time, 0, l)
	values := make([]uint32, 0, l)
	qualifiers := make([][]string, 0, l)
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err, l, nil, nil, nil
		}
		switch typ := record.Value.Kind.(type) {
		case *Value_Uint32Value:
			values = append(values, typ.Uint32Value)
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}, l, nil, nil, nil
		}
		times = append(times, t)
		qualifiers = append(qualifiers, record.Qualifiers)
	}

	return nil, l, times, values, qualifiers
}

func IterUint64TimeSeries(series *TimeSeries) (error, int, []time.Time, []uint64, [][]string) {
	l := len(series.Data)
	times := make([]time.Time, 0, l)
	values := make([]uint64, 0, l)
	qualifiers := make([][]string, 0, l)
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err, l, nil, nil, nil
		}
		switch typ := record.Value.Kind.(type) {
		case *Value_Uint64Value:
			values = append(values, typ.Uint64Value)
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}, l, nil, nil, nil
		}
		times = append(times, t)
		qualifiers = append(qualifiers, record.Qualifiers)
	}

	return nil, l, times, values, qualifiers
}

func IterFloat32TimeSeries(series *TimeSeries) (error, int, []time.Time, []float32, [][]string) {
	l := len(series.Data)
	times := make([]time.Time, 0, l)
	values := make([]float32, 0, l)
	qualifiers := make([][]string, 0, l)
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err, l, nil, nil, nil
		}
		switch typ := record.Value.Kind.(type) {
		case *Value_FloatValue:
			values = append(values, typ.FloatValue)
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}, l, nil, nil, nil
		}
		times = append(times, t)
		qualifiers = append(qualifiers, record.Qualifiers)
	}

	return nil, l, times, values, qualifiers
}

func IterFloat64TimeSeries(series *TimeSeries) (error, int, []time.Time, []float64, [][]string) {
	l := len(series.Data)
	times := make([]time.Time, 0, l)
	values := make([]float64, 0, l)
	qualifiers := make([][]string, 0, l)
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return err, l, nil, nil, nil
		}
		switch typ := record.Value.Kind.(type) {
		case *Value_DoubleValue:
			values = append(values, typ.DoubleValue)
		default:
			return valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}, l, nil, nil, nil
		}
		times = append(times, t)
		qualifiers = append(qualifiers, record.Qualifiers)
	}

	return nil, l, times, values, qualifiers
}

type properties map[string]interface{}

type TimeSeriesIterator interface {
	Len() int
	Times() []time.Time
	Next() bool
	Qualifiers() [][]string
	Properties() []properties
	GetProperties(idx int) properties
	GetQualifiers(idx int) []string
	GetTime(idx int) time.Time
}

type TimeSeriesStringIterator interface {
	TimeSeriesIterator
	GetValue(idx int) string
	Values() []string
}

type TimeSeriesData struct {
	len        int
	itrCount   int
	times      []time.Time
	qualifiers [][]string
	properties []properties
}

func (t TimeSeriesData) Times() []time.Time {
	return t.times
}

func (t TimeSeriesData) Qualifiers() [][]string {
	return t.qualifiers
}

func (t TimeSeriesData) Properties() []properties {
	return t.properties
}

func (t TimeSeriesData) Len() int {
	return t.len
}

func (t TimeSeriesData) GetTime(idx int) time.Time {
	return t.times[idx]
}

func (t TimeSeriesData) GetQualifiers(idx int) []string {
	return t.qualifiers[idx]
}

func (t TimeSeriesData) GetProperties(idx int) properties {
	return t.properties[idx]
}

type TimeSeriesStringData struct {
	TimeSeriesData
	values []string
}

func (t TimeSeriesStringData) GetValue(idx int) string {
	return t.values[idx]
}

func (t TimeSeriesStringData) Values() []string {
	return t.values
}

func IterStringTimeSeries(series *TimeSeries) (*TimeSeriesStringData, error) {
	l := len(series.Data)
	data := &TimeSeriesStringData{
		TimeSeriesData: TimeSeriesData{
			len:        l,
			times:      make([]time.Time, 0, l),
			qualifiers: make([][]string, 0, l),
			properties: make([]properties, 0, l),
		},
		values: make([]string, 0, l),
	}
	for idx, record := range series.Data {
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return nil, err
		}

		if record.Value == nil {
			record.Value = &Value{Kind: &Value_NullValue{NullValue_NULL_VALUE}}
		}
		switch typ := record.Value.Kind.(type) {
		case *Value_StringValue:
			data.values = append(data.values, typ.StringValue)
		case *Value_NullValue:
			var nilString string
			data.values = append(data.values, nilString)
		default:
			return nil, valuesNotAllSameTypeError{
				val:   record.Value,
				tsId:  series.MetaInfo.Id,
				index: idx,
			}
		}
		data.times = append(data.times, t)
		data.qualifiers = append(data.qualifiers, record.Qualifiers)
	}

	return data, nil
}
