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

type TradeConfirm struct {
	OrderId    int64
	LinkId     uint64
	AcctId     int32
	CtrAcctId  int32
	InstId     uint16
	DealQty    float64
	CtrQty     float64
	PbDealQty  float64
	PbCtrQty   float64
	Price      float64
	PbPrice    float64
	TradeDate  int32
	ValueDate  int32
	TradeTime  uint64
	OrderQty   float64
	LimitPrice float64
	StopPrice  float64
	OpenQty    float64
	AvgPrice   float64
	OrderType  uint8
	Brokerage  float64
	CreatedBy  int32
	PlatformId int32
	ProxyIndex uint16
	Complete   uint8
	RequestId  []uint8
	TradeId    []uint8
	OrderCurr  []uint8
}

func (t *TradeConfirm) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, t.OrderId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, t.LinkId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.AcctId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.CtrAcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, t.InstId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.DealQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.CtrQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.PbDealQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.PbCtrQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.PbPrice); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.TradeDate); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.ValueDate); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, t.TradeTime); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.LimitPrice); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.OpenQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.AvgPrice); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, t.OrderType); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.Brokerage); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.CreatedBy); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.PlatformId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, t.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, t.Complete); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(t.RequestId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.RequestId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(t.TradeId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.TradeId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(t.OrderCurr))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.OrderCurr); err != nil {
		return err
	}
	return nil
}

func (t *TradeConfirm) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !t.OrderIdInActingVersion(actingVersion) {
		t.OrderId = t.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &t.OrderId); err != nil {
			return err
		}
	}
	if !t.LinkIdInActingVersion(actingVersion) {
		t.LinkId = t.LinkIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &t.LinkId); err != nil {
			return err
		}
	}
	if !t.AcctIdInActingVersion(actingVersion) {
		t.AcctId = t.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.AcctId); err != nil {
			return err
		}
	}
	if !t.CtrAcctIdInActingVersion(actingVersion) {
		t.CtrAcctId = t.CtrAcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.CtrAcctId); err != nil {
			return err
		}
	}
	if !t.InstIdInActingVersion(actingVersion) {
		t.InstId = t.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &t.InstId); err != nil {
			return err
		}
	}
	if !t.DealQtyInActingVersion(actingVersion) {
		t.DealQty = t.DealQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.DealQty); err != nil {
			return err
		}
	}
	if !t.CtrQtyInActingVersion(actingVersion) {
		t.CtrQty = t.CtrQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.CtrQty); err != nil {
			return err
		}
	}
	if !t.PbDealQtyInActingVersion(actingVersion) {
		t.PbDealQty = t.PbDealQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.PbDealQty); err != nil {
			return err
		}
	}
	if !t.PbCtrQtyInActingVersion(actingVersion) {
		t.PbCtrQty = t.PbCtrQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.PbCtrQty); err != nil {
			return err
		}
	}
	if !t.PriceInActingVersion(actingVersion) {
		t.Price = t.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.Price); err != nil {
			return err
		}
	}
	if !t.PbPriceInActingVersion(actingVersion) {
		t.PbPrice = t.PbPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.PbPrice); err != nil {
			return err
		}
	}
	if !t.TradeDateInActingVersion(actingVersion) {
		t.TradeDate = t.TradeDateNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.TradeDate); err != nil {
			return err
		}
	}
	if !t.ValueDateInActingVersion(actingVersion) {
		t.ValueDate = t.ValueDateNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.ValueDate); err != nil {
			return err
		}
	}
	if !t.TradeTimeInActingVersion(actingVersion) {
		t.TradeTime = t.TradeTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &t.TradeTime); err != nil {
			return err
		}
	}
	if !t.OrderQtyInActingVersion(actingVersion) {
		t.OrderQty = t.OrderQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.OrderQty); err != nil {
			return err
		}
	}
	if !t.LimitPriceInActingVersion(actingVersion) {
		t.LimitPrice = t.LimitPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.LimitPrice); err != nil {
			return err
		}
	}
	if !t.StopPriceInActingVersion(actingVersion) {
		t.StopPrice = t.StopPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.StopPrice); err != nil {
			return err
		}
	}
	if !t.OpenQtyInActingVersion(actingVersion) {
		t.OpenQty = t.OpenQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.OpenQty); err != nil {
			return err
		}
	}
	if !t.AvgPriceInActingVersion(actingVersion) {
		t.AvgPrice = t.AvgPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.AvgPrice); err != nil {
			return err
		}
	}
	if !t.OrderTypeInActingVersion(actingVersion) {
		t.OrderType = t.OrderTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &t.OrderType); err != nil {
			return err
		}
	}
	if !t.BrokerageInActingVersion(actingVersion) {
		t.Brokerage = t.BrokerageNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.Brokerage); err != nil {
			return err
		}
	}
	if !t.CreatedByInActingVersion(actingVersion) {
		t.CreatedBy = t.CreatedByNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.CreatedBy); err != nil {
			return err
		}
	}
	if !t.PlatformIdInActingVersion(actingVersion) {
		t.PlatformId = t.PlatformIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.PlatformId); err != nil {
			return err
		}
	}
	if !t.ProxyIndexInActingVersion(actingVersion) {
		t.ProxyIndex = t.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &t.ProxyIndex); err != nil {
			return err
		}
	}
	if !t.CompleteInActingVersion(actingVersion) {
		t.Complete = t.CompleteNullValue()
	} else {
		if err := _m.ReadUint8(_r, &t.Complete); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.RequestIdInActingVersion(actingVersion) {
		var RequestIdLength uint16
		if err := _m.ReadUint16(_r, &RequestIdLength); err != nil {
			return err
		}
		if cap(t.RequestId) < int(RequestIdLength) {
			t.RequestId = make([]uint8, RequestIdLength)
		}
		t.RequestId = t.RequestId[:RequestIdLength]
		if err := _m.ReadBytes(_r, t.RequestId); err != nil {
			return err
		}
	}

	if t.TradeIdInActingVersion(actingVersion) {
		var TradeIdLength uint16
		if err := _m.ReadUint16(_r, &TradeIdLength); err != nil {
			return err
		}
		if cap(t.TradeId) < int(TradeIdLength) {
			t.TradeId = make([]uint8, TradeIdLength)
		}
		t.TradeId = t.TradeId[:TradeIdLength]
		if err := _m.ReadBytes(_r, t.TradeId); err != nil {
			return err
		}
	}

	if t.OrderCurrInActingVersion(actingVersion) {
		var OrderCurrLength uint16
		if err := _m.ReadUint16(_r, &OrderCurrLength); err != nil {
			return err
		}
		if cap(t.OrderCurr) < int(OrderCurrLength) {
			t.OrderCurr = make([]uint8, OrderCurrLength)
		}
		t.OrderCurr = t.OrderCurr[:OrderCurrLength]
		if err := _m.ReadBytes(_r, t.OrderCurr); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TradeConfirm) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.OrderIdInActingVersion(actingVersion) {
		if t.OrderId < t.OrderIdMinValue() || t.OrderId > t.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on t.OrderId (%v < %v > %v)", t.OrderIdMinValue(), t.OrderId, t.OrderIdMaxValue())
		}
	}
	if t.LinkIdInActingVersion(actingVersion) {
		if t.LinkId < t.LinkIdMinValue() || t.LinkId > t.LinkIdMaxValue() {
			return fmt.Errorf("Range check failed on t.LinkId (%v < %v > %v)", t.LinkIdMinValue(), t.LinkId, t.LinkIdMaxValue())
		}
	}
	if t.AcctIdInActingVersion(actingVersion) {
		if t.AcctId < t.AcctIdMinValue() || t.AcctId > t.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on t.AcctId (%v < %v > %v)", t.AcctIdMinValue(), t.AcctId, t.AcctIdMaxValue())
		}
	}
	if t.CtrAcctIdInActingVersion(actingVersion) {
		if t.CtrAcctId < t.CtrAcctIdMinValue() || t.CtrAcctId > t.CtrAcctIdMaxValue() {
			return fmt.Errorf("Range check failed on t.CtrAcctId (%v < %v > %v)", t.CtrAcctIdMinValue(), t.CtrAcctId, t.CtrAcctIdMaxValue())
		}
	}
	if t.InstIdInActingVersion(actingVersion) {
		if t.InstId < t.InstIdMinValue() || t.InstId > t.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on t.InstId (%v < %v > %v)", t.InstIdMinValue(), t.InstId, t.InstIdMaxValue())
		}
	}
	if t.DealQtyInActingVersion(actingVersion) {
		if t.DealQty < t.DealQtyMinValue() || t.DealQty > t.DealQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.DealQty (%v < %v > %v)", t.DealQtyMinValue(), t.DealQty, t.DealQtyMaxValue())
		}
	}
	if t.CtrQtyInActingVersion(actingVersion) {
		if t.CtrQty < t.CtrQtyMinValue() || t.CtrQty > t.CtrQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.CtrQty (%v < %v > %v)", t.CtrQtyMinValue(), t.CtrQty, t.CtrQtyMaxValue())
		}
	}
	if t.PbDealQtyInActingVersion(actingVersion) {
		if t.PbDealQty < t.PbDealQtyMinValue() || t.PbDealQty > t.PbDealQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.PbDealQty (%v < %v > %v)", t.PbDealQtyMinValue(), t.PbDealQty, t.PbDealQtyMaxValue())
		}
	}
	if t.PbCtrQtyInActingVersion(actingVersion) {
		if t.PbCtrQty < t.PbCtrQtyMinValue() || t.PbCtrQty > t.PbCtrQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.PbCtrQty (%v < %v > %v)", t.PbCtrQtyMinValue(), t.PbCtrQty, t.PbCtrQtyMaxValue())
		}
	}
	if t.PriceInActingVersion(actingVersion) {
		if t.Price < t.PriceMinValue() || t.Price > t.PriceMaxValue() {
			return fmt.Errorf("Range check failed on t.Price (%v < %v > %v)", t.PriceMinValue(), t.Price, t.PriceMaxValue())
		}
	}
	if t.PbPriceInActingVersion(actingVersion) {
		if t.PbPrice < t.PbPriceMinValue() || t.PbPrice > t.PbPriceMaxValue() {
			return fmt.Errorf("Range check failed on t.PbPrice (%v < %v > %v)", t.PbPriceMinValue(), t.PbPrice, t.PbPriceMaxValue())
		}
	}
	if t.TradeDateInActingVersion(actingVersion) {
		if t.TradeDate < t.TradeDateMinValue() || t.TradeDate > t.TradeDateMaxValue() {
			return fmt.Errorf("Range check failed on t.TradeDate (%v < %v > %v)", t.TradeDateMinValue(), t.TradeDate, t.TradeDateMaxValue())
		}
	}
	if t.ValueDateInActingVersion(actingVersion) {
		if t.ValueDate < t.ValueDateMinValue() || t.ValueDate > t.ValueDateMaxValue() {
			return fmt.Errorf("Range check failed on t.ValueDate (%v < %v > %v)", t.ValueDateMinValue(), t.ValueDate, t.ValueDateMaxValue())
		}
	}
	if t.TradeTimeInActingVersion(actingVersion) {
		if t.TradeTime < t.TradeTimeMinValue() || t.TradeTime > t.TradeTimeMaxValue() {
			return fmt.Errorf("Range check failed on t.TradeTime (%v < %v > %v)", t.TradeTimeMinValue(), t.TradeTime, t.TradeTimeMaxValue())
		}
	}
	if t.OrderQtyInActingVersion(actingVersion) {
		if t.OrderQty < t.OrderQtyMinValue() || t.OrderQty > t.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.OrderQty (%v < %v > %v)", t.OrderQtyMinValue(), t.OrderQty, t.OrderQtyMaxValue())
		}
	}
	if t.LimitPriceInActingVersion(actingVersion) {
		if t.LimitPrice < t.LimitPriceMinValue() || t.LimitPrice > t.LimitPriceMaxValue() {
			return fmt.Errorf("Range check failed on t.LimitPrice (%v < %v > %v)", t.LimitPriceMinValue(), t.LimitPrice, t.LimitPriceMaxValue())
		}
	}
	if t.StopPriceInActingVersion(actingVersion) {
		if t.StopPrice < t.StopPriceMinValue() || t.StopPrice > t.StopPriceMaxValue() {
			return fmt.Errorf("Range check failed on t.StopPrice (%v < %v > %v)", t.StopPriceMinValue(), t.StopPrice, t.StopPriceMaxValue())
		}
	}
	if t.OpenQtyInActingVersion(actingVersion) {
		if t.OpenQty < t.OpenQtyMinValue() || t.OpenQty > t.OpenQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.OpenQty (%v < %v > %v)", t.OpenQtyMinValue(), t.OpenQty, t.OpenQtyMaxValue())
		}
	}
	if t.AvgPriceInActingVersion(actingVersion) {
		if t.AvgPrice < t.AvgPriceMinValue() || t.AvgPrice > t.AvgPriceMaxValue() {
			return fmt.Errorf("Range check failed on t.AvgPrice (%v < %v > %v)", t.AvgPriceMinValue(), t.AvgPrice, t.AvgPriceMaxValue())
		}
	}
	if t.OrderTypeInActingVersion(actingVersion) {
		if t.OrderType < t.OrderTypeMinValue() || t.OrderType > t.OrderTypeMaxValue() {
			return fmt.Errorf("Range check failed on t.OrderType (%v < %v > %v)", t.OrderTypeMinValue(), t.OrderType, t.OrderTypeMaxValue())
		}
	}
	if t.BrokerageInActingVersion(actingVersion) {
		if t.Brokerage < t.BrokerageMinValue() || t.Brokerage > t.BrokerageMaxValue() {
			return fmt.Errorf("Range check failed on t.Brokerage (%v < %v > %v)", t.BrokerageMinValue(), t.Brokerage, t.BrokerageMaxValue())
		}
	}
	if t.CreatedByInActingVersion(actingVersion) {
		if t.CreatedBy < t.CreatedByMinValue() || t.CreatedBy > t.CreatedByMaxValue() {
			return fmt.Errorf("Range check failed on t.CreatedBy (%v < %v > %v)", t.CreatedByMinValue(), t.CreatedBy, t.CreatedByMaxValue())
		}
	}
	if t.PlatformIdInActingVersion(actingVersion) {
		if t.PlatformId < t.PlatformIdMinValue() || t.PlatformId > t.PlatformIdMaxValue() {
			return fmt.Errorf("Range check failed on t.PlatformId (%v < %v > %v)", t.PlatformIdMinValue(), t.PlatformId, t.PlatformIdMaxValue())
		}
	}
	if t.ProxyIndexInActingVersion(actingVersion) {
		if t.ProxyIndex < t.ProxyIndexMinValue() || t.ProxyIndex > t.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on t.ProxyIndex (%v < %v > %v)", t.ProxyIndexMinValue(), t.ProxyIndex, t.ProxyIndexMaxValue())
		}
	}
	if t.CompleteInActingVersion(actingVersion) {
		if t.Complete < t.CompleteMinValue() || t.Complete > t.CompleteMaxValue() {
			return fmt.Errorf("Range check failed on t.Complete (%v < %v > %v)", t.CompleteMinValue(), t.Complete, t.CompleteMaxValue())
		}
	}
	if !utf8.Valid(t.RequestId[:]) {
		return errors.New("t.RequestId failed UTF-8 validation")
	}
	if !utf8.Valid(t.TradeId[:]) {
		return errors.New("t.TradeId failed UTF-8 validation")
	}
	if !utf8.Valid(t.OrderCurr[:]) {
		return errors.New("t.OrderCurr failed UTF-8 validation")
	}
	return nil
}

func TradeConfirmInit(t *TradeConfirm) {
	return
}

func (*TradeConfirm) SbeBlockLength() (blockLength uint16) {
	return 150
}

func (*TradeConfirm) SbeTemplateId() (templateId uint16) {
	return 71
}

func (*TradeConfirm) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*TradeConfirm) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*TradeConfirm) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TradeConfirm) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*TradeConfirm) OrderIdId() uint16 {
	return 72
}

func (*TradeConfirm) OrderIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderIdSinceVersion()
}

func (*TradeConfirm) OrderIdDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) OrderIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*TradeConfirm) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*TradeConfirm) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*TradeConfirm) LinkIdId() uint16 {
	return 73
}

func (*TradeConfirm) LinkIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) LinkIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.LinkIdSinceVersion()
}

func (*TradeConfirm) LinkIdDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) LinkIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) LinkIdMinValue() uint64 {
	return 0
}

func (*TradeConfirm) LinkIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*TradeConfirm) LinkIdNullValue() uint64 {
	return math.MaxUint64
}

func (*TradeConfirm) AcctIdId() uint16 {
	return 74
}

func (*TradeConfirm) AcctIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.AcctIdSinceVersion()
}

func (*TradeConfirm) AcctIdDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) AcctIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeConfirm) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeConfirm) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*TradeConfirm) CtrAcctIdId() uint16 {
	return 75
}

func (*TradeConfirm) CtrAcctIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) CtrAcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.CtrAcctIdSinceVersion()
}

func (*TradeConfirm) CtrAcctIdDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) CtrAcctIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) CtrAcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeConfirm) CtrAcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeConfirm) CtrAcctIdNullValue() int32 {
	return math.MinInt32
}

func (*TradeConfirm) InstIdId() uint16 {
	return 76
}

func (*TradeConfirm) InstIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.InstIdSinceVersion()
}

func (*TradeConfirm) InstIdDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) InstIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) InstIdMinValue() uint16 {
	return 0
}

func (*TradeConfirm) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*TradeConfirm) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*TradeConfirm) DealQtyId() uint16 {
	return 77
}

func (*TradeConfirm) DealQtySinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) DealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.DealQtySinceVersion()
}

func (*TradeConfirm) DealQtyDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) DealQtyMetaAttribute(meta int) string {
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

func (*TradeConfirm) DealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) DealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) DealQtyNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) CtrQtyId() uint16 {
	return 78
}

func (*TradeConfirm) CtrQtySinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) CtrQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.CtrQtySinceVersion()
}

func (*TradeConfirm) CtrQtyDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) CtrQtyMetaAttribute(meta int) string {
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

func (*TradeConfirm) CtrQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) CtrQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) CtrQtyNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) PbDealQtyId() uint16 {
	return 79
}

func (*TradeConfirm) PbDealQtySinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) PbDealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PbDealQtySinceVersion()
}

func (*TradeConfirm) PbDealQtyDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) PbDealQtyMetaAttribute(meta int) string {
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

func (*TradeConfirm) PbDealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) PbDealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) PbDealQtyNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) PbCtrQtyId() uint16 {
	return 80
}

func (*TradeConfirm) PbCtrQtySinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) PbCtrQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PbCtrQtySinceVersion()
}

func (*TradeConfirm) PbCtrQtyDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) PbCtrQtyMetaAttribute(meta int) string {
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

func (*TradeConfirm) PbCtrQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) PbCtrQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) PbCtrQtyNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) PriceId() uint16 {
	return 81
}

func (*TradeConfirm) PriceSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PriceSinceVersion()
}

func (*TradeConfirm) PriceDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) PriceMetaAttribute(meta int) string {
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

func (*TradeConfirm) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) PriceNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) PbPriceId() uint16 {
	return 82
}

func (*TradeConfirm) PbPriceSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) PbPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PbPriceSinceVersion()
}

func (*TradeConfirm) PbPriceDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) PbPriceMetaAttribute(meta int) string {
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

func (*TradeConfirm) PbPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) PbPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) PbPriceNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) TradeDateId() uint16 {
	return 83
}

func (*TradeConfirm) TradeDateSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) TradeDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TradeDateSinceVersion()
}

func (*TradeConfirm) TradeDateDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) TradeDateMetaAttribute(meta int) string {
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

func (*TradeConfirm) TradeDateMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeConfirm) TradeDateMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeConfirm) TradeDateNullValue() int32 {
	return math.MinInt32
}

func (*TradeConfirm) ValueDateId() uint16 {
	return 84
}

func (*TradeConfirm) ValueDateSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) ValueDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ValueDateSinceVersion()
}

func (*TradeConfirm) ValueDateDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) ValueDateMetaAttribute(meta int) string {
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

func (*TradeConfirm) ValueDateMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeConfirm) ValueDateMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeConfirm) ValueDateNullValue() int32 {
	return math.MinInt32
}

func (*TradeConfirm) TradeTimeId() uint16 {
	return 85
}

func (*TradeConfirm) TradeTimeSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) TradeTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TradeTimeSinceVersion()
}

func (*TradeConfirm) TradeTimeDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) TradeTimeMetaAttribute(meta int) string {
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

func (*TradeConfirm) TradeTimeMinValue() uint64 {
	return 0
}

func (*TradeConfirm) TradeTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*TradeConfirm) TradeTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*TradeConfirm) OrderQtyId() uint16 {
	return 86
}

func (*TradeConfirm) OrderQtySinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderQtySinceVersion()
}

func (*TradeConfirm) OrderQtyDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) OrderQtyMetaAttribute(meta int) string {
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

func (*TradeConfirm) OrderQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) OrderQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) OrderQtyNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) LimitPriceId() uint16 {
	return 87
}

func (*TradeConfirm) LimitPriceSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) LimitPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.LimitPriceSinceVersion()
}

func (*TradeConfirm) LimitPriceDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) LimitPriceMetaAttribute(meta int) string {
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

func (*TradeConfirm) LimitPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) LimitPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) LimitPriceNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) StopPriceId() uint16 {
	return 88
}

func (*TradeConfirm) StopPriceSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.StopPriceSinceVersion()
}

func (*TradeConfirm) StopPriceDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) StopPriceMetaAttribute(meta int) string {
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

func (*TradeConfirm) StopPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) StopPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) StopPriceNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) OpenQtyId() uint16 {
	return 89
}

func (*TradeConfirm) OpenQtySinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) OpenQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OpenQtySinceVersion()
}

func (*TradeConfirm) OpenQtyDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) OpenQtyMetaAttribute(meta int) string {
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

func (*TradeConfirm) OpenQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) OpenQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) OpenQtyNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) AvgPriceId() uint16 {
	return 90
}

func (*TradeConfirm) AvgPriceSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) AvgPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.AvgPriceSinceVersion()
}

func (*TradeConfirm) AvgPriceDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) AvgPriceMetaAttribute(meta int) string {
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

func (*TradeConfirm) AvgPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) AvgPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) AvgPriceNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) OrderTypeId() uint16 {
	return 91
}

func (*TradeConfirm) OrderTypeSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderTypeSinceVersion()
}

func (*TradeConfirm) OrderTypeDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) OrderTypeMetaAttribute(meta int) string {
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

func (*TradeConfirm) OrderTypeMinValue() uint8 {
	return 0
}

func (*TradeConfirm) OrderTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*TradeConfirm) OrderTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*TradeConfirm) BrokerageId() uint16 {
	return 92
}

func (*TradeConfirm) BrokerageSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) BrokerageInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.BrokerageSinceVersion()
}

func (*TradeConfirm) BrokerageDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) BrokerageMetaAttribute(meta int) string {
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

func (*TradeConfirm) BrokerageMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeConfirm) BrokerageMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeConfirm) BrokerageNullValue() float64 {
	return math.NaN()
}

func (*TradeConfirm) CreatedById() uint16 {
	return 93
}

func (*TradeConfirm) CreatedBySinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) CreatedByInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.CreatedBySinceVersion()
}

func (*TradeConfirm) CreatedByDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) CreatedByMetaAttribute(meta int) string {
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

func (*TradeConfirm) CreatedByMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeConfirm) CreatedByMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeConfirm) CreatedByNullValue() int32 {
	return math.MinInt32
}

func (*TradeConfirm) PlatformIdId() uint16 {
	return 94
}

func (*TradeConfirm) PlatformIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) PlatformIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.PlatformIdSinceVersion()
}

func (*TradeConfirm) PlatformIdDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) PlatformIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) PlatformIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeConfirm) PlatformIdMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeConfirm) PlatformIdNullValue() int32 {
	return math.MinInt32
}

func (*TradeConfirm) ProxyIndexId() uint16 {
	return 95
}

func (*TradeConfirm) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ProxyIndexSinceVersion()
}

func (*TradeConfirm) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) ProxyIndexMetaAttribute(meta int) string {
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

func (*TradeConfirm) ProxyIndexMinValue() uint16 {
	return 0
}

func (*TradeConfirm) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*TradeConfirm) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*TradeConfirm) CompleteId() uint16 {
	return 96
}

func (*TradeConfirm) CompleteSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) CompleteInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.CompleteSinceVersion()
}

func (*TradeConfirm) CompleteDeprecated() uint16 {
	return 0
}

func (*TradeConfirm) CompleteMetaAttribute(meta int) string {
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

func (*TradeConfirm) CompleteMinValue() uint8 {
	return 0
}

func (*TradeConfirm) CompleteMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*TradeConfirm) CompleteNullValue() uint8 {
	return math.MaxUint8
}

func (*TradeConfirm) RequestIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) RequestIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.RequestIdSinceVersion()
}

func (*TradeConfirm) RequestIdDeprecated() uint16 {
	return 0
}

func (TradeConfirm) RequestIdCharacterEncoding() string {
	return "UTF-8"
}

func (TradeConfirm) RequestIdHeaderLength() uint64 {
	return 2
}

func (*TradeConfirm) TradeIdMetaAttribute(meta int) string {
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

func (*TradeConfirm) TradeIdSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) TradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TradeIdSinceVersion()
}

func (*TradeConfirm) TradeIdDeprecated() uint16 {
	return 0
}

func (TradeConfirm) TradeIdCharacterEncoding() string {
	return "UTF-8"
}

func (TradeConfirm) TradeIdHeaderLength() uint64 {
	return 2
}

func (*TradeConfirm) OrderCurrMetaAttribute(meta int) string {
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

func (*TradeConfirm) OrderCurrSinceVersion() uint16 {
	return 0
}

func (t *TradeConfirm) OrderCurrInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.OrderCurrSinceVersion()
}

func (*TradeConfirm) OrderCurrDeprecated() uint16 {
	return 0
}

func (TradeConfirm) OrderCurrCharacterEncoding() string {
	return "UTF-8"
}

func (TradeConfirm) OrderCurrHeaderLength() uint64 {
	return 2
}
