// Generated SBE (Simple Binary Encoding) message codec

package small

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SmallStruct struct {
	Name     [50]byte
	BirthDay int64
	Phone    [20]byte
	Siblings int32
	Spouse   uint8
	Money    float64
}

func (s *SmallStruct) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteBytes(_w, s.Name[:]); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, s.BirthDay); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, s.Phone[:]); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.Siblings); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.Spouse); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, s.Money); err != nil {
		return err
	}
	return nil
}

func (s *SmallStruct) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.NameInActingVersion(actingVersion) {
		for idx := 0; idx < 50; idx++ {
			s.Name[idx] = s.NameNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.Name[:]); err != nil {
			return err
		}
	}
	if !s.BirthDayInActingVersion(actingVersion) {
		s.BirthDay = s.BirthDayNullValue()
	} else {
		if err := _m.ReadInt64(_r, &s.BirthDay); err != nil {
			return err
		}
	}
	if !s.PhoneInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			s.Phone[idx] = s.PhoneNullValue()
		}
	} else {
		if err := _m.ReadBytes(_r, s.Phone[:]); err != nil {
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
	if !s.SpouseInActingVersion(actingVersion) {
		s.Spouse = s.SpouseNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.Spouse); err != nil {
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
	if doRangeCheck {
		if err := s.RangeCheck(actingVersion, s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (s *SmallStruct) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.NameInActingVersion(actingVersion) {
		for idx := 0; idx < 50; idx++ {
			if s.Name[idx] < s.NameMinValue() || s.Name[idx] > s.NameMaxValue() {
				return fmt.Errorf("Range check failed on s.Name[%d] (%v < %v > %v)", idx, s.NameMinValue(), s.Name[idx], s.NameMaxValue())
			}
		}
	}
	for idx, ch := range s.Name {
		if ch > 127 {
			return fmt.Errorf("s.Name[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if s.BirthDayInActingVersion(actingVersion) {
		if s.BirthDay < s.BirthDayMinValue() || s.BirthDay > s.BirthDayMaxValue() {
			return fmt.Errorf("Range check failed on s.BirthDay (%v < %v > %v)", s.BirthDayMinValue(), s.BirthDay, s.BirthDayMaxValue())
		}
	}
	if s.PhoneInActingVersion(actingVersion) {
		for idx := 0; idx < 20; idx++ {
			if s.Phone[idx] < s.PhoneMinValue() || s.Phone[idx] > s.PhoneMaxValue() {
				return fmt.Errorf("Range check failed on s.Phone[%d] (%v < %v > %v)", idx, s.PhoneMinValue(), s.Phone[idx], s.PhoneMaxValue())
			}
		}
	}
	for idx, ch := range s.Phone {
		if ch > 127 {
			return fmt.Errorf("s.Phone[%d]=%d failed ASCII validation", idx, ch)
		}
	}
	if s.SiblingsInActingVersion(actingVersion) {
		if s.Siblings < s.SiblingsMinValue() || s.Siblings > s.SiblingsMaxValue() {
			return fmt.Errorf("Range check failed on s.Siblings (%v < %v > %v)", s.SiblingsMinValue(), s.Siblings, s.SiblingsMaxValue())
		}
	}
	if s.SpouseInActingVersion(actingVersion) {
		if s.Spouse < s.SpouseMinValue() || s.Spouse > s.SpouseMaxValue() {
			return fmt.Errorf("Range check failed on s.Spouse (%v < %v > %v)", s.SpouseMinValue(), s.Spouse, s.SpouseMaxValue())
		}
	}
	if s.MoneyInActingVersion(actingVersion) {
		if s.Money < s.MoneyMinValue() || s.Money > s.MoneyMaxValue() {
			return fmt.Errorf("Range check failed on s.Money (%v < %v > %v)", s.MoneyMinValue(), s.Money, s.MoneyMaxValue())
		}
	}
	return nil
}

func SmallStructInit(s *SmallStruct) {
	return
}

func (*SmallStruct) SbeBlockLength() (blockLength uint16) {
	return 91
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

func (*SmallStruct) NameId() uint16 {
	return 1
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

func (*SmallStruct) NameMinValue() byte {
	return byte(32)
}

func (*SmallStruct) NameMaxValue() byte {
	return byte(126)
}

func (*SmallStruct) NameNullValue() byte {
	return 0
}

func (s *SmallStruct) NameCharacterEncoding() string {
	return "ASCII"
}

func (*SmallStruct) BirthDayId() uint16 {
	return 2
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

func (*SmallStruct) PhoneId() uint16 {
	return 3
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

func (*SmallStruct) PhoneMinValue() byte {
	return byte(32)
}

func (*SmallStruct) PhoneMaxValue() byte {
	return byte(126)
}

func (*SmallStruct) PhoneNullValue() byte {
	return 0
}

func (s *SmallStruct) PhoneCharacterEncoding() string {
	return "ASCII"
}

func (*SmallStruct) SiblingsId() uint16 {
	return 4
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
	return 5
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

func (*SmallStruct) SpouseMinValue() uint8 {
	return 0
}

func (*SmallStruct) SpouseMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SmallStruct) SpouseNullValue() uint8 {
	return math.MaxUint8
}

func (*SmallStruct) MoneyId() uint16 {
	return 6
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
