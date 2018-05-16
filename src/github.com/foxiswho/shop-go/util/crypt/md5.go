package crypt

import (
	"crypto/md5"
	"fmt"
	"io"
)

func Md5(msg string) string {
	h := md5.New()
	io.WriteString(h, msg)
	return fmt.Sprintf("%x", h.Sum(nil))
	//srcData := []byte(msg)
	//h.Write(srcData)
	//cipherText := h.Sum(nil)
	//hexText := make([]byte, 32)
	//hex.Encode(hexText, cipherText)
	//return string(hexText)
}
