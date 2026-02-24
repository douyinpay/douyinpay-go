package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"

	"github.com/tjfoc/gmsm/sm4"
)

// 使用 AEAD_AES_256_GCM 算法进行解密
func DecryptAES256GCM(aesKey, associatedData, nonce, ciphertext string) (plaintext string, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	c, err := aes.NewCipher([]byte(aesKey))
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCMWithNonceSize(c, len(nonce))
	if err != nil {
		return "", err
	}
	dataBytes, err := gcm.Open(nil, []byte(nonce), decodedCiphertext, []byte(associatedData))
	if err != nil {
		return "", err
	}
	return string(dataBytes), nil
}

// 使用 SM4 算法进行解密
func DecryptSM4(sm4Key, associatedData, nonce, ciphertext string) (plaintext string, err error) {
	decodedCiphertext, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}
	block, err := sm4.NewCipher([]byte(sm4Key))
	if err != nil {
		return "", err
	}
	dataBytes := make([]byte, len(decodedCiphertext))
	blockMode := cipher.NewCBCDecrypter(block, []byte(nonce))
	blockMode.CryptBlocks(dataBytes, decodedCiphertext)

	// unpadding 处理
	length := len(dataBytes)
	padding := int(dataBytes[length-1])
	dst := dataBytes[:(length - padding)]

	return string(dst), nil
}
