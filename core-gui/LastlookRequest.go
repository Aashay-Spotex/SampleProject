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

type LastlookRequest struct {
	AcctId     int32
	OrderQty   float64
	Price      float64
	InstId     uint16
	TradeDate  int32
	ValueDate  int32
	CtrAcctId  int32
	TradeTime  uint64
	OrderType  uint16
	ProxyIndex uint16
	Fok        uint8
	Closing    uint8
	RequestId  []uint8
	TradeId    []uint8
	OrderCurr  []uint8
}

func (l *LastlookRequest) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, l.AcctId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, l.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, l.Price); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, l.InstId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, l.TradeDate); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, l.ValueDate); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, l.CtrAcctId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, l.TradeTime); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, l.OrderType); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, l.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, l.Fok); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, l.Closing); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.RequestId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.RequestId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.TradeId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.TradeId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.OrderCurr))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.OrderCurr); err != nil {
		return err
	}
	return nil
}

func (l *LastlookRequest) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !l.AcctIdInActingVersion(actingVersion) {
		l.AcctId = l.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.AcctId); err != nil {
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
	if !l.PriceInActingVersion(actingVersion) {
		l.Price = l.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &l.Price); err != nil {
			return err
		}
	}
	if !l.InstIdInActingVersion(actingVersion) {
		l.InstId = l.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &l.InstId); err != nil {
			return err
		}
	}
	if !l.TradeDateInActingVersion(actingVersion) {
		l.TradeDate = l.TradeDateNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.TradeDate); err != nil {
			return err
		}
	}
	if !l.ValueDateInActingVersion(actingVersion) {
		l.ValueDate = l.ValueDateNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.ValueDate); err != nil {
			return err
		}
	}
	if !l.CtrAcctIdInActingVersion(actingVersion) {
		l.CtrAcctId = l.CtrAcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.CtrAcctId); err != nil {
			return err
		}
	}
	if !l.TradeTimeInActingVersion(actingVersion) {
		l.TradeTime = l.TradeTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &l.TradeTime); err != nil {
			return err
		}
	}
	if !l.OrderTypeInActingVersion(actingVersion) {
		l.OrderType = l.OrderTypeNullValue()
	} else {
		if err := _m.ReadUint16(_r, &l.OrderType); err != nil {
			return err
		}
	}
	if !l.ProxyIndexInActingVersion(actingVersion) {
		l.ProxyIndex = l.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &l.ProxyIndex); err != nil {
			return err
		}
	}
	if !l.FokInActingVersion(actingVersion) {
		l.Fok = l.FokNullValue()
	} else {
		if err := _m.ReadUint8(_r, &l.Fok); err != nil {
			return err
		}
	}
	if !l.ClosingInActingVersion(actingVersion) {
		l.Closing = l.ClosingNullValue()
	} else {
		if err := _m.ReadUint8(_r, &l.Closing); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.RequestIdInActingVersion(actingVersion) {
		var RequestIdLength uint16
		if err := _m.ReadUint16(_r, &RequestIdLength); err != nil {
			return err
		}
		if cap(l.RequestId) < int(RequestIdLength) {
			l.RequestId = make([]uint8, RequestIdLength)
		}
		l.RequestId = l.RequestId[:RequestIdLength]
		if err := _m.ReadBytes(_r, l.RequestId); err != nil {
			return err
		}
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

	if l.OrderCurrInActingVersion(actingVersion) {
		var OrderCurrLength uint16
		if err := _m.ReadUint16(_r, &OrderCurrLength); err != nil {
			return err
		}
		if cap(l.OrderCurr) < int(OrderCurrLength) {
			l.OrderCurr = make([]uint8, OrderCurrLength)
		}
		l.OrderCurr = l.OrderCurr[:OrderCurrLength]
		if err := _m.ReadBytes(_r, l.OrderCurr); err != nil {
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

func (l *LastlookRequest) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.AcctIdInActingVersion(actingVersion) {
		if l.AcctId < l.AcctIdMinValue() || l.AcctId > l.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on l.AcctId (%v < %v > %v)", l.AcctIdMinValue(), l.AcctId, l.AcctIdMaxValue())
		}
	}
	if l.OrderQtyInActingVersion(actingVersion) {
		if l.OrderQty < l.OrderQtyMinValue() || l.OrderQty > l.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on l.OrderQty (%v < %v > %v)", l.OrderQtyMinValue(), l.OrderQty, l.OrderQtyMaxValue())
		}
	}
	if l.PriceInActingVersion(actingVersion) {
		if l.Price < l.PriceMinValue() || l.Price > l.PriceMaxValue() {
			return fmt.Errorf("Range check failed on l.Price (%v < %v > %v)", l.PriceMinValue(), l.Price, l.PriceMaxValue())
		}
	}
	if l.InstIdInActingVersion(actingVersion) {
		if l.InstId < l.InstIdMinValue() || l.InstId > l.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on l.InstId (%v < %v > %v)", l.InstIdMinValue(), l.InstId, l.InstIdMaxValue())
		}
	}
	if l.TradeDateInActingVersion(actingVersion) {
		if l.TradeDate < l.TradeDateMinValue() || l.TradeDate > l.TradeDateMaxValue() {
			return fmt.Errorf("Range check failed on l.TradeDate (%v < %v > %v)", l.TradeDateMinValue(), l.TradeDate, l.TradeDateMaxValue())
		}
	}
	if l.ValueDateInActingVersion(actingVersion) {
		if l.ValueDate < l.ValueDateMinValue() || l.ValueDate > l.ValueDateMaxValue() {
			return fmt.Errorf("Range check failed on l.ValueDate (%v < %v > %v)", l.ValueDateMinValue(), l.ValueDate, l.ValueDateMaxValue())
		}
	}
	if l.CtrAcctIdInActingVersion(actingVersion) {
		if l.CtrAcctId < l.CtrAcctIdMinValue() || l.CtrAcctId > l.CtrAcctIdMaxValue() {
			return fmt.Errorf("Range check failed on l.CtrAcctId (%v < %v > %v)", l.CtrAcctIdMinValue(), l.CtrAcctId, l.CtrAcctIdMaxValue())
		}
	}
	if l.TradeTimeInActingVersion(actingVersion) {
		if l.TradeTime < l.TradeTimeMinValue() || l.TradeTime > l.TradeTimeMaxValue() {
			return fmt.Errorf("Range check failed on l.TradeTime (%v < %v > %v)", l.TradeTimeMinValue(), l.TradeTime, l.TradeTimeMaxValue())
		}
	}
	if l.OrderTypeInActingVersion(actingVersion) {
		if l.OrderType < l.OrderTypeMinValue() || l.OrderType > l.OrderTypeMaxValue() {
			return fmt.Errorf("Range check failed on l.OrderType (%v < %v > %v)", l.OrderTypeMinValue(), l.OrderType, l.OrderTypeMaxValue())
		}
	}
	if l.ProxyIndexInActingVersion(actingVersion) {
		if l.ProxyIndex < l.ProxyIndexMinValue() || l.ProxyIndex > l.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on l.ProxyIndex (%v < %v > %v)", l.ProxyIndexMinValue(), l.ProxyIndex, l.ProxyIndexMaxValue())
		}
	}
	if l.FokInActingVersion(actingVersion) {
		if l.Fok < l.FokMinValue() || l.Fok > l.FokMaxValue() {
			return fmt.Errorf("Range check failed on l.Fok (%v < %v > %v)", l.FokMinValue(), l.Fok, l.FokMaxValue())
		}
	}
	if l.ClosingInActingVersion(actingVersion) {
		if l.Closing < l.ClosingMinValue() || l.Closing > l.ClosingMaxValue() {
			return fmt.Errorf("Range check failed on l.Closing (%v < %v > %v)", l.ClosingMinValue(), l.Closing, l.ClosingMaxValue())
		}
	}
	if !utf8.Valid(l.RequestId[:]) {
		return errors.New("l.RequestId failed UTF-8 validation")
	}
	if !utf8.Valid(l.TradeId[:]) {
		return errors.New("l.TradeId failed UTF-8 validation")
	}
	if !utf8.Valid(l.OrderCurr[:]) {
		return errors.New("l.OrderCurr failed UTF-8 validation")
	}
	return nil
}

func LastlookRequestInit(l *LastlookRequest) {
	return
}

func (*LastlookRequest) SbeBlockLength() (blockLength uint16) {
	return 48
}

func (*LastlookRequest) SbeTemplateId() (templateId uint16) {
	return 43
}

func (*LastlookRequest) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*LastlookRequest) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*LastlookRequest) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*LastlookRequest) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*LastlookRequest) AcctIdId() uint16 {
	return 44
}

func (*LastlookRequest) AcctIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.AcctIdSinceVersion()
}

func (*LastlookRequest) AcctIdDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) AcctIdMetaAttribute(meta int) string {
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

func (*LastlookRequest) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookRequest) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookRequest) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*LastlookRequest) OrderQtyId() uint16 {
	return 45
}

func (*LastlookRequest) OrderQtySinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrderQtySinceVersion()
}

func (*LastlookRequest) OrderQtyDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) OrderQtyMetaAttribute(meta int) string {
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

func (*LastlookRequest) OrderQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*LastlookRequest) OrderQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*LastlookRequest) OrderQtyNullValue() float64 {
	return math.NaN()
}

func (*LastlookRequest) PriceId() uint16 {
	return 46
}

func (*LastlookRequest) PriceSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.PriceSinceVersion()
}

func (*LastlookRequest) PriceDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) PriceMetaAttribute(meta int) string {
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

func (*LastlookRequest) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*LastlookRequest) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*LastlookRequest) PriceNullValue() float64 {
	return math.NaN()
}

func (*LastlookRequest) InstIdId() uint16 {
	return 47
}

func (*LastlookRequest) InstIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.InstIdSinceVersion()
}

func (*LastlookRequest) InstIdDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) InstIdMetaAttribute(meta int) string {
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

func (*LastlookRequest) InstIdMinValue() uint16 {
	return 0
}

func (*LastlookRequest) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*LastlookRequest) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*LastlookRequest) TradeDateId() uint16 {
	return 48
}

func (*LastlookRequest) TradeDateSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TradeDateSinceVersion()
}

func (*LastlookRequest) TradeDateDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) TradeDateMetaAttribute(meta int) string {
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

func (*LastlookRequest) TradeDateMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookRequest) TradeDateMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookRequest) TradeDateNullValue() int32 {
	return math.MinInt32
}

func (*LastlookRequest) ValueDateId() uint16 {
	return 49
}

func (*LastlookRequest) ValueDateSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) ValueDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ValueDateSinceVersion()
}

func (*LastlookRequest) ValueDateDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) ValueDateMetaAttribute(meta int) string {
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

func (*LastlookRequest) ValueDateMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookRequest) ValueDateMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookRequest) ValueDateNullValue() int32 {
	return math.MinInt32
}

func (*LastlookRequest) CtrAcctIdId() uint16 {
	return 50
}

func (*LastlookRequest) CtrAcctIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) CtrAcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.CtrAcctIdSinceVersion()
}

func (*LastlookRequest) CtrAcctIdDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) CtrAcctIdMetaAttribute(meta int) string {
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

func (*LastlookRequest) CtrAcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookRequest) CtrAcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookRequest) CtrAcctIdNullValue() int32 {
	return math.MinInt32
}

func (*LastlookRequest) TradeTimeId() uint16 {
	return 51
}

func (*LastlookRequest) TradeTimeSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) TradeTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TradeTimeSinceVersion()
}

func (*LastlookRequest) TradeTimeDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) TradeTimeMetaAttribute(meta int) string {
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

func (*LastlookRequest) TradeTimeMinValue() uint64 {
	return 0
}

func (*LastlookRequest) TradeTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*LastlookRequest) TradeTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*LastlookRequest) OrderTypeId() uint16 {
	return 52
}

func (*LastlookRequest) OrderTypeSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrderTypeSinceVersion()
}

func (*LastlookRequest) OrderTypeDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) OrderTypeMetaAttribute(meta int) string {
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

func (*LastlookRequest) OrderTypeMinValue() uint16 {
	return 0
}

func (*LastlookRequest) OrderTypeMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*LastlookRequest) OrderTypeNullValue() uint16 {
	return math.MaxUint16
}

func (*LastlookRequest) ProxyIndexId() uint16 {
	return 53
}

func (*LastlookRequest) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ProxyIndexSinceVersion()
}

func (*LastlookRequest) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) ProxyIndexMetaAttribute(meta int) string {
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

func (*LastlookRequest) ProxyIndexMinValue() uint16 {
	return 0
}

func (*LastlookRequest) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*LastlookRequest) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*LastlookRequest) FokId() uint16 {
	return 54
}

func (*LastlookRequest) FokSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) FokInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.FokSinceVersion()
}

func (*LastlookRequest) FokDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) FokMetaAttribute(meta int) string {
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

func (*LastlookRequest) FokMinValue() uint8 {
	return 0
}

func (*LastlookRequest) FokMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*LastlookRequest) FokNullValue() uint8 {
	return math.MaxUint8
}

func (*LastlookRequest) ClosingId() uint16 {
	return 55
}

func (*LastlookRequest) ClosingSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) ClosingInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ClosingSinceVersion()
}

func (*LastlookRequest) ClosingDeprecated() uint16 {
	return 0
}

func (*LastlookRequest) ClosingMetaAttribute(meta int) string {
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

func (*LastlookRequest) ClosingMinValue() uint8 {
	return 0
}

func (*LastlookRequest) ClosingMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*LastlookRequest) ClosingNullValue() uint8 {
	return math.MaxUint8
}

func (*LastlookRequest) RequestIdMetaAttribute(meta int) string {
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

func (*LastlookRequest) RequestIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.RequestIdSinceVersion()
}

func (*LastlookRequest) RequestIdDeprecated() uint16 {
	return 0
}

func (LastlookRequest) RequestIdCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookRequest) RequestIdHeaderLength() uint64 {
	return 2
}

func (*LastlookRequest) TradeIdMetaAttribute(meta int) string {
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

func (*LastlookRequest) TradeIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) TradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TradeIdSinceVersion()
}

func (*LastlookRequest) TradeIdDeprecated() uint16 {
	return 0
}

func (LastlookRequest) TradeIdCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookRequest) TradeIdHeaderLength() uint64 {
	return 2
}

func (*LastlookRequest) OrderCurrMetaAttribute(meta int) string {
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

func (*LastlookRequest) OrderCurrSinceVersion() uint16 {
	return 0
}

func (l *LastlookRequest) OrderCurrInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.OrderCurrSinceVersion()
}

func (*LastlookRequest) OrderCurrDeprecated() uint16 {
	return 0
}

func (LastlookRequest) OrderCurrCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookRequest) OrderCurrHeaderLength() uint64 {
	return 2
}
