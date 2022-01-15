package rand

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
)

type UUID struct {
	bytes []byte
}

func NewUUID() *UUID {
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err.Error())
	}
	bytes[6] = byte(0x40 | (int(bytes[6] & 0xf)))
	bytes[8] = byte(0x80 | (int(bytes[8]) & 0x3f))
	return &UUID{bytes: bytes}
}

func (uuid *UUID) ToString() string {
	return hex.EncodeToString(uuid.bytes)
}

func (uuid *UUID) Bytes() []byte {
	return uuid.bytes
}

func ParserString(str string) (uuid *UUID, err error) {
	if len(str) != 32 {
		return nil, errors.New(str + " is not a valid UUID")
	}
	if str[12] != '4' {
		return nil, errors.New(str + " is not a valid UUID. character 13 should be '4'.")
	}

	if str[16] != '8' && str[16] != '9' && str[16] != 'a' && str[16] != 'b' {
		return nil, errors.New(str + " is not a valid UUID. character 17 should be '8', '9', 'a', or 'b'")
	}

	bytes := make([]byte, 18)
	bytes, err = hex.DecodeString(str)

	if err != nil {
		return nil, err
	}

	return &UUID{bytes: bytes}, nil
}
