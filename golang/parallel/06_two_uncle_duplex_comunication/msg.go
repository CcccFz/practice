package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

type Msg interface {
	Encode() []byte
	Type() msgType
	Serial() uint32
}

type (
	msgType    uint8
	EatZReq    struct{ *msg }
	PlaceLReq  struct{ *msg }
	PleaseLReq struct{ *msg }
	EatLRes    struct{ *msg }
	PlaceZRes  struct{ *msg }
	PleaseZRes struct{ *msg }
)

const (
	_              = iota
	msgTypeEatZReq // 1
	msgTypePlaceLReq
	msgTypePleaseLReq
	msgTypeEatLRes // 4
	msgTypePlaceZRes
	msgTypePleaseZRes
)

func NewMsg(serial uint32, typ msgType) (*msg, error) {
	var payload string

	switch typ {
	case msgTypeEatZReq:
		payload = "吃了没，您吶?"
	case msgTypePlaceLReq:
		payload = "您这，嘛去？"
	case msgTypePleaseLReq:
		payload = "有空家里坐坐啊。"
	case msgTypeEatLRes:
		payload = "刚吃。"
	case msgTypePlaceZRes:
		payload = "嗨！吃饱了溜溜弯儿。"
	case msgTypePleaseZRes:
		payload = "回头去给老太太请安！"
	default:
		return nil, fmt.Errorf("不支持的消息类型: %d", typ)
	}

	return &msg{serial: serial, typ: typ, payload: []byte(payload)}, nil
}

type msg struct {
	serial  uint32
	typ     msgType
	payload []byte
}

func (m *msg) Serial() uint32 {
	return m.serial
}

// S-T-L-P
func (m *msg) Encode() []byte {
	serialBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(serialBytes, m.serial)

	lengthBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthBytes, uint32(len(m.payload)))

	return bytesCombine(serialBytes, []byte{byte(m.typ)}, lengthBytes, m.payload)
}

func (m *msg) Type() msgType {
	return m.typ
}

func Decode(r io.Reader) (Msg, error) {
	serialBytes, typBytes, lengthBytes := make([]byte, 4), make([]byte, 1), make([]byte, 4)

	_, err := io.ReadFull(r, serialBytes)
	if err != nil {
		return nil, err
	}
	if _, err = io.ReadFull(r, typBytes); err != nil {
		return nil, err
	}
	if _, err = io.ReadFull(r, lengthBytes); err != nil {
		return nil, err
	}

	serial := binary.BigEndian.Uint32(serialBytes)
	typ := msgType(typBytes[0])
	length := binary.BigEndian.Uint32(lengthBytes)

	payload := make([]byte, length)
	if _, err = io.ReadFull(r, payload); err != nil {
		return nil, err
	}

	message := &msg{serial: serial, typ: typ, payload: payload}

	switch typ {
	case msgTypeEatZReq:
		return &EatZReq{msg: message}, nil
	case msgTypePlaceLReq:
		return &PlaceLReq{msg: message}, nil
	case msgTypePleaseLReq:
		return &PleaseLReq{msg: message}, nil
	case msgTypeEatLRes:
		return &EatLRes{msg: message}, nil
	case msgTypePlaceZRes:
		return &PlaceZRes{msg: message}, nil
	case msgTypePleaseZRes:
		return &PleaseZRes{msg: message}, nil
	default:
		return nil, fmt.Errorf("不支持的消息类型: %d", typ)
	}
}

func bytesCombine(_bytes ...[]byte) []byte {
	return bytes.Join(_bytes, []byte(""))
}
