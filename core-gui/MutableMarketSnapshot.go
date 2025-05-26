// Generated SBE (Simple Binary Encoding) message codec

package core_gui

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type MutableMarketSnapshot struct {
	BidSize               uint16
	MutableQuoteOrderList []MutableMarketSnapshotMutableQuoteOrderList
}
type MutableMarketSnapshotMutableQuoteOrderList struct {
	AcctId  int32
	DealQty float64
	Price   float64
}

func (m *MutableMarketSnapshot) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := m.RangeCheck(m.SbeSchemaVersion(), m.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint16(_w, m.BidSize); err != nil {
		return err
	}
	var MutableQuoteOrderListBlockLength uint16 = 20
	if err := _m.WriteUint16(_w, MutableQuoteOrderListBlockLength); err != nil {
		return err
	}
	var MutableQuoteOrderListNumInGroup uint16 = uint16(len(m.MutableQuoteOrderList))
	if err := _m.WriteUint16(_w, MutableQuoteOrderListNumInGroup); err != nil {
		return err
	}
	for i := range m.MutableQuoteOrderList {
		if err := m.MutableQuoteOrderList[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (m *MutableMarketSnapshot) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !m.BidSizeInActingVersion(actingVersion) {
		m.BidSize = m.BidSizeNullValue()
	} else {
		if err := _m.ReadUint16(_r, &m.BidSize); err != nil {
			return err
		}
	}
	if actingVersion > m.SbeSchemaVersion() && blockLength > m.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-m.SbeBlockLength()))
	}

	if m.MutableQuoteOrderListInActingVersion(actingVersion) {
		var MutableQuoteOrderListBlockLength uint16
		if err := _m.ReadUint16(_r, &MutableQuoteOrderListBlockLength); err != nil {
			return err
		}
		var MutableQuoteOrderListNumInGroup uint16
		if err := _m.ReadUint16(_r, &MutableQuoteOrderListNumInGroup); err != nil {
			return err
		}
		if cap(m.MutableQuoteOrderList) < int(MutableQuoteOrderListNumInGroup) {
			m.MutableQuoteOrderList = make([]MutableMarketSnapshotMutableQuoteOrderList, MutableQuoteOrderListNumInGroup)
		}
		m.MutableQuoteOrderList = m.MutableQuoteOrderList[:MutableQuoteOrderListNumInGroup]
		for i := range m.MutableQuoteOrderList {
			if err := m.MutableQuoteOrderList[i].Decode(_m, _r, actingVersion, uint(MutableQuoteOrderListBlockLength)); err != nil {
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

func (m *MutableMarketSnapshot) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if m.BidSizeInActingVersion(actingVersion) {
		if m.BidSize < m.BidSizeMinValue() || m.BidSize > m.BidSizeMaxValue() {
			return fmt.Errorf("Range check failed on m.BidSize (%v < %v > %v)", m.BidSizeMinValue(), m.BidSize, m.BidSizeMaxValue())
		}
	}
	for i := range m.MutableQuoteOrderList {
		if err := m.MutableQuoteOrderList[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func MutableMarketSnapshotInit(m *MutableMarketSnapshot) {
	return
}

func (m *MutableMarketSnapshotMutableQuoteOrderList) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
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

func (m *MutableMarketSnapshotMutableQuoteOrderList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
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

func (m *MutableMarketSnapshotMutableQuoteOrderList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
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

func MutableMarketSnapshotMutableQuoteOrderListInit(m *MutableMarketSnapshotMutableQuoteOrderList) {
	return
}

func (*MutableMarketSnapshot) SbeBlockLength() (blockLength uint16) {
	return 2
}

func (*MutableMarketSnapshot) SbeTemplateId() (templateId uint16) {
	return 230
}

func (*MutableMarketSnapshot) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*MutableMarketSnapshot) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*MutableMarketSnapshot) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*MutableMarketSnapshot) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*MutableMarketSnapshot) BidSizeId() uint16 {
	return 234
}

func (*MutableMarketSnapshot) BidSizeSinceVersion() uint16 {
	return 0
}

func (m *MutableMarketSnapshot) BidSizeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.BidSizeSinceVersion()
}

func (*MutableMarketSnapshot) BidSizeDeprecated() uint16 {
	return 0
}

func (*MutableMarketSnapshot) BidSizeMetaAttribute(meta int) string {
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

func (*MutableMarketSnapshot) BidSizeMinValue() uint16 {
	return 0
}

func (*MutableMarketSnapshot) BidSizeMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*MutableMarketSnapshot) BidSizeNullValue() uint16 {
	return math.MaxUint16
}

func (*MutableMarketSnapshotMutableQuoteOrderList) AcctIdId() uint16 {
	return 237
}

func (*MutableMarketSnapshotMutableQuoteOrderList) AcctIdSinceVersion() uint16 {
	return 0
}

func (m *MutableMarketSnapshotMutableQuoteOrderList) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.AcctIdSinceVersion()
}

func (*MutableMarketSnapshotMutableQuoteOrderList) AcctIdDeprecated() uint16 {
	return 0
}

func (*MutableMarketSnapshotMutableQuoteOrderList) AcctIdMetaAttribute(meta int) string {
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

func (*MutableMarketSnapshotMutableQuoteOrderList) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*MutableMarketSnapshotMutableQuoteOrderList) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*MutableMarketSnapshotMutableQuoteOrderList) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*MutableMarketSnapshotMutableQuoteOrderList) DealQtyId() uint16 {
	return 239
}

func (*MutableMarketSnapshotMutableQuoteOrderList) DealQtySinceVersion() uint16 {
	return 0
}

func (m *MutableMarketSnapshotMutableQuoteOrderList) DealQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.DealQtySinceVersion()
}

func (*MutableMarketSnapshotMutableQuoteOrderList) DealQtyDeprecated() uint16 {
	return 0
}

func (*MutableMarketSnapshotMutableQuoteOrderList) DealQtyMetaAttribute(meta int) string {
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

func (*MutableMarketSnapshotMutableQuoteOrderList) DealQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*MutableMarketSnapshotMutableQuoteOrderList) DealQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*MutableMarketSnapshotMutableQuoteOrderList) DealQtyNullValue() float64 {
	return math.NaN()
}

func (*MutableMarketSnapshotMutableQuoteOrderList) PriceId() uint16 {
	return 240
}

func (*MutableMarketSnapshotMutableQuoteOrderList) PriceSinceVersion() uint16 {
	return 0
}

func (m *MutableMarketSnapshotMutableQuoteOrderList) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.PriceSinceVersion()
}

func (*MutableMarketSnapshotMutableQuoteOrderList) PriceDeprecated() uint16 {
	return 0
}

func (*MutableMarketSnapshotMutableQuoteOrderList) PriceMetaAttribute(meta int) string {
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

func (*MutableMarketSnapshotMutableQuoteOrderList) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*MutableMarketSnapshotMutableQuoteOrderList) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*MutableMarketSnapshotMutableQuoteOrderList) PriceNullValue() float64 {
	return math.NaN()
}

func (*MutableMarketSnapshot) MutableQuoteOrderListId() uint16 {
	return 236
}

func (*MutableMarketSnapshot) MutableQuoteOrderListSinceVersion() uint16 {
	return 0
}

func (m *MutableMarketSnapshot) MutableQuoteOrderListInActingVersion(actingVersion uint16) bool {
	return actingVersion >= m.MutableQuoteOrderListSinceVersion()
}

func (*MutableMarketSnapshot) MutableQuoteOrderListDeprecated() uint16 {
	return 0
}

func (*MutableMarketSnapshotMutableQuoteOrderList) SbeBlockLength() (blockLength uint) {
	return 20
}

func (*MutableMarketSnapshotMutableQuoteOrderList) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}
