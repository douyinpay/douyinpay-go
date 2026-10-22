package option

import (
	"crypto/rsa"

	"github.com/douyinpay/douyinpay-go/client"
)

// sensitiveFieldOption 配置敏感字段加解密所需的密钥。
// platformPublicKey 为抖音支付平台证书中的公钥，用于上送敏感字段加密；
// merchantPrivateKey 为服务商私钥，用于下行敏感字段密文解密。
type sensitiveFieldOption struct {
	platformPublicKey  *rsa.PublicKey
	merchantPrivateKey *rsa.PrivateKey
}

// Apply 将敏感字段密钥写入 client.DialSettings。
func (w sensitiveFieldOption) Apply(o *client.DialSettings) error {
	o.PlatformPublicKey = w.platformPublicKey
	o.MerchantPrivateKey = w.merchantPrivateKey
	return nil
}

// OptionSensitiveFieldKeys 使用抖音支付平台证书公钥和服务商私钥初始化敏感字段加解密能力。
// 上送证件持有人姓名、证件号码等敏感字段时使用 platformPublicKey 加密，
// 解析下行密文时使用 merchantPrivateKey 解密。
func OptionSensitiveFieldKeys(platformPublicKey *rsa.PublicKey, merchantPrivateKey *rsa.PrivateKey) client.ClientOption {
	return sensitiveFieldOption{
		platformPublicKey:  platformPublicKey,
		merchantPrivateKey: merchantPrivateKey,
	}
}
