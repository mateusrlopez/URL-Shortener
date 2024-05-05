package utils

import (
	"math"
	"strings"
)

type Encryptor interface {
	EncryptId(id int64) string
}

type base62Encryptor struct {
	base       int64
	characters string
}

func NewBase62Encryptor() Encryptor {
	return &base62Encryptor{
		base:       62,
		characters: "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz",
	}
}

func (enc base62Encryptor) EncryptId(id int64) string {
	var builder strings.Builder

	for id > 0 {
		remainder := int(math.Mod(float64(id), float64(enc.base)))
		builder.WriteByte(enc.characters[remainder])
		id /= enc.base
	}

	return builder.String()
}
