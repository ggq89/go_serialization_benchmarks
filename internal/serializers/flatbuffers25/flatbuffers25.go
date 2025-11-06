package flatbuffers25

import (
	"time"
	"unsafe"

	"github.com/alecthomas/go_serialization_benchmarks/goserbench"
	flatbuffers "github.com/google/flatbuffers/go"
)

type FlatBuffer25Serializer struct {
	unsafeReuse bool
	builder     *flatbuffers.Builder
}

func (s *FlatBuffer25Serializer) Marshal(o interface{}) ([]byte, error) {
	a := o.(*goserbench.SmallStruct)
	builder := s.builder
	if !s.unsafeReuse {
		builder.Bytes = nil // free
	}
	builder.Reset()

	name := builder.CreateString(a.Name)
	phone := builder.CreateString(a.Phone)

	FlatBuffer25AStart(builder)
	FlatBuffer25AAddName(builder, name)
	FlatBuffer25AAddPhone(builder, phone)
	FlatBuffer25AAddBirthDay(builder, a.BirthDay.UnixNano())
	FlatBuffer25AAddSiblings(builder, int32(a.Siblings))
	FlatBuffer25AAddSpouse(builder, a.Spouse)
	FlatBuffer25AAddMoney(builder, a.Money)
	builder.Finish(FlatBuffer25AEnd(builder))
	return builder.FinishedBytes(), nil
}

func (s *FlatBuffer25Serializer) Unmarshal(d []byte, i interface{}) error {
	a := i.(*goserbench.SmallStruct)
	o := FlatBuffer25A{}
	o.Init(d, flatbuffers.GetUOffsetT(d))
	if s.unsafeReuse {
		a.Name = unsafeSliceToString(o.Name())
		a.Phone = unsafeSliceToString(o.Phone())
	} else {
		a.Name = string(o.Name())
		a.Phone = string(o.Phone())
	}
	a.BirthDay = time.Unix(0, o.BirthDay())
	a.Siblings = int(o.Siblings())
	a.Spouse = o.Spouse()
	a.Money = o.Money()
	return nil
}

func NewFlatBuffers25Serializer() goserbench.Serializer {
	return &FlatBuffer25Serializer{builder: flatbuffers.NewBuilder(0), unsafeReuse: false}
}

func NewFlatBuffers25UnsafeReuseSerializer() goserbench.Serializer {
	const maxSerSize = 128
	return &FlatBuffer25Serializer{builder: flatbuffers.NewBuilder(maxSerSize), unsafeReuse: true}
}

// unsafeSliceToString converts a byte slice to a string in an unsafe way
// (modifications to the byte slice modify the string).
func unsafeSliceToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}
