from timeseries import amanzi_pb2
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson, Parse
import datetime
import isodate
import timeseries_generator
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
    """Creates 2 lists of timeseries data.
        First list has datetime objects, second list has values
        Value type used in first timeseries message assumed for all subsequent messages

    :param ts: an amanzi_pb2.TimeSeries object
    :return: a tuple of two lists.  the first is datetime objects and the second is the values
    """
    dts = []
    vals = []

    if ts.data[0].value.double_value:
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            vals.append(d.value.double_value)
    elif ts.data[0].value.float_value:
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            vals.append(d.value.float_value)
    elif ts.data[0].value.int64_value:
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            vals.append(d.value.int64_value)
    elif ts.data[0].value.int32_value:
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
        vals.append(d.value.int32_value)
    elif ts.data[0].value.uint64_value:
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            vals.append(d.value.uint64_value)
    elif ts.data[0].value.uint32_value:
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            vals.append(d.value.uint32_value)
    elif ts.data[0].value.string_value:
        for d in ts.data:
            dts.append(d.datetime.ToDatetime())
            vals.append(d.value.string_value)

    return dts, vals


def ts_to_message_arrays(ts: amanzi_pb2.TimeSeries) -> list:
    """

    :param ts: an amanzi_pb2.TimeSeries object
    :return: a list of lists [[datetime, value],[datetime, value], ...]
    """

    data = []

    if ts.data[0].value.double_value:
        for d in ts.data:
            data.append([d.datetime.ToDatetime(), d.value.double_value])
    elif ts.data[0].value.float_value:
        for d in ts.data:
            data.append([d.datetime.ToDatetime(), d.value.float_value])
    elif ts.data[0].value.int64_value:
        for d in ts.data:
            data.append([d.datetime.ToDatetime(), d.value.int64_value])
    elif ts.data[0].value.int32_value:
        for d in ts.data:
            data.append([d.datetime.ToDatetime(), d.value.int32_value])
    elif ts.data[0].value.uint64_value:
        for d in ts.data:
            data.append([d.datetime.ToDatetime(), d.value.uint64_value])
    elif ts.data[0].value.uint32_value:
        for d in ts.data:
            data.append([d.datetime.ToDatetime(), d.value.uint32_value])
    elif ts.data[0].value.string_value:
        for d in ts.data:
            data.append([d.datetime.ToDatetime(), d.value.string_value])

    return data
