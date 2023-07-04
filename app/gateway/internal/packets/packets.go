package packets

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

const (
	PACKET_MAX_SIZE = 1 << 23 //8M
)

var (
	OVERFLOW_ERR        = errors.New("Packet RemainingLength overflow")
	DAMAGE_DATA_ERR     = errors.New("Bad data from client")
	READ_DATA_SHORT_ERR = errors.New("Failed to read expected data")
)

const (
	RESERVED = iota
	BIND
	BINDACK
	AUTH
	AUTHACK
	SYNC
	SYNCACK
	SYNCCONFIRM
	SUBMIT
	SUBMITACK
	DELIVER
	DELIVERACK
	ACTION
	ACTIONACK
	PING
	PONG
	QUIT
	RESERVED2
)

var PacketNames = map[uint8]string{
	0:  "RESERVED",
	1:  "BIND",
	2:  "BINDACK",
	3:  "AUTH",
	4:  "AUTHACK",
	5:  "SYNC",
	6:  "SYNCACK",
	7:  "SYNCCONFIRM",
	8:  "SUBMIT",
	9:  "SUBMITACK",
	10: "DELIVER",
	11: "DELIVERACK",
	12: "ACTION",
	13: "ACTIONACK",
	14: "PING",
	15: "PONG",
	16: "QUIT",
	17: "RESERVED2",
}

type ControlPacket interface {
	Write(io.Writer) error
	Unpack([]byte) error
	String() string
}

type FixedHeader struct {
	MessageType           byte
	Retain                byte
	RemainingEncodeLength int
	RemainingLength       int
}

func (fh FixedHeader) String() string {
	return fmt.Sprintf("%s: retain: %b remainingLength: %d", PacketNames[fh.MessageType], fh.Retain, fh.RemainingLength)
}

func (fh *FixedHeader) unpack(typeAndFlags byte, r io.Reader) {
	fh.MessageType = typeAndFlags >> 4
	fh.Retain = typeAndFlags & 0x0f
	fh.RemainingLength, fh.RemainingEncodeLength = decodeLength(r)
}

func (fh *FixedHeader) pack(buffer *bytes.Buffer) *bytes.Buffer {
	buffer.WriteByte(fh.MessageType<<4 | fh.Retain&0xf)
	fh.RemainingEncodeLength, _ = buffer.Write(encodeLength(fh.RemainingLength))
	return buffer
}

func encodeLength(length int) []byte {
	var encLength []byte
	for {
		digit := byte(length % 128)
		length /= 128
		if length > 0 {
			digit |= 0x80
		}
		encLength = append(encLength, digit)
		if length == 0 {
			break
		}
	}
	return encLength
}

func decodeLength(r io.Reader) (int, int) {
	var rLength uint32
	var multiplier uint32
	b := make([]byte, 1)
	for multiplier < 27 { //fix: Infinite '(digit & 128) == 1' will cause the dead loop
		io.ReadFull(r, b)
		digit := b[0]
		rLength |= uint32(digit&127) << multiplier
		if (digit & 128) == 0 {
			break
		}
		multiplier += 7
	}
	return int(rLength), int(multiplier/7 + 1)
}

func NewControlPacketWithHeader(fh FixedHeader) (cp ControlPacket) {
	switch fh.MessageType {
	case BIND:
		cp = &BindPacket{FixedHeader: fh}
	case BINDACK:
		cp = &BindAckPacket{FixedHeader: fh}
	case AUTH:
		cp = &AUTHPacket{FixedHeader: fh}
	case AUTHACK:
		cp = &ActionAckPacket{FixedHeader: fh}
	case SYNC:
		cp = &SyncPacket{FixedHeader: fh}
	case SYNCACK:
		cp = &SyncAckPacket{FixedHeader: fh}
	case SYNCCONFIRM:
		cp = &SyncConfirmPacket{FixedHeader: fh}
	case SUBMIT:
		cp = &SubmitPacket{FixedHeader: fh}
	case SUBMITACK:
		cp = &SubmitAckPacket{FixedHeader: fh}
	case DELIVER:
		cp = &DeliverPacket{FixedHeader: fh}
	case DELIVERACK:
		cp = &DeliverAckPacket{FixedHeader: fh}
	case ACTION:
		cp = &ActionPacket{FixedHeader: fh}
	case ACTIONACK:
		cp = &AuthAckPacket{FixedHeader: fh}
	case PING:
		cp = &PingPacket{FixedHeader: fh}
	case PONG:
		cp = &PongPacket{FixedHeader: fh}
	case QUIT:
		cp = &QuitPacket{FixedHeader: fh}
	}
	return cp
}

func ReadPacket(r io.Reader) (cp ControlPacket, err error) {
	var fh FixedHeader
	b := make([]byte, 1)
	_, err = io.ReadFull(r, b)
	if err != nil {
		return nil, err
	}
	fh.unpack(b[0], r)
	if fh.RemainingLength > PACKET_MAX_SIZE {
		return nil, OVERFLOW_ERR
	}
	cp = NewControlPacketWithHeader(fh)
	if cp == nil {
		return nil, DAMAGE_DATA_ERR
	}
	packetBytes := make([]byte, fh.RemainingLength)
	n, err := io.ReadFull(r, packetBytes)
	if err != nil {
		return nil, err
	}
	if n != fh.RemainingLength {
		return nil, READ_DATA_SHORT_ERR
	}

	err = cp.Unpack(packetBytes)
	return cp, err
}
