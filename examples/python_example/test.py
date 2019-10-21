from timeseries import amanzi_pb2
from uuid import uuid4
from protobuf_to_dict import protobuf_to_dict, dict_to_protobuf
from google.protobuf.struct_pb2 import Struct
from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson
import datetime
import random
import json
source = amanzi_pb2.Source()
# set simple values
source.name = "Lake Bisqcuit"
source.code = "LB"
src_dict = protobuf_to_dict(source)
print(src_dict) # lets get a normal python dictionary, potentially more useful than a json_str
# set a more complex value

parameter = amanzi_pb2.Parameter()
parameter.name = "Streamflow"
parameter.code = "QIN"
parameter.units = "cfs"

meta = amanzi_pb2.TimeSeriesMetaInfo()
meta.id = str(uuid4())
meta.name = "test ts"
meta.code = "wow cool"
meta.type = "a bullshit forecast XD"
properties = Struct()
properties.update(
    {
        "key2": [1,2,3,4],
        "key": 1,
        "fab": "string",
        "boo": None,
        "dam": {
            "beta": "biscuit"
        }
    }
)
reference = datetime.datetime.utcnow().timestamp()
start = datetime.datetime(year=2000,month=1,day=1,hour=0, minute=0, second=0, tzinfo=datetime.timezone.utc)
end = start + datetime.timedelta(days=10)
duration = datetime.timedelta(hours=1)
print(reference)

# set up timestamp pbs
reference_pb = Timestamp()
reference_pb.seconds = int(reference)

start_pb = Timestamp()
start_pb.seconds = int(start.timestamp())
end_pb = Timestamp()
end_pb.seconds = int(end.timestamp())


interval = datetime.timedelta(hours=1)
timeInfo = amanzi_pb2.TimeInfo()
timeInfo.referenceTime.CopyFrom(reference_pb)
timeInfo.start.CopyFrom(start_pb)
timeInfo.end.CopyFrom(end_pb)
timeInfo.interval = "PT1H"

meta.source.CopyFrom(source)
meta.sourceLocation.name = "Corner of Lake Bisqcuit"
meta.sourceLocation.code = "315315"
meta.sourceLocation.elevation.value = 100
meta.sourceLocation.elevation.units = "ft"
meta.sourceLocation.elevation.datum = "NAVD88"
meta.timeInfo.CopyFrom(timeInfo)
meta.parameter.CopyFrom(parameter)
meta.properties.CopyFrom(properties)

ts = amanzi_pb2.TimeSeries()
ts.metaInfo.CopyFrom(meta)

current = start

while current.timestamp() <= end.timestamp():
    record = ts.data.add()
    record_dt_pb = Timestamp()
    record_dt_pb.seconds = int(current.timestamp())

    record.datetime.CopyFrom(record_dt_pb)
    record.qualifiers.extend("p")
    record.value.double_value = random.random()
    current = current + duration


print(ts.SerializeToString())
d = protobuf_to_dict(ts)
j = MessageToJson(ts)
print(j)
print(protobuf_to_dict(ts))



