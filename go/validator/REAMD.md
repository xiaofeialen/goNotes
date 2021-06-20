### parameter verifier
1. use

```
1. InitValidator() //must init

var data TestDataStruct
data.Name = "大白"
ok, err := CheckStructParam(data) // ok is false or true; err is error massage

```
2. info
goos: darwin
goarch: amd64
free: 8 GB 2133 MHz LPDDR3
cpu: 2.4 GHz 四核Intel Core i5
BenchmarkCheckDataTest
BenchmarkCheckDataTest-8   	1000000000	         0.000008 ns/op