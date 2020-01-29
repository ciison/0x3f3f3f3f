package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"hash"
)

func main() {
	fmt.Println(hashSerial(md5.New(), "hello", false))
	fmt.Println(hashSerial(sha256.New(), "hello", false))
	fmt.Println(hashSerial(sha1.New(), "hello", false))

}

//
func hashSerial(h hash.Hash, text string, hEx bool) string {
	if h == nil {
		return ""
	}

	if hEx {
		decodeString, _ := hex.DecodeString(text)
		h.Write(decodeString)
	} else {
		h.Write([]byte(text))
	}

	// 选择将生成的序列数据追加到一个空的切片中， 返回追加之后的数据序列
	b := h.Sum(nil)
	return fmt.Sprintf("%x", b)
}
