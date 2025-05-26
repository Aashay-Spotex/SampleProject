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

type ActiveOrderSet struct {
	AcctId          int32
	ProxyIndex      uint16
	CreatedBy       int32
	ActiveOrderList []ActiveOrderSetActiveOrderList
	RequestId       []uint8
}
type ActiveOrderSetActiveOrderList struct {
	OrderQty    float64
	OpenQty     float64
	Price       float64
	StopPrice   float64
	OrderType   uint8
	TimeInForce uint8
	OrderStatus uint8
	Time        int64
	RequestId   []uint8
	OrderId     []uint8
	CurrPair    []uint8
	OrderCurr   []uint8
}

func (a *ActiveOrderSet) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := a.RangeCheck(a.SbeSchemaVersion(), a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, a.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, a.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, a.CreatedBy); err != nil {
		return err
	}
	var ActiveOrderListBlockLength uint16 = 43
	if err := _m.WriteUint16(_w, ActiveOrderListBlockLength); err != nil {
		return err
	}
	var ActiveOrderListNumInGroup uint16 = uint16(len(a.ActiveOrderList))
	if err := _m.WriteUint16(_w, ActiveOrderListNumInGroup); err != nil {
		return err
	}
	for i := range a.ActiveOrderList {
		if err := a.ActiveOrderList[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	if err := _m.WriteUint16(_w, uint16(len(a.RequestId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, a.RequestId); err != nil {
		return err
	}
	return nil
}

func (a *ActiveOrderSet) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !a.AcctIdInActingVersion(actingVersion) {
		a.AcctId = a.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &a.AcctId); err != nil {
			return err
		}
	}
	if !a.ProxyIndexInActingVersion(actingVersion) {
		a.ProxyIndex = a.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &a.ProxyIndex); err != nil {
			return err
		}
	}
	if !a.CreatedByInActingVersion(actingVersion) {
		a.CreatedBy = a.CreatedByNullValue()
	} else {
		if err := _m.ReadInt32(_r, &a.CreatedBy); err != nil {
			return err
		}
	}
	if actingVersion > a.SbeSchemaVersion() && blockLength > a.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-a.SbeBlockLength()))
	}

	if a.ActiveOrderListInActingVersion(actingVersion) {
		var ActiveOrderListBlockLength uint16
		if err := _m.ReadUint16(_r, &ActiveOrderListBlockLength); err != nil {
			return err
		}
		var ActiveOrderListNumInGroup uint16
		if err := _m.ReadUint16(_r, &ActiveOrderListNumInGroup); err != nil {
			return err
		}
		if cap(a.ActiveOrderList) < int(ActiveOrderListNumInGroup) {
			a.ActiveOrderList = make([]ActiveOrderSetActiveOrderList, ActiveOrderListNumInGroup)
		}
		a.ActiveOrderList = a.ActiveOrderList[:ActiveOrderListNumInGroup]
		for i := range a.ActiveOrderList {
			if err := a.ActiveOrderList[i].Decode(_m, _r, actingVersion, uint(ActiveOrderListBlockLength)); err != nil {
				return err
			}
		}
	}

	if a.RequestIdInActingVersion(actingVersion) {
		var RequestIdLength uint16
		if err := _m.ReadUint16(_r, &RequestIdLength); err != nil {
			return err
		}
		if cap(a.RequestId) < int(RequestIdLength) {
			a.RequestId = make([]uint8, RequestIdLength)
		}
		a.RequestId = a.RequestId[:RequestIdLength]
		if err := _m.ReadBytes(_r, a.RequestId); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := a.RangeCheck(actingVersion, a.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (a *ActiveOrderSet) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if a.AcctIdInActingVersion(actingVersion) {
		if a.AcctId < a.AcctIdMinValue() || a.AcctId > a.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on a.AcctId (%v < %v > %v)", a.AcctIdMinValue(), a.AcctId, a.AcctIdMaxValue())
		}
	}
	if a.ProxyIndexInActingVersion(actingVersion) {
		if a.ProxyIndex < a.ProxyIndexMinValue() || a.ProxyIndex > a.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on a.ProxyIndex (%v < %v > %v)", a.ProxyIndexMinValue(), a.ProxyIndex, a.ProxyIndexMaxValue())
		}
	}
	if a.CreatedByInActingVersion(actingVersion) {
		if a.CreatedBy < a.CreatedByMinValue() || a.CreatedBy > a.CreatedByMaxValue() {
			return fmt.Errorf("Range check failed on a.CreatedBy (%v < %v > %v)", a.CreatedByMinValue(), a.CreatedBy, a.CreatedByMaxValue())
		}
	}
	for i := range a.ActiveOrderList {
		if err := a.ActiveOrderList[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	if !utf8.Valid(a.RequestId[:]) {
		return errors.New("a.RequestId failed UTF-8 validation")
	}
	return nil
}

func ActiveOrderSetInit(a *ActiveOrderSet) {
	return
}

func (a *ActiveOrderSetActiveOrderList) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteFloat64(_w, a.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, a.OpenQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, a.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, a.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, a.OrderType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, a.TimeInForce); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, a.OrderStatus); err != nil {
		return err
	}
	if err := _m.WriteInt64(_w, a.Time); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(a.RequestId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, a.RequestId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(a.OrderId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, a.OrderId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(a.CurrPair))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, a.CurrPair); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(a.OrderCurr))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, a.OrderCurr); err != nil {
		return err
	}
	return nil
}

func (a *ActiveOrderSetActiveOrderList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !a.OrderQtyInActingVersion(actingVersion) {
		a.OrderQty = a.OrderQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &a.OrderQty); err != nil {
			return err
		}
	}
	if !a.OpenQtyInActingVersion(actingVersion) {
		a.OpenQty = a.OpenQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &a.OpenQty); err != nil {
			return err
		}
	}
	if !a.PriceInActingVersion(actingVersion) {
		a.Price = a.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &a.Price); err != nil {
			return err
		}
	}
	if !a.StopPriceInActingVersion(actingVersion) {
		a.StopPrice = a.StopPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &a.StopPrice); err != nil {
			return err
		}
	}
	if !a.OrderTypeInActingVersion(actingVersion) {
		a.OrderType = a.OrderTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &a.OrderType); err != nil {
			return err
		}
	}
	if !a.TimeInForceInActingVersion(actingVersion) {
		a.TimeInForce = a.TimeInForceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &a.TimeInForce); err != nil {
			return err
		}
	}
	if !a.OrderStatusInActingVersion(actingVersion) {
		a.OrderStatus = a.OrderStatusNullValue()
	} else {
		if err := _m.ReadUint8(_r, &a.OrderStatus); err != nil {
			return err
		}
	}
	if !a.TimeInActingVersion(actingVersion) {
		a.Time = a.TimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &a.Time); err != nil {
			return err
		}
	}
	if actingVersion > a.SbeSchemaVersion() && blockLength > a.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-a.SbeBlockLength()))
	}

	if a.RequestIdInActingVersion(actingVersion) {
		var RequestIdLength uint16
		if err := _m.ReadUint16(_r, &RequestIdLength); err != nil {
			return err
		}
		if cap(a.RequestId) < int(RequestIdLength) {
			a.RequestId = make([]uint8, RequestIdLength)
		}
		a.RequestId = a.RequestId[:RequestIdLength]
		if err := _m.ReadBytes(_r, a.RequestId); err != nil {
			return err
		}
	}

	if a.OrderIdInActingVersion(actingVersion) {
		var OrderIdLength uint16
		if err := _m.ReadUint16(_r, &OrderIdLength); err != nil {
			return err
		}
		if cap(a.OrderId) < int(OrderIdLength) {
			a.OrderId = make([]uint8, OrderIdLength)
		}
		a.OrderId = a.OrderId[:OrderIdLength]
		if err := _m.ReadBytes(_r, a.OrderId); err != nil {
			return err
		}
	}

	if a.CurrPairInActingVersion(actingVersion) {
		var CurrPairLength uint16
		if err := _m.ReadUint16(_r, &CurrPairLength); err != nil {
			return err
		}
		if cap(a.CurrPair) < int(CurrPairLength) {
			a.CurrPair = make([]uint8, CurrPairLength)
		}
		a.CurrPair = a.CurrPair[:CurrPairLength]
		if err := _m.ReadBytes(_r, a.CurrPair); err != nil {
			return err
		}
	}

	if a.OrderCurrInActingVersion(actingVersion) {
		var OrderCurrLength uint16
		if err := _m.ReadUint16(_r, &OrderCurrLength); err != nil {
			return err
		}
		if cap(a.OrderCurr) < int(OrderCurrLength) {
			a.OrderCurr = make([]uint8, OrderCurrLength)
		}
		a.OrderCurr = a.OrderCurr[:OrderCurrLength]
		if err := _m.ReadBytes(_r, a.OrderCurr); err != nil {
			return err
		}
	}
	return nil
}

func (a *ActiveOrderSetActiveOrderList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if a.OrderQtyInActingVersion(actingVersion) {
		if a.OrderQty < a.OrderQtyMinValue() || a.OrderQty > a.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on a.OrderQty (%v < %v > %v)", a.OrderQtyMinValue(), a.OrderQty, a.OrderQtyMaxValue())
		}
	}
	if a.OpenQtyInActingVersion(actingVersion) {
		if a.OpenQty < a.OpenQtyMinValue() || a.OpenQty > a.OpenQtyMaxValue() {
			return fmt.Errorf("Range check failed on a.OpenQty (%v < %v > %v)", a.OpenQtyMinValue(), a.OpenQty, a.OpenQtyMaxValue())
		}
	}
	if a.PriceInActingVersion(actingVersion) {
		if a.Price < a.PriceMinValue() || a.Price > a.PriceMaxValue() {
			return fmt.Errorf("Range check failed on a.Price (%v < %v > %v)", a.PriceMinValue(), a.Price, a.PriceMaxValue())
		}
	}
	if a.StopPriceInActingVersion(actingVersion) {
		if a.StopPrice < a.StopPriceMinValue() || a.StopPrice > a.StopPriceMaxValue() {
			return fmt.Errorf("Range check failed on a.StopPrice (%v < %v > %v)", a.StopPriceMinValue(), a.StopPrice, a.StopPriceMaxValue())
		}
	}
	if a.OrderTypeInActingVersion(actingVersion) {
		if a.OrderType < a.OrderTypeMinValue() || a.OrderType > a.OrderTypeMaxValue() {
			return fmt.Errorf("Range check failed on a.OrderType (%v < %v > %v)", a.OrderTypeMinValue(), a.OrderType, a.OrderTypeMaxValue())
		}
	}
	if a.TimeInForceInActingVersion(actingVersion) {
		if a.TimeInForce < a.TimeInForceMinValue() || a.TimeInForce > a.TimeInForceMaxValue() {
			return fmt.Errorf("Range check failed on a.TimeInForce (%v < %v > %v)", a.TimeInForceMinValue(), a.TimeInForce, a.TimeInForceMaxValue())
		}
	}
	if a.OrderStatusInActingVersion(actingVersion) {
		if a.OrderStatus < a.OrderStatusMinValue() || a.OrderStatus > a.OrderStatusMaxValue() {
			return fmt.Errorf("Range check failed on a.OrderStatus (%v < %v > %v)", a.OrderStatusMinValue(), a.OrderStatus, a.OrderStatusMaxValue())
		}
	}
	if a.TimeInActingVersion(actingVersion) {
		if a.Time < a.TimeMinValue() || a.Time > a.TimeMaxValue() {
			return fmt.Errorf("Range check failed on a.Time (%v < %v > %v)", a.TimeMinValue(), a.Time, a.TimeMaxValue())
		}
	}
	if !utf8.Valid(a.RequestId[:]) {
		return errors.New("a.RequestId failed UTF-8 validation")
	}
	if !utf8.Valid(a.OrderId[:]) {
		return errors.New("a.OrderId failed UTF-8 validation")
	}
	if !utf8.Valid(a.CurrPair[:]) {
		return errors.New("a.CurrPair failed UTF-8 validation")
	}
	if !utf8.Valid(a.OrderCurr[:]) {
		return errors.New("a.OrderCurr failed UTF-8 validation")
	}
	return nil
}

func ActiveOrderSetActiveOrderListInit(a *ActiveOrderSetActiveOrderList) {
	return
}

func (*ActiveOrderSet) SbeBlockLength() (blockLength uint16) {
	return 10
}

func (*ActiveOrderSet) SbeTemplateId() (templateId uint16) {
	return 145
}

func (*ActiveOrderSet) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*ActiveOrderSet) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*ActiveOrderSet) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ActiveOrderSet) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*ActiveOrderSet) AcctIdId() uint16 {
	return 146
}

func (*ActiveOrderSet) AcctIdSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSet) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.AcctIdSinceVersion()
}

func (*ActiveOrderSet) AcctIdDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSet) AcctIdMetaAttribute(meta int) string {
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

func (*ActiveOrderSet) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ActiveOrderSet) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*ActiveOrderSet) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*ActiveOrderSet) ProxyIndexId() uint16 {
	return 147
}

func (*ActiveOrderSet) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSet) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.ProxyIndexSinceVersion()
}

func (*ActiveOrderSet) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSet) ProxyIndexMetaAttribute(meta int) string {
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

func (*ActiveOrderSet) ProxyIndexMinValue() uint16 {
	return 0
}

func (*ActiveOrderSet) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ActiveOrderSet) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*ActiveOrderSet) CreatedById() uint16 {
	return 148
}

func (*ActiveOrderSet) CreatedBySinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSet) CreatedByInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.CreatedBySinceVersion()
}

func (*ActiveOrderSet) CreatedByDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSet) CreatedByMetaAttribute(meta int) string {
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

func (*ActiveOrderSet) CreatedByMinValue() int32 {
	return math.MinInt32 + 1
}

func (*ActiveOrderSet) CreatedByMaxValue() int32 {
	return math.MaxInt32
}

func (*ActiveOrderSet) CreatedByNullValue() int32 {
	return math.MinInt32
}

func (*ActiveOrderSetActiveOrderList) OrderQtyId() uint16 {
	return 150
}

func (*ActiveOrderSetActiveOrderList) OrderQtySinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.OrderQtySinceVersion()
}

func (*ActiveOrderSetActiveOrderList) OrderQtyDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) OrderQtyMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) OrderQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) OrderQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) OrderQtyNullValue() float64 {
	return math.NaN()
}

func (*ActiveOrderSetActiveOrderList) OpenQtyId() uint16 {
	return 151
}

func (*ActiveOrderSetActiveOrderList) OpenQtySinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) OpenQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.OpenQtySinceVersion()
}

func (*ActiveOrderSetActiveOrderList) OpenQtyDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) OpenQtyMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) OpenQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) OpenQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) OpenQtyNullValue() float64 {
	return math.NaN()
}

func (*ActiveOrderSetActiveOrderList) PriceId() uint16 {
	return 152
}

func (*ActiveOrderSetActiveOrderList) PriceSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.PriceSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) PriceDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) PriceMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) PriceNullValue() float64 {
	return math.NaN()
}

func (*ActiveOrderSetActiveOrderList) StopPriceId() uint16 {
	return 153
}

func (*ActiveOrderSetActiveOrderList) StopPriceSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.StopPriceSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) StopPriceDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) StopPriceMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) StopPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) StopPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*ActiveOrderSetActiveOrderList) StopPriceNullValue() float64 {
	return math.NaN()
}

func (*ActiveOrderSetActiveOrderList) OrderTypeId() uint16 {
	return 154
}

func (*ActiveOrderSetActiveOrderList) OrderTypeSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.OrderTypeSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) OrderTypeDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) OrderTypeMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) OrderTypeMinValue() uint8 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) OrderTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ActiveOrderSetActiveOrderList) OrderTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*ActiveOrderSetActiveOrderList) TimeInForceId() uint16 {
	return 155
}

func (*ActiveOrderSetActiveOrderList) TimeInForceSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.TimeInForceSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) TimeInForceDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) TimeInForceMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) TimeInForceMinValue() uint8 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) TimeInForceMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ActiveOrderSetActiveOrderList) TimeInForceNullValue() uint8 {
	return math.MaxUint8
}

func (*ActiveOrderSetActiveOrderList) OrderStatusId() uint16 {
	return 156
}

func (*ActiveOrderSetActiveOrderList) OrderStatusSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) OrderStatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.OrderStatusSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) OrderStatusDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) OrderStatusMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) OrderStatusMinValue() uint8 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) OrderStatusMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*ActiveOrderSetActiveOrderList) OrderStatusNullValue() uint8 {
	return math.MaxUint8
}

func (*ActiveOrderSetActiveOrderList) TimeId() uint16 {
	return 157
}

func (*ActiveOrderSetActiveOrderList) TimeSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) TimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.TimeSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) TimeDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) TimeMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) TimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*ActiveOrderSetActiveOrderList) TimeMaxValue() int64 {
	return math.MaxInt64
}

func (*ActiveOrderSetActiveOrderList) TimeNullValue() int64 {
	return math.MinInt64
}

func (*ActiveOrderSetActiveOrderList) RequestIdMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) RequestIdSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.RequestIdSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) RequestIdDeprecated() uint16 {
	return 0
}

func (ActiveOrderSetActiveOrderList) RequestIdCharacterEncoding() string {
	return "UTF-8"
}

func (ActiveOrderSetActiveOrderList) RequestIdHeaderLength() uint64 {
	return 2
}

func (*ActiveOrderSetActiveOrderList) OrderIdMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) OrderIdSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.OrderIdSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) OrderIdDeprecated() uint16 {
	return 0
}

func (ActiveOrderSetActiveOrderList) OrderIdCharacterEncoding() string {
	return "UTF-8"
}

func (ActiveOrderSetActiveOrderList) OrderIdHeaderLength() uint64 {
	return 2
}

func (*ActiveOrderSetActiveOrderList) CurrPairMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) CurrPairSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) CurrPairInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.CurrPairSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) CurrPairDeprecated() uint16 {
	return 0
}

func (ActiveOrderSetActiveOrderList) CurrPairCharacterEncoding() string {
	return "UTF-8"
}

func (ActiveOrderSetActiveOrderList) CurrPairHeaderLength() uint64 {
	return 2
}

func (*ActiveOrderSetActiveOrderList) OrderCurrMetaAttribute(meta int) string {
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

func (*ActiveOrderSetActiveOrderList) OrderCurrSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSetActiveOrderList) OrderCurrInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.OrderCurrSinceVersion()
}

func (*ActiveOrderSetActiveOrderList) OrderCurrDeprecated() uint16 {
	return 0
}

func (ActiveOrderSetActiveOrderList) OrderCurrCharacterEncoding() string {
	return "UTF-8"
}

func (ActiveOrderSetActiveOrderList) OrderCurrHeaderLength() uint64 {
	return 2
}

func (*ActiveOrderSet) ActiveOrderListId() uint16 {
	return 149
}

func (*ActiveOrderSet) ActiveOrderListSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSet) ActiveOrderListInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.ActiveOrderListSinceVersion()
}

func (*ActiveOrderSet) ActiveOrderListDeprecated() uint16 {
	return 0
}

func (*ActiveOrderSetActiveOrderList) SbeBlockLength() (blockLength uint) {
	return 43
}

func (*ActiveOrderSetActiveOrderList) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*ActiveOrderSet) RequestIdMetaAttribute(meta int) string {
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

func (*ActiveOrderSet) RequestIdSinceVersion() uint16 {
	return 0
}

func (a *ActiveOrderSet) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= a.RequestIdSinceVersion()
}

func (*ActiveOrderSet) RequestIdDeprecated() uint16 {
	return 0
}

func (ActiveOrderSet) RequestIdCharacterEncoding() string {
	return "UTF-8"
}

func (ActiveOrderSet) RequestIdHeaderLength() uint64 {
	return 2
}
