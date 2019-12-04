package controller

import (
	"crypto/sha256"
)

func HashPass(str string) {
	_=sha256.Sum256([]byte(str))
}