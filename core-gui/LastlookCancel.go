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

type LastlookCancel struct {
	ProxyIndex int32
	AcctId     int32
	InstId     int32
	OrderQty   float64
	TradeId    []uint8
}

func (l *LastlookCancel) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, l.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, l.AcctId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, l.InstId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, l.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.TradeId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.TradeId); err != nil {
		return err
	}
	return nil
}

func (l *LastlookCancel) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !l.ProxyIndexInActingVersion(actingVersion) {
		l.ProxyIndex = l.ProxyIndexNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.ProxyIndex); err != nil {
			return err
		}
	}
	if !l.AcctIdInActingVersion(actingVersion) {
		l.AcctId = l.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.AcctId); err != nil {
			return err
		}
	}
	if !l.InstIdInActingVersion(actingVersion) {
		l.InstId = l.InstIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.InstId); err != nil {
			return err
		}
	}
	if !l.OrderQtyInActingVersion(actingVersion) {
		l.OrderQty = l.OrderQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &l.OrderQty); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.TradeIdInActingVersion(actingVersion) {
		var TradeIdLength uint16
		if err := _m.ReadUint16(_r, &TradeIdLength); err != nil {
			return err
		}
		if cap(l.TradeId) < int(TradeIdLength) {
			l.TradeId = make([]uint8, TradeIdLength)
		}
		l.TradeId = l.TradeId[:TradeIdLength]
		if err := _m.ReadBytes(_r, l.TradeId); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := l.RangeCheck(actingVersion, l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (l *LastlookCancel) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.ProxyIndexInActingVersion(actingVersion) {
		if l.ProxyIndex < l.ProxyIndexMinValue() || l.ProxyIndex > l.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on l.ProxyIndex (%v < %v > %v)", l.ProxyIndexMinValue(), l.ProxyIndex, l.ProxyIndexMaxValue())
		}
	}
	if l.AcctIdInActingVersion(actingVersion) {
		if l.AcctId < l.AcctIdMinValue() || l.AcctId > l.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on l.AcctId (%v < %v > %v)", l.AcctIdMinValue(), l.AcctId, l.AcctIdMaxValue())
		}
	}
	if l.InstIdInActingVersion(actingVersion) {
		if l.InstId < l.InstIdMinValue() || l.InstId > l.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on l.InstId (%v < %v > %v)", l.InstIdMinValue(), l.InstId, l.InstIdMaxValue())
		}
	}
	if l.OrderQtyInActingVersion(actingVersion) {
		if l.OrderQty < l.OrderQtyMinValue() || l.OrderQty > l.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on l.OrderQty (%v < %v > %v)", l.OrderQtyMinValue(), l.OrderQty, l.OrderQtyMaxValue())
		}
	}
	if !utf8.Valid(l.TradeId[:]) {
		return errors.New("l.TradeId failed UTF-8 validation")
	}
	return nil
}

func LastlookCancelInit(l *LastlookCancel) {
	return
}

func (*LastlookCancel) SbeBlockLength() (blockLength uint16) {
	return 20
}

func (*LastlookCancel) SbeTemplateId() (templateId uint16) {
	return 270
}

func (*LastlookCancel) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*LastlookCancel) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*LastlookCancel) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*LastlookCancel) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*LastlookCancel) ProxyIndexId() uint16 {
	return 271
}

func (*LastlookCancel) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (l *LastlookCancel) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ProxyIndexSinceVersion()
}

func (*LastlookCancel) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*LastlookCancel) ProxyIndexMetaAttribute(meta int) string {
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

func (*LastlookCancel) ProxyIndexMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookCancel) ProxyIndexMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookCancel) ProxyIndexNullValue() int32 {
	return math.MinInt32
}

func (*LastlookCancel) AcctIdId() uint16 {
	return 272
}

func (*LastlookCancel) AcctIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookCancel) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.AcctIdSinceVersion()
}

func (*LastlookCancel) AcctIdDeprecated() uint16 {
	return 0
}

func (*LastlookCancel) AcctIdMetaAttribute(meta int) string {
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

func (*LastlookCancel) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookCancel) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookCancel) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*LastlookCancel) InstIdId() uint16 {
	return 273
}

func (*LastlookCancel) InstIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookCancel) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.InstIdSinceVersion()
}

func (*LastlookCancel) InstIdDeprecated() uint16 {
	return 0
}

func (*LastlookCancel) InstIdMetaAttribute(meta int) string {
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

func (*LastlookCancel) InstIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookCancel) InstIdMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookCancel) InstIdNullValue() int32 {
	return math.MinInt32
}

func (*LastlookCancel) OrderQtyId() uint16 {
	return 274
}

func (*LastlookCancel) OrderQtySinceVersion() uint16 {
	return 0
}

func (l *LastlookCancel) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrderQtySinceVersion()
}

func (*LastlookCancel) OrderQtyDeprecated() uint16 {
	return 0
}

func (*LastlookCancel) OrderQtyMetaAttribute(meta int) string {
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

func (*LastlookCancel) OrderQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*LastlookCancel) OrderQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*LastlookCancel) OrderQtyNullValue() float64 {
	return math.NaN()
}

func (*LastlookCancel) TradeIdMetaAttribute(meta int) string {
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

func (*LastlookCancel) TradeIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookCancel) TradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TradeIdSinceVersion()
}

func (*LastlookCancel) TradeIdDeprecated() uint16 {
	return 0
}

func (LastlookCancel) TradeIdCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookCancel) TradeIdHeaderLength() uint64 {
	return 2
}
