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

type OrderResponse struct {
	OrderId     int64
	BrokerId    int32
	AcctId      int32
	Price       float64
	StopPrice   float64
	AvgPrice    float64
	OpenQty     float64
	OrderQty    float64
	OrderType   uint8
	OrderStatus uint8
	TimeInForce uint8
	Time        uint64
	CreatedBy   int32
	InstId      uint16
	RequestId   []uint8
	OrderCurr   []uint8
	Message     []uint8
	OrigOrderId []uint8
}

func (o *OrderResponse) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, o.OrderId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.BrokerId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.AcctId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.AvgPrice); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.OpenQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.OrderType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.OrderStatus); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.TimeInForce); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.Time); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.CreatedBy); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.InstId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.RequestId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.RequestId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.OrderCurr))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrderCurr); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.Message))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.Message); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.OrigOrderId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrigOrderId); err != nil {
		return err
	}
	return nil
}

func (o *OrderResponse) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.OrderIdInActingVersion(actingVersion) {
		o.OrderId = o.OrderIdNullValue()
	} else {
		if err := _m.ReadInt64(_r, &o.OrderId); err != nil {
			return err
		}
	}
	if !o.BrokerIdInActingVersion(actingVersion) {
		o.BrokerId = o.BrokerIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.BrokerId); err != nil {
			return err
		}
	}
	if !o.AcctIdInActingVersion(actingVersion) {
		o.AcctId = o.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.AcctId); err != nil {
			return err
		}
	}
	if !o.PriceInActingVersion(actingVersion) {
		o.Price = o.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.Price); err != nil {
			return err
		}
	}
	if !o.StopPriceInActingVersion(actingVersion) {
		o.StopPrice = o.StopPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.StopPrice); err != nil {
			return err
		}
	}
	if !o.AvgPriceInActingVersion(actingVersion) {
		o.AvgPrice = o.AvgPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.AvgPrice); err != nil {
			return err
		}
	}
	if !o.OpenQtyInActingVersion(actingVersion) {
		o.OpenQty = o.OpenQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.OpenQty); err != nil {
			return err
		}
	}
	if !o.OrderQtyInActingVersion(actingVersion) {
		o.OrderQty = o.OrderQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.OrderQty); err != nil {
			return err
		}
	}
	if !o.OrderTypeInActingVersion(actingVersion) {
		o.OrderType = o.OrderTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.OrderType); err != nil {
			return err
		}
	}
	if !o.OrderStatusInActingVersion(actingVersion) {
		o.OrderStatus = o.OrderStatusNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.OrderStatus); err != nil {
			return err
		}
	}
	if !o.TimeInForceInActingVersion(actingVersion) {
		o.TimeInForce = o.TimeInForceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.TimeInForce); err != nil {
			return err
		}
	}
	if !o.TimeInActingVersion(actingVersion) {
		o.Time = o.TimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.Time); err != nil {
			return err
		}
	}
	if !o.CreatedByInActingVersion(actingVersion) {
		o.CreatedBy = o.CreatedByNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.CreatedBy); err != nil {
			return err
		}
	}
	if !o.InstIdInActingVersion(actingVersion) {
		o.InstId = o.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.InstId); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}

	if o.RequestIdInActingVersion(actingVersion) {
		var RequestIdLength uint16
		if err := _m.ReadUint16(_r, &RequestIdLength); err != nil {
			return err
		}
		if cap(o.RequestId) < int(RequestIdLength) {
			o.RequestId = make([]uint8, RequestIdLength)
		}
		o.RequestId = o.RequestId[:RequestIdLength]
		if err := _m.ReadBytes(_r, o.RequestId); err != nil {
			return err
		}
	}

	if o.OrderCurrInActingVersion(actingVersion) {
		var OrderCurrLength uint16
		if err := _m.ReadUint16(_r, &OrderCurrLength); err != nil {
			return err
		}
		if cap(o.OrderCurr) < int(OrderCurrLength) {
			o.OrderCurr = make([]uint8, OrderCurrLength)
		}
		o.OrderCurr = o.OrderCurr[:OrderCurrLength]
		if err := _m.ReadBytes(_r, o.OrderCurr); err != nil {
			return err
		}
	}

	if o.MessageInActingVersion(actingVersion) {
		var MessageLength uint16
		if err := _m.ReadUint16(_r, &MessageLength); err != nil {
			return err
		}
		if cap(o.Message) < int(MessageLength) {
			o.Message = make([]uint8, MessageLength)
		}
		o.Message = o.Message[:MessageLength]
		if err := _m.ReadBytes(_r, o.Message); err != nil {
			return err
		}
	}

	if o.OrigOrderIdInActingVersion(actingVersion) {
		var OrigOrderIdLength uint16
		if err := _m.ReadUint16(_r, &OrigOrderIdLength); err != nil {
			return err
		}
		if cap(o.OrigOrderId) < int(OrigOrderIdLength) {
			o.OrigOrderId = make([]uint8, OrigOrderIdLength)
		}
		o.OrigOrderId = o.OrigOrderId[:OrigOrderIdLength]
		if err := _m.ReadBytes(_r, o.OrigOrderId); err != nil {
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

func (o *OrderResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderIdInActingVersion(actingVersion) {
		if o.OrderId < o.OrderIdMinValue() || o.OrderId > o.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderId (%v < %v > %v)", o.OrderIdMinValue(), o.OrderId, o.OrderIdMaxValue())
		}
	}
	if o.BrokerIdInActingVersion(actingVersion) {
		if o.BrokerId < o.BrokerIdMinValue() || o.BrokerId > o.BrokerIdMaxValue() {
			return fmt.Errorf("Range check failed on o.BrokerId (%v < %v > %v)", o.BrokerIdMinValue(), o.BrokerId, o.BrokerIdMaxValue())
		}
	}
	if o.AcctIdInActingVersion(actingVersion) {
		if o.AcctId < o.AcctIdMinValue() || o.AcctId > o.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on o.AcctId (%v < %v > %v)", o.AcctIdMinValue(), o.AcctId, o.AcctIdMaxValue())
		}
	}
	if o.PriceInActingVersion(actingVersion) {
		if o.Price < o.PriceMinValue() || o.Price > o.PriceMaxValue() {
			return fmt.Errorf("Range check failed on o.Price (%v < %v > %v)", o.PriceMinValue(), o.Price, o.PriceMaxValue())
		}
	}
	if o.StopPriceInActingVersion(actingVersion) {
		if o.StopPrice < o.StopPriceMinValue() || o.StopPrice > o.StopPriceMaxValue() {
			return fmt.Errorf("Range check failed on o.StopPrice (%v < %v > %v)", o.StopPriceMinValue(), o.StopPrice, o.StopPriceMaxValue())
		}
	}
	if o.AvgPriceInActingVersion(actingVersion) {
		if o.AvgPrice < o.AvgPriceMinValue() || o.AvgPrice > o.AvgPriceMaxValue() {
			return fmt.Errorf("Range check failed on o.AvgPrice (%v < %v > %v)", o.AvgPriceMinValue(), o.AvgPrice, o.AvgPriceMaxValue())
		}
	}
	if o.OpenQtyInActingVersion(actingVersion) {
		if o.OpenQty < o.OpenQtyMinValue() || o.OpenQty > o.OpenQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.OpenQty (%v < %v > %v)", o.OpenQtyMinValue(), o.OpenQty, o.OpenQtyMaxValue())
		}
	}
	if o.OrderQtyInActingVersion(actingVersion) {
		if o.OrderQty < o.OrderQtyMinValue() || o.OrderQty > o.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderQty (%v < %v > %v)", o.OrderQtyMinValue(), o.OrderQty, o.OrderQtyMaxValue())
		}
	}
	if o.OrderTypeInActingVersion(actingVersion) {
		if o.OrderType < o.OrderTypeMinValue() || o.OrderType > o.OrderTypeMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderType (%v < %v > %v)", o.OrderTypeMinValue(), o.OrderType, o.OrderTypeMaxValue())
		}
	}
	if o.OrderStatusInActingVersion(actingVersion) {
		if o.OrderStatus < o.OrderStatusMinValue() || o.OrderStatus > o.OrderStatusMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderStatus (%v < %v > %v)", o.OrderStatusMinValue(), o.OrderStatus, o.OrderStatusMaxValue())
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if o.TimeInForce < o.TimeInForceMinValue() || o.TimeInForce > o.TimeInForceMaxValue() {
			return fmt.Errorf("Range check failed on o.TimeInForce (%v < %v > %v)", o.TimeInForceMinValue(), o.TimeInForce, o.TimeInForceMaxValue())
		}
	}
	if o.TimeInActingVersion(actingVersion) {
		if o.Time < o.TimeMinValue() || o.Time > o.TimeMaxValue() {
			return fmt.Errorf("Range check failed on o.Time (%v < %v > %v)", o.TimeMinValue(), o.Time, o.TimeMaxValue())
		}
	}
	if o.CreatedByInActingVersion(actingVersion) {
		if o.CreatedBy < o.CreatedByMinValue() || o.CreatedBy > o.CreatedByMaxValue() {
			return fmt.Errorf("Range check failed on o.CreatedBy (%v < %v > %v)", o.CreatedByMinValue(), o.CreatedBy, o.CreatedByMaxValue())
		}
	}
	if o.InstIdInActingVersion(actingVersion) {
		if o.InstId < o.InstIdMinValue() || o.InstId > o.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on o.InstId (%v < %v > %v)", o.InstIdMinValue(), o.InstId, o.InstIdMaxValue())
		}
	}
	if !utf8.Valid(o.RequestId[:]) {
		return errors.New("o.RequestId failed UTF-8 validation")
	}
	if !utf8.Valid(o.OrderCurr[:]) {
		return errors.New("o.OrderCurr failed UTF-8 validation")
	}
	if !utf8.Valid(o.Message[:]) {
		return errors.New("o.Message failed UTF-8 validation")
	}
	if !utf8.Valid(o.OrigOrderId[:]) {
		return errors.New("o.OrigOrderId failed UTF-8 validation")
	}
	return nil
}

func OrderResponseInit(o *OrderResponse) {
	return
}

func (*OrderResponse) SbeBlockLength() (blockLength uint16) {
	return 73
}

func (*OrderResponse) SbeTemplateId() (templateId uint16) {
	return 100
}

func (*OrderResponse) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OrderResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OrderResponse) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderResponse) OrderIdId() uint16 {
	return 101
}

func (*OrderResponse) OrderIdSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIdSinceVersion()
}

func (*OrderResponse) OrderIdDeprecated() uint16 {
	return 0
}

func (*OrderResponse) OrderIdMetaAttribute(meta int) string {
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

func (*OrderResponse) OrderIdMinValue() int64 {
	return math.MinInt64 + 1
}

func (*OrderResponse) OrderIdMaxValue() int64 {
	return math.MaxInt64
}

func (*OrderResponse) OrderIdNullValue() int64 {
	return math.MinInt64
}

func (*OrderResponse) BrokerIdId() uint16 {
	return 102
}

func (*OrderResponse) BrokerIdSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) BrokerIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BrokerIdSinceVersion()
}

func (*OrderResponse) BrokerIdDeprecated() uint16 {
	return 0
}

func (*OrderResponse) BrokerIdMetaAttribute(meta int) string {
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

func (*OrderResponse) BrokerIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderResponse) BrokerIdMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderResponse) BrokerIdNullValue() int32 {
	return math.MinInt32
}

func (*OrderResponse) AcctIdId() uint16 {
	return 103
}

func (*OrderResponse) AcctIdSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AcctIdSinceVersion()
}

func (*OrderResponse) AcctIdDeprecated() uint16 {
	return 0
}

func (*OrderResponse) AcctIdMetaAttribute(meta int) string {
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

func (*OrderResponse) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderResponse) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderResponse) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*OrderResponse) PriceId() uint16 {
	return 104
}

func (*OrderResponse) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderResponse) PriceDeprecated() uint16 {
	return 0
}

func (*OrderResponse) PriceMetaAttribute(meta int) string {
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

func (*OrderResponse) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderResponse) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderResponse) PriceNullValue() float64 {
	return math.NaN()
}

func (*OrderResponse) StopPriceId() uint16 {
	return 105
}

func (*OrderResponse) StopPriceSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPriceSinceVersion()
}

func (*OrderResponse) StopPriceDeprecated() uint16 {
	return 0
}

func (*OrderResponse) StopPriceMetaAttribute(meta int) string {
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

func (*OrderResponse) StopPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderResponse) StopPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderResponse) StopPriceNullValue() float64 {
	return math.NaN()
}

func (*OrderResponse) AvgPriceId() uint16 {
	return 106
}

func (*OrderResponse) AvgPriceSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) AvgPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AvgPriceSinceVersion()
}

func (*OrderResponse) AvgPriceDeprecated() uint16 {
	return 0
}

func (*OrderResponse) AvgPriceMetaAttribute(meta int) string {
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

func (*OrderResponse) AvgPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderResponse) AvgPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderResponse) AvgPriceNullValue() float64 {
	return math.NaN()
}

func (*OrderResponse) OpenQtyId() uint16 {
	return 107
}

func (*OrderResponse) OpenQtySinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) OpenQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OpenQtySinceVersion()
}

func (*OrderResponse) OpenQtyDeprecated() uint16 {
	return 0
}

func (*OrderResponse) OpenQtyMetaAttribute(meta int) string {
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

func (*OrderResponse) OpenQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderResponse) OpenQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderResponse) OpenQtyNullValue() float64 {
	return math.NaN()
}

func (*OrderResponse) OrderQtyId() uint16 {
	return 108
}

func (*OrderResponse) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderResponse) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderResponse) OrderQtyMetaAttribute(meta int) string {
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

func (*OrderResponse) OrderQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderResponse) OrderQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderResponse) OrderQtyNullValue() float64 {
	return math.NaN()
}

func (*OrderResponse) OrderTypeId() uint16 {
	return 109
}

func (*OrderResponse) OrderTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderTypeSinceVersion()
}

func (*OrderResponse) OrderTypeDeprecated() uint16 {
	return 0
}

func (*OrderResponse) OrderTypeMetaAttribute(meta int) string {
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

func (*OrderResponse) OrderTypeMinValue() uint8 {
	return 0
}

func (*OrderResponse) OrderTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderResponse) OrderTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderResponse) OrderStatusId() uint16 {
	return 110
}

func (*OrderResponse) OrderStatusSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) OrderStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderStatusSinceVersion()
}

func (*OrderResponse) OrderStatusDeprecated() uint16 {
	return 0
}

func (*OrderResponse) OrderStatusMetaAttribute(meta int) string {
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

func (*OrderResponse) OrderStatusMinValue() uint8 {
	return 0
}

func (*OrderResponse) OrderStatusMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderResponse) OrderStatusNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderResponse) TimeInForceId() uint16 {
	return 111
}

func (*OrderResponse) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderResponse) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderResponse) TimeInForceMetaAttribute(meta int) string {
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

func (*OrderResponse) TimeInForceMinValue() uint8 {
	return 0
}

func (*OrderResponse) TimeInForceMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderResponse) TimeInForceNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderResponse) TimeId() uint16 {
	return 112
}

func (*OrderResponse) TimeSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) TimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeSinceVersion()
}

func (*OrderResponse) TimeDeprecated() uint16 {
	return 0
}

func (*OrderResponse) TimeMetaAttribute(meta int) string {
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

func (*OrderResponse) TimeMinValue() uint64 {
	return 0
}

func (*OrderResponse) TimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderResponse) TimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderResponse) CreatedById() uint16 {
	return 113
}

func (*OrderResponse) CreatedBySinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) CreatedByInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CreatedBySinceVersion()
}

func (*OrderResponse) CreatedByDeprecated() uint16 {
	return 0
}

func (*OrderResponse) CreatedByMetaAttribute(meta int) string {
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

func (*OrderResponse) CreatedByMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderResponse) CreatedByMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderResponse) CreatedByNullValue() int32 {
	return math.MinInt32
}

func (*OrderResponse) InstIdId() uint16 {
	return 115
}

func (*OrderResponse) InstIdSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstIdSinceVersion()
}

func (*OrderResponse) InstIdDeprecated() uint16 {
	return 0
}

func (*OrderResponse) InstIdMetaAttribute(meta int) string {
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

func (*OrderResponse) InstIdMinValue() uint16 {
	return 0
}

func (*OrderResponse) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderResponse) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderResponse) RequestIdMetaAttribute(meta int) string {
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

func (*OrderResponse) RequestIdSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestIdSinceVersion()
}

func (*OrderResponse) RequestIdDeprecated() uint16 {
	return 0
}

func (OrderResponse) RequestIdCharacterEncoding() string {
	return "UTF-8"
}

func (OrderResponse) RequestIdHeaderLength() uint64 {
	return 2
}

func (*OrderResponse) OrderCurrMetaAttribute(meta int) string {
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

func (*OrderResponse) OrderCurrSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) OrderCurrInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderCurrSinceVersion()
}

func (*OrderResponse) OrderCurrDeprecated() uint16 {
	return 0
}

func (OrderResponse) OrderCurrCharacterEncoding() string {
	return "UTF-8"
}

func (OrderResponse) OrderCurrHeaderLength() uint64 {
	return 2
}

func (*OrderResponse) MessageMetaAttribute(meta int) string {
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

func (*OrderResponse) MessageSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) MessageInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MessageSinceVersion()
}

func (*OrderResponse) MessageDeprecated() uint16 {
	return 0
}

func (OrderResponse) MessageCharacterEncoding() string {
	return "UTF-8"
}

func (OrderResponse) MessageHeaderLength() uint64 {
	return 2
}

func (*OrderResponse) OrigOrderIdMetaAttribute(meta int) string {
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

func (*OrderResponse) OrigOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OrderResponse) OrigOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigOrderIdSinceVersion()
}

func (*OrderResponse) OrigOrderIdDeprecated() uint16 {
	return 0
}

func (OrderResponse) OrigOrderIdCharacterEncoding() string {
	return "UTF-8"
}

func (OrderResponse) OrigOrderIdHeaderLength() uint64 {
	return 2
}
