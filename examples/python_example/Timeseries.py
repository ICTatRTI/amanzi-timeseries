from timeseries import amanzi_pb2
# from google.protobuf.timestamp_pb2 import Timestamp
from google.protobuf.json_format import MessageToJson, Parse
from datetime import datetime
# import datetime
# import isodate
# import timeseries_generator
# from uuid import uuid4
# from typing import Dict, AnyStr, Any, List


class Timeseries:
    """A class to work with the Amanzi Timeseries Protobuf
        Can be passed a valid protobuf when initiated.  If not, a blank object will be created.
        If a valid protobuf is not provided when object is initiated, properties should be set manually.
    """
    def __init__(self, ts: amanzi_pb2.TimeSeries = amanzi_pb2.TimeSeries()):
        self.ts = ts

    def load_json(self, json: str):
        """Loads a JSON into timeseries proto"""
        self.ts = Parse(json, amanzi_pb2.TimeSeries())

    def load_binary(self, binary: str):
        """Loads a binary string into timeseries proto"""
        self.ts.ParseFromString(binary)

    def to_json(self):
        """Returns a JSON object of the timeseries object"""
        return MessageToJson(self.ts)

    def to_proto(self):
        """Returns the timeseries object as a Protobuf"""
        return self.ts

    def to_data_arrays(self):
        """Creates 2 lists of timeseries data.
            First list has datetime objects, second list has values
            Value type used in first timeseries message assumed for all subsequent messages

        :return:
        """
        dts = []
        vals = []

        if self.ts.data[0].value.double_value:
            for d in self.ts.data:
                dts.append(d.datetime.ToDatetime())
                vals.append(d.value.double_value)
        elif self.ts.data[0].value.float_value:
            for d in self.ts.data:
                dts.append(d.datetime.ToDatetime())
                vals.append(d.value.float_value)
        elif self.ts.data[0].value.int64_value:
            for d in self.ts.data:
                dts.append(d.datetime.ToDatetime())
                vals.append(d.value.int64_value)
        elif self.ts.data[0].value.int32_value:
            for d in self.ts.data:
                dts.append(d.datetime.ToDatetime())
            vals.append(d.value.int32_value)
        elif self.ts.data[0].value.uint64_value:
            for d in self.ts.data:
                dts.append(d.datetime.ToDatetime())
                vals.append(d.value.uint64_value)
        elif self.ts.data[0].value.uint32_value:
            for d in self.ts.data:
                dts.append(d.datetime.ToDatetime())
                vals.append(d.value.uint32_value)
        elif self.ts.data[0].value.string_value:
            for d in self.ts.data:
                dts.append(d.datetime.ToDatetime())
                vals.append(d.value.string_value)

        return dts, vals

    def to_message_arrays(self):

        data = []

        if self.ts.data[0].value.double_value:
            for d in self.ts.data:
                data.append([d.datetime.ToDatetime(), d.value.double_value])
        elif self.ts.data[0].value.float_value:
            for d in self.ts.data:
                data.append([d.datetime.ToDatetime(), d.value.float_value])
        elif self.ts.data[0].value.int64_value:
            for d in self.ts.data:
                data.append([d.datetime.ToDatetime(), d.value.int64_value])
        elif self.ts.data[0].value.int32_value:
            for d in self.ts.data:
                data.append([d.datetime.ToDatetime(), d.value.int32_value])
        elif self.ts.data[0].value.uint64_value:
            for d in self.ts.data:
                data.append([d.datetime.ToDatetime(), d.value.uint64_value])
        elif self.ts.data[0].value.uint32_value:
            for d in self.ts.data:
                data.append([d.datetime.ToDatetime(), d.value.uint32_value])
        elif self.ts.data[0].value.string_value:
            for d in self.ts.data:
                data.append([d.datetime.ToDatetime(), d.value.string_value])

        return data

    def get_message_by_datetime(self, dt: datetime):
        data = self.ts.data

    @property
    def id(self) -> str:
        return self.ts.metaInfo.id

    @id.setter
    def id(self, id: str):
        self.ts.metaInfo.id = id

    @property
    def name(self) -> str:
        return self.ts.metaInfo.name

    @name.setter
    def name(self, name: str):
        self.ts.metaInfo.name = name

    @property
    def code(self) -> str:
        return self.ts.metaInfo.code

    @code.setter
    def code(self, code: str):
        self.ts.metaInfo.code = code

    @property
    def statistic(self) -> str:
            return self.ts.metaInfo.statistic

    @statistic.setter
    def statistic(self, statistic: str):
        self.ts.metaInfo.statistic = statistic

    @property
    def type(self) -> str:
        return self.ts.metaInfo.type

    @type.setter
    def type(self, type: str):
        self.ts.metaInfo.type = type

    @property
    def source(self) -> amanzi_pb2.Source:
        return self.ts.metaInfo.source

    @source.setter
    def source(self, source: amanzi_pb2.Source):
        self.ts.metaInfo.source.CopyFrom(source)

    @property
    def interval(self) -> str:
        return self.ts.metaInfo.timeInfo.interval

    @interval.setter
    def interval(self, interval):
        self.ts.metaInfo.timeInfo.interval = interval


