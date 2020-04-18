package util

import (
	"crypto/sha1"
	"encoding/hex"

	// "hash"
	"io"
	"os"
)

// type Sha1Stream struct {
// 	_sha1 hash.Hash
// }

// func (obj *Sha1Stream) Update(data []byte) {

// }

func FileSha1(file *os.File) string {
	_sha1 := sha1.New()
	io.Copy(_sha1, file)
	return hex.EncodeToString(_sha1.Sum(nil))
}

func Sha1(data []byte) string {
	_sha1 := sha1.New()
	_sha1.Write(data)
	return hex.EncodeToString(_sha1.Sum([]byte("")))
}