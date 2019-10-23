from timeseries import amanzi_pb2
from uuid import uuid4
from google.protobuf.timestamp_pb2 import Timestamp
import datetime
import random
import isodate
from typing import Dict, AnyStr, Any, List


def meta(
        name: str,
        code: str,
        statistic: str,
        type: str,
        properties: Dict[AnyStr, Any] = None,
        source: amanzi_pb2.Source = amanzi_pb2.Source(),
        source_location: amanzi_pb2.SourceLocation = amanzi_pb2.SourceLocation(),
        parameter: amanzi_pb2.Parameter = amanzi_pb2.Parameter(),
        time_info: amanzi_pb2.TimeInfo = amanzi_pb2.TimeInfo(),
        origin: amanzi_pb2.Origin = amanzi_pb2.Origin(),
        groups: List[amanzi_pb2.Group] = None
) -> amanzi_pb2.TimeSeriesMetaInfo:
    m = amanzi_pb2.TimeSeriesMetaInfo()
    m.id = str(uuid4())
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


def generate_timeseries(meta_info: amanzi_pb2.TimeSeriesMetaInfo) -> amanzi_pb2.TimeSeries:
    ts = amanzi_pb2.TimeSeries()
    ts.metaInfo.CopyFrom(meta_info)

    current = datetime.datetime.utcfromtimestamp(ts.metaInfo.timeInfo.start.seconds)
    end = ts.metaInfo.timeInfo.end.seconds
    duration = isodate.parse_duration(ts.metaInfo.timeInfo.interval)
    while current.timestamp() <= end:
        record = ts.data.add()
        record_dt_pb = Timestamp()
        record_dt_pb.seconds = int(current.timestamp())

        record.datetime.CopyFrom(record_dt_pb)
        record.qualifiers.extend("p")
        record.value.double_value = random.random()
        current = current + duration
    return ts


def generate_timeseries2(meta_info: amanzi_pb2.TimeSeriesMetaInfo) -> amanzi_pb2.TimeSeries:
    ts = amanzi_pb2.TimeSeries()
    ts.metaInfo.CopyFrom(meta_info)

    current = ts.metaInfo.timeInfo.start.ToDatetime()
    end = ts.metaInfo.timeInfo.end.ToDatetime()
    duration = isodate.parse_duration(ts.metaInfo.timeInfo.interval)
    while current <= end:
        record = ts.data.add()
        record.datetime.FromDatetime(current)
        record.qualifiers.extend("p")
        record.value.string_value = str(random.random() * 100)
        current = current + duration
    return ts


