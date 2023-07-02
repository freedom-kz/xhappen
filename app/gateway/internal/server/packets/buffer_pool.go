package packets

import (
	"bytes"
	"sync"
)

const (
	MAX_BUFFER_SIZE = 1024 * 4 //可循环用buf大小限制
)

var bp sync.Pool

func init() {
	bp.New = func() interface{} {
		return &bytes.Buffer{}
	}
}

func bufferPoolGet() *bytes.Buffer {
	return bp.Get().(*bytes.Buffer)
}

func bufferPoolPut(b *bytes.Buffer) {
	if b.Cap() > MAX_BUFFER_SIZE {
		return
	}

	b.Reset()
	bp.Put(b)
}
