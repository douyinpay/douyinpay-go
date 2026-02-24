package utils

import (
	"crypto/ecdsa"
	"encoding/base64"
	"fmt"
	"io/ioutil"

	"github.com/tjfoc/gmsm/sm2"
	"github.com/tjfoc/gmsm/x509"
)

func LoadSm2PrivateKeyWithPath(privateKeyPath string) (*sm2.PrivateKey, error) {
	certificateBytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		return nil, fmt.Errorf("read privateKey pem file err:%s", err.Error())
	}
	return LoadSm2PrivateKey(string(certificateBytes))
}

func LoadSm2PrivateKey(privateKeyString string) (*sm2.PrivateKey, error) {
	privateKeyDecode, err := base64.StdEncoding.DecodeString(TrimPrivateKey(privateKeyString))
	if err != nil {
		return nil, fmt.Errorf("decode sm2 privateKey error:%s", err.Error())
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(privateKeyDecode, nil)
	if err != nil {
		return nil, fmt.Errorf("parse sm2 privateKey error:%s", err.Error())
	}
	return privateKey, err
}

func LoadSm2CertificateWithPath(cetificatePath string) (*sm2.PrivateKey, error) {
	certificateBytes, err := ioutil.ReadFile(cetificatePath)
	if err != nil {
		return nil, fmt.Errorf("read privateKey pem file err:%s", err.Error())
	}
	return LoadSm2PrivateKey(string(certificateBytes))
}

func LoadSm2Certificate(certificateString string) (*x509.Certificate, error) {
	certificateDecode, err := base64.StdEncoding.DecodeString(TrimCertificate(certificateString))
	if err != nil {
		return nil, fmt.Errorf("decode sm2 certificate error:%s", err.Error())
	}
	certificate, err := x509.ParseCertificate([]byte(certificateDecode))
	if err != nil {
		return nil, fmt.Errorf("parse sm2 certificate error:%s", err.Error())
	}
	return certificate, nil
}

func CovertToSm2PublicKey(cert *x509.Certificate) *sm2.PublicKey {
	if cert == nil || cert.PublicKey == nil {
		return nil
	}

	switch pub := cert.PublicKey.(type) {
	case *ecdsa.PublicKey:
		switch pub.Curve {
		case sm2.P256Sm2():
			sm2PublicKey := &sm2.PublicKey{
				Curve: pub.Curve,
				X:     pub.X,
				Y:     pub.Y,
			}
			return sm2PublicKey
		default:
			return nil
		}
	}
	return nil
}
