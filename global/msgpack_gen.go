package global

// NOTE: THIS FILE WAS PRODUCED BY THE
// MSGP CODE GENERATION TOOL (github.com/tinylib/msgp)
// DO NOT EDIT

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *DelToken) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zxvk uint32
	zxvk, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zxvk > 0 {
		zxvk--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, err = dc.ReadBytes(z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, err = dc.ReadBytes(z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, err = dc.ReadBytes(z.Token)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *DelToken) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "acc"
	err = en.Append(0x83, 0xa3, 0x61, 0x63, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Acc)
	if err != nil {
		return
	}
	// write "ai"
	err = en.Append(0xa2, 0x61, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.AppID)
	if err != nil {
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Token)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *DelToken) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "acc"
	o = append(o, 0x83, 0xa3, 0x61, 0x63, 0x63)
	o = msgp.AppendBytes(o, z.Acc)
	// string "ai"
	o = append(o, 0xa2, 0x61, 0x69)
	o = msgp.AppendBytes(o, z.AppID)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendBytes(o, z.Token)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *DelToken) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zbzg uint32
	zbzg, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zbzg > 0 {
		zbzg--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, bts, err = msgp.ReadBytesBytes(bts, z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, bts, err = msgp.ReadBytesBytes(bts, z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, bts, err = msgp.ReadBytesBytes(bts, z.Token)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *DelToken) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.Acc) + 3 + msgp.BytesPrefixSize + len(z.AppID) + 2 + msgp.BytesPrefixSize + len(z.Token)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *ProtoBufMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zajw uint32
	zajw, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zajw > 0 {
		zajw--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "q":
			var zwht uint32
			zwht, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Qos) >= int(zwht) {
				z.Qos = (z.Qos)[:zwht]
			} else {
				z.Qos = make([]int32, zwht)
			}
			for zbai := range z.Qos {
				z.Qos[zbai], err = dc.ReadInt32()
				if err != nil {
					return
				}
			}
		case "mi":
			var zhct uint32
			zhct, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.MsgIDs) >= int(zhct) {
				z.MsgIDs = (z.MsgIDs)[:zhct]
			} else {
				z.MsgIDs = make([][]byte, zhct)
			}
			for zcmr := range z.MsgIDs {
				z.MsgIDs[zcmr], err = dc.ReadBytes(z.MsgIDs[zcmr])
				if err != nil {
					return
				}
			}
		case "m":
			z.Msg, err = dc.ReadBytes(z.Msg)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *ProtoBufMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 3
	// write "q"
	err = en.Append(0x83, 0xa1, 0x71)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Qos)))
	if err != nil {
		return
	}
	for zbai := range z.Qos {
		err = en.WriteInt32(z.Qos[zbai])
		if err != nil {
			return
		}
	}
	// write "mi"
	err = en.Append(0xa2, 0x6d, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.MsgIDs)))
	if err != nil {
		return
	}
	for zcmr := range z.MsgIDs {
		err = en.WriteBytes(z.MsgIDs[zcmr])
		if err != nil {
			return
		}
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Msg)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *ProtoBufMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 3
	// string "q"
	o = append(o, 0x83, 0xa1, 0x71)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Qos)))
	for zbai := range z.Qos {
		o = msgp.AppendInt32(o, z.Qos[zbai])
	}
	// string "mi"
	o = append(o, 0xa2, 0x6d, 0x69)
	o = msgp.AppendArrayHeader(o, uint32(len(z.MsgIDs)))
	for zcmr := range z.MsgIDs {
		o = msgp.AppendBytes(o, z.MsgIDs[zcmr])
	}
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendBytes(o, z.Msg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *ProtoBufMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zcua uint32
	zcua, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zcua > 0 {
		zcua--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "q":
			var zxhx uint32
			zxhx, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Qos) >= int(zxhx) {
				z.Qos = (z.Qos)[:zxhx]
			} else {
				z.Qos = make([]int32, zxhx)
			}
			for zbai := range z.Qos {
				z.Qos[zbai], bts, err = msgp.ReadInt32Bytes(bts)
				if err != nil {
					return
				}
			}
		case "mi":
			var zlqf uint32
			zlqf, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.MsgIDs) >= int(zlqf) {
				z.MsgIDs = (z.MsgIDs)[:zlqf]
			} else {
				z.MsgIDs = make([][]byte, zlqf)
			}
			for zcmr := range z.MsgIDs {
				z.MsgIDs[zcmr], bts, err = msgp.ReadBytesBytes(bts, z.MsgIDs[zcmr])
				if err != nil {
					return
				}
			}
		case "m":
			z.Msg, bts, err = msgp.ReadBytesBytes(bts, z.Msg)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *ProtoBufMsg) Msgsize() (s int) {
	s = 1 + 2 + msgp.ArrayHeaderSize + (len(z.Qos) * (msgp.Int32Size)) + 3 + msgp.ArrayHeaderSize
	for zcmr := range z.MsgIDs {
		s += msgp.BytesPrefixSize + len(z.MsgIDs[zcmr])
	}
	s += 2 + msgp.BytesPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *PubApns) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdaf uint32
	zdaf, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdaf > 0 {
		zdaf--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, err = dc.ReadBytes(z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, err = dc.ReadBytes(z.AppID)
			if err != nil {
				return
			}
		case "mi":
			z.MsgID, err = dc.ReadBytes(z.MsgID)
			if err != nil {
				return
			}
		case "m":
			z.Msg, err = dc.ReadBytes(z.Msg)
			if err != nil {
				return
			}
		case "jm":
			z.JsonMsg, err = dc.ReadBytes(z.JsonMsg)
			if err != nil {
				return
			}
		case "s":
			z.Sound, err = dc.ReadBytes(z.Sound)
			if err != nil {
				return
			}
		case "b":
			z.Badge, err = dc.ReadInt()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *PubApns) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 7
	// write "acc"
	err = en.Append(0x87, 0xa3, 0x61, 0x63, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Acc)
	if err != nil {
		return
	}
	// write "ai"
	err = en.Append(0xa2, 0x61, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.AppID)
	if err != nil {
		return
	}
	// write "mi"
	err = en.Append(0xa2, 0x6d, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.MsgID)
	if err != nil {
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Msg)
	if err != nil {
		return
	}
	// write "jm"
	err = en.Append(0xa2, 0x6a, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.JsonMsg)
	if err != nil {
		return
	}
	// write "s"
	err = en.Append(0xa1, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Sound)
	if err != nil {
		return
	}
	// write "b"
	err = en.Append(0xa1, 0x62)
	if err != nil {
		return err
	}
	err = en.WriteInt(z.Badge)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *PubApns) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 7
	// string "acc"
	o = append(o, 0x87, 0xa3, 0x61, 0x63, 0x63)
	o = msgp.AppendBytes(o, z.Acc)
	// string "ai"
	o = append(o, 0xa2, 0x61, 0x69)
	o = msgp.AppendBytes(o, z.AppID)
	// string "mi"
	o = append(o, 0xa2, 0x6d, 0x69)
	o = msgp.AppendBytes(o, z.MsgID)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendBytes(o, z.Msg)
	// string "jm"
	o = append(o, 0xa2, 0x6a, 0x6d)
	o = msgp.AppendBytes(o, z.JsonMsg)
	// string "s"
	o = append(o, 0xa1, 0x73)
	o = msgp.AppendBytes(o, z.Sound)
	// string "b"
	o = append(o, 0xa1, 0x62)
	o = msgp.AppendInt(o, z.Badge)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *PubApns) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpks uint32
	zpks, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpks > 0 {
		zpks--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, bts, err = msgp.ReadBytesBytes(bts, z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, bts, err = msgp.ReadBytesBytes(bts, z.AppID)
			if err != nil {
				return
			}
		case "mi":
			z.MsgID, bts, err = msgp.ReadBytesBytes(bts, z.MsgID)
			if err != nil {
				return
			}
		case "m":
			z.Msg, bts, err = msgp.ReadBytesBytes(bts, z.Msg)
			if err != nil {
				return
			}
		case "jm":
			z.JsonMsg, bts, err = msgp.ReadBytesBytes(bts, z.JsonMsg)
			if err != nil {
				return
			}
		case "s":
			z.Sound, bts, err = msgp.ReadBytesBytes(bts, z.Sound)
			if err != nil {
				return
			}
		case "b":
			z.Badge, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *PubApns) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.Acc) + 3 + msgp.BytesPrefixSize + len(z.AppID) + 3 + msgp.BytesPrefixSize + len(z.MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msg) + 3 + msgp.BytesPrefixSize + len(z.JsonMsg) + 2 + msgp.BytesPrefixSize + len(z.Sound) + 2 + msgp.IntSize
	return
}

// DecodeMsg implements msgp.Decodable
func (z *SetToken) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zcxo uint32
	zcxo, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zcxo > 0 {
		zcxo--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, err = dc.ReadBytes(z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, err = dc.ReadBytes(z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, err = dc.ReadBytes(z.Token)
			if err != nil {
				return
			}
		case "tps":
			var zeff uint32
			zeff, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Topics) >= int(zeff) {
				z.Topics = (z.Topics)[:zeff]
			} else {
				z.Topics = make([][]byte, zeff)
			}
			for zjfb := range z.Topics {
				z.Topics[zjfb], err = dc.ReadBytes(z.Topics[zjfb])
				if err != nil {
					return
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *SetToken) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 4
	// write "acc"
	err = en.Append(0x84, 0xa3, 0x61, 0x63, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Acc)
	if err != nil {
		return
	}
	// write "ai"
	err = en.Append(0xa2, 0x61, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.AppID)
	if err != nil {
		return
	}
	// write "t"
	err = en.Append(0xa1, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Token)
	if err != nil {
		return
	}
	// write "tps"
	err = en.Append(0xa3, 0x74, 0x70, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Topics)))
	if err != nil {
		return
	}
	for zjfb := range z.Topics {
		err = en.WriteBytes(z.Topics[zjfb])
		if err != nil {
			return
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *SetToken) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 4
	// string "acc"
	o = append(o, 0x84, 0xa3, 0x61, 0x63, 0x63)
	o = msgp.AppendBytes(o, z.Acc)
	// string "ai"
	o = append(o, 0xa2, 0x61, 0x69)
	o = msgp.AppendBytes(o, z.AppID)
	// string "t"
	o = append(o, 0xa1, 0x74)
	o = msgp.AppendBytes(o, z.Token)
	// string "tps"
	o = append(o, 0xa3, 0x74, 0x70, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Topics)))
	for zjfb := range z.Topics {
		o = msgp.AppendBytes(o, z.Topics[zjfb])
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *SetToken) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zrsw uint32
	zrsw, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zrsw > 0 {
		zrsw--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "acc":
			z.Acc, bts, err = msgp.ReadBytesBytes(bts, z.Acc)
			if err != nil {
				return
			}
		case "ai":
			z.AppID, bts, err = msgp.ReadBytesBytes(bts, z.AppID)
			if err != nil {
				return
			}
		case "t":
			z.Token, bts, err = msgp.ReadBytesBytes(bts, z.Token)
			if err != nil {
				return
			}
		case "tps":
			var zxpk uint32
			zxpk, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Topics) >= int(zxpk) {
				z.Topics = (z.Topics)[:zxpk]
			} else {
				z.Topics = make([][]byte, zxpk)
			}
			for zjfb := range z.Topics {
				z.Topics[zjfb], bts, err = msgp.ReadBytesBytes(bts, z.Topics[zjfb])
				if err != nil {
					return
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *SetToken) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.Acc) + 3 + msgp.BytesPrefixSize + len(z.AppID) + 2 + msgp.BytesPrefixSize + len(z.Token) + 4 + msgp.ArrayHeaderSize
	for zjfb := range z.Topics {
		s += msgp.BytesPrefixSize + len(z.Topics[zjfb])
	}
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TextMsg) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zdnj uint32
	zdnj, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zdnj > 0 {
		zdnj--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "fac":
			z.FAcc, err = dc.ReadBytes(z.FAcc)
			if err != nil {
				return
			}
		case "ft":
			z.FTopic, err = dc.ReadBytes(z.FTopic)
			if err != nil {
				return
			}
		case "rc":
			z.RetryCount, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "q":
			z.Qos, err = dc.ReadInt32()
			if err != nil {
				return
			}
		case "mi":
			z.MsgID, err = dc.ReadBytes(z.MsgID)
			if err != nil {
				return
			}
		case "m":
			z.Msg, err = dc.ReadBytes(z.Msg)
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TextMsg) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 6
	// write "fac"
	err = en.Append(0x86, 0xa3, 0x66, 0x61, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.FAcc)
	if err != nil {
		return
	}
	// write "ft"
	err = en.Append(0xa2, 0x66, 0x74)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.FTopic)
	if err != nil {
		return
	}
	// write "rc"
	err = en.Append(0xa2, 0x72, 0x63)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.RetryCount)
	if err != nil {
		return
	}
	// write "q"
	err = en.Append(0xa1, 0x71)
	if err != nil {
		return err
	}
	err = en.WriteInt32(z.Qos)
	if err != nil {
		return
	}
	// write "mi"
	err = en.Append(0xa2, 0x6d, 0x69)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.MsgID)
	if err != nil {
		return
	}
	// write "m"
	err = en.Append(0xa1, 0x6d)
	if err != nil {
		return err
	}
	err = en.WriteBytes(z.Msg)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TextMsg) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 6
	// string "fac"
	o = append(o, 0x86, 0xa3, 0x66, 0x61, 0x63)
	o = msgp.AppendBytes(o, z.FAcc)
	// string "ft"
	o = append(o, 0xa2, 0x66, 0x74)
	o = msgp.AppendBytes(o, z.FTopic)
	// string "rc"
	o = append(o, 0xa2, 0x72, 0x63)
	o = msgp.AppendInt32(o, z.RetryCount)
	// string "q"
	o = append(o, 0xa1, 0x71)
	o = msgp.AppendInt32(o, z.Qos)
	// string "mi"
	o = append(o, 0xa2, 0x6d, 0x69)
	o = msgp.AppendBytes(o, z.MsgID)
	// string "m"
	o = append(o, 0xa1, 0x6d)
	o = msgp.AppendBytes(o, z.Msg)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TextMsg) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zobc uint32
	zobc, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zobc > 0 {
		zobc--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "fac":
			z.FAcc, bts, err = msgp.ReadBytesBytes(bts, z.FAcc)
			if err != nil {
				return
			}
		case "ft":
			z.FTopic, bts, err = msgp.ReadBytesBytes(bts, z.FTopic)
			if err != nil {
				return
			}
		case "rc":
			z.RetryCount, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "q":
			z.Qos, bts, err = msgp.ReadInt32Bytes(bts)
			if err != nil {
				return
			}
		case "mi":
			z.MsgID, bts, err = msgp.ReadBytesBytes(bts, z.MsgID)
			if err != nil {
				return
			}
		case "m":
			z.Msg, bts, err = msgp.ReadBytesBytes(bts, z.Msg)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TextMsg) Msgsize() (s int) {
	s = 1 + 4 + msgp.BytesPrefixSize + len(z.FAcc) + 3 + msgp.BytesPrefixSize + len(z.FTopic) + 3 + msgp.Int32Size + 2 + msgp.Int32Size + 3 + msgp.BytesPrefixSize + len(z.MsgID) + 2 + msgp.BytesPrefixSize + len(z.Msg)
	return
}

// DecodeMsg implements msgp.Decodable
func (z *TextMsgs) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zkgt uint32
	zkgt, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zkgt > 0 {
		zkgt--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ms":
			var zema uint32
			zema, err = dc.ReadArrayHeader()
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zema) {
				z.Msgs = (z.Msgs)[:zema]
			} else {
				z.Msgs = make([]*TextMsg, zema)
			}
			for zsnv := range z.Msgs {
				if dc.IsNil() {
					err = dc.ReadNil()
					if err != nil {
						return
					}
					z.Msgs[zsnv] = nil
				} else {
					if z.Msgs[zsnv] == nil {
						z.Msgs[zsnv] = new(TextMsg)
					}
					err = z.Msgs[zsnv].DecodeMsg(dc)
					if err != nil {
						return
					}
				}
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *TextMsgs) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 1
	// write "ms"
	err = en.Append(0x81, 0xa2, 0x6d, 0x73)
	if err != nil {
		return err
	}
	err = en.WriteArrayHeader(uint32(len(z.Msgs)))
	if err != nil {
		return
	}
	for zsnv := range z.Msgs {
		if z.Msgs[zsnv] == nil {
			err = en.WriteNil()
			if err != nil {
				return
			}
		} else {
			err = z.Msgs[zsnv].EncodeMsg(en)
			if err != nil {
				return
			}
		}
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *TextMsgs) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 1
	// string "ms"
	o = append(o, 0x81, 0xa2, 0x6d, 0x73)
	o = msgp.AppendArrayHeader(o, uint32(len(z.Msgs)))
	for zsnv := range z.Msgs {
		if z.Msgs[zsnv] == nil {
			o = msgp.AppendNil(o)
		} else {
			o, err = z.Msgs[zsnv].MarshalMsg(o)
			if err != nil {
				return
			}
		}
	}
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *TextMsgs) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zpez uint32
	zpez, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zpez > 0 {
		zpez--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "ms":
			var zqke uint32
			zqke, bts, err = msgp.ReadArrayHeaderBytes(bts)
			if err != nil {
				return
			}
			if cap(z.Msgs) >= int(zqke) {
				z.Msgs = (z.Msgs)[:zqke]
			} else {
				z.Msgs = make([]*TextMsg, zqke)
			}
			for zsnv := range z.Msgs {
				if msgp.IsNil(bts) {
					bts, err = msgp.ReadNilBytes(bts)
					if err != nil {
						return
					}
					z.Msgs[zsnv] = nil
				} else {
					if z.Msgs[zsnv] == nil {
						z.Msgs[zsnv] = new(TextMsg)
					}
					bts, err = z.Msgs[zsnv].UnmarshalMsg(bts)
					if err != nil {
						return
					}
				}
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *TextMsgs) Msgsize() (s int) {
	s = 1 + 3 + msgp.ArrayHeaderSize
	for zsnv := range z.Msgs {
		if z.Msgs[zsnv] == nil {
			s += msgp.NilSize
		} else {
			s += z.Msgs[zsnv].Msgsize()
		}
	}
	return
}