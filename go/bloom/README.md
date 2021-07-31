#### bloom

##### Use bloom filters to solve cache penetration
1. use package 
 ```json
"github.com/bits-and-blooms/bloom/v3"
```
2. create bloom
```go
bloomFilter:= bloom.NewWithEstimates(1000000, 0.3)
```
3. add data
```go
 bloomFilter.Add(key)
```
4. checkData
```go
bloomFilter.Test(key)
```

#### Actual use demo
1. see 
    ./demo/penetrate.go
    ./demo/penetrate_test.go
