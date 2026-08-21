package downloader

import (
	"encoding/json"
	"fmt"
)

// Certificate 平台证书信息
type Certificate struct {
	// 字段含义：证书序列号。
	// 格式规则：string。
	// 示例：763CC8F6EF3A8802。
	CertNo *string `json:"cert_no"`
	// 字段含义：证书生效时间。
	// 格式规则：string，格式 YYMMDDHHMMSS。
	// 业务规则：加密请求中的敏感信息时，应使用生效时间较晚（最新）的平台证书。
	// 示例：20230322042245。
	EffectiveTime *string `json:"effective_time"`
	// 字段含义：证书失效时间。
	// 格式规则：string，格式 YYMMDDHHMMSS。
	// 业务规则：证书过期后将失效；抖音支付会在过期前提前把新证书加入平台证书查询列表。
	// 示例：20280320042245。
	ExpireTime *string `json:"expire_time"`
	// 字段含义：证书类型。
	// 格式规则：string。
	// 业务规则：目前取值为 RSA。
	// 示例：RSA。
	CertType *string `json:"cert_type"`
	// 字段含义：证书加密后的内容。
	// 格式规则：object。
	// 业务规则：为保证安全，证书内容经加密返回；调用前需在“产品中心-密钥管理”完成“接口加密密钥”设置，以解密证书密文。
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

// DownloadCertificatesResponse 平台证书查询接口响应
type DownloadCertificatesResponse struct {
	// 字段含义：平台证书信息。
	// 格式规则：array。
	// 业务规则：调用方应定期（建议 6~12 小时）获取并更新平台证书，不要硬编码用于验签的平台证书；加密请求中的敏感信息时使用生效时间较晚（最新）的证书。
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

// EncryptCertificate 证书加密后的内容
type EncryptCertificate struct {
	// 字段含义：证书密文。
	// 格式规则：string。
	// 业务规则：使用接口加密密钥解密后可获得证书完整内容。
	// 示例：lRatST1Wlxoxxxxxxxxxxxxxxxx。
	Ciphertext *string `json:"cipher_text"`
	// 字段含义：加密算法。
	// 格式规则：string。
	// 业务规则：目前取值仅为 AEAD-AES-256-GCM。
	// 示例：AEAD-AES-256-GCM。
	Algorithm *string `json:"algorithm"`
	// 字段含义：随机串。
	// 格式规则：string。
	// 业务规则：对应到加密算法中的 IV。
	// 示例：6tKL7i5sEaO4。
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
