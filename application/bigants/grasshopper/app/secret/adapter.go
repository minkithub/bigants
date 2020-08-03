// Package secret 은 Argon2di 기반의 암호화, cryptographically secure한 랜덤 패스워드 생성 API를 제공한다.
package secret

import (
	"encoding/base64"
)

type Adapter struct{}

func (*Adapter) Hash(plain string) (encoded string, err error) {
	return CreateHashFromPassword(plain)
}

func (*Adapter) Compare(encoded string, plain string) (ok bool, err error) {
	return ComparePasswordAndHash(plain, encoded)

}
func (a *Adapter) RandomPassword() (plain string, err error) {
	b, err := generateRandomBytes(33)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}
