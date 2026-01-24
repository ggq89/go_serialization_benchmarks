package fory

import (
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/apache/fory/go/fory"
)

var f *fory.Fory

func init() {
	// Create a Fory instance
	f = fory.New()

	// Register struct with a type ID
	if err := f.RegisterStruct(goserbench.SmallStruct{}, 1); err != nil {
		panic(err)
	}
}

type ForySerializer struct{}

func (fs ForySerializer) Marshal(o interface{}) ([]byte, error) {
	return f.Serialize(o)
}

func (fs ForySerializer) Unmarshal(d []byte, o interface{}) error {
	return f.Deserialize(d, o)
}

func NewForySerializer() goserbench.Serializer {
	return ForySerializer{}
}
