package capnprotov3

import (
	"time"

	capnp "capnproto.org/go/capnp/v3"
	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
)

type CapNProtoSerializer struct{}

func (x CapNProtoSerializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*goserbench.SmallStruct)
	arena := capnp.SingleSegment(nil)
	msg, seg, err := capnp.NewMessage(arena)
	if err != nil {
		return nil, err
	}

	c, err := NewRootCapnpV3(seg)
	if err != nil {
		return nil, err
	}

	c.SetName(a.Name)
	c.SetBirthDay(a.BirthDay.UnixNano())
	c.SetPhone(a.Phone)
	c.SetSiblings(int32(a.Siblings))
	c.SetSpouse(a.Spouse)
	c.SetMoney(a.Money)

	return msg.Marshal()
}

func (x CapNProtoSerializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*goserbench.SmallStruct)

	msg, err := capnp.Unmarshal(d)
	if err != nil {
		return err
	}

	o, err := ReadRootCapnpV3(msg)
	if err != nil {
		return err
	}

	a.Name, _ = o.Name()
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Phone, _ = o.Phone()
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()

	return nil
}

func NewCapNProtoSerializer() goserbench.Serializer {
	return CapNProtoSerializer{}
}
