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

type PriceSetBuffer struct {
	AcctId        int32
	ProxyIndex    uint16
	CreateTime    uint64
	InstId        uint16
	ValueDate     int32
	BidChanged    uint8
	OfferChanged  uint8
	ExpireTime    uint64
	PricePropList []PriceSetBufferPricePropList
}
type PriceSetBufferPricePropList struct {
	OrderQty  float64
	Price     float64
	RequestId []uint8
}

func (p *PriceSetBuffer) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, p.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, p.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.CreateTime); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, p.InstId); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, p.ValueDate); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.BidChanged); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, p.OfferChanged); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.ExpireTime); err != nil {
		return err
	}
	var PricePropListBlockLength uint16 = 16
	if err := _m.WriteUint16(_w, PricePropListBlockLength); err != nil {
		return err
	}
	var PricePropListNumInGroup uint16 = uint16(len(p.PricePropList))
	if err := _m.WriteUint16(_w, PricePropListNumInGroup); err != nil {
		return err
	}
	for i := range p.PricePropList {
		if err := p.PricePropList[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (p *PriceSetBuffer) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !p.AcctIdInActingVersion(actingVersion) {
		p.AcctId = p.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &p.AcctId); err != nil {
			return err
		}
	}
	if !p.ProxyIndexInActingVersion(actingVersion) {
		p.ProxyIndex = p.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &p.ProxyIndex); err != nil {
			return err
		}
	}
	if !p.CreateTimeInActingVersion(actingVersion) {
		p.CreateTime = p.CreateTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.CreateTime); err != nil {
			return err
		}
	}
	if !p.InstIdInActingVersion(actingVersion) {
		p.InstId = p.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &p.InstId); err != nil {
			return err
		}
	}
	if !p.ValueDateInActingVersion(actingVersion) {
		p.ValueDate = p.ValueDateNullValue()
	} else {
		if err := _m.ReadInt32(_r, &p.ValueDate); err != nil {
			return err
		}
	}
	if !p.BidChangedInActingVersion(actingVersion) {
		p.BidChanged = p.BidChangedNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.BidChanged); err != nil {
			return err
		}
	}
	if !p.OfferChangedInActingVersion(actingVersion) {
		p.OfferChanged = p.OfferChangedNullValue()
	} else {
		if err := _m.ReadUint8(_r, &p.OfferChanged); err != nil {
			return err
		}
	}
	if !p.ExpireTimeInActingVersion(actingVersion) {
		p.ExpireTime = p.ExpireTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.ExpireTime); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}

	if p.PricePropListInActingVersion(actingVersion) {
		var PricePropListBlockLength uint16
		if err := _m.ReadUint16(_r, &PricePropListBlockLength); err != nil {
			return err
		}
		var PricePropListNumInGroup uint16
		if err := _m.ReadUint16(_r, &PricePropListNumInGroup); err != nil {
			return err
		}
		if cap(p.PricePropList) < int(PricePropListNumInGroup) {
			p.PricePropList = make([]PriceSetBufferPricePropList, PricePropListNumInGroup)
		}
		p.PricePropList = p.PricePropList[:PricePropListNumInGroup]
		for i := range p.PricePropList {
			if err := p.PricePropList[i].Decode(_m, _r, actingVersion, uint(PricePropListBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := p.RangeCheck(actingVersion, p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (p *PriceSetBuffer) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.AcctIdInActingVersion(actingVersion) {
		if p.AcctId < p.AcctIdMinValue() || p.AcctId > p.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on p.AcctId (%v < %v > %v)", p.AcctIdMinValue(), p.AcctId, p.AcctIdMaxValue())
		}
	}
	if p.ProxyIndexInActingVersion(actingVersion) {
		if p.ProxyIndex < p.ProxyIndexMinValue() || p.ProxyIndex > p.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on p.ProxyIndex (%v < %v > %v)", p.ProxyIndexMinValue(), p.ProxyIndex, p.ProxyIndexMaxValue())
		}
	}
	if p.CreateTimeInActingVersion(actingVersion) {
		if p.CreateTime < p.CreateTimeMinValue() || p.CreateTime > p.CreateTimeMaxValue() {
			return fmt.Errorf("Range check failed on p.CreateTime (%v < %v > %v)", p.CreateTimeMinValue(), p.CreateTime, p.CreateTimeMaxValue())
		}
	}
	if p.InstIdInActingVersion(actingVersion) {
		if p.InstId < p.InstIdMinValue() || p.InstId > p.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on p.InstId (%v < %v > %v)", p.InstIdMinValue(), p.InstId, p.InstIdMaxValue())
		}
	}
	if p.ValueDateInActingVersion(actingVersion) {
		if p.ValueDate < p.ValueDateMinValue() || p.ValueDate > p.ValueDateMaxValue() {
			return fmt.Errorf("Range check failed on p.ValueDate (%v < %v > %v)", p.ValueDateMinValue(), p.ValueDate, p.ValueDateMaxValue())
		}
	}
	if p.BidChangedInActingVersion(actingVersion) {
		if p.BidChanged < p.BidChangedMinValue() || p.BidChanged > p.BidChangedMaxValue() {
			return fmt.Errorf("Range check failed on p.BidChanged (%v < %v > %v)", p.BidChangedMinValue(), p.BidChanged, p.BidChangedMaxValue())
		}
	}
	if p.OfferChangedInActingVersion(actingVersion) {
		if p.OfferChanged < p.OfferChangedMinValue() || p.OfferChanged > p.OfferChangedMaxValue() {
			return fmt.Errorf("Range check failed on p.OfferChanged (%v < %v > %v)", p.OfferChangedMinValue(), p.OfferChanged, p.OfferChangedMaxValue())
		}
	}
	if p.ExpireTimeInActingVersion(actingVersion) {
		if p.ExpireTime < p.ExpireTimeMinValue() || p.ExpireTime > p.ExpireTimeMaxValue() {
			return fmt.Errorf("Range check failed on p.ExpireTime (%v < %v > %v)", p.ExpireTimeMinValue(), p.ExpireTime, p.ExpireTimeMaxValue())
		}
	}
	for i := range p.PricePropList {
		if err := p.PricePropList[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func PriceSetBufferInit(p *PriceSetBuffer) {
	return
}

func (p *PriceSetBufferPricePropList) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteFloat64(_w, p.OrderQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, p.Price); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(p.RequestId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.RequestId); err != nil {
		return err
	}
	return nil
}

func (p *PriceSetBufferPricePropList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if !p.OrderQtyInActingVersion(actingVersion) {
		p.OrderQty = p.OrderQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &p.OrderQty); err != nil {
			return err
		}
	}
	if !p.PriceInActingVersion(actingVersion) {
		p.Price = p.PriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &p.Price); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}

	if p.RequestIdInActingVersion(actingVersion) {
		var RequestIdLength uint16
		if err := _m.ReadUint16(_r, &RequestIdLength); err != nil {
			return err
		}
		if cap(p.RequestId) < int(RequestIdLength) {
			p.RequestId = make([]uint8, RequestIdLength)
		}
		p.RequestId = p.RequestId[:RequestIdLength]
		if err := _m.ReadBytes(_r, p.RequestId); err != nil {
			return err
		}
	}
	return nil
}

func (p *PriceSetBufferPricePropList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.OrderQtyInActingVersion(actingVersion) {
		if p.OrderQty < p.OrderQtyMinValue() || p.OrderQty > p.OrderQtyMaxValue() {
			return fmt.Errorf("Range check failed on p.OrderQty (%v < %v > %v)", p.OrderQtyMinValue(), p.OrderQty, p.OrderQtyMaxValue())
		}
	}
	if p.PriceInActingVersion(actingVersion) {
		if p.Price < p.PriceMinValue() || p.Price > p.PriceMaxValue() {
			return fmt.Errorf("Range check failed on p.Price (%v < %v > %v)", p.PriceMinValue(), p.Price, p.PriceMaxValue())
		}
	}
	if !utf8.Valid(p.RequestId[:]) {
		return errors.New("p.RequestId failed UTF-8 validation")
	}
	return nil
}

func PriceSetBufferPricePropListInit(p *PriceSetBufferPricePropList) {
	return
}

func (*PriceSetBuffer) SbeBlockLength() (blockLength uint16) {
	return 30
}

func (*PriceSetBuffer) SbeTemplateId() (templateId uint16) {
	return 250
}

func (*PriceSetBuffer) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*PriceSetBuffer) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*PriceSetBuffer) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*PriceSetBuffer) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*PriceSetBuffer) AcctIdId() uint16 {
	return 251
}

func (*PriceSetBuffer) AcctIdSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.AcctIdSinceVersion()
}

func (*PriceSetBuffer) AcctIdDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) AcctIdMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*PriceSetBuffer) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*PriceSetBuffer) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*PriceSetBuffer) ProxyIndexId() uint16 {
	return 252
}

func (*PriceSetBuffer) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ProxyIndexSinceVersion()
}

func (*PriceSetBuffer) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) ProxyIndexMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) ProxyIndexMinValue() uint16 {
	return 0
}

func (*PriceSetBuffer) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*PriceSetBuffer) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*PriceSetBuffer) CreateTimeId() uint16 {
	return 253
}

func (*PriceSetBuffer) CreateTimeSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) CreateTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.CreateTimeSinceVersion()
}

func (*PriceSetBuffer) CreateTimeDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) CreateTimeMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) CreateTimeMinValue() uint64 {
	return 0
}

func (*PriceSetBuffer) CreateTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PriceSetBuffer) CreateTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*PriceSetBuffer) InstIdId() uint16 {
	return 254
}

func (*PriceSetBuffer) InstIdSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.InstIdSinceVersion()
}

func (*PriceSetBuffer) InstIdDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) InstIdMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) InstIdMinValue() uint16 {
	return 0
}

func (*PriceSetBuffer) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*PriceSetBuffer) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*PriceSetBuffer) ValueDateId() uint16 {
	return 255
}

func (*PriceSetBuffer) ValueDateSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) ValueDateInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ValueDateSinceVersion()
}

func (*PriceSetBuffer) ValueDateDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) ValueDateMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) ValueDateMinValue() int32 {
	return math.MinInt32 + 1
}

func (*PriceSetBuffer) ValueDateMaxValue() int32 {
	return math.MaxInt32
}

func (*PriceSetBuffer) ValueDateNullValue() int32 {
	return math.MinInt32
}

func (*PriceSetBuffer) BidChangedId() uint16 {
	return 256
}

func (*PriceSetBuffer) BidChangedSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) BidChangedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.BidChangedSinceVersion()
}

func (*PriceSetBuffer) BidChangedDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) BidChangedMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) BidChangedMinValue() uint8 {
	return 0
}

func (*PriceSetBuffer) BidChangedMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PriceSetBuffer) BidChangedNullValue() uint8 {
	return math.MaxUint8
}

func (*PriceSetBuffer) OfferChangedId() uint16 {
	return 257
}

func (*PriceSetBuffer) OfferChangedSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) OfferChangedInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.OfferChangedSinceVersion()
}

func (*PriceSetBuffer) OfferChangedDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) OfferChangedMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) OfferChangedMinValue() uint8 {
	return 0
}

func (*PriceSetBuffer) OfferChangedMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*PriceSetBuffer) OfferChangedNullValue() uint8 {
	return math.MaxUint8
}

func (*PriceSetBuffer) ExpireTimeId() uint16 {
	return 258
}

func (*PriceSetBuffer) ExpireTimeSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) ExpireTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ExpireTimeSinceVersion()
}

func (*PriceSetBuffer) ExpireTimeDeprecated() uint16 {
	return 0
}

func (*PriceSetBuffer) ExpireTimeMetaAttribute(meta int) string {
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

func (*PriceSetBuffer) ExpireTimeMinValue() uint64 {
	return 0
}

func (*PriceSetBuffer) ExpireTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*PriceSetBuffer) ExpireTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*PriceSetBufferPricePropList) OrderQtyId() uint16 {
	return 262
}

func (*PriceSetBufferPricePropList) OrderQtySinceVersion() uint16 {
	return 0
}

func (p *PriceSetBufferPricePropList) OrderQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.OrderQtySinceVersion()
}

func (*PriceSetBufferPricePropList) OrderQtyDeprecated() uint16 {
	return 0
}

func (*PriceSetBufferPricePropList) OrderQtyMetaAttribute(meta int) string {
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

func (*PriceSetBufferPricePropList) OrderQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*PriceSetBufferPricePropList) OrderQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*PriceSetBufferPricePropList) OrderQtyNullValue() float64 {
	return math.NaN()
}

func (*PriceSetBufferPricePropList) PriceId() uint16 {
	return 263
}

func (*PriceSetBufferPricePropList) PriceSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBufferPricePropList) PriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PriceSinceVersion()
}

func (*PriceSetBufferPricePropList) PriceDeprecated() uint16 {
	return 0
}

func (*PriceSetBufferPricePropList) PriceMetaAttribute(meta int) string {
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

func (*PriceSetBufferPricePropList) PriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*PriceSetBufferPricePropList) PriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*PriceSetBufferPricePropList) PriceNullValue() float64 {
	return math.NaN()
}

func (*PriceSetBufferPricePropList) RequestIdMetaAttribute(meta int) string {
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

func (*PriceSetBufferPricePropList) RequestIdSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBufferPricePropList) RequestIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.RequestIdSinceVersion()
}

func (*PriceSetBufferPricePropList) RequestIdDeprecated() uint16 {
	return 0
}

func (PriceSetBufferPricePropList) RequestIdCharacterEncoding() string {
	return "UTF-8"
}

func (PriceSetBufferPricePropList) RequestIdHeaderLength() uint64 {
	return 2
}

func (*PriceSetBuffer) PricePropListId() uint16 {
	return 259
}

func (*PriceSetBuffer) PricePropListSinceVersion() uint16 {
	return 0
}

func (p *PriceSetBuffer) PricePropListInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.PricePropListSinceVersion()
}

func (*PriceSetBuffer) PricePropListDeprecated() uint16 {
	return 0
}

func (*PriceSetBufferPricePropList) SbeBlockLength() (blockLength uint) {
	return 16
}

func (*PriceSetBufferPricePropList) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}
