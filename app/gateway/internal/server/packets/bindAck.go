package packets

import (
	"fmt"
	"io"

	v1 "xhappen/api/protocol/v1"

	"google.golang.org/protobuf/proto"
)

type BindAckPacket struct {
	FixedHeader
	v1.BindAck
}

func (c *BindAckPacket) String() string {
	str := fmt.Sprintf("Header %s >>> Body: %s", c.FixedHeader, &c.BindAck)
	return str
}

func (c *BindAckPacket) Write(w io.Writer) error {
	buf := bufferPoolGet()
	buf.Reset()
	data, err := proto.Marshal(&c.BindAck)
	if err != nil {
		return err
	}

	c.RemainingLength = len(data)
	buf = c.FixedHeader.pack(buf)
	buf.Write(data)
	_, err = w.Write(buf.Bytes())
	bufferPoolPut(buf)
	return err
}

func (c *BindAckPacket) Unpack(payload []byte) error {
	return proto.Unmarshal(payload, &c.BindAck)
}
