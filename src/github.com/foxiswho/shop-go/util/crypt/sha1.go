package crypt

import (
	"crypto/sha1"
	"fmt"
	"io"
)

func Sha1(msg string) string {
	h := sha1.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
}
