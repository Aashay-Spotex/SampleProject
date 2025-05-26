// Generated SBE (Simple Binary Encoding) message codec

package core_gui

import (
	"fmt"
	"io"
	"io/ioutil"
	"math"
)

type Heartbeat struct {
	Time int64
}

func (h *Heartbeat) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := h.RangeCheck(h.SbeSchemaVersion(), h.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteInt64(_w, h.Time); err != nil {
		return err
	}
	return nil
}

func (h *Heartbeat) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !h.TimeInActingVersion(actingVersion) {
		h.Time = h.TimeNullValue()
	} else {
		if err := _m.ReadInt64(_r, &h.Time); err != nil {
			return err
		}
	}
	if actingVersion > h.SbeSchemaVersion() && blockLength > h.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-h.SbeBlockLength()))
	}
	if doRangeCheck {
		if err := h.RangeCheck(actingVersion, h.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (h *Heartbeat) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if h.TimeInActingVersion(actingVersion) {
		if h.Time < h.TimeMinValue() || h.Time > h.TimeMaxValue() {
			return fmt.Errorf("Range check failed on h.Time (%v < %v > %v)", h.TimeMinValue(), h.Time, h.TimeMaxValue())
		}
	}
	return nil
}

func HeartbeatInit(h *Heartbeat) {
	return
}

func (*Heartbeat) SbeBlockLength() (blockLength uint16) {
	return 8
}

func (*Heartbeat) SbeTemplateId() (templateId uint16) {
	return 141
}

func (*Heartbeat) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Heartbeat) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Heartbeat) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Heartbeat) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*Heartbeat) TimeId() uint16 {
	return 142
}

func (*Heartbeat) TimeSinceVersion() uint16 {
	return 0
}

func (h *Heartbeat) TimeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= h.TimeSinceVersion()
}

func (*Heartbeat) TimeDeprecated() uint16 {
	return 0
}

func (*Heartbeat) TimeMetaAttribute(meta int) string {
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

func (*Heartbeat) TimeMinValue() int64 {
	return math.MinInt64 + 1
}

func (*Heartbeat) TimeMaxValue() int64 {
	return math.MaxInt64
}

func (*Heartbeat) TimeNullValue() int64 {
	return math.MinInt64
}
