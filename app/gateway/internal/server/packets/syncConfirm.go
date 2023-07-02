package packets

import (
	"fmt"
	"io"

	v1 "xhappen/api/protocol/v1"

	"google.golang.org/protobuf/proto"
)

type SyncConfirmPacket struct {
	FixedHeader
	v1.SyncConfirm
}

func (c *SyncConfirmPacket) String() string {
	str := fmt.Sprintf("Header %s >>> Body: %s", c.FixedHeader, &c.SyncConfirm)
	return str
}

func (c *SyncConfirmPacket) Write(w io.Writer) error {
	buf := bufferPoolGet()
	buf.Reset()
	data, err := proto.Marshal(&c.SyncConfirm)
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

func (c *SyncConfirmPacket) Unpack(payload []byte) error {
	return proto.Unmarshal(payload, &c.SyncConfirm)
}
