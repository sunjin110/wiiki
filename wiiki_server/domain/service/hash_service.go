package service

import (
	"wiiki_server/common/wiikierr"

	"golang.org/x/crypto/bcrypt"
)

type Hash interface {
	Generate(str string) (string, error)
	Compare(hash string, str string) error
}

func NewHash(cost int) Hash {
	return &hashImpl{
		cost: cost,
	}
}

type hashImpl struct {
	cost int
}

func (impl *hashImpl) Generate(str string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(str), impl.cost)
	if err != nil {
		return "", wiikierr.Bind(err, wiikierr.FailedGenerateHash, "str=%s", str)
	}

	return string(hashed), nil
}

func (impl *hashImpl) Compare(hash string, str string) error {

	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	if err != nil {

		if err == bcrypt.ErrMismatchedHashAndPassword {
			return wiikierr.Bind(err, wiikierr.FailedMismatchHashAndPassword, "hash=%s, str=%s", hash, str)
		}

		return wiikierr.Bind(err, wiikierr.FailedCompareHashAndPassword, "hash=%s, str=%s", hash, str)
	}

	return nil
}
