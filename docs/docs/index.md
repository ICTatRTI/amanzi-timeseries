# Amanzi™ Timeseries Specifcation

The Amanzi™ Time Series Specification (also referred to as the TSAPI) is a development standard for the storage, processing,
transmitting and consuming of Time Series data throughout the Amanzi™ system. The Amanzi™ Time series specification is intended
to be a fully comprehensive format supporting the following kinds of datasets: 

* Real time (e.g sensor data), 
* Historical  (both modeled and observed)
* Projections (forecasts into the future).


## Storing Time Series data efficiently
Time series data is a challenging beast to handle efficiently at scale and there have been numerous implementations attempting
to solve the problem. The Amanzi™ specification for storage has six major objectives:

* Keep the storage solution simple but efficient.
* Easily allow for local (users harddisk) or remote storage options.
* Store in a format that is commonly used for processing the data.
* Always preserve raw data from an origin datasource for a specified period of time.
* Store time series records in such a way that allows for partial reads/pagination.
* Flexible enough to handle very structured data or completely unstructured data.


## Providing sensible ways to process time series data

Through the implementations of this specification the goal is to provide a default set of operations
that are commonly performed on time series data. See [Implementations](implementations/create.md) for a list of 
implementations of this specification.

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

Unfortunately there is not a one-size fits all solution for transmitting time series data. However [this]() will
list some best practices around transporting time series data around for common use cases. 


## Consuming Time series data as an end-user

TODO
