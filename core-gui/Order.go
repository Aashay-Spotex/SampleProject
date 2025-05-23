// Generated SBE (Simple Binary Encoding) message codec

package core_gui

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"unicode/utf8"
)

type Order struct {
	AcctId   int32
	DealQty  float64
	CtrQty   float64
	CurrPair []uint8
}

func (o *Order) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, o.AcctId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.DealQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.CtrQty); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.CurrPair))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.CurrPair); err != nil {
		return err
	}
	return nil
}

func (o *Order) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.AcctIdInActingVersion(actingVersion) {
		o.AcctId = o.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.AcctId); err != nil {
			return err
		}
	}
	if !o.DealQtyInActingVersion(actingVersion) {
		o.DealQty = o.DealQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.DealQty); err != nil {
			return err
		}
	}
	if !o.CtrQtyInActingVersion(actingVersion) {
		o.CtrQty = o.CtrQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.CtrQty); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}

	if o.CurrPairInActingVersion(actingVersion) {
		var CurrPairLength uint16
		if err := _m.ReadUint16(_r, &CurrPairLength); err != nil {
			return err
		}
		if cap(o.CurrPair) < int(CurrPairLength) {
			o.CurrPair = make([]uint8, CurrPairLength)
		}
		o.CurrPair = o.CurrPair[:CurrPairLength]
		if err := _m.ReadBytes(_r, o.CurrPair); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := o.RangeCheck(actingVersion, o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (o *Order) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.AcctIdInActingVersion(actingVersion) {
		if o.AcctId < o.AcctIdMinValue() || o.AcctId > o.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on o.AcctId (%v < %v > %v)", o.AcctIdMinValue(), o.AcctId, o.AcctIdMaxValue())
		}
	}
	if o.DealQtyInActingVersion(actingVersion) {
		if o.DealQty < o.DealQtyMinValue() || o.DealQty > o.DealQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.DealQty (%v < %v > %v)", o.DealQtyMinValue(), o.DealQty, o.DealQtyMaxValue())
		}
	}
	if o.CtrQtyInActingVersion(actingVersion) {
		if o.CtrQty < o.CtrQtyMinValue() || o.CtrQty > o.CtrQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.CtrQty (%v < %v > %v)", o.CtrQtyMinValue(), o.CtrQty, o.CtrQtyMaxValue())
		}
	}
	if !utf8.Valid(o.CurrPair[:]) {
		return errors.New("o.CurrPair failed UTF-8 validation")
	}
	return nil
}

func OrderInit(o *Order) {
	return
}

func (*Order) SbeBlockLength() (blockLength uint16) {
	return 20
}

func (*Order) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*Order) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Order) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Order) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Order) SbeSemanticVersion() (semanticVersion string) {
	return "1.0"
}

func (*Order) AcctIdId() uint16 {
	return 2
}

func (*Order) AcctIdSinceVersion() uint16 {
	return 0
}

func (o *Order) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AcctIdSinceVersion()
}

func (*Order) AcctIdDeprecated() uint16 {
	return 0
}

func (*Order) AcctIdMetaAttribute(meta int) string {
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

func (*Order) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*Order) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*Order) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*Order) DealQtyId() uint16 {
	return 3
}

func (*Order) DealQtySinceVersion() uint16 {
	return 0
}

func (o *Order) DealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.DealQtySinceVersion()
}

func (*Order) DealQtyDeprecated() uint16 {
	return 0
}

func (*Order) DealQtyMetaAttribute(meta int) string {
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

func (*Order) DealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*Order) DealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*Order) DealQtyNullValue() float64 {
	return math.NaN()
}

func (*Order) CtrQtyId() uint16 {
	return 4
}

func (*Order) CtrQtySinceVersion() uint16 {
	return 0
}

func (o *Order) CtrQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CtrQtySinceVersion()
}

func (*Order) CtrQtyDeprecated() uint16 {
	return 0
}

func (*Order) CtrQtyMetaAttribute(meta int) string {
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

func (*Order) CtrQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*Order) CtrQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*Order) CtrQtyNullValue() float64 {
	return math.NaN()
}

func (*Order) CurrPairMetaAttribute(meta int) string {
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

func (*Order) CurrPairSinceVersion() uint16 {
	return 0
}

func (o *Order) CurrPairInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CurrPairSinceVersion()
}

func (*Order) CurrPairDeprecated() uint16 {
	return 0
}

func (Order) CurrPairCharacterEncoding() string {
	return "UTF-8"
}

func (Order) CurrPairHeaderLength() uint64 {
	return 2
}
