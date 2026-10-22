package indirectmerchant

import (
	"context"
	"crypto/rand"
	"crypto/rsa"
	"testing"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/services"
	"github.com/douyinpay/douyinpay-go/tools/auth"
)

// stubSigner 为测试提供满足 DialSettings.Validate 的签名器占位实现。
type stubSigner struct{}

func (stubSigner) Sign(ctx context.Context, message string) (*auth.SignatureResult, error) {
	return &auth.SignatureResult{}, nil
}
func (stubSigner) Algorithm() string { return "RSA" }

// stubVerifier 为测试提供满足 DialSettings.Validate 的验签器占位实现。
type stubVerifier struct{}

func (stubVerifier) Verify(ctx context.Context, serial, message, signature string) error { return nil }
func (stubVerifier) GetSerial(ctx context.Context) (string, error)                       { return "", nil }

func newServiceWithKeys(t *testing.T) (*ApiService, *rsa.PrivateKey) {
	t.Helper()
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatalf("generate rsa key: %v", err)
	}
	cli, err := client.NewClientWithDialSettings(context.Background(), &client.DialSettings{
		Signer:             stubSigner{},
		Verifier:           stubVerifier{},
		PlatformPublicKey:  &key.PublicKey,
		MerchantPrivateKey: key,
	})
	if err != nil {
		t.Fatalf("new client: %v", err)
	}
	svc := (*ApiService)(&services.Service{Client: cli})
	return svc, key
}

// TestEncryptIdentificationDoesNotMutateCaller 验证加密作用于副本，不修改调用方传入的身份信息对象。
func TestEncryptIdentificationDoesNotMutateCaller(t *testing.T) {
	svc, _ := newServiceWithKeys(t)
	original := IdentificationInfo{
		IdentificationType:   "IDENTIFICATION_TYPE_IDCARD",
		IdentificationName:   "张三",
		IdentificationNumber: "110101199001011234",
	}

	encrypted, err := svc.encryptIdentification(original)
	if err != nil {
		t.Fatalf("encrypt identification: %v", err)
	}
	if original.IdentificationName != "张三" || original.IdentificationNumber != "110101199001011234" {
		t.Fatalf("caller object was mutated: %+v", original)
	}
	if encrypted.IdentificationName == original.IdentificationName {
		t.Fatal("identification_name should be encrypted")
	}
	if encrypted.IdentificationNumber == original.IdentificationNumber {
		t.Fatal("identification_number should be encrypted")
	}
	// 非敏感字段应原样保留。
	if encrypted.IdentificationType != original.IdentificationType {
		t.Fatal("identification_type should be preserved")
	}
}

// TestEncryptThenDecryptIdentification 验证服务层加密后再解密能够还原敏感字段。
func TestEncryptThenDecryptIdentification(t *testing.T) {
	svc, _ := newServiceWithKeys(t)
	encrypted, err := svc.encryptIdentification(IdentificationInfo{
		IdentificationName:   "李四",
		IdentificationNumber: "310101199203034567",
	})
	if err != nil {
		t.Fatalf("encrypt: %v", err)
	}
	info := encrypted
	if err := svc.decryptIdentification(&info); err != nil {
		t.Fatalf("decrypt: %v", err)
	}
	if info.IdentificationName != "李四" {
		t.Fatalf("identification_name mismatch: got %q", info.IdentificationName)
	}
	if info.IdentificationNumber != "310101199203034567" {
		t.Fatalf("identification_number mismatch: got %q", info.IdentificationNumber)
	}
}

// TestDecryptIdentificationInvalidCipher 验证解密非法密文时返回带上下文错误。
func TestDecryptIdentificationInvalidCipher(t *testing.T) {
	svc, _ := newServiceWithKeys(t)
	info := IdentificationInfo{IdentificationName: "not-a-valid-cipher"}
	if err := svc.decryptIdentification(&info); err == nil {
		t.Fatal("expected error decrypting invalid cipher")
	}
}

// TestEncryptIdentificationEmptyFields 验证敏感字段为空时跳过加密且不报错。
func TestEncryptIdentificationEmptyFields(t *testing.T) {
	svc, _ := newServiceWithKeys(t)
	info := IdentificationInfo{IdentificationType: "IDENTIFICATION_TYPE_IDCARD"}
	got, err := svc.encryptIdentification(info)
	if err != nil {
		t.Fatalf("encrypt empty sensitive fields: %v", err)
	}
	if got.IdentificationName != "" || got.IdentificationNumber != "" {
		t.Fatal("empty sensitive fields should remain empty")
	}
}
