package ptypes

import (
	"fmt"
	"time"
)

type arraysNotMatchError struct {
	timeLength       int
	valuesLength     int
	qualifierLength  int
	propertiesLength int
}

func (a arraysNotMatchError) Error() string {
	return fmt.Sprintf("ptypes: array lengths do not match\n times %d\n values %d\n qualifiers %d\n properties %d\n", a.timeLength, a.valuesLength, a.qualifierLength, a.propertiesLength)
}

//
func NewIntTimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []int, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewInt32TimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []int32, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewUintTimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []uint, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewInt64TimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []int64, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewUint64TimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []uint64, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewUint32TimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []uint32, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewStringTimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []string, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewFloat32TimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []float32, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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

func NewFloat64TimeSeries(meta *TimeSeriesMetaInfo, times []time.Time, values []float64, qualifiers [][]string, properties []Properties) (*TimeSeries, error) {
	timesL := len(times)
	valuesL := len(values)
	qualL := len(qualifiers)
	propL := len(properties)
	lengthsNotEqual := !(timesL == valuesL && valuesL == qualL && qualL == propL)
	if lengthsNotEqual {
		return nil, arraysNotMatchError{
			timeLength:       timesL,
			valuesLength:     valuesL,
			qualifierLength:  qualL,
			propertiesLength: propL,
		}
	}
	data := make([]*TimeRecord, 0, len(values))
	for i := 0; i < timesL; i++ {
		record, err := NewTimeRecord(times[i], values[i], properties[i], qualifiers[i]...)
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
