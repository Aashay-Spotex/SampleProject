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

type OrderBuffer struct {
	TradeSource   uint8
	AcctId        int32
	ProxyIndex    uint16
	CreateTime    uint64
	InstId        uint16
	LinkId        uint64
	CreatedBy     int32
	OrderQty      float64
	MinQty        float64
	Price         float64
	StopPrice     float64
	OrderType     uint8
	TimeInForce   uint8
	CancelReplace uint8
	ExpireTime    uint64
	BaseOrder     uint8
	RequestId     []uint8
	OrigOrderId   []uint8
}

func (o *OrderBuffer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, o.TradeSource); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.CreateTime); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.InstId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.LinkId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.CreatedBy); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.MinQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, o.StopPrice); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.OrderType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.TimeInForce); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.CancelReplace); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, o.ExpireTime); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.BaseOrder); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.RequestId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.RequestId); err != nil {
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

func (o *OrderBuffer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.TradeSourceInActingVersion(actingVersion) {
		o.TradeSource = o.TradeSourceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.TradeSource); err != nil {
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
	if !o.ProxyIndexInActingVersion(actingVersion) {
		o.ProxyIndex = o.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.ProxyIndex); err != nil {
			return err
		}
	}
	if !o.CreateTimeInActingVersion(actingVersion) {
		o.CreateTime = o.CreateTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.CreateTime); err != nil {
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
	if !o.LinkIdInActingVersion(actingVersion) {
		o.LinkId = o.LinkIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.LinkId); err != nil {
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
	if !o.OrderQtyInActingVersion(actingVersion) {
		o.OrderQty = o.OrderQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.OrderQty); err != nil {
			return err
		}
	}
	if !o.MinQtyInActingVersion(actingVersion) {
		o.MinQty = o.MinQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &o.MinQty); err != nil {
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
	if !o.OrderTypeInActingVersion(actingVersion) {
		o.OrderType = o.OrderTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.OrderType); err != nil {
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
	if !o.CancelReplaceInActingVersion(actingVersion) {
		o.CancelReplace = o.CancelReplaceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.CancelReplace); err != nil {
			return err
		}
	}
	if !o.ExpireTimeInActingVersion(actingVersion) {
		o.ExpireTime = o.ExpireTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &o.ExpireTime); err != nil {
			return err
		}
	}
	if !o.BaseOrderInActingVersion(actingVersion) {
		o.BaseOrder = o.BaseOrderNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.BaseOrder); err != nil {
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

func (o *OrderBuffer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.TradeSourceInActingVersion(actingVersion) {
		if o.TradeSource < o.TradeSourceMinValue() || o.TradeSource > o.TradeSourceMaxValue() {
			return fmt.Errorf("Range check failed on o.TradeSource (%v < %v > %v)", o.TradeSourceMinValue(), o.TradeSource, o.TradeSourceMaxValue())
		}
	}
	if o.AcctIdInActingVersion(actingVersion) {
		if o.AcctId < o.AcctIdMinValue() || o.AcctId > o.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on o.AcctId (%v < %v > %v)", o.AcctIdMinValue(), o.AcctId, o.AcctIdMaxValue())
		}
	}
	if o.ProxyIndexInActingVersion(actingVersion) {
		if o.ProxyIndex < o.ProxyIndexMinValue() || o.ProxyIndex > o.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on o.ProxyIndex (%v < %v > %v)", o.ProxyIndexMinValue(), o.ProxyIndex, o.ProxyIndexMaxValue())
		}
	}
	if o.CreateTimeInActingVersion(actingVersion) {
		if o.CreateTime < o.CreateTimeMinValue() || o.CreateTime > o.CreateTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.CreateTime (%v < %v > %v)", o.CreateTimeMinValue(), o.CreateTime, o.CreateTimeMaxValue())
		}
	}
	if o.InstIdInActingVersion(actingVersion) {
		if o.InstId < o.InstIdMinValue() || o.InstId > o.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on o.InstId (%v < %v > %v)", o.InstIdMinValue(), o.InstId, o.InstIdMaxValue())
		}
	}
	if o.LinkIdInActingVersion(actingVersion) {
		if o.LinkId < o.LinkIdMinValue() || o.LinkId > o.LinkIdMaxValue() {
			return fmt.Errorf("Range check failed on o.LinkId (%v < %v > %v)", o.LinkIdMinValue(), o.LinkId, o.LinkIdMaxValue())
		}
	}
	if o.CreatedByInActingVersion(actingVersion) {
		if o.CreatedBy < o.CreatedByMinValue() || o.CreatedBy > o.CreatedByMaxValue() {
			return fmt.Errorf("Range check failed on o.CreatedBy (%v < %v > %v)", o.CreatedByMinValue(), o.CreatedBy, o.CreatedByMaxValue())
		}
	}
	if o.OrderQtyInActingVersion(actingVersion) {
		if o.OrderQty < o.OrderQtyMinValue() || o.OrderQty > o.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderQty (%v < %v > %v)", o.OrderQtyMinValue(), o.OrderQty, o.OrderQtyMaxValue())
		}
	}
	if o.MinQtyInActingVersion(actingVersion) {
		if o.MinQty < o.MinQtyMinValue() || o.MinQty > o.MinQtyMaxValue() {
			return fmt.Errorf("Range check failed on o.MinQty (%v < %v > %v)", o.MinQtyMinValue(), o.MinQty, o.MinQtyMaxValue())
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
	if o.OrderTypeInActingVersion(actingVersion) {
		if o.OrderType < o.OrderTypeMinValue() || o.OrderType > o.OrderTypeMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderType (%v < %v > %v)", o.OrderTypeMinValue(), o.OrderType, o.OrderTypeMaxValue())
		}
	}
	if o.TimeInForceInActingVersion(actingVersion) {
		if o.TimeInForce < o.TimeInForceMinValue() || o.TimeInForce > o.TimeInForceMaxValue() {
			return fmt.Errorf("Range check failed on o.TimeInForce (%v < %v > %v)", o.TimeInForceMinValue(), o.TimeInForce, o.TimeInForceMaxValue())
		}
	}
	if o.CancelReplaceInActingVersion(actingVersion) {
		if o.CancelReplace < o.CancelReplaceMinValue() || o.CancelReplace > o.CancelReplaceMaxValue() {
			return fmt.Errorf("Range check failed on o.CancelReplace (%v < %v > %v)", o.CancelReplaceMinValue(), o.CancelReplace, o.CancelReplaceMaxValue())
		}
	}
	if o.ExpireTimeInActingVersion(actingVersion) {
		if o.ExpireTime < o.ExpireTimeMinValue() || o.ExpireTime > o.ExpireTimeMaxValue() {
			return fmt.Errorf("Range check failed on o.ExpireTime (%v < %v > %v)", o.ExpireTimeMinValue(), o.ExpireTime, o.ExpireTimeMaxValue())
		}
	}
	if o.BaseOrderInActingVersion(actingVersion) {
		if o.BaseOrder < o.BaseOrderMinValue() || o.BaseOrder > o.BaseOrderMaxValue() {
			return fmt.Errorf("Range check failed on o.BaseOrder (%v < %v > %v)", o.BaseOrderMinValue(), o.BaseOrder, o.BaseOrderMaxValue())
		}
	}
	if !utf8.Valid(o.RequestId[:]) {
		return errors.New("o.RequestId failed UTF-8 validation")
	}
	if !utf8.Valid(o.OrigOrderId[:]) {
		return errors.New("o.OrigOrderId failed UTF-8 validation")
	}
	return nil
}

func OrderBufferInit(o *OrderBuffer) {
	return
}

func (*OrderBuffer) SbeBlockLength() (blockLength uint16) {
	return 73
}

func (*OrderBuffer) SbeTemplateId() (templateId uint16) {
	return 17
}

func (*OrderBuffer) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderBuffer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OrderBuffer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OrderBuffer) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderBuffer) TradeSourceId() uint16 {
	return 18
}

func (*OrderBuffer) TradeSourceSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) TradeSourceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TradeSourceSinceVersion()
}

func (*OrderBuffer) TradeSourceDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) TradeSourceMetaAttribute(meta int) string {
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

func (*OrderBuffer) TradeSourceMinValue() uint8 {
	return 0
}

func (*OrderBuffer) TradeSourceMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderBuffer) TradeSourceNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderBuffer) AcctIdId() uint16 {
	return 19
}

func (*OrderBuffer) AcctIdSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AcctIdSinceVersion()
}

func (*OrderBuffer) AcctIdDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) AcctIdMetaAttribute(meta int) string {
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

func (*OrderBuffer) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderBuffer) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderBuffer) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*OrderBuffer) ProxyIndexId() uint16 {
	return 20
}

func (*OrderBuffer) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ProxyIndexSinceVersion()
}

func (*OrderBuffer) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) ProxyIndexMetaAttribute(meta int) string {
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

func (*OrderBuffer) ProxyIndexMinValue() uint16 {
	return 0
}

func (*OrderBuffer) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderBuffer) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderBuffer) CreateTimeId() uint16 {
	return 21
}

func (*OrderBuffer) CreateTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) CreateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CreateTimeSinceVersion()
}

func (*OrderBuffer) CreateTimeDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) CreateTimeMetaAttribute(meta int) string {
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

func (*OrderBuffer) CreateTimeMinValue() uint64 {
	return 0
}

func (*OrderBuffer) CreateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderBuffer) CreateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderBuffer) InstIdId() uint16 {
	return 22
}

func (*OrderBuffer) InstIdSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstIdSinceVersion()
}

func (*OrderBuffer) InstIdDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) InstIdMetaAttribute(meta int) string {
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

func (*OrderBuffer) InstIdMinValue() uint16 {
	return 0
}

func (*OrderBuffer) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderBuffer) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderBuffer) LinkIdId() uint16 {
	return 24
}

func (*OrderBuffer) LinkIdSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) LinkIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.LinkIdSinceVersion()
}

func (*OrderBuffer) LinkIdDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) LinkIdMetaAttribute(meta int) string {
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

func (*OrderBuffer) LinkIdMinValue() uint64 {
	return 0
}

func (*OrderBuffer) LinkIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderBuffer) LinkIdNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderBuffer) CreatedById() uint16 {
	return 25
}

func (*OrderBuffer) CreatedBySinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) CreatedByInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CreatedBySinceVersion()
}

func (*OrderBuffer) CreatedByDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) CreatedByMetaAttribute(meta int) string {
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

func (*OrderBuffer) CreatedByMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderBuffer) CreatedByMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderBuffer) CreatedByNullValue() int32 {
	return math.MinInt32
}

func (*OrderBuffer) OrderQtyId() uint16 {
	return 28
}

func (*OrderBuffer) OrderQtySinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderQtySinceVersion()
}

func (*OrderBuffer) OrderQtyDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) OrderQtyMetaAttribute(meta int) string {
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

func (*OrderBuffer) OrderQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderBuffer) OrderQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderBuffer) OrderQtyNullValue() float64 {
	return math.NaN()
}

func (*OrderBuffer) MinQtyId() uint16 {
	return 29
}

func (*OrderBuffer) MinQtySinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.MinQtySinceVersion()
}

func (*OrderBuffer) MinQtyDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) MinQtyMetaAttribute(meta int) string {
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

func (*OrderBuffer) MinQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderBuffer) MinQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderBuffer) MinQtyNullValue() float64 {
	return math.NaN()
}

func (*OrderBuffer) PriceId() uint16 {
	return 30
}

func (*OrderBuffer) PriceSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.PriceSinceVersion()
}

func (*OrderBuffer) PriceDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) PriceMetaAttribute(meta int) string {
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

func (*OrderBuffer) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderBuffer) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderBuffer) PriceNullValue() float64 {
	return math.NaN()
}

func (*OrderBuffer) StopPriceId() uint16 {
	return 31
}

func (*OrderBuffer) StopPriceSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) StopPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.StopPriceSinceVersion()
}

func (*OrderBuffer) StopPriceDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) StopPriceMetaAttribute(meta int) string {
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

func (*OrderBuffer) StopPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*OrderBuffer) StopPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*OrderBuffer) StopPriceNullValue() float64 {
	return math.NaN()
}

func (*OrderBuffer) OrderTypeId() uint16 {
	return 32
}

func (*OrderBuffer) OrderTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderTypeSinceVersion()
}

func (*OrderBuffer) OrderTypeDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) OrderTypeMetaAttribute(meta int) string {
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

func (*OrderBuffer) OrderTypeMinValue() uint8 {
	return 0
}

func (*OrderBuffer) OrderTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderBuffer) OrderTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderBuffer) TimeInForceId() uint16 {
	return 33
}

func (*OrderBuffer) TimeInForceSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.TimeInForceSinceVersion()
}

func (*OrderBuffer) TimeInForceDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) TimeInForceMetaAttribute(meta int) string {
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

func (*OrderBuffer) TimeInForceMinValue() uint8 {
	return 0
}

func (*OrderBuffer) TimeInForceMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderBuffer) TimeInForceNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderBuffer) CancelReplaceId() uint16 {
	return 34
}

func (*OrderBuffer) CancelReplaceSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) CancelReplaceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelReplaceSinceVersion()
}

func (*OrderBuffer) CancelReplaceDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) CancelReplaceMetaAttribute(meta int) string {
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

func (*OrderBuffer) CancelReplaceMinValue() uint8 {
	return 0
}

func (*OrderBuffer) CancelReplaceMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderBuffer) CancelReplaceNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderBuffer) ExpireTimeId() uint16 {
	return 35
}

func (*OrderBuffer) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ExpireTimeSinceVersion()
}

func (*OrderBuffer) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) ExpireTimeMetaAttribute(meta int) string {
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

func (*OrderBuffer) ExpireTimeMinValue() uint64 {
	return 0
}

func (*OrderBuffer) ExpireTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*OrderBuffer) ExpireTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*OrderBuffer) BaseOrderId() uint16 {
	return 37
}

func (*OrderBuffer) BaseOrderSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) BaseOrderInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.BaseOrderSinceVersion()
}

func (*OrderBuffer) BaseOrderDeprecated() uint16 {
	return 0
}

func (*OrderBuffer) BaseOrderMetaAttribute(meta int) string {
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

func (*OrderBuffer) BaseOrderMinValue() uint8 {
	return 0
}

func (*OrderBuffer) BaseOrderMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderBuffer) BaseOrderNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderBuffer) RequestIdMetaAttribute(meta int) string {
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

func (*OrderBuffer) RequestIdSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RequestIdSinceVersion()
}

func (*OrderBuffer) RequestIdDeprecated() uint16 {
	return 0
}

func (OrderBuffer) RequestIdCharacterEncoding() string {
	return "UTF-8"
}

func (OrderBuffer) RequestIdHeaderLength() uint64 {
	return 2
}

func (*OrderBuffer) OrigOrderIdMetaAttribute(meta int) string {
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

func (*OrderBuffer) OrigOrderIdSinceVersion() uint16 {
	return 0
}

func (o *OrderBuffer) OrigOrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrigOrderIdSinceVersion()
}

func (*OrderBuffer) OrigOrderIdDeprecated() uint16 {
	return 0
}

func (OrderBuffer) OrigOrderIdCharacterEncoding() string {
	return "UTF-8"
}

func (OrderBuffer) OrigOrderIdHeaderLength() uint64 {
	return 2
}
