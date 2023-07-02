package packets

import (
	"fmt"
	"io"

	v1 "xhappen/api/protocol/v1"

	"google.golang.org/protobuf/proto"
)

type SyncPacket struct {
	FixedHeader
	v1.Sync
}

func (c *SyncPacket) String() string {
	str := fmt.Sprintf("Header %s >>> Body: %s", c.FixedHeader, &c.Sync)
	return str
}

func (c *SyncPacket) Write(w io.Writer) error {
	buf := bufferPoolGet()
	buf.Reset()
	data, err := proto.Marshal(&c.Sync)
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

func (c *SyncPacket) Unpack(payload []byte) error {
	return proto.Unmarshal(payload, &c.Sync)
}
