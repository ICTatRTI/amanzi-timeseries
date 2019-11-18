from timeseries import amanzi_pb2
from google.protobuf.json_format import MessageToJson, Parse
from ts_utils import ts_to_data_arrays, ts_to_message_arrays

# Read JSON
with open("missing_value_ts.json", "r") as f:
    new_ts = Parse(f.read(), amanzi_pb2.TimeSeries())

dts, vals = ts_to_data_arrays(new_ts)
print(dts)
print(vals)

msgs = ts_to_message_arrays(new_ts)
for msg in msgs:
    print(msg)
