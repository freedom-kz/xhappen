package packets

import (
	"fmt"
	"io"

	v1 "xhappen/api/protocol/v1"

	"google.golang.org/protobuf/proto"
)

type DeliverAckPacket struct {
	FixedHeader
	v1.DeliverAck
}

func (c *DeliverAckPacket) String() string {
	str := fmt.Sprintf("Header %s >>> Body: %s", c.FixedHeader, &c.DeliverAck)
	return str
}

func (c *DeliverAckPacket) Write(w io.Writer) error {
	buf := bufferPoolGet()
	buf.Reset()
	data, err := proto.Marshal(&c.DeliverAck)
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

func (c *DeliverAckPacket) Unpack(payload []byte) error {
	return proto.Unmarshal(payload, &c.DeliverAck)
}
