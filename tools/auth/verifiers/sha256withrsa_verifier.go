package verifiers

import (
	"context"
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type SHA256WithRSAVerifier struct {
	PublicKey    *rsa.PublicKey
	SerialNumber string
}

// Verify 对数字签名信息进行验证
func (v *SHA256WithRSAVerifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	err := checkParameter(ctx, serialNumber, message, signature)
	if err != nil {
		return err
	}

	//比对序列号
	if serialNumber != v.SerialNumber {
		return fmt.Errorf("verify failed: serialNumber not match")
	}

	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("verify failed: signature not base64 encoded")
	}
	hashed := sha256.Sum256([]byte(message))

	err = rsa.VerifyPKCS1v15(v.PublicKey, crypto.SHA256, hashed[:], sigBytes)
	if err != nil {
		return fmt.Errorf("verifty signature with public key err:%s", err.Error())
	}
	return nil
}

func checkParameter(ctx context.Context, serialNumber, message, signature string) error {
	if ctx == nil {
		return fmt.Errorf("context is nil, verifier need input context.Context")
	}
	if strings.TrimSpace(serialNumber) == "" {
		return fmt.Errorf("serialNumber is empty, verifier need input serialNumber")
	}
	if strings.TrimSpace(message) == "" {
		return fmt.Errorf("message is empty, verifier need input message")
	}
	if strings.TrimSpace(signature) == "" {
		return fmt.Errorf("signature is empty, verifier need input signature")
	}
	return nil
}

// GetSerial 获取可验签的平台证书序列号
func (verifier *SHA256WithRSAVerifier) GetSerial(ctx context.Context) (string, error) {
	return verifier.SerialNumber, nil
}
