package downloader

import (
	"context"
	"crypto/rsa"
	"sync"
	"time"

	"github.com/douyinpay/douyinpay-go/client"
	"github.com/douyinpay/douyinpay-go/tools/auth/signers"
	"github.com/douyinpay/douyinpay-go/tools/auth/verifiers"
	"github.com/douyinpay/douyinpay-go/tools/consts"
	"github.com/douyinpay/douyinpay-go/utils"

	"github.com/tjfoc/gmsm/sm2"
)

const (
	// DefaultDownloadInterval 默认微信支付平台证书更新间隔
	DefaultDownloadInterval = 24 * time.Hour
)

// CertificateDownloaderMgr 证书下载器管理器
// 可挂载证书下载器 CertificateDownloader，会定时调用 CertificateDownloader 下载最新的证书
//
// CertificateDownloaderMgr 不会被 GoGC 自动回收，不再使用时应调用 Stop 方法，防止发生资源泄漏
type CertificateDownloaderMgr struct {
	ctx           context.Context
	task          *utils.RepeatedTask
	downloaderMap map[string]*CertificateDownloader
	lock          sync.RWMutex
}

// GetCertificateVisitor 获取某个商户的平台证书访问器
func (mgr *CertificateDownloaderMgr) GetCertificateVisitor(mchID string) client.CertificateVisitor {
	// 访问使用读锁
	mgr.lock.RLock()
	defer mgr.lock.RUnlock()
	return mgr.downloaderMap[mchID]
}

// Stop 停止 CertificateDownloaderMgr 的自动下载 Goroutine
// 当且仅当不再需要当前管理器自动下载后调用
// 一旦调用成功，当前管理器无法再次启动
func (mgr *CertificateDownloaderMgr) Stop() {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.task.Stop()
}

func (mgr *CertificateDownloaderMgr) getTickHandler() func(time.Time) {
	return func(time.Time) {
		mgr.DownloadCertificates(mgr.ctx)
	}
}

// DownloadCertificates 让所有已注册下载器均进行一次下载
func (mgr *CertificateDownloaderMgr) DownloadCertificates(ctx context.Context) {
	tmpDownloaderMap := make(map[string]*CertificateDownloader)

	mgr.lock.RLock()
	for key, downloader := range mgr.downloaderMap {
		tmpDownloaderMap[key] = downloader
	}
	mgr.lock.RUnlock()

	for _, downloader := range tmpDownloaderMap {
		_ = downloader.DownloadCertificates(ctx)
	}
}

// RemoveDownloader 移除商户的平台证书下载器
// 移除后从 GetCertificateVisitor 接口获得的对应商户的 CertificateVisitor 将会失效，
// 请确认不再需要该商户的证书后再行移除，如果下载器存在，本接口将会返回该下载器。
func (mgr *CertificateDownloaderMgr) RemoveDownloader(_ context.Context, mchID string) *CertificateDownloader {
	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	downloader, ok := mgr.downloaderMap[mchID]
	if !ok {
		return nil
	}

	delete(mgr.downloaderMap, mchID)
	return downloader
}

// HasDownloader 检查是否已经注册过 mchID 这个商户的下载器
func (mgr *CertificateDownloaderMgr) HasDownloader(_ context.Context, mchID string) bool {
	mgr.lock.RLock()
	defer mgr.lock.RUnlock()

	_, ok := mgr.downloaderMap[mchID]
	return ok
}

// RegisterRSADownloaderWithPrivateKey 向 Mgr 注册商户的RSA平台证书下载器
func (mgr *CertificateDownloaderMgr) RegisterRSADownloaderWithPrivateKey(
	ctx context.Context, mchID string,
	mchCertificateSerialNo string, mchPrimaryKey *rsa.PrivateKey, encryptKey string,
) error {
	downloader, err := NewRSACertificateDownloader(ctx, mchID, mchCertificateSerialNo, mchPrimaryKey, encryptKey)
	if err != nil {
		return err
	}
	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// 重新设置一下
	mgr.downloaderMap[mchID] = downloader
	return nil
}

// RegisterSM2DownloaderWithPrivateKey 向 Mgr 注册商户的SM2平台证书下载器
func (mgr *CertificateDownloaderMgr) RegisterSM2DownloaderWithPrivateKey(
	ctx context.Context, mchID string,
	mchCertificateSerialNo string, mchPrimaryKey *sm2.PrivateKey, encryptKey string,
) error {
	downloader, err := NewSM2CertificateDownloader(ctx, mchID, mchCertificateSerialNo, mchPrimaryKey, encryptKey)
	if err != nil {
		return err
	}

	mgr.lock.Lock()
	defer mgr.lock.Unlock()
	// 重新设置一下
	mgr.downloaderMap[mchID] = downloader
	return nil
}

// RegisterDownloaderWithClient 向 Mgr 注册商户的平台证书下载器
func (mgr *CertificateDownloaderMgr) RegisterDownloaderWithClient(
	ctx context.Context, client *client.Client, mchID string, signType string, encryptKey string,
) error {

	downloader, err := NewCertificateDownloaderWithClient(ctx, client, signType, encryptKey)
	if err != nil {
		return err
	}

	mgr.lock.Lock()
	defer mgr.lock.Unlock()

	mgr.downloaderMap[mchID] = downloader
	return nil
}

// NewRSACertificateDownloader 使用商户号/商户私钥等信息初始化商户的RSA平台证书下载器 CertificateDownloader
// 初始化完成后会立即发起一次下载，确保下载器被正确初始化。
func NewRSACertificateDownloader(
	ctx context.Context, mchID string, mchCertificateSerialNo string, mchPrivateKey *rsa.PrivateKey, encryptKey string,
) (*CertificateDownloader, error) {
	client, err := client.NewClientWithDialSettings(ctx,
		&client.DialSettings{
			Signer: &signers.SHA256WithRSASigner{
				MchID:               mchID,
				CertificateSerialNo: mchCertificateSerialNo,
				PrivateKey:          mchPrivateKey,
			},
			Verifier: &verifiers.NilVerifier{},
		},
	)
	if err != nil {
		return nil, err
	}
	return NewCertificateDownloaderWithClient(ctx, client, consts.CRYPTO_TYPE_RSA, encryptKey)
}

// NewSM2CertificateDownloader 使用商户号/商户私钥等信息初始化商户的SM2平台证书下载器 CertificateDownloader
// 初始化完成后会立即发起一次下载，确保下载器被正确初始化。
func NewSM2CertificateDownloader(
	ctx context.Context, mchID string, mchCertificateSerialNo string, mchPrivateKey *sm2.PrivateKey, encryptKey string,
) (*CertificateDownloader, error) {
	client, err := client.NewClientWithDialSettings(ctx,
		&client.DialSettings{
			Signer: &signers.Sm2Signer{
				MchID:               mchID,
				CertificateSerialNo: mchCertificateSerialNo,
				PrivateKey:          mchPrivateKey,
			},
			Verifier: &verifiers.NilVerifier{},
		},
	)
	if err != nil {
		return nil, err
	}
	return NewCertificateDownloaderWithClient(ctx, client, consts.CRYPTO_TYPE_SM2, encryptKey)
}

// NewCertificateDownloaderWithClient 使用 core.Client 初始化商户的平台证书下载器 CertificateDownloader
// 初始化完成后会立即发起一次下载，确保下载器被正确初始化。
func NewCertificateDownloaderWithClient(
	ctx context.Context, client *client.Client, signType string, encryptKey string,
) (*CertificateDownloader, error) {
	d := CertificateDownloader{
		client:       client,
		encryptKey:   encryptKey,
		signType:     signType,
		certificates: NewCertificateMap(nil),
	}
	if err := d.DownloadCertificates(ctx); err != nil {
		return nil, err
	}
	return &d, nil
}

// NewCertificateDownloaderMgr 以默认间隔 DefaultDownloadInterval 创建证书下载管理器
// 该管理器将以 DefaultDownloadInterval 的间隔定期调度所有 Downloader 进行证书下载。
// 证书管理器一旦创建即启动，使用完毕请调用 Stop() 防止发生资源泄漏
func NewCertificateDownloaderMgr(ctx context.Context) *CertificateDownloaderMgr {
	return NewCertificateDownloaderMgrWithInterval(ctx, DefaultDownloadInterval)
}

// NewCertificateDownloaderMgrWithInterval 创建一个空证书下载管理器（自定义更新间隔）
//
// 更新间隔最大不建议超过 2 天，以免错过平台证书平滑切换窗口；
// 同时亦不建议小于 1 小时，以避免过多请求导致浪费
func NewCertificateDownloaderMgrWithInterval(
	ctx context.Context, downloadInterval time.Duration,
) *CertificateDownloaderMgr {
	if downloadInterval <= 0 {
		downloadInterval = DefaultDownloadInterval
	}

	downloader := CertificateDownloaderMgr{
		ctx:           ctx,
		downloaderMap: make(map[string]*CertificateDownloader),
	}
	downloader.task = utils.NewRepeatedTask(downloadInterval, downloader.getTickHandler())
	downloader.task.Start()
	return &downloader
}
