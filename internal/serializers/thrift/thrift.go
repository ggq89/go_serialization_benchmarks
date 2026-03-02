package thrift

import (
	"context"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	genserbench "github.com/alecthomas/go_serialization_benchmarks/internal/serializers/thrift/gen-go/goserbench"
	"github.com/apache/thrift/lib/go/thrift"
)

type ThriftSerializer struct{}

func (s ThriftSerializer) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	thriftObj := &genserbench.ThriftStruct{
		Name:     v.Name,
		BirthDay: v.BirthDay.UnixNano(),
		Phone:    v.Phone,
		Siblings: int32(v.Siblings),
		Spouse:   v.Spouse,
		Money:    v.Money,
	}

	// Use TMemoryBuffer to write binary data
	transport := thrift.NewTMemoryBuffer()
	protocol := thrift.NewTBinaryProtocol(transport, true, true)

	if err := thriftObj.Write(context.Background(), protocol); err != nil {
		return nil, err
	}

	return transport.Bytes(), nil
}

func (s ThriftSerializer) Unmarshal(d []byte, o interface{}) error {
	// Use TMemoryBuffer to read binary data
	transport := thrift.NewTMemoryBufferLen(len(d))
	transport.Write(d)
	protocol := thrift.NewTBinaryProtocol(transport, true, true)

	thriftObj := &genserbench.ThriftStruct{}
	if err := thriftObj.Read(context.Background(), protocol); err != nil {
		return err
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = thriftObj.Name
	v.BirthDay = time.Unix(0, thriftObj.BirthDay)
	v.Phone = thriftObj.Phone
	v.Siblings = int(thriftObj.Siblings)
	v.Spouse = thriftObj.Spouse
	v.Money = thriftObj.Money
	return nil
}

func NewThriftSerializer() goserbench.Serializer {
	return ThriftSerializer{}
}
