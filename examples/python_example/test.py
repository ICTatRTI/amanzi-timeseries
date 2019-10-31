from timeseries import amanzi_pb2
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson, Parse
import datetime
import isodate
import timeseries_generator
from uuid import uuid4


# create and set source values
source = amanzi_pb2.Source()
source.name = "Lake Biscuit Operating Authority"
source.code = "LBOA"

# create and set elevation values
elevation = amanzi_pb2.Elevation()
elevation.value = 100
elevation.units = "ft"
elevation.datum = "NAVD88"

# create and set source_location values
source_location = amanzi_pb2.SourceLocation()
source_location.name = "Lake Biscuit Dam"
source_location.code = "5315315"
source_location.elevation.CopyFrom(elevation)

# create and set parameter values
parameter = amanzi_pb2.Parameter()
parameter.name = "Streamflow"
parameter.code = "QIN"
parameter.units = "cfs"

# create and set time info values
reference = datetime.datetime.utcnow()
start = datetime.datetime(year=2000,
                          month=1,
                          day=1,
                          hour=0,
                          minute=0,
                          second=0,
                          tzinfo=datetime.timezone.utc)
end = start + datetime.timedelta(days=2)
duration = datetime.timedelta(hours=1)
# set up timestamp pbs
dur_str = isodate.duration_isoformat(duration)

timeInfo = amanzi_pb2.TimeInfo()
timeInfo.referenceTime.FromDatetime(reference)
timeInfo.start.FromDatetime(start)
timeInfo.end.FromDatetime(end)
timeInfo.interval = dur_str

# create and set groups
groups = []
group1 = amanzi_pb2.Group()
group1.groupId = str(uuid4())
group1.description = "Ensemble 1"
groups.append(group1)

group2 = amanzi_pb2.Group()
group2.groupId = str(uuid4())
group2.description = "HMS Model Simulation"
groups.append(group2)

# create and set origin
origin = amanzi_pb2.Origin()
origin.description = "Calculated hourly mean from raw timeseries"
origin.referenceIds.extend([str(uuid4()), str(uuid4())])

# create metaInfo object with data created above
meta = timeseries_generator.meta(
    name="Lake Biscuit Release",
    code="LBR",
    statistic="Instantaneous",
    type="forecast",
    properties={
        "test": "foo"
    },
    source=source,
    source_location=source_location,
    parameter=parameter,
    time_info=timeInfo,
    groups=groups,
    origin=origin
)

# generate a random time series
ts = timeseries_generator.generate_timeseries2(meta)

# to binary string
s = ts.SerializeToString()
# print(s)

# to json
j = MessageToJson(ts)
# print(j)

# parse to go from json to proto TS
new_ts = Parse(j, amanzi_pb2.TimeSeries())
print(new_ts)
# print(protobuf_to_dict(ts))

# if ts.data[0].value.string_value:
#     print("String: ", ts.data[0].value.string_value)
# if ts.data[0].value.double_value:
#     print("Double: ", ts.data[0].value.double_value)

# for dvp in ts.data:
#     print(dvp.datetime.ToDatetime(), dvp.value.double_value)



