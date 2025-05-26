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

type ProxyStatus struct {
	ProxyIndex uint16
	EventTime  uint64
	ProxyType  []uint8
	Hostname   []uint8
	Status     []uint8
}

func (p *ProxyStatus) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := p.RangeCheck(p.SbeSchemaVersion(), p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint16(_w, p.ProxyIndex); err != nil {
		return err
	}
	if err := _m.WriteUint64(_w, p.EventTime); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(p.ProxyType))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.ProxyType); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(p.Hostname))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.Hostname); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(p.Status))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, p.Status); err != nil {
		return err
	}
	return nil
}

func (p *ProxyStatus) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !p.ProxyIndexInActingVersion(actingVersion) {
		p.ProxyIndex = p.ProxyIndexNullValue()
	} else {
		if err := _m.ReadUint16(_r, &p.ProxyIndex); err != nil {
			return err
		}
	}
	if !p.EventTimeInActingVersion(actingVersion) {
		p.EventTime = p.EventTimeNullValue()
	} else {
		if err := _m.ReadUint64(_r, &p.EventTime); err != nil {
			return err
		}
	}
	if actingVersion > p.SbeSchemaVersion() && blockLength > p.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-p.SbeBlockLength()))
	}

	if p.ProxyTypeInActingVersion(actingVersion) {
		var ProxyTypeLength uint16
		if err := _m.ReadUint16(_r, &ProxyTypeLength); err != nil {
			return err
		}
		if cap(p.ProxyType) < int(ProxyTypeLength) {
			p.ProxyType = make([]uint8, ProxyTypeLength)
		}
		p.ProxyType = p.ProxyType[:ProxyTypeLength]
		if err := _m.ReadBytes(_r, p.ProxyType); err != nil {
			return err
		}
	}

	if p.HostnameInActingVersion(actingVersion) {
		var HostnameLength uint16
		if err := _m.ReadUint16(_r, &HostnameLength); err != nil {
			return err
		}
		if cap(p.Hostname) < int(HostnameLength) {
			p.Hostname = make([]uint8, HostnameLength)
		}
		p.Hostname = p.Hostname[:HostnameLength]
		if err := _m.ReadBytes(_r, p.Hostname); err != nil {
			return err
		}
	}

	if p.StatusInActingVersion(actingVersion) {
		var StatusLength uint16
		if err := _m.ReadUint16(_r, &StatusLength); err != nil {
			return err
		}
		if cap(p.Status) < int(StatusLength) {
			p.Status = make([]uint8, StatusLength)
		}
		p.Status = p.Status[:StatusLength]
		if err := _m.ReadBytes(_r, p.Status); err != nil {
			return err
		}
	}
	if doRangeCheck {
		if err := p.RangeCheck(actingVersion, p.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (p *ProxyStatus) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if p.ProxyIndexInActingVersion(actingVersion) {
		if p.ProxyIndex < p.ProxyIndexMinValue() || p.ProxyIndex > p.ProxyIndexMaxValue() {
			return fmt.Errorf("Range check failed on p.ProxyIndex (%v < %v > %v)", p.ProxyIndexMinValue(), p.ProxyIndex, p.ProxyIndexMaxValue())
		}
	}
	if p.EventTimeInActingVersion(actingVersion) {
		if p.EventTime < p.EventTimeMinValue() || p.EventTime > p.EventTimeMaxValue() {
			return fmt.Errorf("Range check failed on p.EventTime (%v < %v > %v)", p.EventTimeMinValue(), p.EventTime, p.EventTimeMaxValue())
		}
	}
	if !utf8.Valid(p.ProxyType[:]) {
		return errors.New("p.ProxyType failed UTF-8 validation")
	}
	if !utf8.Valid(p.Hostname[:]) {
		return errors.New("p.Hostname failed UTF-8 validation")
	}
	if !utf8.Valid(p.Status[:]) {
		return errors.New("p.Status failed UTF-8 validation")
	}
	return nil
}

func ProxyStatusInit(p *ProxyStatus) {
	return
}

func (*ProxyStatus) SbeBlockLength() (blockLength uint16) {
	return 10
}

func (*ProxyStatus) SbeTemplateId() (templateId uint16) {
	return 220
}

func (*ProxyStatus) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*ProxyStatus) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*ProxyStatus) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*ProxyStatus) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*ProxyStatus) ProxyIndexId() uint16 {
	return 221
}

func (*ProxyStatus) ProxyIndexSinceVersion() uint16 {
	return 0
}

func (p *ProxyStatus) ProxyIndexInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ProxyIndexSinceVersion()
}

func (*ProxyStatus) ProxyIndexDeprecated() uint16 {
	return 0
}

func (*ProxyStatus) ProxyIndexMetaAttribute(meta int) string {
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

func (*ProxyStatus) ProxyIndexMinValue() uint16 {
	return 0
}

func (*ProxyStatus) ProxyIndexMaxValue() uint16 {
	return math.MaxUint16 - 1
}

func (*ProxyStatus) ProxyIndexNullValue() uint16 {
	return math.MaxUint16
}

func (*ProxyStatus) EventTimeId() uint16 {
	return 222
}

func (*ProxyStatus) EventTimeSinceVersion() uint16 {
	return 0
}

func (p *ProxyStatus) EventTimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.EventTimeSinceVersion()
}

func (*ProxyStatus) EventTimeDeprecated() uint16 {
	return 0
}

func (*ProxyStatus) EventTimeMetaAttribute(meta int) string {
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

func (*ProxyStatus) EventTimeMinValue() uint64 {
	return 0
}

func (*ProxyStatus) EventTimeMaxValue() uint64 {
	return math.MaxUint64 - 1
}

func (*ProxyStatus) EventTimeNullValue() uint64 {
	return math.MaxUint64
}

func (*ProxyStatus) ProxyTypeMetaAttribute(meta int) string {
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

func (*ProxyStatus) ProxyTypeSinceVersion() uint16 {
	return 0
}

func (p *ProxyStatus) ProxyTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.ProxyTypeSinceVersion()
}

func (*ProxyStatus) ProxyTypeDeprecated() uint16 {
	return 0
}

func (ProxyStatus) ProxyTypeCharacterEncoding() string {
	return "UTF-8"
}

func (ProxyStatus) ProxyTypeHeaderLength() uint64 {
	return 2
}

func (*ProxyStatus) HostnameMetaAttribute(meta int) string {
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

func (*ProxyStatus) HostnameSinceVersion() uint16 {
	return 0
}

func (p *ProxyStatus) HostnameInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.HostnameSinceVersion()
}

func (*ProxyStatus) HostnameDeprecated() uint16 {
	return 0
}

func (ProxyStatus) HostnameCharacterEncoding() string {
	return "UTF-8"
}

func (ProxyStatus) HostnameHeaderLength() uint64 {
	return 2
}

func (*ProxyStatus) StatusMetaAttribute(meta int) string {
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

func (*ProxyStatus) StatusSinceVersion() uint16 {
	return 0
}

func (p *ProxyStatus) StatusInActingVersion(actingVersion uint16) bool {
	return actingVersion >= p.StatusSinceVersion()
}

func (*ProxyStatus) StatusDeprecated() uint16 {
	return 0
}

func (ProxyStatus) StatusCharacterEncoding() string {
	return "UTF-8"
}

func (ProxyStatus) StatusHeaderLength() uint64 {
	return 2
}
