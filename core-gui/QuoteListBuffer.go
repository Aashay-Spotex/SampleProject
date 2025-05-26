// Generated SBE (Simple Binary Encoding) message codec

package core_gui

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type QuoteListBuffer struct {
	AcctId         int32
	MarketSequence uint64
	InstId         uint16
	QuoteTime      uint64
	QuoteEventList []QuoteListBufferQuoteEventList
}
type QuoteListBufferQuoteEventList struct {
	OrderId        uint64
	BuySell        uint8
	Price          float64
	DealQty        float64
	MinQty         int32
	QuoteEventType uint8
	OrderType      uint8
	TimeInForce    uint8
	Aggressor      uint8
}

func (q *QuoteListBuffer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := q.RangeCheck(q.SbeSchemaVersion(), q.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, q.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.MarketSequence); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, q.InstId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, q.QuoteTime); err != nil {
		return err
	}
	var QuoteEventListBlockLength uint16 = 33
	if err := _m.WriteUint16(_w, QuoteEventListBlockLength); err != nil {
		return err
	}
	var QuoteEventListNumInGroup uint16 = uint16(len(q.QuoteEventList))
	if err := _m.WriteUint16(_w, QuoteEventListNumInGroup); err != nil {
		return err
	}
	for i := range q.QuoteEventList {
		if err := q.QuoteEventList[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (q *QuoteListBuffer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !q.AcctIdInActingVersion(actingVersion) {
		q.AcctId = q.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &q.AcctId); err != nil {
			return err
		}
	}
	if !q.MarketSequenceInActingVersion(actingVersion) {
		q.MarketSequence = q.MarketSequenceNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.MarketSequence); err != nil {
			return err
		}
	}
	if !q.InstIdInActingVersion(actingVersion) {
		q.InstId = q.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &q.InstId); err != nil {
			return err
		}
	}
	if !q.QuoteTimeInActingVersion(actingVersion) {
		q.QuoteTime = q.QuoteTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.QuoteTime); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}

	if q.QuoteEventListInActingVersion(actingVersion) {
		var QuoteEventListBlockLength uint16
		if err := _m.ReadUint16(_r, &QuoteEventListBlockLength); err != nil {
			return err
		}
		var QuoteEventListNumInGroup uint16
		if err := _m.ReadUint16(_r, &QuoteEventListNumInGroup); err != nil {
			return err
		}
		if cap(q.QuoteEventList) < int(QuoteEventListNumInGroup) {
			q.QuoteEventList = make([]QuoteListBufferQuoteEventList, QuoteEventListNumInGroup)
		}
		q.QuoteEventList = q.QuoteEventList[:QuoteEventListNumInGroup]
		for i := range q.QuoteEventList {
			if err := q.QuoteEventList[i].Decode(_m, _r, actingVersion, uint(QuoteEventListBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := q.RangeCheck(actingVersion, q.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (q *QuoteListBuffer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.AcctIdInActingVersion(actingVersion) {
		if q.AcctId < q.AcctIdMinValue() || q.AcctId > q.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on q.AcctId (%v < %v > %v)", q.AcctIdMinValue(), q.AcctId, q.AcctIdMaxValue())
		}
	}
	if q.MarketSequenceInActingVersion(actingVersion) {
		if q.MarketSequence < q.MarketSequenceMinValue() || q.MarketSequence > q.MarketSequenceMaxValue() {
			return fmt.Errorf("Range check failed on q.MarketSequence (%v < %v > %v)", q.MarketSequenceMinValue(), q.MarketSequence, q.MarketSequenceMaxValue())
		}
	}
	if q.InstIdInActingVersion(actingVersion) {
		if q.InstId < q.InstIdMinValue() || q.InstId > q.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on q.InstId (%v < %v > %v)", q.InstIdMinValue(), q.InstId, q.InstIdMaxValue())
		}
	}
	if q.QuoteTimeInActingVersion(actingVersion) {
		if q.QuoteTime < q.QuoteTimeMinValue() || q.QuoteTime > q.QuoteTimeMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteTime (%v < %v > %v)", q.QuoteTimeMinValue(), q.QuoteTime, q.QuoteTimeMaxValue())
		}
	}
	for i := range q.QuoteEventList {
		if err := q.QuoteEventList[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func QuoteListBufferInit(q *QuoteListBuffer) {
	return
}

func (q *QuoteListBufferQuoteEventList) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint64(_w, q.OrderId); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.BuySell); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, q.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, q.DealQty); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, q.MinQty); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.QuoteEventType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.OrderType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.TimeInForce); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, q.Aggressor); err != nil {
		return err
	}
	return nil
}

func (q *QuoteListBufferQuoteEventList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !q.OrderIdInActingVersion(actingVersion) {
		q.OrderId = q.OrderIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &q.OrderId); err != nil {
			return err
		}
	}
	if !q.BuySellInActingVersion(actingVersion) {
		q.BuySell = q.BuySellNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.BuySell); err != nil {
			return err
		}
	}
	if !q.PriceInActingVersion(actingVersion) {
		q.Price = q.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &q.Price); err != nil {
			return err
		}
	}
	if !q.DealQtyInActingVersion(actingVersion) {
		q.DealQty = q.DealQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &q.DealQty); err != nil {
			return err
		}
	}
	if !q.MinQtyInActingVersion(actingVersion) {
		q.MinQty = q.MinQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &q.MinQty); err != nil {
			return err
		}
	}
	if !q.QuoteEventTypeInActingVersion(actingVersion) {
		q.QuoteEventType = q.QuoteEventTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.QuoteEventType); err != nil {
			return err
		}
	}
	if !q.OrderTypeInActingVersion(actingVersion) {
		q.OrderType = q.OrderTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.OrderType); err != nil {
			return err
		}
	}
	if !q.TimeInForceInActingVersion(actingVersion) {
		q.TimeInForce = q.TimeInForceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.TimeInForce); err != nil {
			return err
		}
	}
	if !q.AggressorInActingVersion(actingVersion) {
		q.Aggressor = q.AggressorNullValue()
	} else {
		if err := _m.ReadUint8(_r, &q.Aggressor); err != nil {
			return err
		}
	}
	if actingVersion > q.SbeSchemaVersion() && blockLength > q.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-q.SbeBlockLength()))
	}
	return nil
}

func (q *QuoteListBufferQuoteEventList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if q.OrderIdInActingVersion(actingVersion) {
		if q.OrderId < q.OrderIdMinValue() || q.OrderId > q.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on q.OrderId (%v < %v > %v)", q.OrderIdMinValue(), q.OrderId, q.OrderIdMaxValue())
		}
	}
	if q.BuySellInActingVersion(actingVersion) {
		if q.BuySell < q.BuySellMinValue() || q.BuySell > q.BuySellMaxValue() {
			return fmt.Errorf("Range check failed on q.BuySell (%v < %v > %v)", q.BuySellMinValue(), q.BuySell, q.BuySellMaxValue())
		}
	}
	if q.PriceInActingVersion(actingVersion) {
		if q.Price < q.PriceMinValue() || q.Price > q.PriceMaxValue() {
			return fmt.Errorf("Range check failed on q.Price (%v < %v > %v)", q.PriceMinValue(), q.Price, q.PriceMaxValue())
		}
	}
	if q.DealQtyInActingVersion(actingVersion) {
		if q.DealQty < q.DealQtyMinValue() || q.DealQty > q.DealQtyMaxValue() {
			return fmt.Errorf("Range check failed on q.DealQty (%v < %v > %v)", q.DealQtyMinValue(), q.DealQty, q.DealQtyMaxValue())
		}
	}
	if q.MinQtyInActingVersion(actingVersion) {
		if q.MinQty < q.MinQtyMinValue() || q.MinQty > q.MinQtyMaxValue() {
			return fmt.Errorf("Range check failed on q.MinQty (%v < %v > %v)", q.MinQtyMinValue(), q.MinQty, q.MinQtyMaxValue())
		}
	}
	if q.QuoteEventTypeInActingVersion(actingVersion) {
		if q.QuoteEventType < q.QuoteEventTypeMinValue() || q.QuoteEventType > q.QuoteEventTypeMaxValue() {
			return fmt.Errorf("Range check failed on q.QuoteEventType (%v < %v > %v)", q.QuoteEventTypeMinValue(), q.QuoteEventType, q.QuoteEventTypeMaxValue())
		}
	}
	if q.OrderTypeInActingVersion(actingVersion) {
		if q.OrderType < q.OrderTypeMinValue() || q.OrderType > q.OrderTypeMaxValue() {
			return fmt.Errorf("Range check failed on q.OrderType (%v < %v > %v)", q.OrderTypeMinValue(), q.OrderType, q.OrderTypeMaxValue())
		}
	}
	if q.TimeInForceInActingVersion(actingVersion) {
		if q.TimeInForce < q.TimeInForceMinValue() || q.TimeInForce > q.TimeInForceMaxValue() {
			return fmt.Errorf("Range check failed on q.TimeInForce (%v < %v > %v)", q.TimeInForceMinValue(), q.TimeInForce, q.TimeInForceMaxValue())
		}
	}
	if q.AggressorInActingVersion(actingVersion) {
		if q.Aggressor < q.AggressorMinValue() || q.Aggressor > q.AggressorMaxValue() {
			return fmt.Errorf("Range check failed on q.Aggressor (%v < %v > %v)", q.AggressorMinValue(), q.Aggressor, q.AggressorMaxValue())
		}
	}
	return nil
}

func QuoteListBufferQuoteEventListInit(q *QuoteListBufferQuoteEventList) {
	return
}

func (*QuoteListBuffer) SbeBlockLength() (blockLength uint16) {
	return 22
}

func (*QuoteListBuffer) SbeTemplateId() (templateId uint16) {
	return 1
}

func (*QuoteListBuffer) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*QuoteListBuffer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*QuoteListBuffer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*QuoteListBuffer) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*QuoteListBuffer) AcctIdId() uint16 {
	return 2
}

func (*QuoteListBuffer) AcctIdSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBuffer) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.AcctIdSinceVersion()
}

func (*QuoteListBuffer) AcctIdDeprecated() uint16 {
	return 0
}

func (*QuoteListBuffer) AcctIdMetaAttribute(meta int) string {
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

func (*QuoteListBuffer) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*QuoteListBuffer) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*QuoteListBuffer) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*QuoteListBuffer) MarketSequenceId() uint16 {
	return 3
}

func (*QuoteListBuffer) MarketSequenceSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBuffer) MarketSequenceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.MarketSequenceSinceVersion()
}

func (*QuoteListBuffer) MarketSequenceDeprecated() uint16 {
	return 0
}

func (*QuoteListBuffer) MarketSequenceMetaAttribute(meta int) string {
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

func (*QuoteListBuffer) MarketSequenceMinValue() uint64 {
	return 0
}

func (*QuoteListBuffer) MarketSequenceMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteListBuffer) MarketSequenceNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteListBuffer) InstIdId() uint16 {
	return 4
}

func (*QuoteListBuffer) InstIdSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBuffer) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.InstIdSinceVersion()
}

func (*QuoteListBuffer) InstIdDeprecated() uint16 {
	return 0
}

func (*QuoteListBuffer) InstIdMetaAttribute(meta int) string {
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

func (*QuoteListBuffer) InstIdMinValue() uint16 {
	return 0
}

func (*QuoteListBuffer) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*QuoteListBuffer) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*QuoteListBuffer) QuoteTimeId() uint16 {
	return 5
}

func (*QuoteListBuffer) QuoteTimeSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBuffer) QuoteTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteTimeSinceVersion()
}

func (*QuoteListBuffer) QuoteTimeDeprecated() uint16 {
	return 0
}

func (*QuoteListBuffer) QuoteTimeMetaAttribute(meta int) string {
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

func (*QuoteListBuffer) QuoteTimeMinValue() uint64 {
	return 0
}

func (*QuoteListBuffer) QuoteTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteListBuffer) QuoteTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteListBufferQuoteEventList) OrderIdId() uint16 {
	return 7
}

func (*QuoteListBufferQuoteEventList) OrderIdSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.OrderIdSinceVersion()
}

func (*QuoteListBufferQuoteEventList) OrderIdDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) OrderIdMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) OrderIdMinValue() uint64 {
	return 0
}

func (*QuoteListBufferQuoteEventList) OrderIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*QuoteListBufferQuoteEventList) OrderIdNullValue() uint64 {
	return math.MaxUint64
}

func (*QuoteListBufferQuoteEventList) BuySellId() uint16 {
	return 8
}

func (*QuoteListBufferQuoteEventList) BuySellSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) BuySellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.BuySellSinceVersion()
}

func (*QuoteListBufferQuoteEventList) BuySellDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) BuySellMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) BuySellMinValue() uint8 {
	return 0
}

func (*QuoteListBufferQuoteEventList) BuySellMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*QuoteListBufferQuoteEventList) BuySellNullValue() uint8 {
	return math.MaxUint8
}

func (*QuoteListBufferQuoteEventList) PriceId() uint16 {
	return 9
}

func (*QuoteListBufferQuoteEventList) PriceSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.PriceSinceVersion()
}

func (*QuoteListBufferQuoteEventList) PriceDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) PriceMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*QuoteListBufferQuoteEventList) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*QuoteListBufferQuoteEventList) PriceNullValue() float64 {
	return math.NaN()
}

func (*QuoteListBufferQuoteEventList) DealQtyId() uint16 {
	return 10
}

func (*QuoteListBufferQuoteEventList) DealQtySinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) DealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.DealQtySinceVersion()
}

func (*QuoteListBufferQuoteEventList) DealQtyDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) DealQtyMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) DealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*QuoteListBufferQuoteEventList) DealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*QuoteListBufferQuoteEventList) DealQtyNullValue() float64 {
	return math.NaN()
}

func (*QuoteListBufferQuoteEventList) MinQtyId() uint16 {
	return 12
}

func (*QuoteListBufferQuoteEventList) MinQtySinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.MinQtySinceVersion()
}

func (*QuoteListBufferQuoteEventList) MinQtyDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) MinQtyMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) MinQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*QuoteListBufferQuoteEventList) MinQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*QuoteListBufferQuoteEventList) MinQtyNullValue() int32 {
	return math.MinInt32
}

func (*QuoteListBufferQuoteEventList) QuoteEventTypeId() uint16 {
	return 13
}

func (*QuoteListBufferQuoteEventList) QuoteEventTypeSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) QuoteEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteEventTypeSinceVersion()
}

func (*QuoteListBufferQuoteEventList) QuoteEventTypeDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) QuoteEventTypeMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) QuoteEventTypeMinValue() uint8 {
	return 0
}

func (*QuoteListBufferQuoteEventList) QuoteEventTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*QuoteListBufferQuoteEventList) QuoteEventTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*QuoteListBufferQuoteEventList) OrderTypeId() uint16 {
	return 14
}

func (*QuoteListBufferQuoteEventList) OrderTypeSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.OrderTypeSinceVersion()
}

func (*QuoteListBufferQuoteEventList) OrderTypeDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) OrderTypeMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) OrderTypeMinValue() uint8 {
	return 0
}

func (*QuoteListBufferQuoteEventList) OrderTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*QuoteListBufferQuoteEventList) OrderTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*QuoteListBufferQuoteEventList) TimeInForceId() uint16 {
	return 15
}

func (*QuoteListBufferQuoteEventList) TimeInForceSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.TimeInForceSinceVersion()
}

func (*QuoteListBufferQuoteEventList) TimeInForceDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) TimeInForceMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) TimeInForceMinValue() uint8 {
	return 0
}

func (*QuoteListBufferQuoteEventList) TimeInForceMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*QuoteListBufferQuoteEventList) TimeInForceNullValue() uint8 {
	return math.MaxUint8
}

func (*QuoteListBufferQuoteEventList) AggressorId() uint16 {
	return 16
}

func (*QuoteListBufferQuoteEventList) AggressorSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBufferQuoteEventList) AggressorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.AggressorSinceVersion()
}

func (*QuoteListBufferQuoteEventList) AggressorDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) AggressorMetaAttribute(meta int) string {
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

func (*QuoteListBufferQuoteEventList) AggressorMinValue() uint8 {
	return 0
}

func (*QuoteListBufferQuoteEventList) AggressorMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*QuoteListBufferQuoteEventList) AggressorNullValue() uint8 {
	return math.MaxUint8
}

func (*QuoteListBuffer) QuoteEventListId() uint16 {
	return 6
}

func (*QuoteListBuffer) QuoteEventListSinceVersion() uint16 {
	return 0
}

func (q *QuoteListBuffer) QuoteEventListInActingVersion(actingVersion uint16) bool {
	return actingVersion >= q.QuoteEventListSinceVersion()
}

func (*QuoteListBuffer) QuoteEventListDeprecated() uint16 {
	return 0
}

func (*QuoteListBufferQuoteEventList) SbeBlockLength() (blockLength uint) {
	return 33
}

func (*QuoteListBufferQuoteEventList) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}
