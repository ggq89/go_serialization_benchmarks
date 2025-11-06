package protobuf3

import (
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type Protobuf3 struct {
	a Protobuf3Go

	// marshaller and unmarshaller are set on creation to either binary
	// or json marshallers.
	marshaller   func(proto.Message) ([]byte, error)
	unmarshaller func([]byte, proto.Message) error
}

func (s *Protobuf3) Marshal(o interface{}) ([]byte, error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a
	a.Name = v.Name
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = v.Phone
	a.Siblings = int32(v.Siblings)
	a.Spouse = v.Spouse
	a.Money = v.Money
	return s.marshaller(a)
}

func (s *Protobuf3) Unmarshal(d []byte, o interface{}) error {
	a := &s.a
	*a = Protobuf3Go{}

	err := s.unmarshaller(d, a)
	if err != nil {
		return err
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = a.Name
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = a.Phone
	v.Siblings = int(a.Siblings)
	v.Spouse = a.Spouse
	v.Money = a.Money
	return nil
}

func NewProtobufSerializer() goserbench.Serializer {
	return &Protobuf3{
		marshaller:   proto.Marshal,
		unmarshaller: proto.Unmarshal,
	}
}
