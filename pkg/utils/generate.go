package utils

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"strings"
	"time"

	uuid "github.com/satori/go.uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func GenerateId() string {
	u1 := uuid.NewV4()
	u := strings.ReplaceAll(u1.String(), "-", "")
	return u
}

func GenerateAuthCode(width int) string {
	numeric := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	r := len(numeric)
	rand.Seed(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < width; i++ {
		fmt.Fprintf(&sb, "%d", numeric[rand.Intn(r)])
	}
	return sb.String()
}

func Hash(s string) uint64 {
	h := fnv.New64a()
	h.Write([]byte(s))
	return h.Sum64()
}

func HashWith2String(first string, second string) uint64 {
	hf := Hash(first)
	hs := Hash(second)
	if hf <= hs {
		return Hash(first + second)
	} else {
		return Hash(second + first)
	}
}
