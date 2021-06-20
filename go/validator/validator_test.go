package validator

import "testing"

type TestDataStruct struct {
	Name string `validate:"required" label:"名称"`
	Age   int    `validate:"required,gte=0,lte=130"`
}

func TestCheckDataTest(t *testing.T) {
	InitValidator()
	var data TestDataStruct
	data.Name = "大白"
	ok, err := CheckStructParam(data)
	if !ok {
		t.Fatal(err)
	}

}

func BenchmarkCheckDataTest(b *testing.B)  {
	var data TestDataStruct
	data.Name = "大白"
	data.Age=27
	ok, err := CheckStructParam(data)
	if !ok {
		b.Fatal(err)
	}

}
func init()  {
	InitValidator()

}