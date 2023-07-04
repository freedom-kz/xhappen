package packets

import (
	"fmt"
	"io"

	v1 "xhappen/api/protocol/v1"

	"google.golang.org/protobuf/proto"
)

type SubmitAckPacket struct {
	FixedHeader
	v1.SubmitAck
}

func (c *SubmitAckPacket) String() string {
	str := fmt.Sprintf("Header %s >>> Body: %s", c.FixedHeader, &c.SubmitAck)
	return str
}

func (c *SubmitAckPacket) Write(w io.Writer) error {
	buf := bufferPoolGet()
	buf.Reset()
	data, err := proto.Marshal(&c.SubmitAck)
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

func (c *SubmitAckPacket) Unpack(payload []byte) error {
	return proto.Unmarshal(payload, &c.SubmitAck)
}
