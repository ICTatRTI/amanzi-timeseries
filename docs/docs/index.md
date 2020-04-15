# Amanzi™ Timeseries Specifcation

The Amanzi™ Time Series Specification is a on-the-wire format standard for transmitting 
Time Series data. The specification is based on googles [Protobuf Protocols](https://developers.google.com/protocol-buffers/).
In addition to the on-the-wire format specification, wrapper libraries that provide quality of life improvements for working with
the protobufs. 

## Transmitting time series data 

To transmit time series data to multiple components in a cloud based system a common structure is needed. There are many
things to consider when deciding on a format for transport such as: 

* Size of the data  
* The client device that will eventually receive the data(e.g browser, your harddrive, or another server) 
* The transport protocol used to transmit the data to its destination. 

For most Browser based applications HTTP using a transport format of [json](https://www.json.org/) is recommended, however its important that
the data is appropriately set up for partial reads, or pagination as very large json data sets will reduce your users 
performance significantly.

For server to server communication ideally HTTP2 with a binary [protobuf](https://developers.google.com/protocol-buffers/) format is recommended.

Implementations of this specifications **must** support both of these formats

