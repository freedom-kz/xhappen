package packets

import (
	"fmt"
	"io"

	v1 "xhappen/api/protocol/v1"

	"google.golang.org/protobuf/proto"
)

type ActionAckPacket struct {
	FixedHeader
	v1.ActionAck
}

func (c *ActionAckPacket) String() string {
	str := fmt.Sprintf("Header %s >>> Body: %s", c.FixedHeader, &c.ActionAck)
	return str
}

func (c *ActionAckPacket) Write(w io.Writer) error {
	buf := bufferPoolGet()
	buf.Reset()
	data, err := proto.Marshal(&c.ActionAck)
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

func (c *ActionAckPacket) Unpack(payload []byte) error {
	return proto.Unmarshal(payload, &c.ActionAck)
}
