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

type Command struct {
	CommandType uint8
	ParamList   []CommandParamList
}
type CommandParamList struct {
	Key   []uint8
	Value []uint8
}

func (c *Command) Encode(_m *SbeGoMarshaller, _w io.Writer, doRangeCheck bool) error {
	if doRangeCheck {
		if err := c.RangeCheck(c.SbeSchemaVersion(), c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	if err := _m.WriteUint8(_w, c.CommandType); err != nil {
		return err
	}
	var ParamListBlockLength uint16 = 0
	if err := _m.WriteUint16(_w, ParamListBlockLength); err != nil {
		return err
	}
	var ParamListNumInGroup uint16 = uint16(len(c.ParamList))
	if err := _m.WriteUint16(_w, ParamListNumInGroup); err != nil {
		return err
	}
	for i := range c.ParamList {
		if err := c.ParamList[i].Encode(_m, _w); err != nil {
			return err
		}
	}
	return nil
}

func (c *Command) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint16, doRangeCheck bool) error {
	if !c.CommandTypeInActingVersion(actingVersion) {
		c.CommandType = c.CommandTypeNullValue()
	} else {
		if err := _m.ReadUint8(_r, &c.CommandType); err != nil {
			return err
		}
	}
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}

	if c.ParamListInActingVersion(actingVersion) {
		var ParamListBlockLength uint16
		if err := _m.ReadUint16(_r, &ParamListBlockLength); err != nil {
			return err
		}
		var ParamListNumInGroup uint16
		if err := _m.ReadUint16(_r, &ParamListNumInGroup); err != nil {
			return err
		}
		if cap(c.ParamList) < int(ParamListNumInGroup) {
			c.ParamList = make([]CommandParamList, ParamListNumInGroup)
		}
		c.ParamList = c.ParamList[:ParamListNumInGroup]
		for i := range c.ParamList {
			if err := c.ParamList[i].Decode(_m, _r, actingVersion, uint(ParamListBlockLength)); err != nil {
				return err
			}
		}
	}
	if doRangeCheck {
		if err := c.RangeCheck(actingVersion, c.SbeSchemaVersion()); err != nil {
			return err
		}
	}
	return nil
}

func (c *Command) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if c.CommandTypeInActingVersion(actingVersion) {
		if c.CommandType < c.CommandTypeMinValue() || c.CommandType > c.CommandTypeMaxValue() {
			return fmt.Errorf("Range check failed on c.CommandType (%v < %v > %v)", c.CommandTypeMinValue(), c.CommandType, c.CommandTypeMaxValue())
		}
	}
	for i := range c.ParamList {
		if err := c.ParamList[i].RangeCheck(actingVersion, schemaVersion); err != nil {
			return err
		}
	}
	return nil
}

func CommandInit(c *Command) {
	return
}

func (c *CommandParamList) Encode(_m *SbeGoMarshaller, _w io.Writer) error {
	if err := _m.WriteUint16(_w, uint16(len(c.Key))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, c.Key); err != nil {
		return err
	}
	if err := _m.WriteUint16(_w, uint16(len(c.Value))); err != nil {
		return err
	}
	if err := _m.WriteBytes(_w, c.Value); err != nil {
		return err
	}
	return nil
}

func (c *CommandParamList) Decode(_m *SbeGoMarshaller, _r io.Reader, actingVersion uint16, blockLength uint) error {
	if actingVersion > c.SbeSchemaVersion() && blockLength > c.SbeBlockLength() {
		io.CopyN(ioutil.Discard, _r, int64(blockLength-c.SbeBlockLength()))
	}

	if c.KeyInActingVersion(actingVersion) {
		var KeyLength uint16
		if err := _m.ReadUint16(_r, &KeyLength); err != nil {
			return err
		}
		if cap(c.Key) < int(KeyLength) {
			c.Key = make([]uint8, KeyLength)
		}
		c.Key = c.Key[:KeyLength]
		if err := _m.ReadBytes(_r, c.Key); err != nil {
			return err
		}
	}

	if c.ValueInActingVersion(actingVersion) {
		var ValueLength uint16
		if err := _m.ReadUint16(_r, &ValueLength); err != nil {
			return err
		}
		if cap(c.Value) < int(ValueLength) {
			c.Value = make([]uint8, ValueLength)
		}
		c.Value = c.Value[:ValueLength]
		if err := _m.ReadBytes(_r, c.Value); err != nil {
			return err
		}
	}
	return nil
}

func (c *CommandParamList) RangeCheck(actingVersion uint16, schemaVersion uint16) error {
	if !utf8.Valid(c.Key[:]) {
		return errors.New("c.Key failed UTF-8 validation")
	}
	if !utf8.Valid(c.Value[:]) {
		return errors.New("c.Value failed UTF-8 validation")
	}
	return nil
}

func CommandParamListInit(c *CommandParamList) {
	return
}

func (*Command) SbeBlockLength() (blockLength uint16) {
	return 1
}

func (*Command) SbeTemplateId() (templateId uint16) {
	return 175
}

func (*Command) SbeSchemaId() (schemaId uint16) {
	return 1
}

func (*Command) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}

func (*Command) SbeSemanticType() (semanticType []byte) {
	return []byte("")
}

func (*Command) SbeSemanticVersion() (semanticVersion string) {
	return "5.2"
}

func (*Command) CommandTypeId() uint16 {
	return 176
}

func (*Command) CommandTypeSinceVersion() uint16 {
	return 0
}

func (c *Command) CommandTypeInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.CommandTypeSinceVersion()
}

func (*Command) CommandTypeDeprecated() uint16 {
	return 0
}

func (*Command) CommandTypeMetaAttribute(meta int) string {
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

func (*Command) CommandTypeMinValue() uint8 {
	return 0
}

func (*Command) CommandTypeMaxValue() uint8 {
	return math.MaxUint8 - 1
}

func (*Command) CommandTypeNullValue() uint8 {
	return math.MaxUint8
}

func (*CommandParamList) KeyMetaAttribute(meta int) string {
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

func (*CommandParamList) KeySinceVersion() uint16 {
	return 0
}

func (c *CommandParamList) KeyInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.KeySinceVersion()
}

func (*CommandParamList) KeyDeprecated() uint16 {
	return 0
}

func (CommandParamList) KeyCharacterEncoding() string {
	return "UTF-8"
}

func (CommandParamList) KeyHeaderLength() uint64 {
	return 2
}

func (*CommandParamList) ValueMetaAttribute(meta int) string {
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

func (*CommandParamList) ValueSinceVersion() uint16 {
	return 0
}

func (c *CommandParamList) ValueInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ValueSinceVersion()
}

func (*CommandParamList) ValueDeprecated() uint16 {
	return 0
}

func (CommandParamList) ValueCharacterEncoding() string {
	return "UTF-8"
}

func (CommandParamList) ValueHeaderLength() uint64 {
	return 2
}

func (*Command) ParamListId() uint16 {
	return 177
}

func (*Command) ParamListSinceVersion() uint16 {
	return 0
}

func (c *Command) ParamListInActingVersion(actingVersion uint16) bool {
	return actingVersion >= c.ParamListSinceVersion()
}

func (*Command) ParamListDeprecated() uint16 {
	return 0
}

func (*CommandParamList) SbeBlockLength() (blockLength uint) {
	return 0
}

func (*CommandParamList) SbeSchemaVersion() (schemaVersion uint16) {
	return 0
}
