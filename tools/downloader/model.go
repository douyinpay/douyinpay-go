package downloader

import (
	"encoding/json"
	"fmt"
)

// Certificate 微信支付平台证书信息
type Certificate struct {
	// 证书序列号
	CertNo *string `json:"cert_no"`
	// 证书有效期开始时间
	EffectiveTime *string `json:"effective_time"`
	// 证书过期时间
	ExpireTime *string `json:"expire_time"`
	// 证书类型
	CertType *string `json:"cert_type"`
	// 为了保证安全性，微信支付在回调通知和平台证书下载接口中，对关键信息进行了AES-256-GCM加密
	EncryptCertificate *EncryptCertificate `json:"encrypt_certificate"`
}

func (o *Certificate) Validate() error {
	if o == nil {
		return fmt.Errorf("field `Certificate` is required and must be specified in Certificate")
	}
	if o.CertNo == nil {
		return fmt.Errorf("field `CertNo` is required and must be specified in Certificate")
	}
	if o.CertType == nil {
		return fmt.Errorf("field `CertType` is required and must be specified in Certificate")
	}
	if o.EncryptCertificate == nil {
		return fmt.Errorf("field `EncryptCertificate` is required and must be specified in Certificate")
	}
	return nil
}

func (o *Certificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.CertNo == nil {
		return nil, fmt.Errorf("field `CertNo` is required and must be specified in Certificate")
	}
	toSerialize["cert_no"] = o.CertNo

	if o.EffectiveTime == nil {
		return nil, fmt.Errorf("field `EffectiveTime` is required and must be specified in Certificate")
	}
	toSerialize["effective_time"] = o.EffectiveTime

	if o.ExpireTime == nil {
		return nil, fmt.Errorf("field `ExpireTime` is required and must be specified in Certificate")
	}
	toSerialize["expire_time"] = o.ExpireTime

	if o.CertType == nil {
		return nil, fmt.Errorf("field `CertType` is required and must be specified in Certificate")
	}
	toSerialize["cert_type"] = o.CertType

	if o.EncryptCertificate == nil {
		return nil, fmt.Errorf("field `EncryptCertificate` is required and must be specified in Certificate")
	}
	toSerialize["encrypt_certificate"] = o.EncryptCertificate
	return json.Marshal(toSerialize)
}

func (o *Certificate) String() string {
	var ret string
	if o.CertNo == nil {
		ret += "CertNo:<nil>, "
	} else {
		ret += fmt.Sprintf("CertNo:%v, ", *o.CertNo)
	}

	if o.CertType == nil {
		ret += "CertType:<nil>, "
	} else {
		ret += fmt.Sprintf("CertType:%v, ", *o.CertType)
	}

	if o.EffectiveTime == nil {
		ret += "EffectiveTime:<nil>, "
	} else {
		ret += fmt.Sprintf("EffectiveTime:%v, ", *o.EffectiveTime)
	}

	if o.ExpireTime == nil {
		ret += "ExpireTime:<nil>, "
	} else {
		ret += fmt.Sprintf("ExpireTime:%v, ", *o.ExpireTime)
	}

	ret += fmt.Sprintf("EncryptCertificate:%v", o.EncryptCertificate)

	return fmt.Sprintf("Certificate{%s}", ret)
}

func (o *Certificate) Clone() *Certificate {
	ret := Certificate{}

	if o.CertNo != nil {
		ret.CertNo = new(string)
		*ret.CertNo = *o.CertNo
	}

	if o.EffectiveTime != nil {
		ret.EffectiveTime = new(string)
		*ret.EffectiveTime = *o.EffectiveTime
	}

	if o.ExpireTime != nil {
		ret.ExpireTime = new(string)
		*ret.ExpireTime = *o.ExpireTime
	}

	if o.CertType != nil {
		ret.CertType = new(string)
		*ret.CertType = *o.CertType
	}

	if o.EncryptCertificate != nil {
		ret.EncryptCertificate = o.EncryptCertificate.Clone()
	}

	return &ret
}

// DownloadCertificatesResponse
type DownloadCertificatesResponse struct {
	// 平台证书列表
	Certificates []Certificate `json:"certificates,omitempty"`
}

func (o *DownloadCertificatesResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Certificates != nil {
		toSerialize["Certificates"] = o.Certificates
	}
	return json.Marshal(toSerialize)
}

func (o *DownloadCertificatesResponse) String() string {
	var ret string
	ret += fmt.Sprintf("Certificates:%v", o.Certificates)

	return fmt.Sprintf("DownloadCertificatesResponse{%s}", ret)
}

func (o *DownloadCertificatesResponse) Clone() *DownloadCertificatesResponse {
	ret := DownloadCertificatesResponse{}

	if o.Certificates != nil {
		ret.Certificates = make([]Certificate, len(o.Certificates))
		for i, item := range o.Certificates {
			ret.Certificates[i] = *item.Clone()
		}
	}

	return &ret
}

// EncryptCertificate 加密证书信息
type EncryptCertificate struct {
	// 证书内容密文，解密后会获得证书完整内容
	Ciphertext *string `json:"cipher_text"`
	// 加密所使用的算法，目前可能取值仅为 AEAD_AES_256_GCM
	Algorithm *string `json:"algorithm"`
	// 加密所使用的随机字符串
	Nonce *string `json:"nonce"`
}

func (o *EncryptCertificate) Validate() error {
	if o == nil {
		return fmt.Errorf("field `EncryptCertificate` is required and must be specified in EncryptCertificate")
	}
	if o.Algorithm == nil {
		return fmt.Errorf("field `Algorithm` is required and must be specified in EncryptCertificate")
	}
	if o.Nonce == nil {
		return fmt.Errorf("field `Nonce` is required and must be specified in EncryptCertificate")
	}
	if o.Ciphertext == nil {
		return fmt.Errorf("field `Ciphertext` is required and must be specified in EncryptCertificate")
	}
	return nil
}
func (o *EncryptCertificate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}

	if o.Algorithm == nil {
		return nil, fmt.Errorf("field `Algorithm` is required and must be specified in EncryptCertificate")
	}
	toSerialize["algorithm"] = o.Algorithm

	if o.Nonce == nil {
		return nil, fmt.Errorf("field `Nonce` is required and must be specified in EncryptCertificate")
	}
	toSerialize["nonce"] = o.Nonce
	if o.Ciphertext == nil {
		return nil, fmt.Errorf("field `Ciphertext` is required and must be specified in EncryptCertificate")
	}
	toSerialize["ciphertext"] = o.Ciphertext
	return json.Marshal(toSerialize)
}

func (o *EncryptCertificate) String() string {
	var ret string
	if o.Algorithm == nil {
		ret += "Algorithm:<nil>, "
	} else {
		ret += fmt.Sprintf("Algorithm:%v, ", *o.Algorithm)
	}

	if o.Nonce == nil {
		ret += "Nonce:<nil>, "
	} else {
		ret += fmt.Sprintf("Nonce:%v, ", *o.Nonce)
	}

	if o.Ciphertext == nil {
		ret += "Ciphertext:<nil>"
	} else {
		ret += fmt.Sprintf("Ciphertext:%v", *o.Ciphertext)
	}

	return fmt.Sprintf("EncryptCertificate{%s}", ret)
}

func (o *EncryptCertificate) Clone() *EncryptCertificate {
	ret := EncryptCertificate{}

	if o.Algorithm != nil {
		ret.Algorithm = new(string)
		*ret.Algorithm = *o.Algorithm
	}

	if o.Nonce != nil {
		ret.Nonce = new(string)
		*ret.Nonce = *o.Nonce
	}

	if o.Ciphertext != nil {
		ret.Ciphertext = new(string)
		*ret.Ciphertext = *o.Ciphertext
	}

	return &ret
}
