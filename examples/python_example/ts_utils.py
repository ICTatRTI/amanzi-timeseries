from timeseries import amanzi_pb2
import pandas as pd
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson, Parse
import datetime
import isodate
from uuid import uuid4
from typing import Dict, AnyStr, Any, List


def create_source(name: str,
                  code: str
                  ) -> amanzi_pb2.Source:
    """Create a source message Amanzi TS metaInfo

    :param name:
    :param code:
    :return:
    """
    source = amanzi_pb2.Source()
    source.name = name
    source.code = code
    return source


def create_elevation(value: float,
                     units: str,
                     datum: str
                     ) -> amanzi_pb2.Elevation:
    """Create elevation message for Amanzi TS metaInfo

    :param value:
    :param units:
    :param datum:
    :return:
    """
    elevation = amanzi_pb2.Elevation()
    elevation.value = value
    elevation.units = units
    elevation.datum = datum
    return elevation


def create_source_location(name: str,
                           code: str,
                           elevation: amanzi_pb2.Elevation = amanzi_pb2.Elevation()
                           ) -> amanzi_pb2.SourceLocation:
    """Create source location message for Amanzi TS metaInfo

    :param name:
    :param code:
    :param elevation:
    :return:
    """
    source_location = amanzi_pb2.SourceLocation()
    source_location.name = name
    source_location.code = code
    source_location.elevation.CopyFrom(elevation)


def create_parameter(name: str,
                     code: str,
                     units: str
                     ) -> amanzi_pb2.Parameter:
    """Create parameter message for Amanzi TS metaInfo

    :param name:
    :param code:
    :param units:
    :return:
    """
    parameter = amanzi_pb2.Parameter()
    parameter.name = name
    parameter.code = code
    parameter.units = units
    return parameter


def create_time_info(reference: datetime.datetime,
                     start: datetime.datetime,
                     end: datetime.datetime,
                     windowed_interval:  datetime.timedelta,
                     interval:  datetime.timedelta
                     ) -> amanzi_pb2.TimeInfo:
    """Create time info message

    :param reference:
    :param start:
    :param end:
    :param windowed_interval:
    :param interval:
    :return:
    """

    interval_str = isodate.duration_isoformat(interval)
    windowed_interval_str = isodate.duration_isoformat(windowed_interval)

    timeInfo = amanzi_pb2.TimeInfo()
    timeInfo.referenceTime.FromDatetime(reference)
    timeInfo.start.FromDatetime(start)
    timeInfo.end.FromDatetime(end)
    timeInfo.interval = interval_str
    timeInfo.windowed_interval = windowed_interval_str
    return timeInfo


def create_origin(description: str,
                  reference_ids: List[str]
                  ) -> amanzi_pb2.Origin:
    """Create origin message

    :param description:
    :param reference_ids:
    :return:
    """
    origin = amanzi_pb2.Origin()
    origin.description = description
    origin.referenceIds.extend(reference_ids)


def create_meta(
        id: str = None,
        name: str = None,
        code: str = None,
        statistic: str = None,
        type: str = None,
        properties: Dict[AnyStr, Any] = {},
        source: amanzi_pb2.Source = amanzi_pb2.Source(),
        source_location: amanzi_pb2.SourceLocation = amanzi_pb2.SourceLocation(),
        parameter: amanzi_pb2.Parameter = amanzi_pb2.Parameter(),
        time_info: amanzi_pb2.TimeInfo = amanzi_pb2.TimeInfo(),
        origin: amanzi_pb2.Origin = amanzi_pb2.Origin(),
        groups: List[amanzi_pb2.Group] = []
) -> amanzi_pb2.TimeSeriesMetaInfo:
    """Create meta info object
    """
    m = amanzi_pb2.TimeSeriesMetaInfo()
    if not id:
        id = str(uuid4())

    m.id = id
    m.statistic = statistic
    m.type = type
    m.name = name
    m.code = code
    m.properties.update(properties)
    m.source.CopyFrom(source)
    m.sourceLocation.CopyFrom(source_location)
    m.parameter.CopyFrom(parameter)
    m.timeInfo.CopyFrom(time_info)
    m.origin.CopyFrom(origin)
    m.groups.extend(groups)
    return m


def ts_to_data_arrays(ts: amanzi_pb2.TimeSeries) -> tuple:
    """Creates 2 lists from an Amanzi Timeseries.
        First list has datetime objects, second list has values
        Value type used in first timeseries message assumed for all subsequent messages

    Parameters
    ----------
    ts: an amanzi_pb2.TimeSeries object

    Return
    ------
    A tuple of two lists:
        The first is datetime objects and the second is the values
    """
    dts = []
    vals = []
    val_field = ts.data[0].value.ListFields()[0][0].name

    if val_field == 'double_value':
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            if d.value.double_value:
                vals.append(d.value.double_value)
            else:
                vals.append(None)
    elif val_field == 'float_value':
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            if d.value.float_value:
                vals.append(d.value.float_value)
            else:
                vals.append(None)
    elif val_field == 'int64_value':
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            if d.value.int64_value:
                vals.append(d.value.int64_value)
            else:
                vals.append(None)
    elif val_field == 'int32_value':
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            if d.value.int32_value:
                vals.append(d.value.int32_value)
            else:
                vals.append(None)
    elif val_field == 'uint64_value':
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            if d.value.uint64_value:
                vals.append(d.value.uint64_value)
            else:
                vals.append(None)
    elif val_field == 'uint32_value':
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            if d.value.uint32_value:
                vals.append(d.value.uint32_value)
            else:
                vals.append(None)
    elif val_field == 'string_value':
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            if d.value.string_value:
                vals.append(d.value.string_value)
            else:
                vals.append(None)

    return dts, vals


def ts_to_message_arrays(ts: amanzi_pb2.TimeSeries) -> list:
    """Takes an Amanzi Timeseries and converts it to a list of lists

    Parameters
    ----------
    ts: an amanzi_pb2.TimeSeries object

    Returns
    -------
    data: list
        A list of lists with pattern [[datetime, value],[datetime, value], ...]
    """

    data = []
    val_field = ts.data[0].value.ListFields()[0][0].name

    if val_field == 'double_value':
        for d in ts.data:
            if d.value.double_value:
                data.append([d.datetime.ToDatetime(), d.value.double_value])
            else:
                data.append([d.datetime.ToDatetime(), None])
    elif val_field == 'float_value':
        for d in ts.data:
            if d.value.float_value:
                data.append([d.datetime.ToDatetime(), d.value.float_value])
            else:
                data.append([d.datetime.ToDatetime(), None])
    elif val_field == 'int64_value':
        for d in ts.data:
            if d.value.int64_value:
                data.append([d.datetime.ToDatetime(), d.value.int64_value])
            else:
                data.append([d.datetime.ToDatetime(), None])
    elif val_field == 'int32_value':
        for d in ts.data:
            if d.value.int32_value:
                data.append([d.datetime.ToDatetime(), d.value.int32_value])
            else:
                data.append([d.datetime.ToDatetime(), None])
    elif val_field == 'uint64_value':
        for d in ts.data:
            if d.value.uint64_value:
                data.append([d.datetime.ToDatetime(), d.value.uint64_value])
            else:
                data.append([d.datetime.ToDatetime(), None])
    elif val_field == 'uint32_value':
        for d in ts.data:
            if d.value.uint32_value:
                data.append([d.datetime.ToDatetime(), d.value.uint32_value])
            else:
                data.append([d.datetime.ToDatetime(), None])
    elif val_field == 'string_value':
        for d in ts.data:
            if d.value.string_value:
                data.append([d.datetime.ToDatetime(), d.value.string_value])
            else:
                data.append([d.datetime.ToDatetime(), None])

    return data


def ts_to_df(ts: amanzi_pb2.TimeSeries) -> pd.DataFrame:
    """Amanzi protobuf timeseries to Pandas DataFrame
    :param ts: json timeseries (amanzi.ts.spec)
    :return: dataframe
    """
    dts, vals = ts_to_data_arrays(ts)
    df = pd.DataFrame({'data': vals}, index=dts)
    return df


def write_arrays_to_ts(meta: amanzi_pb2.TimeSeriesMetaInfo,
                       datetime: List[datetime.datetime],
                       values: list,
                       qualifiers: List[list] = None) -> amanzi_pb2.TimeSeries:
    """Takes meta data, a list of datetime and a list of values and returns a timesseries

    Parameters
    ----------
    meta: amanzi_pb2.TimeSeriesMetaInfo
        metaInfo that will be used for the returned TS
    datetime: List[datetime.datetime]
        List of datetime values that will be in returned TS
        Must be same length as values list
    values: list
        List of values that will be used in the returned TS
        Value type used in returned TS is selected based on type of first value in list
        int -> int64_value
        float -> float_value
        all others -> string_value
    qualifiers: List[list] = None
        Optional list of lists for qualifier data

    Returns
    -------
    ts: amanzi_pb2.TimeSeries
        an Amanzi TS that has the provided metadata and the date, value and qualifier data
    """

    ts = amanzi_pb2.TimeSeries()
    ts.metaInfo.CopyFrom(meta)

    if len(datetime) != len(values):
        raise IndexError("Length of datetime and values list must be the same")

    if qualifiers:
        if len(datetime) != len(qualifiers):
            raise IndexError("Length of qualifiers and datetime list must be the same")

    # ToDo: I'd rather check for value type once but if first value is None that doesn't work.
    # This can result in more than one value type in a TS
    for i, d in enumerate(datetime):
        record = ts.data.add()
        record.datetime.FromDatetime(d)
        if values[i]:
            if isinstance(values[0], int):
                record.value.int64_value = values[i]
            elif isinstance(values[0], float):
                record.value.float_value = values[i]
            else:
                record.value.string_value = str(values[i])
        if qualifiers:
            record.qualifiers.extend(qualifiers[i])

    return ts
