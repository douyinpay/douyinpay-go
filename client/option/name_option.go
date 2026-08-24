package option

import "github.com/douyinpay/douyinpay-go/client"

const (
	// 每次sdk升级，需要更新版本号
	ClientVersion = "GO-v1.0.7"
)

func WithClientAgentName(agentName string, mchID string) client.ClientOption {
	return clientOption{
		settings: client.DialSettings{
			AgentName: agentName + "-" + ClientVersion + "-" + mchID,
		},
	}
}
