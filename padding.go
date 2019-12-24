// Padding package implements padding methods
package padding

import (
	"bytes"
	"errors"
)

type PKCS5 struct{}

func (PKCS5) Pad(src []byte, blockSize int) []byte {
	p := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{byte(p)}, p)...)
}

func (PKCS5) UnPad(src []byte) ([]byte, error) {
	length := len(src)
	unp := int(src[length-1])
	if length < unp {
		return nil, errors.New("unpadding error")
	}
	return src[:(length - unp)], nil
}

func NewPKCS5() PKCS5 {
	return PKCS5{}
}

type Zero struct{}

func (Zero) Pad(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	return append(src, bytes.Repeat([]byte{0}, padding)...)
}

func (Zero) UnPad(src []byte) ([]byte, error) {
	return bytes.TrimFunc(src,
		func(r rune) bool {
			return r == rune(0)
		}), nil
}

func NewZero() Zero {
	return Zero{}
}
