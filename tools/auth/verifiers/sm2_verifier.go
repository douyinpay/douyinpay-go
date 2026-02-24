package verifiers

import (
	"context"
	"encoding/asn1"
	"encoding/base64"
	"fmt"
	"math/big"
	"strings"

	"github.com/douyinpay/douyinpay-go/utils"
	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

// Sm2Verifier 数字签名验证器
type Sm2Verifier struct {
	PublicKey    *sm2.PublicKey
	SerialNumber string
}

// Verify 对数字签名信息进行验证
func (v *Sm2Verifier) Verify(ctx context.Context, serialNumber, message, signature string) error {
	err := checkParameter(ctx, serialNumber, message, signature)
	if err != nil {
		return err
	}

	// 先自校验 signature
	sigBytes, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		return fmt.Errorf("verify failed: signature not base64 encoded")
	}

	p7, err := x509.ParsePKCS7(sigBytes)
	if err != nil {
		return err
	}

	signerCert, sig, err := getSignerCertAndSignature(p7)
	if err != nil {
		return err
	}
	// 先跟本地的验证一遍
	if serialNumber != v.SerialNumber {
		return fmt.Errorf("verify failed: serialNumber not match")
	}
	// 再跟返回的验证一遍
	if serialNumber != utils.ConvertSerailNo(signerCert.SerialNumber) {
		return fmt.Errorf("verify failed: serialNumber not match")
	}
	// 与本地的公钥进行验证
	msgBytes := []byte(message)
	if !verifyWithUid(v.PublicKey, msgBytes, sig) {
		return fmt.Errorf("failed to verify with signature")
	}
	// 返回的公钥就不再重复验证了，保证与本地一致的就好了
	if !equalSM2PublicKey(utils.CovertToSm2PublicKey(signerCert), v.PublicKey) {
		return fmt.Errorf("verify failed: publicKey not match")
	}

	return nil
}

func getSignerCertAndSignature(p7 *x509.PKCS7) (*x509.Certificate, []byte, error) {
	if p7 == nil || len(p7.Certificates) == 0 || len(p7.Signers) == 0 {
		return nil, nil, fmt.Errorf("invalid pkcs7 signed data:%+v", p7)
	}
	// 只验证第一个签名者,参考java-cfca
	signer := p7.Signers[0]
	// 找到匹配的证书
	name := string(signer.IssuerAndSerialNumber.IssuerName.Bytes)
	sn := signer.IssuerAndSerialNumber.SerialNumber
	for _, cert := range p7.Certificates {
		if strings.Contains(name, cert.Issuer.CommonName) && cert.SerialNumber.Cmp(sn) == 0 {
			return cert, signer.EncryptedDigest, nil
		}
	}
	return nil, nil, fmt.Errorf("no cert found for signer")
}

func equalSM2PublicKey(a, b *sm2.PublicKey) bool {
	if a == nil || b == nil {
		return false
	}
	// 曲线必须一致
	if a.Curve.Params().Name != b.Curve.Params().Name {
		return false
	}
	// 椭圆曲线点坐标
	return a.X.Cmp(b.X) == 0 && a.Y.Cmp(b.Y) == 0
}

var defaultPubKeyVerifyUID = []byte{49, 50, 51, 52, 53, 54, 55, 56, 49, 50, 51, 52, 53, 54, 55, 56}

func verifyWithUid(pub *sm2.PublicKey, msg, sig []byte) bool {
	r, s, err := getSm2RS(sig)
	if err != nil {
		return false
	}
	return sm2.Sm2Verify(pub, msg, defaultPubKeyVerifyUID, r, s)
}

func getSm2RS(sign []byte) (*big.Int, *big.Int, error) {
	// 规范化 r、s，解决前导0的问题
	var rawOctString asn1.RawValue
	_, err := asn1.Unmarshal(sign, &rawOctString)
	if err != nil {
		return nil, nil, err
	}
	nums := make([]*big.Int, 0, 2)
	data := rawOctString.Bytes
	for i := 0; i < 2; i++ {
		var x *big.Int
		x, data, err = getAsn1Interger(data)
		if err != nil {
			return nil, nil, err
		}
		nums = append(nums, x)
	}
	return nums[0], nums[1], nil
}

func getAsn1Interger(data []byte) (*big.Int, []byte, error) {
	if len(data) < 2 {
		return nil, data, fmt.Errorf("data length not match")
	}
	length := int(data[1])
	data = data[2:]
	if length > len(data) {
		return nil, data, fmt.Errorf("invalid length")
	}
	num := new(big.Int)
	num.SetBytes(data[:length])
	data = data[length:]
	return num, data, nil
}

// GetSerial 获取可验签的平台证书序列号
func (v *Sm2Verifier) GetSerial(ctx context.Context) (string, error) {
	return v.SerialNumber, nil
}
