Amanzi Timeseries Specification Protobuf Description
====================================================

## Background
The TS Spec is implemented using Google's Protocol Buffers https://developers.google.com/protocol-buffers/docs/overview.  

According to the website, Protocol Buffers are a language-neutral, platform-neutral, extensible way of serializing structured data for use in communications protocols, data storage, and more.  The TS Spec Protocol Buffer (proto) described herein should be used to read and write all timeseries data used within Amanzi.

## Purpose
This document is intended to describe the various parts (messages) that make up the timeseries specification (TS Spec) and what each one is intended to be used for.

The following is an example of what a timeseries message looks like in JSON.

```json
{
  "metaInfo": {
    "id": "39909dd0-2f86-4c22-a569-bc6cf0eb6f11",
    "name": "Lake Biscuit Release",
    "code": "LBR",
    "type": "forecast",
    "typeVariant": "raw",
    "properties": {
      "test": "foo"
    },
    "groups": [
      {
        "groupId": "781bccb0-cc19-40bb-b5bf-9b1b6ebe4499",
        "description": "Ensemble 1",
        "groupType": "Ensemble"
      },
      {
        "groupId": "c2a456a4-b923-4504-8a3f-c0a970b86c17",
        "description": "HMS Model Simulation",
        "groupType": "Model"
      }
    ],
    "source": {
      "name": "Lake Biscuit Operating Authority",
      "code": "LBOA"
    },
    "sourceLocation": {
      "name": "Lake Biscuit Dam",
      "code": "5315315",
      "elevation": {
        "value": 100.0,
        "units": "ft",
        "datum": "NAVD88"
      }
    },
    "parameter": {
      "name": "Streamflow",
      "code": "QIN",
      "units": "cfs",
      "statistic": "Instantaneous",
      "windowedInterval": "PT3H"
    },
    "timeInfo": {
      "referenceTime": "2019-10-31T17:12:07.153644Z",
      "start": "2000-01-01T00:00:00Z",
      "end": "2000-01-03T00:00:00Z",
      "interval": "PT1H"
    },
    "origin": {
      "description": "Calculated hourly mean from raw timeseries",
      "referenceIds": [
        "82d81690-ff8f-4378-b714-7539bfdb3937",
        "66674987-c947-432a-970b-25f3a0b0174d"
      ]
    }
  },
  "data": [
    {
      "datetime": "2000-01-01T05:00:00Z",
      "value": {
        "doubleValue": 0.8432695424884182
      },
      "qualifiers": [
        "p"
      ]
    },
    {
      "datetime": "2000-01-01T06:00:00Z",
      "value": {
        "doubleValue": 1.8432695424884182
      },
      "qualifiers": [
        "p"
      ]
    }
  ]
}
  ```
## Notes on Datetime
The datetime fields are stored in the proto as integer seconds since the epoch.  In a JSON representation of proto message they are converted to ISO 8601 string format. This can make it appear as though the datetime is stored as ISO 8601 string but it is not.

When working with the proto in Python there are `ToDatetime()` and `FromDatetime()` methods for the datetime fields that handel the conversion from python datetime object to integer seconds and back.

The TS Spec is a collection of nested messages.  The following sections describe each message, organized by nesting level.


## First Level Messages
| Name      | Type        |    Description                          
|:-----     |:-----             |:------------    
| metaInfo | object |  A collection of messages the describe the data contained in the timeseries
| geoLocations | object | a GeoJSON object that defines the spatial attributes of the timeseries (i.e. what spatial areas/points/lines does the data apply to)
| data | array | The timeseries data (datetime, value and qualifier data)

## Second Level Messages
_metaInfo:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| id        | string    | A universally unique identifier (UUID)    
| name      | string    | A long name for the timeseries           
| code      | string    | A short name for the timeseries           
| type      | string    | The type of value in the timeseries, for example forecast, observed, simulated
| typeVariant | string | Used to further defined the type.  (e.g. raw, adjusted, etc.)
| properties| object    | A generic object to store timeseries properties not captured in the other metaInfo messages.
| groups    | array    | Used to specify what groups the timeseries belongs to.  this could be an ensemble, an ensemble member, a model simulation, or any other group of timeseries.  See below for description of group
| source    | object    | Where the data came from.  What organization, what model, etc. (see message definition below)
| sourceLocation | object | Used to describe the name of the location for the source. (see message definition below)
| parameter | object | Used to describe the variable type of the time series (e.g. streamflow, stage, temperature) (see message definition below)
| timeInfo | object | Contains all the time related info for the timeseries (see message definition below)
| origin | object | Used to describe where the data came from or how it was processed and what timeseries were used to create it.

_geoLocation:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| feature   | object    | needs work

_data:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| id        | array     | An array of data messages

## Third and Higher Level Nested Messages
_group:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| groupId       | string | A universally unique identifier (UUID)    |
| description   | string | A human readable way to identify the group
| groupType | string | The type of group that this is (e.g. Ensemble, Ensemble Member, etc.)

_source:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| name | string | Full name of the source (e.g. United States Geologic Survey)
| code | string | Short name or acronym (e.g. USGS)

_sourceLocation:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| name | string | Name of location used by specified source
| code | string | Short name of location used by source
| elevation | object | An elevation message used to describe the reference elevation for a location (e.g. the datum)

_parameter:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| name | string | Name of the parameter being represented in the timeseries (e.g. Streamflow)
| code | string | Short name or code for the parameter (e.g.QIN)
| units | string | The units of the parameter (e.g. cfs)
| statistic | string    | Describes if and how the data is aggregated.  For example if the data is instantaneous, mean, sum, min, etc.
| windowedInterval | string | The interval of the value.  Would be used primarily when doing moving aggregations, for example.  In ISO 8601 duration string format (e.g. PT1H)

_timeInfo:_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| referenceTime | integer | A reference time for the timeseries. Can be used for time of forecast, or some other reference time.  Datetime represented as seconds since the epoch (1/1/1970 00:00:00 UTC) in the proto.  Converted to an ISO 8601 string format (e.g. "2019-10-31T17:12:07Z") in JSON.
| start | integer | Start time of the timeseries.  Datetime represented as seconds since the epoch (1/1/1970 00:00:00 UTC) in the proto.  Converted to an ISO 8601 string format (e.g. "2019-10-31T17:12:07Z") in JSON
| end | | End time of the timeseries. Datetime represented as seconds since the epoch (1/1/1970 00:00:00 UTC) in the proto.  Converted to an ISO 8601 string format (e.g. "2019-10-31T17:12:07Z") in JSON
| interval | string | The expected interval of the values.  In ISO 8601 duration string format (e.g. PT1H).  If missing treat as irregular.

_origin:_
| Name      | Type      |    Description                            |
|:-----     |:-----     |:------------                              |
| description | string | A description of how the timeseries was created
| referenceIds | array | An array of UUIDs that went into the creation of the timeseries 

_elevation:_
| Name      | Type      |    Description                            |
|:-----     |:-----     |:------------                              |
| value | float | The value of the elevation reference (e.g. the gage datum)
| units | string | The units of the value (e.g. ft)
| datum | string | Datum used for the sourceLocation (e.g. NAVD88)

_data "message":_
| Name      | Type      |    Description                            
|:-----     |:-----     |:------------                              
| datetime | integer | The datetime of the value in the value object. Datetime represented as seconds since the epoch (1/1/1970 00:00:00 UTC) in the proto.  Converted to an ISO 8601 string format (e.g. "2019-10-31T17:12:07Z") in JSON
| value | object | An object that contains a value type key and value.  See below for description
| qualifiers | array | An array of strings containing value qualifiers (e.g. "raw", "verified") 

_value:_
| Name      | Type      |    Description                            |
|:-----     |:-----     |:------------                              |
| doubleValue | double | Double value
| floatValue | float | Float value
| int64Value | int64 | 64-bit Integer value
| int32Value | int32 | 32-bit Integer value
| uint64Value | uint64 | Unsigned 64-bit Integer value
| uint32Value | uint32 | Unsigned 32-bit Integer value
| stringValue |string | String value

## Data Validation
The proto only specifies the value type for each message property.  It does not specify what values are valid.  Determining what values are valid is left up to the user, client or application using the timeseries.  If it encounters an invalid value it should alert the user. All messages that make up the proto are optional, nothing is required.