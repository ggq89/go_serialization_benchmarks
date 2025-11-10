// Generated SBE (Simple Binary Encoding) message codec

package small

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

type SmallStruct struct {
	BirthDay int64
	Siblings int32
	Spouse   BooleanTypeEnum
	Money    float64
	Name     []uint8
	Phone    []uint8
}

func (s *SmallStruct) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, s.BirthDay); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.Siblings); err != nil {
		return err
	}
	if err := s.Spouse.Encode(_m, _w); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, s.Money); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(s.Name))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Name); err != nil {
		return err
	}
	if err := _m.WriteUint32(_w, uint32(len(s.Phone))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Phone); err != nil {
		return err
	}
	return nil
}

func (s *SmallStruct) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.BirthDayInActingVersion(actingVersion) {
		s.BirthDay = s.BirthDayNullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.BirthDay); err != nil {
			return err
		}
	}
	if !s.SiblingsInActingVersion(actingVersion) {
		s.Siblings = s.SiblingsNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.Siblings); err != nil {
			return err
		}
	}
	if s.SpouseInActingVersion(actingVersion) {
		if err := s.Spouse.Decode(_m, _r, actingVersion); err != nil {
			return err
		}
	}
	if !s.MoneyInActingVersion(actingVersion) {
		s.Money = s.MoneyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &s.Money); err != nil {
			return err
		}
	}
	if actingVersion > s.SbeSchemaVersion() && blockLength > s.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-s.SbeBlockLength()))
	}

	if s.NameInActingVersion(actingVersion) {
		var NameLength uint32
		if err := _m.ReadUint32(_r, &NameLength); err != nil {
			return err
		}
		if cap(s.Name) < int(NameLength) {
			s.Name = make([]uint8, NameLength)
		}
		s.Name = s.Name[:NameLength]
		if err := _m.ReadBytes(_r, s.Name); err != nil {
			return err
		}
	}

	if s.PhoneInActingVersion(actingVersion) {
		var PhoneLength uint32
		if err := _m.ReadUint32(_r, &PhoneLength); err != nil {
			return err
		}
		if cap(s.Phone) < int(PhoneLength) {
			s.Phone = make([]uint8, PhoneLength)
		}
		s.Phone = s.Phone[:PhoneLength]
		if err := _m.ReadBytes(_r, s.Phone); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *SmallStruct) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.BirthDayInActingVersion(actingVersion) {
		if s.BirthDay < s.BirthDayMinValue() || s.BirthDay > s.BirthDayMaxValue() {
			return fmt.Errorf("Range check failed on s.BirthDay (%v < %v > %v)", s.BirthDayMinValue(), s.BirthDay, s.BirthDayMaxValue())
		}
	}
	if s.SiblingsInActingVersion(actingVersion) {
		if s.Siblings < s.SiblingsMinValue() || s.Siblings > s.SiblingsMaxValue() {
			return fmt.Errorf("Range check failed on s.Siblings (%v < %v > %v)", s.SiblingsMinValue(), s.Siblings, s.SiblingsMaxValue())
		}
	}
	if err := s.Spouse.RangeCheck(actingVersion, schemaVersion); err != nil {
		return err
	}
	if s.MoneyInActingVersion(actingVersion) {
		if s.Money < s.MoneyMinValue() || s.Money > s.MoneyMaxValue() {
			return fmt.Errorf("Range check failed on s.Money (%v < %v > %v)", s.MoneyMinValue(), s.Money, s.MoneyMaxValue())
		}
	}
	if !utf8.Valid(s.Name[:]) {
		return errors.New("s.Name failed UTF-8 validation")
	}
	if !utf8.Valid(s.Phone[:]) {
		return errors.New("s.Phone failed UTF-8 validation")
	}
	return nil
}

func SmallStructInit(s *SmallStruct) {
	return
}

func (*SmallStruct) SbeBlockLength() (blockLength uint16) {
	return 21
}

func (*SmallStruct) SbeTemplateId() (templateId uint16) {
	return 12
}

func (*SmallStruct) SbeSchemaId() (schemaId uint16) {
	return 4
}

func (*SmallStruct) SbeSchemaVersion() (schemaVersion uint16) {
	return 1
}

func (*SmallStruct) SbeSemanticType() (semanticType []byte) {
	return []byte("small")
}

func (*SmallStruct) SbeSemanticVersion() (semanticVersion string) {
	return "1.0"
}

func (*SmallStruct) BirthDayId() uint16 {
	return 1
}

func (*SmallStruct) BirthDaySinceVersion() uint16 {
	return 0
}

func (s *SmallStruct) BirthDayInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BirthDaySinceVersion()
}

func (*SmallStruct) BirthDayDeprecated() uint16 {
	return 0
}

func (*SmallStruct) BirthDayMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SmallStruct) BirthDayMinValue() int64 {
	return math.MinInt64 + 1
}

func (*SmallStruct) BirthDayMaxValue() int64 {
	return math.MaxInt64
}

func (*SmallStruct) BirthDayNullValue() int64 {
	return math.MinInt64
}

func (*SmallStruct) SiblingsId() uint16 {
	return 2
}

func (*SmallStruct) SiblingsSinceVersion() uint16 {
	return 0
}

func (s *SmallStruct) SiblingsInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SiblingsSinceVersion()
}

func (*SmallStruct) SiblingsDeprecated() uint16 {
	return 0
}

func (*SmallStruct) SiblingsMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SmallStruct) SiblingsMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SmallStruct) SiblingsMaxValue() int32 {
	return math.MaxInt32
}

func (*SmallStruct) SiblingsNullValue() int32 {
	return math.MinInt32
}

func (*SmallStruct) SpouseId() uint16 {
	return 3
}

func (*SmallStruct) SpouseSinceVersion() uint16 {
	return 0
}

func (s *SmallStruct) SpouseInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.SpouseSinceVersion()
}

func (*SmallStruct) SpouseDeprecated() uint16 {
	return 0
}

func (*SmallStruct) SpouseMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SmallStruct) MoneyId() uint16 {
	return 4
}

func (*SmallStruct) MoneySinceVersion() uint16 {
	return 0
}

func (s *SmallStruct) MoneyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MoneySinceVersion()
}

func (*SmallStruct) MoneyDeprecated() uint16 {
	return 0
}

func (*SmallStruct) MoneyMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SmallStruct) MoneyMinValue() float64 {
	return -math.MaxFloat64
}

func (*SmallStruct) MoneyMaxValue() float64 {
	return math.MaxFloat64
}

func (*SmallStruct) MoneyNullValue() float64 {
	return math.NaN()
}

func (*SmallStruct) NameMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SmallStruct) NameSinceVersion() uint16 {
	return 0
}

func (s *SmallStruct) NameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.NameSinceVersion()
}

func (*SmallStruct) NameDeprecated() uint16 {
	return 0
}

func (SmallStruct) NameCharacterEncoding() string {
	return "UTF-8"
}

func (SmallStruct) NameHeaderLength() uint64 {
	return 4
}

func (*SmallStruct) PhoneMetaAttribute(meta int) string {
	switch meta {
	case 1:
		return ""
	case 2:
		return ""
	case 3:
		return ""
	case 4:
		return "required"
	}
	return ""
}

func (*SmallStruct) PhoneSinceVersion() uint16 {
	return 0
}

func (s *SmallStruct) PhoneInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PhoneSinceVersion()
}

func (*SmallStruct) PhoneDeprecated() uint16 {
	return 0
}

func (SmallStruct) PhoneCharacterEncoding() string {
	return "UTF-8"
}

func (SmallStruct) PhoneHeaderLength() uint64 {
	return 4
}
