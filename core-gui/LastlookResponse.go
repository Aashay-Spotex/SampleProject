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

type LastlookResponse struct {
	AcctId     int32
	InstId     uint16
	ExecQty    float64
	ExecPrice  float64
	Accept     uint8
	ProxyIndex uint16
	Complete   uint8
	ClExecId   []uint8
	TradeId    []uint8
	ExecCurr   []uint8
	Message    []uint8
}

func (l *LastlookResponse) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := l.RangeCheck(l.SbeSchemaVersion(), l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, l.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, l.InstId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, l.ExecQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, l.ExecPrice); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, l.Accept); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, l.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, l.Complete); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.ClExecId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.ClExecId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.TradeId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.TradeId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.ExecCurr))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.ExecCurr); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(l.Message))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, l.Message); err != nil {
		return err
	}
	return nil
}

func (l *LastlookResponse) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !l.AcctIdInActingVersion(actingVersion) {
		l.AcctId = l.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &l.AcctId); err != nil {
			return err
		}
	}
	if !l.InstIdInActingVersion(actingVersion) {
		l.InstId = l.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &l.InstId); err != nil {
			return err
		}
	}
	if !l.ExecQtyInActingVersion(actingVersion) {
		l.ExecQty = l.ExecQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &l.ExecQty); err != nil {
			return err
		}
	}
	if !l.ExecPriceInActingVersion(actingVersion) {
		l.ExecPrice = l.ExecPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &l.ExecPrice); err != nil {
			return err
		}
	}
	if !l.AcceptInActingVersion(actingVersion) {
		l.Accept = l.AcceptNullValue()
	} else {
		if err := _m.ReadUint8(_r, &l.Accept); err != nil {
			return err
		}
	}
	if !l.ProxyIndexInActingVersion(actingVersion) {
		l.ProxyIndex = l.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &l.ProxyIndex); err != nil {
			return err
		}
	}
	if !l.CompleteInActingVersion(actingVersion) {
		l.Complete = l.CompleteNullValue()
	} else {
		if err := _m.ReadUint8(_r, &l.Complete); err != nil {
			return err
		}
	}
	if actingVersion > l.SbeSchemaVersion() && blockLength > l.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-l.SbeBlockLength()))
	}

	if l.ClExecIdInActingVersion(actingVersion) {
		var ClExecIdLength uint16
		if err := _m.ReadUint16(_r, &ClExecIdLength); err != nil {
			return err
		}
		if cap(l.ClExecId) < int(ClExecIdLength) {
			l.ClExecId = make([]uint8, ClExecIdLength)
		}
		l.ClExecId = l.ClExecId[:ClExecIdLength]
		if err := _m.ReadBytes(_r, l.ClExecId); err != nil {
			return err
		}
	}

	if l.TradeIdInActingVersion(actingVersion) {
		var TradeIdLength uint16
		if err := _m.ReadUint16(_r, &TradeIdLength); err != nil {
			return err
		}
		if cap(l.TradeId) < int(TradeIdLength) {
			l.TradeId = make([]uint8, TradeIdLength)
		}
		l.TradeId = l.TradeId[:TradeIdLength]
		if err := _m.ReadBytes(_r, l.TradeId); err != nil {
			return err
		}
	}

	if l.ExecCurrInActingVersion(actingVersion) {
		var ExecCurrLength uint16
		if err := _m.ReadUint16(_r, &ExecCurrLength); err != nil {
			return err
		}
		if cap(l.ExecCurr) < int(ExecCurrLength) {
			l.ExecCurr = make([]uint8, ExecCurrLength)
		}
		l.ExecCurr = l.ExecCurr[:ExecCurrLength]
		if err := _m.ReadBytes(_r, l.ExecCurr); err != nil {
			return err
		}
	}

	if l.MessageInActingVersion(actingVersion) {
		var MessageLength uint16
		if err := _m.ReadUint16(_r, &MessageLength); err != nil {
			return err
		}
		if cap(l.Message) < int(MessageLength) {
			l.Message = make([]uint8, MessageLength)
		}
		l.Message = l.Message[:MessageLength]
		if err := _m.ReadBytes(_r, l.Message); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := l.RangeCheck(actingVersion, l.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (l *LastlookResponse) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if l.AcctIdInActingVersion(actingVersion) {
		if l.AcctId < l.AcctIdMinValue() || l.AcctId > l.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on l.AcctId (%v < %v > %v)", l.AcctIdMinValue(), l.AcctId, l.AcctIdMaxValue())
		}
	}
	if l.InstIdInActingVersion(actingVersion) {
		if l.InstId < l.InstIdMinValue() || l.InstId > l.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on l.InstId (%v < %v > %v)", l.InstIdMinValue(), l.InstId, l.InstIdMaxValue())
		}
	}
	if l.ExecQtyInActingVersion(actingVersion) {
		if l.ExecQty < l.ExecQtyMinValue() || l.ExecQty > l.ExecQtyMaxValue() {
			return fmt.Errorf("Range check failed on l.ExecQty (%v < %v > %v)", l.ExecQtyMinValue(), l.ExecQty, l.ExecQtyMaxValue())
		}
	}
	if l.ExecPriceInActingVersion(actingVersion) {
		if l.ExecPrice < l.ExecPriceMinValue() || l.ExecPrice > l.ExecPriceMaxValue() {
			return fmt.Errorf("Range check failed on l.ExecPrice (%v < %v > %v)", l.ExecPriceMinValue(), l.ExecPrice, l.ExecPriceMaxValue())
		}
	}
	if l.AcceptInActingVersion(actingVersion) {
		if l.Accept < l.AcceptMinValue() || l.Accept > l.AcceptMaxValue() {
			return fmt.Errorf("Range check failed on l.Accept (%v < %v > %v)", l.AcceptMinValue(), l.Accept, l.AcceptMaxValue())
		}
	}
	if l.ProxyIndexInActingVersion(actingVersion) {
		if l.ProxyIndex < l.ProxyIndexMinValue() || l.ProxyIndex > l.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on l.ProxyIndex (%v < %v > %v)", l.ProxyIndexMinValue(), l.ProxyIndex, l.ProxyIndexMaxValue())
		}
	}
	if l.CompleteInActingVersion(actingVersion) {
		if l.Complete < l.CompleteMinValue() || l.Complete > l.CompleteMaxValue() {
			return fmt.Errorf("Range check failed on l.Complete (%v < %v > %v)", l.CompleteMinValue(), l.Complete, l.CompleteMaxValue())
		}
	}
	if !utf8.Valid(l.ClExecId[:]) {
		return errors.New("l.ClExecId failed UTF-8 validation")
	}
	if !utf8.Valid(l.TradeId[:]) {
		return errors.New("l.TradeId failed UTF-8 validation")
	}
	if !utf8.Valid(l.ExecCurr[:]) {
		return errors.New("l.ExecCurr failed UTF-8 validation")
	}
	if !utf8.Valid(l.Message[:]) {
		return errors.New("l.Message failed UTF-8 validation")
	}
	return nil
}

func LastlookResponseInit(l *LastlookResponse) {
	return
}

func (*LastlookResponse) SbeBlockLength() (blockLength uint16) {
	return 26
}

func (*LastlookResponse) SbeTemplateId() (templateId uint16) {
	return 59
}

func (*LastlookResponse) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*LastlookResponse) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*LastlookResponse) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*LastlookResponse) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*LastlookResponse) AcctIdId() uint16 {
	return 60
}

func (*LastlookResponse) AcctIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.AcctIdSinceVersion()
}

func (*LastlookResponse) AcctIdDeprecated() uint16 {
	return 0
}

func (*LastlookResponse) AcctIdMetaAttribute(meta int) string {
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

func (*LastlookResponse) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*LastlookResponse) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*LastlookResponse) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*LastlookResponse) InstIdId() uint16 {
	return 61
}

func (*LastlookResponse) InstIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.InstIdSinceVersion()
}

func (*LastlookResponse) InstIdDeprecated() uint16 {
	return 0
}

func (*LastlookResponse) InstIdMetaAttribute(meta int) string {
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

func (*LastlookResponse) InstIdMinValue() uint16 {
	return 0
}

func (*LastlookResponse) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*LastlookResponse) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*LastlookResponse) ExecQtyId() uint16 {
	return 62
}

func (*LastlookResponse) ExecQtySinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) ExecQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ExecQtySinceVersion()
}

func (*LastlookResponse) ExecQtyDeprecated() uint16 {
	return 0
}

func (*LastlookResponse) ExecQtyMetaAttribute(meta int) string {
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

func (*LastlookResponse) ExecQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*LastlookResponse) ExecQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*LastlookResponse) ExecQtyNullValue() float64 {
	return math.NaN()
}

func (*LastlookResponse) ExecPriceId() uint16 {
	return 63
}

func (*LastlookResponse) ExecPriceSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) ExecPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ExecPriceSinceVersion()
}

func (*LastlookResponse) ExecPriceDeprecated() uint16 {
	return 0
}

func (*LastlookResponse) ExecPriceMetaAttribute(meta int) string {
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

func (*LastlookResponse) ExecPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*LastlookResponse) ExecPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*LastlookResponse) ExecPriceNullValue() float64 {
	return math.NaN()
}

func (*LastlookResponse) AcceptId() uint16 {
	return 64
}

func (*LastlookResponse) AcceptSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) AcceptInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.AcceptSinceVersion()
}

func (*LastlookResponse) AcceptDeprecated() uint16 {
	return 0
}

func (*LastlookResponse) AcceptMetaAttribute(meta int) string {
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

func (*LastlookResponse) AcceptMinValue() uint8 {
	return 0
}

func (*LastlookResponse) AcceptMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*LastlookResponse) AcceptNullValue() uint8 {
	return math.MaxUint8
}

func (*LastlookResponse) ProxyIndexId() uint16 {
	return 65
}

func (*LastlookResponse) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ProxyIndexSinceVersion()
}

func (*LastlookResponse) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*LastlookResponse) ProxyIndexMetaAttribute(meta int) string {
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

func (*LastlookResponse) ProxyIndexMinValue() uint16 {
	return 0
}

func (*LastlookResponse) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*LastlookResponse) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*LastlookResponse) CompleteId() uint16 {
	return 66
}

func (*LastlookResponse) CompleteSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) CompleteInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.CompleteSinceVersion()
}

func (*LastlookResponse) CompleteDeprecated() uint16 {
	return 0
}

func (*LastlookResponse) CompleteMetaAttribute(meta int) string {
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

func (*LastlookResponse) CompleteMinValue() uint8 {
	return 0
}

func (*LastlookResponse) CompleteMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*LastlookResponse) CompleteNullValue() uint8 {
	return math.MaxUint8
}

func (*LastlookResponse) ClExecIdMetaAttribute(meta int) string {
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

func (*LastlookResponse) ClExecIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) ClExecIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ClExecIdSinceVersion()
}

func (*LastlookResponse) ClExecIdDeprecated() uint16 {
	return 0
}

func (LastlookResponse) ClExecIdCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookResponse) ClExecIdHeaderLength() uint64 {
	return 2
}

func (*LastlookResponse) TradeIdMetaAttribute(meta int) string {
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

func (*LastlookResponse) TradeIdSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) TradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.TradeIdSinceVersion()
}

func (*LastlookResponse) TradeIdDeprecated() uint16 {
	return 0
}

func (LastlookResponse) TradeIdCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookResponse) TradeIdHeaderLength() uint64 {
	return 2
}

func (*LastlookResponse) ExecCurrMetaAttribute(meta int) string {
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

func (*LastlookResponse) ExecCurrSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) ExecCurrInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.ExecCurrSinceVersion()
}

func (*LastlookResponse) ExecCurrDeprecated() uint16 {
	return 0
}

func (LastlookResponse) ExecCurrCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookResponse) ExecCurrHeaderLength() uint64 {
	return 2
}

func (*LastlookResponse) MessageMetaAttribute(meta int) string {
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

func (*LastlookResponse) MessageSinceVersion() uint16 {
	return 0
}

func (l *LastlookResponse) MessageInActingVersion(actingVersion uint16) bool {
	return actingVersion >= l.MessageSinceVersion()
}

func (*LastlookResponse) MessageDeprecated() uint16 {
	return 0
}

func (LastlookResponse) MessageCharacterEncoding() string {
	return "UTF-8"
}

func (LastlookResponse) MessageHeaderLength() uint64 {
	return 2
}
