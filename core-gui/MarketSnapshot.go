// Generated SBE (Simple Binary Encoding) message codec

package core_gui

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MarketSnapshot struct {
	MarketSequence          uint64
	InstId                  uint16
	Time                    uint64
	OrdersPerPriceLevelList []MarketSnapshotOrdersPerPriceLevelList
}
type MarketSnapshotOrdersPerPriceLevelList struct {
	BidOffer       uint8
	Price          float64
	TotalDealQty   float64
	TotalMMQty     float64
	QuoteOrderList []MarketSnapshotOrdersPerPriceLevelListQuoteOrderList
}
type MarketSnapshotOrdersPerPriceLevelListQuoteOrderList struct {
	AcctId  int32
	DealQty float64
	Price   float64
}

func (m *MarketSnapshot) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint64(_w, m.MarketSequence); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, m.InstId); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, m.Time); err != nil {
		return err
	}
	var OrdersPerPriceLevelListBlockLength uint16 = 25
	if err := _m.WriteUint16(_w, OrdersPerPriceLevelListBlockLength); err != nil {
		return err
	}
	var OrdersPerPriceLevelListNumInGroup uint16 = uint16(len(m.OrdersPerPriceLevelList))
	if err := _m.WriteUint16(_w, OrdersPerPriceLevelListNumInGroup); err != nil {
		return err
	}
	for i := range m.OrdersPerPriceLevelList {
		if err := m.OrdersPerPriceLevelList[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketSnapshot) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.MarketSequenceInActingVersion(actingVersion) {
		m.MarketSequence = m.MarketSequenceNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.MarketSequence); err != nil {
			return err
		}
	}
	if !m.InstIdInActingVersion(actingVersion) {
		m.InstId = m.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.InstId); err != nil {
			return err
		}
	}
	if !m.TimeInActingVersion(actingVersion) {
		m.Time = m.TimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &m.Time); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.OrdersPerPriceLevelListInActingVersion(actingVersion) {
		var OrdersPerPriceLevelListBlockLength uint16
		if err := _m.ReadUint16(_r, &OrdersPerPriceLevelListBlockLength); err != nil {
			return err
		}
		var OrdersPerPriceLevelListNumInGroup uint16
		if err := _m.ReadUint16(_r, &OrdersPerPriceLevelListNumInGroup); err != nil {
			return err
		}
		if cap(m.OrdersPerPriceLevelList) < int(OrdersPerPriceLevelListNumInGroup) {
			m.OrdersPerPriceLevelList = make([]MarketSnapshotOrdersPerPriceLevelList, OrdersPerPriceLevelListNumInGroup)
		}
		m.OrdersPerPriceLevelList = m.OrdersPerPriceLevelList[:OrdersPerPriceLevelListNumInGroup]
		for i := range m.OrdersPerPriceLevelList {
			if err := m.OrdersPerPriceLevelList[i].Decode(_m, _r, actingVersion, uint(OrdersPerPriceLevelListBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := m.RangeCheck(actingVersion, m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketSnapshot) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.MarketSequenceInActingVersion(actingVersion) {
		if m.MarketSequence < m.MarketSequenceMinValue() || m.MarketSequence > m.MarketSequenceMaxValue() {
			return fmt.Errorf("Range check failed on m.MarketSequence (%v < %v > %v)", m.MarketSequenceMinValue(), m.MarketSequence, m.MarketSequenceMaxValue())
		}
	}
	if m.InstIdInActingVersion(actingVersion) {
		if m.InstId < m.InstIdMinValue() || m.InstId > m.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on m.InstId (%v < %v > %v)", m.InstIdMinValue(), m.InstId, m.InstIdMaxValue())
		}
	}
	if m.TimeInActingVersion(actingVersion) {
		if m.Time < m.TimeMinValue() || m.Time > m.TimeMaxValue() {
			return fmt.Errorf("Range check failed on m.Time (%v < %v > %v)", m.TimeMinValue(), m.Time, m.TimeMaxValue())
		}
	}
	for i := range m.OrdersPerPriceLevelList {
		if err := m.OrdersPerPriceLevelList[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MarketSnapshotInit(m *MarketSnapshot) {
	return
}

func (m *MarketSnapshotOrdersPerPriceLevelList) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint8(_w, m.BidOffer); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, m.Price); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, m.TotalDealQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, m.TotalMMQty); err != nil {
		return err
	}
	var QuoteOrderListBlockLength uint16 = 20
	if err := _m.WriteUint16(_w, QuoteOrderListBlockLength); err != nil {
		return err
	}
	var QuoteOrderListNumInGroup uint16 = uint16(len(m.QuoteOrderList))
	if err := _m.WriteUint16(_w, QuoteOrderListNumInGroup); err != nil {
		return err
	}
	for i := range m.QuoteOrderList {
		if err := m.QuoteOrderList[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MarketSnapshotOrdersPerPriceLevelList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.BidOfferInActingVersion(actingVersion) {
		m.BidOffer = m.BidOfferNullValue()
	} else {
		if err := _m.ReadUint8(_r, &m.BidOffer); err != nil {
			return err
		}
	}
	if !m.PriceInActingVersion(actingVersion) {
		m.Price = m.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &m.Price); err != nil {
			return err
		}
	}
	if !m.TotalDealQtyInActingVersion(actingVersion) {
		m.TotalDealQty = m.TotalDealQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &m.TotalDealQty); err != nil {
			return err
		}
	}
	if !m.TotalMMQtyInActingVersion(actingVersion) {
		m.TotalMMQty = m.TotalMMQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &m.TotalMMQty); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.QuoteOrderListInActingVersion(actingVersion) {
		var QuoteOrderListBlockLength uint16
		if err := _m.ReadUint16(_r, &QuoteOrderListBlockLength); err != nil {
			return err
		}
		var QuoteOrderListNumInGroup uint16
		if err := _m.ReadUint16(_r, &QuoteOrderListNumInGroup); err != nil {
			return err
		}
		if cap(m.QuoteOrderList) < int(QuoteOrderListNumInGroup) {
			m.QuoteOrderList = make([]MarketSnapshotOrdersPerPriceLevelListQuoteOrderList, QuoteOrderListNumInGroup)
		}
		m.QuoteOrderList = m.QuoteOrderList[:QuoteOrderListNumInGroup]
		for i := range m.QuoteOrderList {
			if err := m.QuoteOrderList[i].Decode(_m, _r, actingVersion, uint(QuoteOrderListBlockLength)); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *MarketSnapshotOrdersPerPriceLevelList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.BidOfferInActingVersion(actingVersion) {
		if m.BidOffer < m.BidOfferMinValue() || m.BidOffer > m.BidOfferMaxValue() {
			return fmt.Errorf("Range check failed on m.BidOffer (%v < %v > %v)", m.BidOfferMinValue(), m.BidOffer, m.BidOfferMaxValue())
		}
	}
	if m.PriceInActingVersion(actingVersion) {
		if m.Price < m.PriceMinValue() || m.Price > m.PriceMaxValue() {
			return fmt.Errorf("Range check failed on m.Price (%v < %v > %v)", m.PriceMinValue(), m.Price, m.PriceMaxValue())
		}
	}
	if m.TotalDealQtyInActingVersion(actingVersion) {
		if m.TotalDealQty < m.TotalDealQtyMinValue() || m.TotalDealQty > m.TotalDealQtyMaxValue() {
			return fmt.Errorf("Range check failed on m.TotalDealQty (%v < %v > %v)", m.TotalDealQtyMinValue(), m.TotalDealQty, m.TotalDealQtyMaxValue())
		}
	}
	if m.TotalMMQtyInActingVersion(actingVersion) {
		if m.TotalMMQty < m.TotalMMQtyMinValue() || m.TotalMMQty > m.TotalMMQtyMaxValue() {
			return fmt.Errorf("Range check failed on m.TotalMMQty (%v < %v > %v)", m.TotalMMQtyMinValue(), m.TotalMMQty, m.TotalMMQtyMaxValue())
		}
	}
	for i := range m.QuoteOrderList {
		if err := m.QuoteOrderList[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MarketSnapshotOrdersPerPriceLevelListInit(m *MarketSnapshotOrdersPerPriceLevelList) {
	return
}

func (m *MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteInt32(_w, m.AcctId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, m.DealQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, m.Price); err != nil {
		return err
	}
	return nil
}

func (m *MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !m.AcctIdInActingVersion(actingVersion) {
		m.AcctId = m.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &m.AcctId); err != nil {
			return err
		}
	}
	if !m.DealQtyInActingVersion(actingVersion) {
		m.DealQty = m.DealQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &m.DealQty); err != nil {
			return err
		}
	}
	if !m.PriceInActingVersion(actingVersion) {
		m.Price = m.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &m.Price); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}
	return nil
}

func (m *MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.AcctIdInActingVersion(actingVersion) {
		if m.AcctId < m.AcctIdMinValue() || m.AcctId > m.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on m.AcctId (%v < %v > %v)", m.AcctIdMinValue(), m.AcctId, m.AcctIdMaxValue())
		}
	}
	if m.DealQtyInActingVersion(actingVersion) {
		if m.DealQty < m.DealQtyMinValue() || m.DealQty > m.DealQtyMaxValue() {
			return fmt.Errorf("Range check failed on m.DealQty (%v < %v > %v)", m.DealQtyMinValue(), m.DealQty, m.DealQtyMaxValue())
		}
	}
	if m.PriceInActingVersion(actingVersion) {
		if m.Price < m.PriceMinValue() || m.Price > m.PriceMaxValue() {
			return fmt.Errorf("Range check failed on m.Price (%v < %v > %v)", m.PriceMinValue(), m.Price, m.PriceMaxValue())
		}
	}
	return nil
}

func MarketSnapshotOrdersPerPriceLevelListQuoteOrderListInit(m *MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) {
	return
}

func (*MarketSnapshot) SbeBlockLength() (blockLength uint16) {
	return 18
}

func (*MarketSnapshot) SbeTemplateId() (templateId uint16) {
	return 196
}

func (*MarketSnapshot) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MarketSnapshot) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MarketSnapshot) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MarketSnapshot) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MarketSnapshot) MarketSequenceId() uint16 {
	return 197
}

func (*MarketSnapshot) MarketSequenceSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshot) MarketSequenceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MarketSequenceSinceVersion()
}

func (*MarketSnapshot) MarketSequenceDeprecated() uint16 {
	return 0
}

func (*MarketSnapshot) MarketSequenceMetaAttribute(meta int) string {
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

func (*MarketSnapshot) MarketSequenceMinValue() uint64 {
	return 0
}

func (*MarketSnapshot) MarketSequenceMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketSnapshot) MarketSequenceNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketSnapshot) InstIdId() uint16 {
	return 198
}

func (*MarketSnapshot) InstIdSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshot) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.InstIdSinceVersion()
}

func (*MarketSnapshot) InstIdDeprecated() uint16 {
	return 0
}

func (*MarketSnapshot) InstIdMetaAttribute(meta int) string {
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

func (*MarketSnapshot) InstIdMinValue() uint16 {
	return 0
}

func (*MarketSnapshot) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MarketSnapshot) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*MarketSnapshot) TimeId() uint16 {
	return 199
}

func (*MarketSnapshot) TimeSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshot) TimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TimeSinceVersion()
}

func (*MarketSnapshot) TimeDeprecated() uint16 {
	return 0
}

func (*MarketSnapshot) TimeMetaAttribute(meta int) string {
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

func (*MarketSnapshot) TimeMinValue() uint64 {
	return 0
}

func (*MarketSnapshot) TimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*MarketSnapshot) TimeNullValue() uint64 {
	return math.MaxUint64
}

func (*MarketSnapshotOrdersPerPriceLevelList) BidOfferId() uint16 {
	return 207
}

func (*MarketSnapshotOrdersPerPriceLevelList) BidOfferSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelList) BidOfferInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidOfferSinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelList) BidOfferDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelList) BidOfferMetaAttribute(meta int) string {
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

func (*MarketSnapshotOrdersPerPriceLevelList) BidOfferMinValue() uint8 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelList) BidOfferMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*MarketSnapshotOrdersPerPriceLevelList) BidOfferNullValue() uint8 {
	return math.MaxUint8
}

func (*MarketSnapshotOrdersPerPriceLevelList) PriceId() uint16 {
	return 208
}

func (*MarketSnapshotOrdersPerPriceLevelList) PriceSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelList) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceSinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelList) PriceDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelList) PriceMetaAttribute(meta int) string {
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

func (*MarketSnapshotOrdersPerPriceLevelList) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelList) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelList) PriceNullValue() float64 {
	return math.NaN()
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalDealQtyId() uint16 {
	return 209
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalDealQtySinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelList) TotalDealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotalDealQtySinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalDealQtyDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalDealQtyMetaAttribute(meta int) string {
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

func (*MarketSnapshotOrdersPerPriceLevelList) TotalDealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalDealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalDealQtyNullValue() float64 {
	return math.NaN()
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalMMQtyId() uint16 {
	return 210
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalMMQtySinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelList) TotalMMQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.TotalMMQtySinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalMMQtyDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalMMQtyMetaAttribute(meta int) string {
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

func (*MarketSnapshotOrdersPerPriceLevelList) TotalMMQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalMMQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelList) TotalMMQtyNullValue() float64 {
	return math.NaN()
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdId() uint16 {
	return 212
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AcctIdSinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdMetaAttribute(meta int) string {
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

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtyId() uint16 {
	return 214
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtySinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DealQtySinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtyDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtyMetaAttribute(meta int) string {
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

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) DealQtyNullValue() float64 {
	return math.NaN()
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceId() uint16 {
	return 215
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceSinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceMetaAttribute(meta int) string {
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

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) PriceNullValue() float64 {
	return math.NaN()
}

func (*MarketSnapshot) OrdersPerPriceLevelListId() uint16 {
	return 206
}

func (*MarketSnapshot) OrdersPerPriceLevelListSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshot) OrdersPerPriceLevelListInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.OrdersPerPriceLevelListSinceVersion()
}

func (*MarketSnapshot) OrdersPerPriceLevelListDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelList) SbeBlockLength() (blockLength uint) {
	return 25
}

func (*MarketSnapshotOrdersPerPriceLevelList) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelList) QuoteOrderListId() uint16 {
	return 211
}

func (*MarketSnapshotOrdersPerPriceLevelList) QuoteOrderListSinceVersion() uint16 {
	return 0
}

func (m *MarketSnapshotOrdersPerPriceLevelList) QuoteOrderListInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.QuoteOrderListSinceVersion()
}

func (*MarketSnapshotOrdersPerPriceLevelList) QuoteOrderListDeprecated() uint16 {
	return 0
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) SbeBlockLength() (blockLength uint) {
	return 20
}

func (*MarketSnapshotOrdersPerPriceLevelListQuoteOrderList) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}
