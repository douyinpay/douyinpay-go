package utils

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"encoding/base64"
	"fmt"
	"github.com/tjfoc/gmsm/sm2"
)

func SignSHA256WithRSA(source string, privateKey *rsa.PrivateKey) (signature string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key should not be nil")
	}
	h := crypto.Hash.New(crypto.SHA256)
	_, err = h.Write([]byte(source))
	if err != nil {
		return "", nil
	}
	hashed := h.Sum(nil)
	signatureByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}

func SignSm2(source string, privateKey *sm2.PrivateKey) (signature string, err error) {
	if privateKey == nil {
		return "", fmt.Errorf("private key should not be nil")
	}
	signatureByte, err := privateKey.Sign(rand.Reader, []byte(source), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(signatureByte), nil
}
