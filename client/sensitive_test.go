package client

import (
	"crypto/rand"
	"crypto/rsa"
	"testing"
)

func mustGenKey(t *testing.T) *rsa.PrivateKey {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate rsa key: %v", err)
	}
	return key
}

// TestEncryptDecryptSensitiveRoundTrip 验证平台公钥加密、服务商私钥解密能够往返还原明文。
func TestEncryptDecryptSensitiveRoundTrip(t *testing.T) {
	key := mustGenKey(t)
	c := &Client{platformPublicKey: &key.PublicKey, merchantPrivateKey: key}

	plain := "张三"
	cipher, err := c.EncryptSensitive(plain)
	if err != nil {
		t.Fatalf("encrypt: %v", err)
	}
	if cipher == "" || cipher == plain {
		t.Fatalf("cipher should be non-empty base64 and differ from plain, got %q", cipher)
	}

	got, err := c.DecryptSensitive(cipher)
	if err != nil {
		t.Fatalf("decrypt: %v", err)
	}
	if got != plain {
		t.Fatalf("round trip mismatch: got %q want %q", got, plain)
	}
}

// TestEncryptSensitiveWithoutKey 验证未配置平台公钥时加密返回带上下文错误。
func TestEncryptSensitiveWithoutKey(t *testing.T) {
	c := &Client{}
	if _, err := c.EncryptSensitive("张三"); err == nil {
		t.Fatal("expected error when platform public key is not configured")
	}
}

// TestDecryptSensitiveWithoutKey 验证未配置服务商私钥时解密返回带上下文错误。
func TestDecryptSensitiveWithoutKey(t *testing.T) {
	c := &Client{}
	if _, err := c.DecryptSensitive("aGVsbG8="); err == nil {
		t.Fatal("expected error when merchant private key is not configured")
	}
}

// TestDecryptSensitiveInvalidBase64 验证非法 Base64 密文返回带上下文错误。
func TestDecryptSensitiveInvalidBase64(t *testing.T) {
	key := mustGenKey(t)
	c := &Client{merchantPrivateKey: key}
	if _, err := c.DecryptSensitive("not-base64-!!!"); err == nil {
		t.Fatal("expected error on invalid base64 cipher")
	}
}

// TestDecryptSensitiveWrongKey 验证使用错误私钥解密密文时返回错误。
func TestDecryptSensitiveWrongKey(t *testing.T) {
	encKey := mustGenKey(t)
	decKey := mustGenKey(t)
	enc := &Client{platformPublicKey: &encKey.PublicKey}
	dec := &Client{merchantPrivateKey: decKey}

	cipher, err := enc.EncryptSensitive("张三")
	if err != nil {
		t.Fatalf("encrypt: %v", err)
	}
	if _, err := dec.DecryptSensitive(cipher); err == nil {
		t.Fatal("expected error when decrypting with wrong private key")
	}
}

// TestHasSensitiveKeys 验证是否已配置平台公钥的判断。
func TestHasSensitiveKeys(t *testing.T) {
	if (&Client{}).HasSensitiveKeys() {
		t.Fatal("empty client should not report sensitive keys")
	}
	key := mustGenKey(t)
	if !(&Client{platformPublicKey: &key.PublicKey}).HasSensitiveKeys() {
		t.Fatal("client with platform public key should report sensitive keys")
	}
}
