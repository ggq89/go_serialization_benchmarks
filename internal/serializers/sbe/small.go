package sbe

import (
	"bytes"
	"time"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	"github.com/alecthomas/go_serialization_benchmarks/internal/serializers/sbe/small"
)

type SBESerializer struct {
	header small.MessageHeader
	a      small.SmallStruct
	m      *small.SbeGoMarshaller
	buf    *bytes.Buffer
}

func (s *SBESerializer) Marshal(o interface{}) (buf []byte, err error) {
	v := o.(*goserbench.SmallStruct)
	a := &s.a

	small.SmallStructInit(a)
	a.Name = []uint8(v.Name)
	a.BirthDay = v.BirthDay.UnixNano()
	a.Phone = []uint8(v.Phone)
	a.Siblings = int32(v.Siblings)
	a.Spouse = fromBool(v.Spouse)
	a.Money = v.Money

	err = s.header.Encode(s.m, s.buf)
	if err != nil {
		return nil, err
	}

	err = a.Encode(s.m, s.buf, true)
	if err != nil {
		return nil, err
	}

	buf = s.buf.Bytes()
	s.buf.Reset()
	return
}

func fromBool(b bool) small.BooleanTypeEnum {
	if b {
		return 1
	} else {
		return 0
	}
}

func getBool(b small.BooleanTypeEnum) bool {
	if b == 0 {
		return false
	} else {
		return true
	}
}

func (s *SBESerializer) Unmarshal(bs []byte, o interface{}) (err error) {
	a := &s.a
	reader := bytes.NewReader(bs)

	err = s.header.Decode(s.m, reader, a.SbeSchemaVersion())
	if err != nil {
		return err
	}

	err = a.Decode(s.m, reader, a.SbeSchemaVersion(), a.SbeBlockLength(), true)
	if err != nil {
		return err
	}

	v := o.(*goserbench.SmallStruct)
	v.Name = string(a.Name)
	v.BirthDay = time.Unix(0, a.BirthDay)
	v.Phone = string(a.Phone)
	v.Siblings = int(a.Siblings)
	v.Spouse = getBool(a.Spouse)
	v.Money = a.Money
	return
}

func (s *SBESerializer) TimePrecision() time.Duration {
	return time.Nanosecond
}

func NewSBESerializer() goserbench.Serializer {
	var a small.SmallStruct
	header := small.MessageHeader{
		BlockLength: a.SbeBlockLength(),
		TemplateId:  a.SbeTemplateId(),
		SchemaId:    a.SbeSchemaId(),
		Version:     a.SbeSchemaVersion(),
	}

	return &SBESerializer{
		header: header,
		a:      small.SmallStruct{},
		m:      small.NewSbeGoMarshaller(),
		buf:    new(bytes.Buffer),
	}
}

// func unsafeSliceToString(b []byte) string {
// 	return *(*string)(unsafe.Pointer(&b))
// }
