package utils

import (
	"crypto/rsa"
	xx "crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"

	"github.com/tjfoc/gmsm/x509"
)

// LoadCertificate 通过证书的文本内容加载证书
func LoadCertificate(certificateStr string) (certificate *x509.Certificate, err error) {
	block, _ := pem.Decode([]byte(certificateStr))
	if block == nil {
		return nil, fmt.Errorf("decode certificate err")
	}
	if block.Type != "CERTIFICATE" {
		return nil, fmt.Errorf("the kind of PEM should be CERTIFICATE")
	}
	certificate, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse certificate err:%s", err.Error())
	}
	return certificate, nil
}

// LoadPrivateKey 通过私钥的文本内容加载私钥
func LoadPrivateKey(privateKeyStr string) (privateKey *rsa.PrivateKey, err error) {
	block, _ := pem.Decode([]byte(privateKeyStr))
	if block == nil {
		return nil, fmt.Errorf("decode private key err")
	}
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("the kind of PEM should be PRVATE KEY")
	}
	key, err := xx.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("parse private key err:%s", err.Error())
	}
	privateKey, ok := key.(*rsa.PrivateKey)
	if !ok {
		return nil, fmt.Errorf("%s is not rsa private key", privateKeyStr)
	}
	return privateKey, nil
}

// LoadPublicKey 通过公钥的文本内容加载公钥
func LoadPublicKey(publicKeyStr string) (publicKey *rsa.PublicKey, err error) {
	bytes, _ := base64.StdEncoding.DecodeString(publicKeyStr)
	key, err := x509.ParsePKIXPublicKey(bytes)
	if err != nil {
		return nil, fmt.Errorf("parse public key err:%s", err.Error())
	}
	publicKey, ok := key.(*rsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("%s is not rsa public key", publicKeyStr)
	}
	return publicKey, nil
}

// LoadCertificateWithPath  通过证书的文件路径加载证书
func LoadCertificateWithPath(path string) (certificate *x509.Certificate, err error) {
	certificateBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read certificate pem file err:%s", err.Error())
	}
	return LoadCertificate(string(certificateBytes))
}

// LoadPrivateKeyWithPath 通过私钥的文件路径内容加载私钥
func LoadPrivateKeyWithPath(path string) (privateKey *rsa.PrivateKey, err error) {
	privateKeyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read private pem file err:%s", err.Error())
	}
	return LoadPrivateKey(string(privateKeyBytes))
}

// LoadPublicKeyWithPath 通过公钥的文件路径加载公钥
func LoadPublicKeyWithPath(path string) (publicKey *rsa.PublicKey, err error) {
	publicKeyBytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read certificate pem file err:%s", err.Error())
	}
	return LoadPublicKey(string(publicKeyBytes))
}

// GetCertificateSerialNumber 从证书中获取证书序列号
func ConvertSerailNo(SerialNumber *big.Int) string {
	return fmt.Sprintf("%X", SerialNumber.Bytes())
}

func TrimPrivateKey(keyString string) string {
	keyString = strings.Replace(keyString, "-----BEGIN PRIVATE KEY-----", "", -1)
	keyString = strings.Replace(keyString, "-----END PRIVATE KEY-----", "", -1)
	keyString = strings.ReplaceAll(keyString, "\n", "")
	return keyString
}

func TrimCertificate(keyString string) string {
	keyString = strings.Replace(keyString, "-----BEGIN CERTIFICATE-----", "", -1)
	keyString = strings.Replace(keyString, "-----END CERTIFICATE-----", "", -1)
	keyString = strings.ReplaceAll(keyString, "\n", "")
	return keyString
}
