from timeseries import amanzi_pb2
from google.protobuf.json_format import MessageToJson, Parse
from ts_utils import ts_to_data_arrays, ts_to_message_arrays, write_arrays_to_ts
from datetime import datetime

# Read JSON
with open("missing_value_ts.json", "r") as f:
    ts = Parse(f.read(), amanzi_pb2.TimeSeries())

dts, vals = ts_to_data_arrays(ts)
print(dts)
print(vals)

msgs = ts_to_message_arrays(ts)
for msg in msgs:
    print(msg)

# datetime = [datetime(2000, 1, 1, 0, 0), datetime(2000, 1, 1, 1, 0), datetime(2000, 1, 1, 2, 0), datetime(2000, 1, 1, 3, 0)]
# values = [1,2,3,4]
# qualifiers = [[],[],[],[]]
# new_ts = write_arrays_to_ts(meta=ts.metaInfo, datetime=datetime, values=values, qualifiers=qualifiers)
# print(MessageToJson(new_ts))

# datetime = [datetime(2000, 1, 1, 0, 0), datetime(2000, 1, 1, 1, 0), datetime(2000, 1, 1, 2, 0), datetime(2000, 1, 1, 3, 0)]
# values = [1.0, 2.2, 3.0, 4.2]
# qualifiers = [[],[],[],[]]
# new_ts = write_arrays_to_ts(meta=ts.metaInfo, datetime=datetime, values=values, qualifiers=qualifiers)
# print(MessageToJson(new_ts))

# datetime = [datetime(2000, 1, 1, 0, 0), datetime(2000, 1, 1, 1, 0), datetime(2000, 1, 1, 2, 0), datetime(2000, 1, 1, 3, 0)]
# values = ["foo", "bar", "biz", "baz"]
# qualifiers = [[],[],[],[]]
# new_ts = write_arrays_to_ts(meta=ts.metaInfo, datetime=datetime, values=values, qualifiers=qualifiers)
# print(MessageToJson(new_ts))

datetime = [datetime(2000, 1, 1, 0, 0), datetime(2000, 1, 1, 1, 0), datetime(2000, 1, 1, 2, 0), datetime(2000, 1, 1, 3, 0)]
values = [None, "bar", 1, 1.5]
qualifiers = [[],[],[],[]]
new_ts = write_arrays_to_ts(meta=ts.metaInfo, datetime=datetime, values=values, qualifiers=qualifiers)
print(MessageToJson(new_ts))