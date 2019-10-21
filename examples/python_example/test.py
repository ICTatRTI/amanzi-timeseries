from timeseries import amanzi_pb2
from uuid import uuid4
from protobuf_to_dict import protobuf_to_dict, dict_to_protobuf
from google.protobuf.struct_pb2 import Struct
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
# meta.properties = {
#     "foo": 1.0,
#     "bar": 2.0
# }

properties = Struct()
properties_dict = {
        "key2": [1,2,3,4],
        "key": {
            "foo": {
                "bar": {
                    "value": 1
                }
            }
        }
    }
properties.update(properties_dict)
d = properties['key']['foo']['bar']['value']
print("D PRINTED", d)
print(properties.SerializeToString())
d_properties = protobuf_to_dict(properties)
print(d_properties)

meta.source.name = "Lake Biscuit"
meta.source.code = "LB"
meta.sourceLocation.name = "Corner of Lake Bisqcuit"
meta.sourceLocation.code = "315315"
meta.sourceLocation.elevation.value = 100
meta.sourceLocation.elevation.units = "ft"
meta.sourceLocation.elevation.datum = "NAVD88"
meta.parameter.name = "Streamflow"
meta.parameter.code = "QIN"
meta.parameter.units = "cfs"
# meta.properties.update(
#     {
#         "key2": [1,2,3,4],
#         "key": 1
#     }
# )
metaInfo = protobuf_to_dict(meta)
print(metaInfo)
print(source.SerializeToString())
print(meta.SerializeToString())


