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

type OrderCancel struct {
	OrderCancelType uint8
	AcctId          int32
	InstId          uint16
	Respond         uint8
	ProxyIndex      uint16
	CreatedBy       int32
	OrderId         []uint8
	CancelId        []uint8
}

func (o *OrderCancel) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := o.RangeCheck(o.SbeSchemaVersion(), o.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, o.OrderCancelType); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.InstId); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, o.Respond); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, o.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, o.CreatedBy); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.OrderId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.OrderId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(o.CancelId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, o.CancelId); err != nil {
		return err
	}
	return nil
}

func (o *OrderCancel) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !o.OrderCancelTypeInActingVersion(actingVersion) {
		o.OrderCancelType = o.OrderCancelTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.OrderCancelType); err != nil {
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
	if !o.InstIdInActingVersion(actingVersion) {
		o.InstId = o.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &o.InstId); err != nil {
			return err
		}
	}
	if !o.RespondInActingVersion(actingVersion) {
		o.Respond = o.RespondNullValue()
	} else {
		if err := _m.ReadUint8(_r, &o.Respond); err != nil {
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
	if !o.CreatedByInActingVersion(actingVersion) {
		o.CreatedBy = o.CreatedByNullValue()
	} else {
		if err := _m.ReadInt32(_r, &o.CreatedBy); err != nil {
			return err
		}
	}
	if actingVersion > o.SbeSchemaVersion() && blockLength > o.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-o.SbeBlockLength()))
	}

	if o.OrderIdInActingVersion(actingVersion) {
		var OrderIdLength uint16
		if err := _m.ReadUint16(_r, &OrderIdLength); err != nil {
			return err
		}
		if cap(o.OrderId) < int(OrderIdLength) {
			o.OrderId = make([]uint8, OrderIdLength)
		}
		o.OrderId = o.OrderId[:OrderIdLength]
		if err := _m.ReadBytes(_r, o.OrderId); err != nil {
			return err
		}
	}

	if o.CancelIdInActingVersion(actingVersion) {
		var CancelIdLength uint16
		if err := _m.ReadUint16(_r, &CancelIdLength); err != nil {
			return err
		}
		if cap(o.CancelId) < int(CancelIdLength) {
			o.CancelId = make([]uint8, CancelIdLength)
		}
		o.CancelId = o.CancelId[:CancelIdLength]
		if err := _m.ReadBytes(_r, o.CancelId); err != nil {
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

func (o *OrderCancel) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if o.OrderCancelTypeInActingVersion(actingVersion) {
		if o.OrderCancelType < o.OrderCancelTypeMinValue() || o.OrderCancelType > o.OrderCancelTypeMaxValue() {
			return fmt.Errorf("Range check failed on o.OrderCancelType (%v < %v > %v)", o.OrderCancelTypeMinValue(), o.OrderCancelType, o.OrderCancelTypeMaxValue())
		}
	}
	if o.AcctIdInActingVersion(actingVersion) {
		if o.AcctId < o.AcctIdMinValue() || o.AcctId > o.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on o.AcctId (%v < %v > %v)", o.AcctIdMinValue(), o.AcctId, o.AcctIdMaxValue())
		}
	}
	if o.InstIdInActingVersion(actingVersion) {
		if o.InstId < o.InstIdMinValue() || o.InstId > o.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on o.InstId (%v < %v > %v)", o.InstIdMinValue(), o.InstId, o.InstIdMaxValue())
		}
	}
	if o.RespondInActingVersion(actingVersion) {
		if o.Respond < o.RespondMinValue() || o.Respond > o.RespondMaxValue() {
			return fmt.Errorf("Range check failed on o.Respond (%v < %v > %v)", o.RespondMinValue(), o.Respond, o.RespondMaxValue())
		}
	}
	if o.ProxyIndexInActingVersion(actingVersion) {
		if o.ProxyIndex < o.ProxyIndexMinValue() || o.ProxyIndex > o.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on o.ProxyIndex (%v < %v > %v)", o.ProxyIndexMinValue(), o.ProxyIndex, o.ProxyIndexMaxValue())
		}
	}
	if o.CreatedByInActingVersion(actingVersion) {
		if o.CreatedBy < o.CreatedByMinValue() || o.CreatedBy > o.CreatedByMaxValue() {
			return fmt.Errorf("Range check failed on o.CreatedBy (%v < %v > %v)", o.CreatedByMinValue(), o.CreatedBy, o.CreatedByMaxValue())
		}
	}
	if !utf8.Valid(o.OrderId[:]) {
		return errors.New("o.OrderId failed UTF-8 validation")
	}
	if !utf8.Valid(o.CancelId[:]) {
		return errors.New("o.CancelId failed UTF-8 validation")
	}
	return nil
}

func OrderCancelInit(o *OrderCancel) {
	return
}

func (*OrderCancel) SbeBlockLength() (blockLength uint16) {
	return 14
}

func (*OrderCancel) SbeTemplateId() (templateId uint16) {
	return 165
}

func (*OrderCancel) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*OrderCancel) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*OrderCancel) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*OrderCancel) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*OrderCancel) OrderCancelTypeId() uint16 {
	return 166
}

func (*OrderCancel) OrderCancelTypeSinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) OrderCancelTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderCancelTypeSinceVersion()
}

func (*OrderCancel) OrderCancelTypeDeprecated() uint16 {
	return 0
}

func (*OrderCancel) OrderCancelTypeMetaAttribute(meta int) string {
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

func (*OrderCancel) OrderCancelTypeMinValue() uint8 {
	return 0
}

func (*OrderCancel) OrderCancelTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderCancel) OrderCancelTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderCancel) AcctIdId() uint16 {
	return 167
}

func (*OrderCancel) AcctIdSinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.AcctIdSinceVersion()
}

func (*OrderCancel) AcctIdDeprecated() uint16 {
	return 0
}

func (*OrderCancel) AcctIdMetaAttribute(meta int) string {
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

func (*OrderCancel) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderCancel) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderCancel) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*OrderCancel) InstIdId() uint16 {
	return 168
}

func (*OrderCancel) InstIdSinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.InstIdSinceVersion()
}

func (*OrderCancel) InstIdDeprecated() uint16 {
	return 0
}

func (*OrderCancel) InstIdMetaAttribute(meta int) string {
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

func (*OrderCancel) InstIdMinValue() uint16 {
	return 0
}

func (*OrderCancel) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancel) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderCancel) RespondId() uint16 {
	return 169
}

func (*OrderCancel) RespondSinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) RespondInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.RespondSinceVersion()
}

func (*OrderCancel) RespondDeprecated() uint16 {
	return 0
}

func (*OrderCancel) RespondMetaAttribute(meta int) string {
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

func (*OrderCancel) RespondMinValue() uint8 {
	return 0
}

func (*OrderCancel) RespondMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*OrderCancel) RespondNullValue() uint8 {
	return math.MaxUint8
}

func (*OrderCancel) ProxyIndexId() uint16 {
	return 170
}

func (*OrderCancel) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.ProxyIndexSinceVersion()
}

func (*OrderCancel) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*OrderCancel) ProxyIndexMetaAttribute(meta int) string {
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

func (*OrderCancel) ProxyIndexMinValue() uint16 {
	return 0
}

func (*OrderCancel) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*OrderCancel) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*OrderCancel) CreatedById() uint16 {
	return 171
}

func (*OrderCancel) CreatedBySinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) CreatedByInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CreatedBySinceVersion()
}

func (*OrderCancel) CreatedByDeprecated() uint16 {
	return 0
}

func (*OrderCancel) CreatedByMetaAttribute(meta int) string {
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

func (*OrderCancel) CreatedByMinValue() int32 {
	return math.MinInt32 + 1
}

func (*OrderCancel) CreatedByMaxValue() int32 {
	return math.MaxInt32
}

func (*OrderCancel) CreatedByNullValue() int32 {
	return math.MinInt32
}

func (*OrderCancel) OrderIdMetaAttribute(meta int) string {
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

func (*OrderCancel) OrderIdSinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) OrderIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.OrderIdSinceVersion()
}

func (*OrderCancel) OrderIdDeprecated() uint16 {
	return 0
}

func (OrderCancel) OrderIdCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancel) OrderIdHeaderLength() uint64 {
	return 2
}

func (*OrderCancel) CancelIdMetaAttribute(meta int) string {
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

func (*OrderCancel) CancelIdSinceVersion() uint16 {
	return 0
}

func (o *OrderCancel) CancelIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= o.CancelIdSinceVersion()
}

func (*OrderCancel) CancelIdDeprecated() uint16 {
	return 0
}

func (OrderCancel) CancelIdCharacterEncoding() string {
	return "UTF-8"
}

func (OrderCancel) CancelIdHeaderLength() uint64 {
	return 2
}
