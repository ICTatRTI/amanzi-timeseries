package collections

import (
	"fmt"
	aptypes "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/ptypes"
	"time"
)

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

type TimeSeriesIterator interface {
	Next() bool
	Reset()
	StrRecord() aptypes.StringRecord
	Int32Record() aptypes.Int32Record
	Int64Record() aptypes.Int64Record
	Uint32Record() aptypes.Uint32Record
	Uint64Record() aptypes.Uint64Record
	Float32Record() aptypes.Float32Record
	Float64Record() aptypes.Float64Record
	Len() int
}
type LazyIterator struct {
	current int
	meta    *aptypes.TimeSeriesMetaInfo
	data    []*aptypes.TimeRecord
	err     error
}

func NewLazyIterator(series *aptypes.TimeSeries) *LazyIterator {
	return &LazyIterator{
		current: -1,
		meta:    series.MetaInfo,
		data:    series.Data,
		err:     nil,
	}
}

func (itr *LazyIterator) StrRecord() aptypes.StringRecord {
	record := itr.data[itr.current]
	strRecord, err := aptypes.NewStrRecord(record)
	if err != nil {
		itr.err = fmt.Errorf("go-amanzi: %w occured on index %d", err, itr.current)
	}
	return strRecord
}

func (itr *LazyIterator) Int32Record() aptypes.Int32Record {
	record := itr.data[itr.current]
	int32Record, err := aptypes.NewInt32Record(record)
	if err != nil {
		itr.err = fmt.Errorf("go-amanzi: %w occured on index %d", err, itr.current)
	}
	return int32Record
}

func (itr *LazyIterator) Int64Record() aptypes.Int64Record {
	record := itr.data[itr.current]
	int64Record, err := aptypes.NewInt64Record(record)
	if err != nil {
		itr.err = fmt.Errorf("go-amanzi: %w occured on index %d", err, itr.current)
	}
	return int64Record
}

func (itr *LazyIterator) Uint32Record() aptypes.Uint32Record {
	record := itr.data[itr.current]
	uint32Record, err := aptypes.NewUint32Record(record)
	if err != nil {
		itr.err = fmt.Errorf("go-amanzi: %w occured on index %d", err, itr.current)
	}
	return uint32Record
}

func (itr *LazyIterator) Uint64Record() aptypes.Uint64Record {
	record := itr.data[itr.current]
	uint64Record, err := aptypes.NewUint64Record(record)
	if err != nil {
		itr.err = fmt.Errorf("go-amanzi: %w occured on index %d", err, itr.current)
	}
	return uint64Record
}

func (itr *LazyIterator) Float32Record() aptypes.Float32Record {
	record := itr.data[itr.current]
	float32Record, err := aptypes.NewFloat32Record(record)
	if err != nil {
		itr.err = fmt.Errorf("go-amanzi: %w occured on index %d", err, itr.current)
	}
	return float32Record
}

func (itr *LazyIterator) Float64Record() aptypes.Float64Record {
	record := itr.data[itr.current]
	float64Record, err := aptypes.NewFloat64Record(record)
	if err != nil {
		itr.err = fmt.Errorf("go-amanzi: %w occured on index %d", err, itr.current)
	}
	return float64Record
}

func (itr *LazyIterator) Err() error {
	return itr.err
}
func (itr *LazyIterator) Reset() {
	itr.current = -1
}
func (itr *LazyIterator) Len() int {
	return len(itr.data)
}
func (itr *LazyIterator) Next() bool {
	itr.current++
	if itr.current >= len(itr.data) {
		return false
	}
	return true
}

type timeSeriesData struct {
	time         time.Time
	qualifier    []string
	missing      bool
	properties   aptypes.Properties
	strValue     string
	int32Value   int32
	int64Value   int64
	uint32Value  uint32
	uint64Value  uint64
	float32Value float32
	float64Value float64
}

type EagerIterator struct {
	current int
	data    []timeSeriesData
}

func NewEagerIterator(series *aptypes.TimeSeries) (*EagerIterator, error) {
	l := len(series.Data)
	itr := &EagerIterator{
		current: -1,
		data:    make([]timeSeriesData, 0, l),
	}
	for idx, record := range series.Data {
		var val timeSeriesData
		t, err := ptypes.Timestamp(record.Datetime)
		if err != nil {
			return nil, err
		}
		val.time = t

		val.qualifier = record.Qualifiers
		m, err := aptypes.GetMap(record.Properties)
		if err != nil {
			return nil, err
		}
		val.properties = m
		if record.Value == nil {
			record.Value = &aptypes.Value{Kind: &aptypes.Value_NullValue{NullValue: aptypes.NullValue_NULL_VALUE}}
		}
		switch typ := record.Value.Kind.(type) {
		case *aptypes.Value_Int32Value:
			val.int32Value = typ.Int32Value
		case *aptypes.Value_Int64Value:
			val.int64Value = typ.Int64Value
		case *aptypes.Value_FloatValue:
			val.float32Value = typ.FloatValue
		case *aptypes.Value_DoubleValue:
			val.float64Value = typ.DoubleValue
		case *aptypes.Value_Uint32Value:
			val.uint32Value = typ.Uint32Value
		case *aptypes.Value_Uint64Value:
			val.uint64Value = typ.Uint64Value
		case *aptypes.Value_StringValue:
			val.strValue = typ.StringValue
		case *aptypes.Value_NullValue:
			val.missing = true
		default:
			panic(fmt.Errorf("go-amanzi: unkown type %v on index %d", typ, idx))
		}
		itr.data = append(itr.data, val)
	}
	return itr, nil
}

func (e *EagerIterator) Next() bool {
	e.current++
	if e.current >= len(e.data) {
		return false
	}
	return true
}

func (e *EagerIterator) Reset() {
	e.current = -1
}

func (e *EagerIterator) Len() int {
	return len(e.data)
}

func (e *EagerIterator) StrRecord() aptypes.StringRecord {
	return aptypes.StringRecord{
		Record: aptypes.Record{
			Time:       e.data[e.current].time,
			Qualifiers: e.data[e.current].qualifier,
			Properties: e.data[e.current].properties,
			Missing:    e.data[e.current].missing,
		},
		Value: e.data[e.current].strValue,
	}
}

func (e *EagerIterator) Int32Record() aptypes.Int32Record {
	return aptypes.Int32Record{
		Record: aptypes.Record{
			Time:       e.data[e.current].time,
			Qualifiers: e.data[e.current].qualifier,
			Properties: e.data[e.current].properties,
			Missing:    e.data[e.current].missing,
		},
		Value: 0,
	}
}

func (e *EagerIterator) Int64Record() aptypes.Int64Record {
	return aptypes.Int64Record{
		Record: aptypes.Record{
			Time:       e.data[e.current].time,
			Qualifiers: e.data[e.current].qualifier,
			Properties: e.data[e.current].properties,
			Missing:    e.data[e.current].missing,
		},
		Value: e.data[e.current].int64Value,
	}
}

func (e *EagerIterator) Uint32Record() aptypes.Uint32Record {
	return aptypes.Uint32Record{
		Record: aptypes.Record{
			Time:       e.data[e.current].time,
			Qualifiers: e.data[e.current].qualifier,
			Properties: e.data[e.current].properties,
			Missing:    e.data[e.current].missing,
		},
		Value: e.data[e.current].uint32Value,
	}
}

func (e *EagerIterator) Uint64Record() aptypes.Uint64Record {
	return aptypes.Uint64Record{
		Record: aptypes.Record{
			Time:       e.data[e.current].time,
			Qualifiers: e.data[e.current].qualifier,
			Properties: e.data[e.current].properties,
			Missing:    e.data[e.current].missing,
		},
		Value: e.data[e.current].uint64Value,
	}
}

func (e *EagerIterator) Float32Record() aptypes.Float32Record {
	return aptypes.Float32Record{
		Record: aptypes.Record{
			Time:       e.data[e.current].time,
			Qualifiers: e.data[e.current].qualifier,
			Properties: e.data[e.current].properties,
			Missing:    e.data[e.current].missing,
		},
		Value: e.data[e.current].float32Value,
	}
}

func (e *EagerIterator) Float64Record() aptypes.Float64Record {
	return aptypes.Float64Record{
		Record: aptypes.Record{
			Time:       e.data[e.current].time,
			Qualifiers: e.data[e.current].qualifier,
			Properties: e.data[e.current].properties,
			Missing:    e.data[e.current].missing,
		},
		Value: e.data[e.current].float64Value,
	}
}
