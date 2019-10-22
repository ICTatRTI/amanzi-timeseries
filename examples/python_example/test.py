from timeseries import amanzi_pb2
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson
import datetime
import isodate
import timeseries_generator
source = amanzi_pb2.Source()
# set simple values
source.name = "Lake Bisqcuit"
source.code = "LB"

elevation = amanzi_pb2.Elevation()
elevation.value = 100
elevation.units = "ft"
elevation.datum = "NAVD88"

source_location = amanzi_pb2.SourceLocation()
source_location.name = "corner of lake biscuit"
source_location.code = "5315315"
source_location.elevation.CopyFrom(elevation)

parameter = amanzi_pb2.Parameter()
parameter.name = "Streamflow"
parameter.code = "QIN"
parameter.units = "cfs"

reference = datetime.datetime.utcnow().timestamp()
start = datetime.datetime(year=2000,month=1,day=1,hour=0, minute=0, second=0, tzinfo=datetime.timezone.utc)
end = start + datetime.timedelta(days=10)
duration = datetime.timedelta(hours=1)
# set up timestamp pbs
reference_pb = Timestamp()
reference_pb.seconds = int(reference)
start_pb = Timestamp()
start_pb.seconds = int(start.timestamp())
end_pb = Timestamp()
end_pb.seconds = int(end.timestamp())
dur_str = isodate.duration_isoformat(duration)

timeInfo = amanzi_pb2.TimeInfo()
timeInfo.referenceTime.CopyFrom(reference_pb)
timeInfo.start.CopyFrom(start_pb)
timeInfo.end.CopyFrom(end_pb)
timeInfo.interval = dur_str

meta = timeseries_generator.meta(
    name="test",
    code="t1",
    statistic="generated",
    type="forecast",
    properties={
        "test": "foo"
    },
    source=source,
    source_location=source_location,
    parameter=parameter,
    time_info=timeInfo
)

ts = timeseries_generator.generate_timeseries(meta)


print(ts.SerializeToString())
j = MessageToJson(ts)
print(j)
# new_ts = Parse(j, amanzi_pb2.TimeSeries())
# print(new_ts)
# print(protobuf_to_dict(ts))



