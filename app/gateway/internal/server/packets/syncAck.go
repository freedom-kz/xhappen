package packets

import (
	"fmt"
	"io"

	v1 "xhappen/api/protocol/v1"

	"google.golang.org/protobuf/proto"
)

type SyncAckPacket struct {
	FixedHeader
	v1.SyncAck
}

func (c *SyncAckPacket) String() string {
	str := fmt.Sprintf("Header %s >>> Body: %s", c.FixedHeader, &c.SyncAck)
	return str
}

func (c *SyncAckPacket) Write(w io.Writer) error {
	buf := bufferPoolGet()
	buf.Reset()
	data, err := proto.Marshal(&c.SyncAck)
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

func (c *SyncAckPacket) Unpack(payload []byte) error {
	return proto.Unmarshal(payload, &c.SyncAck)
}
