package module

import (
	"crypto/md5"
	"encoding/hex"
)

const salt = "alsfnqljozlwfanls"

type Encryption struct {
	Password string
	Size     int
}

type option func(e *Encryption)

func WithSize(size int) option {
	return func(e *Encryption) {
		e.Size = size
	}
}

func WithPassword(password string) option {
	return func(e *Encryption) {
		e.Password = password
	}
}

func NewEncryption(opts ...option) *Encryption {
	var ep = &Encryption{
		Size: 3,
	}
	for _, o := range opts {
		o(ep)
	}
	return ep
}

func (e *Encryption) EncodePassword() string {
	m := md5.New()
	m.Write([]byte(e.Password))
	for i := 0; i < e.Size; i++ {
		m.Write([]byte(salt))
	}
	m5 := m.Sum(nil)
	return hex.EncodeToString(m5)
}

func (e *Encryption) VerifyPassword(crypt string) bool {
	return e.EncodePassword() == crypt
}
