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

type TradeReject struct {
	AcctId     int32
	InstId     uint16
	ExecQty    float64
	ExecPrice  float64
	Accept     uint8
	RejectTime uint64
	RejectCode int32
	ProxyIndex uint16
	ExecCurr   []uint8
	ClExecId   []uint8
	TradeId    []uint8
	Text       []uint8
}

func (t *TradeReject) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := t.RangeCheck(t.SbeSchemaVersion(), t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt32(_w, t.AcctId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, t.InstId); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.ExecQty); err != nil {
		return err
	}
	if err := _m.WriteFloat64(_w, t.ExecPrice); err != nil {
		return err
	}
	if err := _m.WriteUint8(_w, t.Accept); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, t.RejectTime); err != nil {
		return err
	}
	if err := _m.WriteInt32(_w, t.RejectCode); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, t.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(t.ExecCurr))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.ExecCurr); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(t.ClExecId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.ClExecId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(t.TradeId))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.TradeId); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(t.Text))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, t.Text); err != nil {
		return err
	}
	return nil
}

func (t *TradeReject) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !t.AcctIdInActingVersion(actingVersion) {
		t.AcctId = t.AcctIdNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.AcctId); err != nil {
			return err
		}
	}
	if !t.InstIdInActingVersion(actingVersion) {
		t.InstId = t.InstIdNullValue()
	} else {
		if err := _m.ReadUint16(_r, &t.InstId); err != nil {
			return err
		}
	}
	if !t.ExecQtyInActingVersion(actingVersion) {
		t.ExecQty = t.ExecQtyNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.ExecQty); err != nil {
			return err
		}
	}
	if !t.ExecPriceInActingVersion(actingVersion) {
		t.ExecPrice = t.ExecPriceNullValue()
	} else {
		if err := _m.ReadFloat64(_r, &t.ExecPrice); err != nil {
			return err
		}
	}
	if !t.AcceptInActingVersion(actingVersion) {
		t.Accept = t.AcceptNullValue()
	} else {
		if err := _m.ReadUint8(_r, &t.Accept); err != nil {
			return err
		}
	}
	if !t.RejectTimeInActingVersion(actingVersion) {
		t.RejectTime = t.RejectTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &t.RejectTime); err != nil {
			return err
		}
	}
	if !t.RejectCodeInActingVersion(actingVersion) {
		t.RejectCode = t.RejectCodeNullValue()
	} else {
		if err := _m.ReadInt32(_r, &t.RejectCode); err != nil {
			return err
		}
	}
	if !t.ProxyIndexInActingVersion(actingVersion) {
		t.ProxyIndex = t.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &t.ProxyIndex); err != nil {
			return err
		}
	}
	if actingVersion > t.SbeSchemaVersion() && blockLength > t.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-t.SbeBlockLength()))
	}

	if t.ExecCurrInActingVersion(actingVersion) {
		var ExecCurrLength uint16
		if err := _m.ReadUint16(_r, &ExecCurrLength); err != nil {
			return err
		}
		if cap(t.ExecCurr) < int(ExecCurrLength) {
			t.ExecCurr = make([]uint8, ExecCurrLength)
		}
		t.ExecCurr = t.ExecCurr[:ExecCurrLength]
		if err := _m.ReadBytes(_r, t.ExecCurr); err != nil {
			return err
		}
	}

	if t.ClExecIdInActingVersion(actingVersion) {
		var ClExecIdLength uint16
		if err := _m.ReadUint16(_r, &ClExecIdLength); err != nil {
			return err
		}
		if cap(t.ClExecId) < int(ClExecIdLength) {
			t.ClExecId = make([]uint8, ClExecIdLength)
		}
		t.ClExecId = t.ClExecId[:ClExecIdLength]
		if err := _m.ReadBytes(_r, t.ClExecId); err != nil {
			return err
		}
	}

	if t.TradeIdInActingVersion(actingVersion) {
		var TradeIdLength uint16
		if err := _m.ReadUint16(_r, &TradeIdLength); err != nil {
			return err
		}
		if cap(t.TradeId) < int(TradeIdLength) {
			t.TradeId = make([]uint8, TradeIdLength)
		}
		t.TradeId = t.TradeId[:TradeIdLength]
		if err := _m.ReadBytes(_r, t.TradeId); err != nil {
			return err
		}
	}

	if t.TextInActingVersion(actingVersion) {
		var TextLength uint16
		if err := _m.ReadUint16(_r, &TextLength); err != nil {
			return err
		}
		if cap(t.Text) < int(TextLength) {
			t.Text = make([]uint8, TextLength)
		}
		t.Text = t.Text[:TextLength]
		if err := _m.ReadBytes(_r, t.Text); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := t.RangeCheck(actingVersion, t.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (t *TradeReject) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if t.AcctIdInActingVersion(actingVersion) {
		if t.AcctId < t.AcctIdMinValue() || t.AcctId > t.AcctIdMaxValue() {
			return fmt.Errorf("Range check failed on t.AcctId (%v < %v > %v)", t.AcctIdMinValue(), t.AcctId, t.AcctIdMaxValue())
		}
	}
	if t.InstIdInActingVersion(actingVersion) {
		if t.InstId < t.InstIdMinValue() || t.InstId > t.InstIdMaxValue() {
			return fmt.Errorf("Range check failed on t.InstId (%v < %v > %v)", t.InstIdMinValue(), t.InstId, t.InstIdMaxValue())
		}
	}
	if t.ExecQtyInActingVersion(actingVersion) {
		if t.ExecQty < t.ExecQtyMinValue() || t.ExecQty > t.ExecQtyMaxValue() {
			return fmt.Errorf("Range check failed on t.ExecQty (%v < %v > %v)", t.ExecQtyMinValue(), t.ExecQty, t.ExecQtyMaxValue())
		}
	}
	if t.ExecPriceInActingVersion(actingVersion) {
		if t.ExecPrice < t.ExecPriceMinValue() || t.ExecPrice > t.ExecPriceMaxValue() {
			return fmt.Errorf("Range check failed on t.ExecPrice (%v < %v > %v)", t.ExecPriceMinValue(), t.ExecPrice, t.ExecPriceMaxValue())
		}
	}
	if t.AcceptInActingVersion(actingVersion) {
		if t.Accept < t.AcceptMinValue() || t.Accept > t.AcceptMaxValue() {
			return fmt.Errorf("Range check failed on t.Accept (%v < %v > %v)", t.AcceptMinValue(), t.Accept, t.AcceptMaxValue())
		}
	}
	if t.RejectTimeInActingVersion(actingVersion) {
		if t.RejectTime < t.RejectTimeMinValue() || t.RejectTime > t.RejectTimeMaxValue() {
			return fmt.Errorf("Range check failed on t.RejectTime (%v < %v > %v)", t.RejectTimeMinValue(), t.RejectTime, t.RejectTimeMaxValue())
		}
	}
	if t.RejectCodeInActingVersion(actingVersion) {
		if t.RejectCode < t.RejectCodeMinValue() || t.RejectCode > t.RejectCodeMaxValue() {
			return fmt.Errorf("Range check failed on t.RejectCode (%v < %v > %v)", t.RejectCodeMinValue(), t.RejectCode, t.RejectCodeMaxValue())
		}
	}
	if t.ProxyIndexInActingVersion(actingVersion) {
		if t.ProxyIndex < t.ProxyIndexMinValue() || t.ProxyIndex > t.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on t.ProxyIndex (%v < %v > %v)", t.ProxyIndexMinValue(), t.ProxyIndex, t.ProxyIndexMaxValue())
		}
	}
	if !utf8.Valid(t.ExecCurr[:]) {
		return errors.New("t.ExecCurr failed UTF-8 validation")
	}
	if !utf8.Valid(t.ClExecId[:]) {
		return errors.New("t.ClExecId failed UTF-8 validation")
	}
	if !utf8.Valid(t.TradeId[:]) {
		return errors.New("t.TradeId failed UTF-8 validation")
	}
	if !utf8.Valid(t.Text[:]) {
		return errors.New("t.Text failed UTF-8 validation")
	}
	return nil
}

func TradeRejectInit(t *TradeReject) {
	return
}

func (*TradeReject) SbeBlockLength() (blockLength uint16) {
	return 37
}

func (*TradeReject) SbeTemplateId() (templateId uint16) {
	return 180
}

func (*TradeReject) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*TradeReject) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*TradeReject) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*TradeReject) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*TradeReject) AcctIdId() uint16 {
	return 181
}

func (*TradeReject) AcctIdSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) AcctIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.AcctIdSinceVersion()
}

func (*TradeReject) AcctIdDeprecated() uint16 {
	return 0
}

func (*TradeReject) AcctIdMetaAttribute(meta int) string {
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

func (*TradeReject) AcctIdMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeReject) AcctIdMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeReject) AcctIdNullValue() int32 {
	return math.MinInt32
}

func (*TradeReject) InstIdId() uint16 {
	return 182
}

func (*TradeReject) InstIdSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) InstIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.InstIdSinceVersion()
}

func (*TradeReject) InstIdDeprecated() uint16 {
	return 0
}

func (*TradeReject) InstIdMetaAttribute(meta int) string {
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

func (*TradeReject) InstIdMinValue() uint16 {
	return 0
}

func (*TradeReject) InstIdMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*TradeReject) InstIdNullValue() uint16 {
	return math.MaxUint16
}

func (*TradeReject) ExecQtyId() uint16 {
	return 183
}

func (*TradeReject) ExecQtySinceVersion() uint16 {
	return 0
}

func (t *TradeReject) ExecQtyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ExecQtySinceVersion()
}

func (*TradeReject) ExecQtyDeprecated() uint16 {
	return 0
}

func (*TradeReject) ExecQtyMetaAttribute(meta int) string {
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

func (*TradeReject) ExecQtyMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeReject) ExecQtyMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeReject) ExecQtyNullValue() float64 {
	return math.NaN()
}

func (*TradeReject) ExecPriceId() uint16 {
	return 184
}

func (*TradeReject) ExecPriceSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) ExecPriceInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ExecPriceSinceVersion()
}

func (*TradeReject) ExecPriceDeprecated() uint16 {
	return 0
}

func (*TradeReject) ExecPriceMetaAttribute(meta int) string {
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

func (*TradeReject) ExecPriceMinValue() float64 {
	return -math.MaxFloat64
}

func (*TradeReject) ExecPriceMaxValue() float64 {
	return math.MaxFloat64
}

func (*TradeReject) ExecPriceNullValue() float64 {
	return math.NaN()
}

func (*TradeReject) AcceptId() uint16 {
	return 185
}

func (*TradeReject) AcceptSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) AcceptInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.AcceptSinceVersion()
}

func (*TradeReject) AcceptDeprecated() uint16 {
	return 0
}

func (*TradeReject) AcceptMetaAttribute(meta int) string {
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

func (*TradeReject) AcceptMinValue() uint8 {
	return 0
}

func (*TradeReject) AcceptMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*TradeReject) AcceptNullValue() uint8 {
	return math.MaxUint8
}

func (*TradeReject) RejectTimeId() uint16 {
	return 186
}

func (*TradeReject) RejectTimeSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) RejectTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.RejectTimeSinceVersion()
}

func (*TradeReject) RejectTimeDeprecated() uint16 {
	return 0
}

func (*TradeReject) RejectTimeMetaAttribute(meta int) string {
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

func (*TradeReject) RejectTimeMinValue() uint64 {
	return 0
}

func (*TradeReject) RejectTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*TradeReject) RejectTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*TradeReject) RejectCodeId() uint16 {
	return 187
}

func (*TradeReject) RejectCodeSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) RejectCodeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.RejectCodeSinceVersion()
}

func (*TradeReject) RejectCodeDeprecated() uint16 {
	return 0
}

func (*TradeReject) RejectCodeMetaAttribute(meta int) string {
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

func (*TradeReject) RejectCodeMinValue() int32 {
	return math.MinInt32 + 1
}

func (*TradeReject) RejectCodeMaxValue() int32 {
	return math.MaxInt32
}

func (*TradeReject) RejectCodeNullValue() int32 {
	return math.MinInt32
}

func (*TradeReject) ProxyIndexId() uint16 {
	return 188
}

func (*TradeReject) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ProxyIndexSinceVersion()
}

func (*TradeReject) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*TradeReject) ProxyIndexMetaAttribute(meta int) string {
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

func (*TradeReject) ProxyIndexMinValue() uint16 {
	return 0
}

func (*TradeReject) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*TradeReject) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*TradeReject) ExecCurrMetaAttribute(meta int) string {
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

func (*TradeReject) ExecCurrSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) ExecCurrInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ExecCurrSinceVersion()
}

func (*TradeReject) ExecCurrDeprecated() uint16 {
	return 0
}

func (TradeReject) ExecCurrCharacterEncoding() string {
	return "UTF-8"
}

func (TradeReject) ExecCurrHeaderLength() uint64 {
	return 2
}

func (*TradeReject) ClExecIdMetaAttribute(meta int) string {
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

func (*TradeReject) ClExecIdSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) ClExecIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.ClExecIdSinceVersion()
}

func (*TradeReject) ClExecIdDeprecated() uint16 {
	return 0
}

func (TradeReject) ClExecIdCharacterEncoding() string {
	return "UTF-8"
}

func (TradeReject) ClExecIdHeaderLength() uint64 {
	return 2
}

func (*TradeReject) TradeIdMetaAttribute(meta int) string {
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

func (*TradeReject) TradeIdSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) TradeIdInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TradeIdSinceVersion()
}

func (*TradeReject) TradeIdDeprecated() uint16 {
	return 0
}

func (TradeReject) TradeIdCharacterEncoding() string {
	return "UTF-8"
}

func (TradeReject) TradeIdHeaderLength() uint64 {
	return 2
}

func (*TradeReject) TextMetaAttribute(meta int) string {
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

func (*TradeReject) TextSinceVersion() uint16 {
	return 0
}

func (t *TradeReject) TextInActingVersion(actingVersion uint16) bool {
	return actingVersion >= t.TextSinceVersion()
}

func (*TradeReject) TextDeprecated() uint16 {
	return 0
}

func (TradeReject) TextCharacterEncoding() string {
	return "UTF-8"
}

func (TradeReject) TextHeaderLength() uint64 {
	return 2
}
