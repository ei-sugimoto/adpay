package vo

import (
	"encoding/base64"

	"golang.org/x/crypto/bcrypt"
)

type Password string

func NewPassword(password string) Password {
	return Password(password)
}

func (p Password) String() string {
	return string(p)
}

func (p Password) Compare(password string) bool {
	return p.String() == password
}

func (p Password) Crypto() Password {
	cost := bcrypt.DefaultCost
	hash, _ := bcrypt.GenerateFromPassword([]byte(p.String()), cost)

	return Password(base64.StdEncoding.EncodeToString(hash))

}

func (p Password) Verify(password string) bool {
	hash, _ := base64.StdEncoding.DecodeString(p.String())
	err := bcrypt.CompareHashAndPassword(hash, []byte(password))

	return err == nil
}
