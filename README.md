# oba - Golang Api for One Bus Away
[![Build Status](https://travis-ci.org/Setheck/oba.svg?branch=master)](https://travis-ci.org/Setheck/oba) [![Go Report Card](https://goreportcard.com/badge/github.com/setheck/oba)](https://goreportcard.com/report/github.com/setheck/oba)

[OneBusAway](https://onebusaway.org/)
# Summary
I wanted to write some apps that interface with one bus away and could not find a go api, so I wrote one!
[One Bus Away Documentation](http://developer.onebusaway.org/modules/onebusaway-application-modules/current/api/where/index.html)
It was also a good excuse to have some fun with golang

# Get 
```
$ go get -u github.com/Setheck/oba
```

# Import
``` go
import (
    "fmt"
    "github.com/Setheck/oba"
)
```

# Use
### Agency
```go
```go
func main() {
    client := oba.NewDefaultClientS(server.URL, TestApiKey)
    agency, e := client.Agency("1")
    if e != nil {
        t.Error(e)
    }
}
```
```
## AgenciesWithCoverage
```go
func main() {
    client := oba.NewDefaultClientS(server.URL, TestApiKey)
    awcs, e := client.AgenciesWithCoverage()
    if e != nil {
        t.Error(e)
    }
}
```
## Route
```go
func main() {
    client := oba.NewDefaultClientS("http://api.pugetsound.onebusaway.org/api/where/", "TEST")
    route, err := client.Route("1_100224")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Print(route.ID)
}
```