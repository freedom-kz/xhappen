package util

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

func NewId() string {
	u1 := uuid.NewV4()
	u := strings.ReplaceAll(u1.String(), "-", "")
	return u
}
