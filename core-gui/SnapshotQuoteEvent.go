// Generated SBE (Simple Binary Encoding) message codec

package core_gui

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type SnapshotQuoteEvent struct {
	OrderId        uint64
	QuoteTime      uint64
	BuySell        uint8
	MinQty         int32
	QuoteEventType uint8
	OrderType      uint8
	TimeInForce    uint8
	Aggressor      uint8
	Price          float64
	DealQty        float64
	MarketSequence uint64
	AcctId         int32
	InstId         uint16
}

func (s *SnapshotQuoteEvent) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := s.RangeCheck(s.SbeSchemaVersion(), s.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, s.OrderId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.QuoteTime); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.BuySell); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.MinQty); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.QuoteEventType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.OrderType); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.TimeInForce); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, s.Aggressor); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, s.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, s.DealQty); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, s.MarketSequence); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, s.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, s.InstId); err != nil {
		return err
	}
	return nil
}

func (s *SnapshotQuoteEvent) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !s.OrderIdInActingVersion(actingVersion) {
		s.OrderId = s.OrderIdNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.OrderId); err != nil {
			return err
		}
	}
	if !s.QuoteTimeInActingVersion(actingVersion) {
		s.QuoteTime = s.QuoteTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.QuoteTime); err != nil {
			return err
		}
	}
	if !s.BuySellInActingVersion(actingVersion) {
		s.BuySell = s.BuySellNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.BuySell); err != nil {
			return err
		}
	}
	if !s.MinQtyInActingVersion(actingVersion) {
		s.MinQty = s.MinQtyNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.MinQty); err != nil {
			return err
		}
	}
	if !s.QuoteEventTypeInActingVersion(actingVersion) {
		s.QuoteEventType = s.QuoteEventTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.QuoteEventType); err != nil {
			return err
		}
	}
	if !s.OrderTypeInActingVersion(actingVersion) {
		s.OrderType = s.OrderTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.OrderType); err != nil {
			return err
		}
	}
	if !s.TimeInForceInActingVersion(actingVersion) {
		s.TimeInForce = s.TimeInForceNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.TimeInForce); err != nil {
			return err
		}
	}
	if !s.AggressorInActingVersion(actingVersion) {
		s.Aggressor = s.AggressorNullValue()
	} else {
		if err := _m.ReadUint8(_r, &s.Aggressor); err != nil {
			return err
		}
	}
	if !s.PriceInActingVersion(actingVersion) {
		s.Price = s.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &s.Price); err != nil {
			return err
		}
	}
	if !s.DealQtyInActingVersion(actingVersion) {
		s.DealQty = s.DealQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &s.DealQty); err != nil {
			return err
		}
	}
	if !s.MarketSequenceInActingVersion(actingVersion) {
		s.MarketSequence = s.MarketSequenceNullValue()
	} else {
		if err := _m.ReadUint64(_r, &s.MarketSequence); err != nil {
			return err
		}
	}
	if !s.AcctIdInActingVersion(actingVersion) {
		s.AcctId = s.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &s.AcctId); err != nil {
			return err
		}
	}
	if !s.InstIdInActingVersion(actingVersion) {
		s.InstId = s.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &s.InstId); err != nil {
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

func (s *SnapshotQuoteEvent) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if s.OrderIdInActingVersion(actingVersion) {
		if s.OrderId < s.OrderIdMinValue() || s.OrderId > s.OrderIdMaxValue() {
			return fmt.Errorf("Range check failed on s.OrderId (%v < %v > %v)", s.OrderIdMinValue(), s.OrderId, s.OrderIdMaxValue())
		}
	}
	if s.QuoteTimeInActingVersion(actingVersion) {
		if s.QuoteTime < s.QuoteTimeMinValue() || s.QuoteTime > s.QuoteTimeMaxValue() {
			return fmt.Errorf("Range check failed on s.QuoteTime (%v < %v > %v)", s.QuoteTimeMinValue(), s.QuoteTime, s.QuoteTimeMaxValue())
		}
	}
	if s.BuySellInActingVersion(actingVersion) {
		if s.BuySell < s.BuySellMinValue() || s.BuySell > s.BuySellMaxValue() {
			return fmt.Errorf("Range check failed on s.BuySell (%v < %v > %v)", s.BuySellMinValue(), s.BuySell, s.BuySellMaxValue())
		}
	}
	if s.MinQtyInActingVersion(actingVersion) {
		if s.MinQty < s.MinQtyMinValue() || s.MinQty > s.MinQtyMaxValue() {
			return fmt.Errorf("Range check failed on s.MinQty (%v < %v > %v)", s.MinQtyMinValue(), s.MinQty, s.MinQtyMaxValue())
		}
	}
	if s.QuoteEventTypeInActingVersion(actingVersion) {
		if s.QuoteEventType < s.QuoteEventTypeMinValue() || s.QuoteEventType > s.QuoteEventTypeMaxValue() {
			return fmt.Errorf("Range check failed on s.QuoteEventType (%v < %v > %v)", s.QuoteEventTypeMinValue(), s.QuoteEventType, s.QuoteEventTypeMaxValue())
		}
	}
	if s.OrderTypeInActingVersion(actingVersion) {
		if s.OrderType < s.OrderTypeMinValue() || s.OrderType > s.OrderTypeMaxValue() {
			return fmt.Errorf("Range check failed on s.OrderType (%v < %v > %v)", s.OrderTypeMinValue(), s.OrderType, s.OrderTypeMaxValue())
		}
	}
	if s.TimeInForceInActingVersion(actingVersion) {
		if s.TimeInForce < s.TimeInForceMinValue() || s.TimeInForce > s.TimeInForceMaxValue() {
			return fmt.Errorf("Range check failed on s.TimeInForce (%v < %v > %v)", s.TimeInForceMinValue(), s.TimeInForce, s.TimeInForceMaxValue())
		}
	}
	if s.AggressorInActingVersion(actingVersion) {
		if s.Aggressor < s.AggressorMinValue() || s.Aggressor > s.AggressorMaxValue() {
			return fmt.Errorf("Range check failed on s.Aggressor (%v < %v > %v)", s.AggressorMinValue(), s.Aggressor, s.AggressorMaxValue())
		}
	}
	if s.PriceInActingVersion(actingVersion) {
		if s.Price < s.PriceMinValue() || s.Price > s.PriceMaxValue() {
			return fmt.Errorf("Range check failed on s.Price (%v < %v > %v)", s.PriceMinValue(), s.Price, s.PriceMaxValue())
		}
	}
	if s.DealQtyInActingVersion(actingVersion) {
		if s.DealQty < s.DealQtyMinValue() || s.DealQty > s.DealQtyMaxValue() {
			return fmt.Errorf("Range check failed on s.DealQty (%v < %v > %v)", s.DealQtyMinValue(), s.DealQty, s.DealQtyMaxValue())
		}
	}
	if s.MarketSequenceInActingVersion(actingVersion) {
		if s.MarketSequence < s.MarketSequenceMinValue() || s.MarketSequence > s.MarketSequenceMaxValue() {
			return fmt.Errorf("Range check failed on s.MarketSequence (%v < %v > %v)", s.MarketSequenceMinValue(), s.MarketSequence, s.MarketSequenceMaxValue())
		}
	}
	if s.AcctIdInActingVersion(actingVersion) {
		if s.AcctId < s.AcctIdMinValue() || s.AcctId > s.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on s.AcctId (%v < %v > %v)", s.AcctIdMinValue(), s.AcctId, s.AcctIdMaxValue())
		}
	}
	if s.InstIdInActingVersion(actingVersion) {
		if s.InstId < s.InstIdMinValue() || s.InstId > s.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on s.InstId (%v < %v > %v)", s.InstIdMinValue(), s.InstId, s.InstIdMaxValue())
		}
	}
	return nil
}

func SnapshotQuoteEventInit(s *SnapshotQuoteEvent) {
	return
}

func (*SnapshotQuoteEvent) SbeBlockLength() (blockLength uint16) {
	return 55
}

func (*SnapshotQuoteEvent) SbeTemplateId() (templateId uint16) {
	return 125
}

func (*SnapshotQuoteEvent) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*SnapshotQuoteEvent) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*SnapshotQuoteEvent) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*SnapshotQuoteEvent) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*SnapshotQuoteEvent) OrderIdId() uint16 {
	return 126
}

func (*SnapshotQuoteEvent) OrderIdSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OrderIdSinceVersion()
}

func (*SnapshotQuoteEvent) OrderIdDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) OrderIdMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) OrderIdMinValue() uint64 {
	return 0
}

func (*SnapshotQuoteEvent) OrderIdMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotQuoteEvent) OrderIdNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotQuoteEvent) QuoteTimeId() uint16 {
	return 127
}

func (*SnapshotQuoteEvent) QuoteTimeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) QuoteTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.QuoteTimeSinceVersion()
}

func (*SnapshotQuoteEvent) QuoteTimeDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) QuoteTimeMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) QuoteTimeMinValue() uint64 {
	return 0
}

func (*SnapshotQuoteEvent) QuoteTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotQuoteEvent) QuoteTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotQuoteEvent) BuySellId() uint16 {
	return 128
}

func (*SnapshotQuoteEvent) BuySellSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) BuySellInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.BuySellSinceVersion()
}

func (*SnapshotQuoteEvent) BuySellDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) BuySellMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) BuySellMinValue() uint8 {
	return 0
}

func (*SnapshotQuoteEvent) BuySellMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SnapshotQuoteEvent) BuySellNullValue() uint8 {
	return math.MaxUint8
}

func (*SnapshotQuoteEvent) MinQtyId() uint16 {
	return 129
}

func (*SnapshotQuoteEvent) MinQtySinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) MinQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MinQtySinceVersion()
}

func (*SnapshotQuoteEvent) MinQtyDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) MinQtyMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) MinQtyMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotQuoteEvent) MinQtyMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotQuoteEvent) MinQtyNullValue() int32 {
	return math.MinInt32
}

func (*SnapshotQuoteEvent) QuoteEventTypeId() uint16 {
	return 130
}

func (*SnapshotQuoteEvent) QuoteEventTypeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) QuoteEventTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.QuoteEventTypeSinceVersion()
}

func (*SnapshotQuoteEvent) QuoteEventTypeDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) QuoteEventTypeMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) QuoteEventTypeMinValue() uint8 {
	return 0
}

func (*SnapshotQuoteEvent) QuoteEventTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SnapshotQuoteEvent) QuoteEventTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*SnapshotQuoteEvent) OrderTypeId() uint16 {
	return 131
}

func (*SnapshotQuoteEvent) OrderTypeSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) OrderTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.OrderTypeSinceVersion()
}

func (*SnapshotQuoteEvent) OrderTypeDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) OrderTypeMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) OrderTypeMinValue() uint8 {
	return 0
}

func (*SnapshotQuoteEvent) OrderTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SnapshotQuoteEvent) OrderTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*SnapshotQuoteEvent) TimeInForceId() uint16 {
	return 132
}

func (*SnapshotQuoteEvent) TimeInForceSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) TimeInForceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.TimeInForceSinceVersion()
}

func (*SnapshotQuoteEvent) TimeInForceDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) TimeInForceMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) TimeInForceMinValue() uint8 {
	return 0
}

func (*SnapshotQuoteEvent) TimeInForceMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SnapshotQuoteEvent) TimeInForceNullValue() uint8 {
	return math.MaxUint8
}

func (*SnapshotQuoteEvent) AggressorId() uint16 {
	return 133
}

func (*SnapshotQuoteEvent) AggressorSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) AggressorInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.AggressorSinceVersion()
}

func (*SnapshotQuoteEvent) AggressorDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) AggressorMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) AggressorMinValue() uint8 {
	return 0
}

func (*SnapshotQuoteEvent) AggressorMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*SnapshotQuoteEvent) AggressorNullValue() uint8 {
	return math.MaxUint8
}

func (*SnapshotQuoteEvent) PriceId() uint16 {
	return 134
}

func (*SnapshotQuoteEvent) PriceSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.PriceSinceVersion()
}

func (*SnapshotQuoteEvent) PriceDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) PriceMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*SnapshotQuoteEvent) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*SnapshotQuoteEvent) PriceNullValue() float64 {
	return math.NaN()
}

func (*SnapshotQuoteEvent) DealQtyId() uint16 {
	return 135
}

func (*SnapshotQuoteEvent) DealQtySinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) DealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.DealQtySinceVersion()
}

func (*SnapshotQuoteEvent) DealQtyDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) DealQtyMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) DealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*SnapshotQuoteEvent) DealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*SnapshotQuoteEvent) DealQtyNullValue() float64 {
	return math.NaN()
}

func (*SnapshotQuoteEvent) MarketSequenceId() uint16 {
	return 137
}

func (*SnapshotQuoteEvent) MarketSequenceSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) MarketSequenceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.MarketSequenceSinceVersion()
}

func (*SnapshotQuoteEvent) MarketSequenceDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) MarketSequenceMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) MarketSequenceMinValue() uint64 {
	return 0
}

func (*SnapshotQuoteEvent) MarketSequenceMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*SnapshotQuoteEvent) MarketSequenceNullValue() uint64 {
	return math.MaxUint64
}

func (*SnapshotQuoteEvent) AcctIdId() uint16 {
	return 138
}

func (*SnapshotQuoteEvent) AcctIdSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.AcctIdSinceVersion()
}

func (*SnapshotQuoteEvent) AcctIdDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) AcctIdMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*SnapshotQuoteEvent) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*SnapshotQuoteEvent) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*SnapshotQuoteEvent) InstIdId() uint16 {
	return 139
}

func (*SnapshotQuoteEvent) InstIdSinceVersion() uint16 {
	return 0
}

func (s *SnapshotQuoteEvent) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= s.InstIdSinceVersion()
}

func (*SnapshotQuoteEvent) InstIdDeprecated() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) InstIdMetaAttribute(meta int) string {
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

func (*SnapshotQuoteEvent) InstIdMinValue() uint16 {
	return 0
}

func (*SnapshotQuoteEvent) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*SnapshotQuoteEvent) InstIdNullValue() uint16 {
	return math.MaxUint16
}
