# go-amanzi
A go wrapper for the Amanzi timeseries specification.

Supports the following:
* Reading from  proto into native go types / slices
* Writing to proto from native go types / slices
* Provides useful Iterators, and functions for iterating over a Time series


## Installation
To install on GOPATH
```bash 
$ go get github.com/ICTatRTI/amanzi-timeseries/go-amanzi
```

Or if you are using modules

```go
package main 
import ( 
    "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
    "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/collections"
)
//...
```

## Examples
The full examples can be found [here](https://github.com/ICTatRTI/go-amanzi-examples)
Given an Amanzi Timeseries from a file `example.json` The following would read the file into memory


### Reading from Proto
```go
package main

import (
	"fmt"
	ac "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/collections"
	apt "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"
)
func main() {
	f, err := os.Open("./example.json")
	if err != nil {
		log.Fatal(err)
	}
	var pb apt.TimeSeries
    // were using jsonpb because this file is json, if its a binary could use proto.Unmarshal
	if err :=jsonpb.Unmarshal(f, &pb); err != nil {
		log.Fatal(err)
	}
    // use pb below
}
```

### Eagerly Iterating over an Amanzi Timeseries 
```go
package main

import (
	"fmt"
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/collections"
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./eagerIterator/data/example.json")
	if err != nil {
		log.Fatal(err)
	}
	var pb ptypes.TimeSeries
	if err :=jsonpb.Unmarshal(f, &pb); err != nil {
		log.Fatal(err)
	}

	itr, err := collections.NewEagerIterator(&pb)
	if err != nil {
		log.Fatal(err)
	}
	var sum float32
	for itr.Next() {
		rec := itr.Float32Record()

		if !rec.Missing {
			sum += rec.Value
		}
	}

	fmt.Printf("the sum is %f", sum)
    //sum is printed
}
```

## Lazily iterating over a An Amanzi TimeSeries 
```go
package main

import (
	"fmt"
	ac "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/collections"
	apt "github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"
)
func main() {
	f, err := os.Open("./lazyIterator/data/example.json")
	if err != nil {
		log.Fatal(err)
	}
	var pb apt.TimeSeries
	if err :=jsonpb.Unmarshal(f, &pb); err != nil {
		log.Fatal(err)
	}

	itr, err := ac.NewEagerIterator(&pb)
	if err != nil {
		log.Fatal(err)
	}
	var sum float32
	for itr.Next() {
		rec := itr.Float32Record()
		if !rec.Missing {
			sum += rec.Value
		}
	}

	fmt.Printf("the sum is %f", sum)
}
```

### Grabbing Native go slices from an Amanzi Protobuf
```go
package main

import (
	"fmt"
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/collections"
	"github.com/ICTatRTI/amanzi-timeseries/go-amanzi/ptypes"
	"github.com/golang/protobuf/jsonpb"
	"log"
	"os"
	"time"
)

func main() {
	f, err := os.Open("./slices/data/example.json")
	if err != nil {
		log.Fatal(err)
	}
	var pb ptypes.TimeSeries
	if err :=jsonpb.Unmarshal(f, &pb); err != nil {
		log.Fatal(err)
	}

	data := collections.Float32TsData{
		Times:   make([]time.Time, 0, len(pb.Data)),
		Values:  make([]float32, 0, len(pb.Data)),
		Missing: make(map[int]struct{}),
	}

	if err := collections.Float32Slices(&pb, &data); err != nil {
		log.Fatal(err)
	}
	var sum float32
	for i := 0; i< len(pb.Data); i++ {
		if _, ok := data.Missing[i]; !ok {
			sum += data.Values[i]
		}
	}
	fmt.Printf("the sum is %f", sum)
}
```


