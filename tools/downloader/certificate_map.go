package downloader

import (
	"context"
	"crypto/x509/pkix"
	"encoding/asn1"
	"fmt"
	"math/big"
	"net"

	"github.com/tjfoc/gmsm/x509"
)

// CertificateMap 最简单的证书获取器——证书Map
type CertificateMap struct {
	m      map[string]*x509.Certificate
	newest string
}

// Reset 完整重设 CertificateMap 中存储的证书，并重新选择最新的证书
func (m *CertificateMap) Reset(newCertificates map[string]*x509.Certificate) {
	var (
		newestCert     *x509.Certificate
		newestSerialNo string
	)

	m.m = make(map[string]*x509.Certificate)

	for serialNo, cert := range newCertificates {
		m.m[serialNo] = cert
		if newestSerialNo == "" || newestCert == nil || cert.NotBefore.After(newestCert.NotBefore) {
			newestSerialNo = serialNo
			newestCert = cert
		}
	}

	m.newest = newestSerialNo
}

// Get 获取证书序列号对应的平台证书
func (m *CertificateMap) Get(_ context.Context, serialNumber string) (*x509.Certificate, bool) {
	cert, ok := m.m[serialNumber]
	if !ok {
		return nil, false
	}
	return cloneCertificate(cert), true
}

// GetAll 获取平台证书Map
func (m *CertificateMap) GetAll(_ context.Context) map[string]*x509.Certificate {
	ret := make(map[string]*x509.Certificate)

	for serialNo, cert := range m.m {
		ret[serialNo] = cloneCertificate(cert)
	}

	return ret
}

// GetNewestSerial 获取最新的平台证书的证书序列号
func (m *CertificateMap) GetNewestSerial(_ context.Context) string {
	return m.newest
}

// NewCertificateMap 使用 证书序列号->证书 映射 初始化 CertificateMap
func NewCertificateMap(certificateMap map[string]*x509.Certificate) *CertificateMap {
	m := CertificateMap{}
	m.Reset(certificateMap)
	return &m
}

// NewCertificateMapWithList 使用 证书列表 初始化 CertificateMap
func NewCertificateMapWithList(l []*x509.Certificate) *CertificateMap {
	m := make(map[string]*x509.Certificate)
	for _, c := range l {
		m[GetCertificateSerialNumber(*c)] = c
	}
	return NewCertificateMap(m)
}

// GetCertificateSerialNumber 从证书中获取证书序列号
func GetCertificateSerialNumber(certificate x509.Certificate) string {
	return fmt.Sprintf("%X", certificate.SerialNumber.Bytes())
}

func cloneCertificate(cert *x509.Certificate) *x509.Certificate {
	if cert == nil {
		return nil
	}
	copied := *cert
	copied.SerialNumber = cloneBigInt(cert.SerialNumber)
	copied.Raw = cloneBytes(cert.Raw)
	copied.RawTBSCertificate = cloneBytes(cert.RawTBSCertificate)
	copied.RawSubjectPublicKeyInfo = cloneBytes(cert.RawSubjectPublicKeyInfo)
	copied.RawSubject = cloneBytes(cert.RawSubject)
	copied.RawIssuer = cloneBytes(cert.RawIssuer)
	copied.Signature = cloneBytes(cert.Signature)
	copied.SubjectKeyId = cloneBytes(cert.SubjectKeyId)
	copied.AuthorityKeyId = cloneBytes(cert.AuthorityKeyId)
	copied.Extensions = cloneExtensions(cert.Extensions)
	copied.ExtraExtensions = cloneExtensions(cert.ExtraExtensions)
	copied.UnhandledCriticalExtensions = cloneOIDs(cert.UnhandledCriticalExtensions)
	copied.ExtKeyUsage = append([]x509.ExtKeyUsage(nil), cert.ExtKeyUsage...)
	copied.UnknownExtKeyUsage = cloneOIDs(cert.UnknownExtKeyUsage)
	copied.OCSPServer = append([]string(nil), cert.OCSPServer...)
	copied.IssuingCertificateURL = append([]string(nil), cert.IssuingCertificateURL...)
	copied.DNSNames = append([]string(nil), cert.DNSNames...)
	copied.EmailAddresses = append([]string(nil), cert.EmailAddresses...)
	copied.IPAddresses = cloneIPs(cert.IPAddresses)
	copied.PermittedDNSDomains = append([]string(nil), cert.PermittedDNSDomains...)
	copied.CRLDistributionPoints = append([]string(nil), cert.CRLDistributionPoints...)
	copied.PolicyIdentifiers = cloneOIDs(cert.PolicyIdentifiers)
	return &copied
}

func cloneBytes(in []byte) []byte {
	return append([]byte(nil), in...)
}

func cloneBigInt(in *big.Int) *big.Int {
	if in == nil {
		return nil
	}
	return new(big.Int).Set(in)
}

func cloneExtensions(in []pkix.Extension) []pkix.Extension {
	out := append([]pkix.Extension(nil), in...)
	for i := range out {
		out[i].Id = append(asn1.ObjectIdentifier(nil), in[i].Id...)
		out[i].Value = cloneBytes(in[i].Value)
	}
	return out
}

func cloneOIDs(in []asn1.ObjectIdentifier) []asn1.ObjectIdentifier {
	out := append([]asn1.ObjectIdentifier(nil), in...)
	for i := range out {
		out[i] = append(asn1.ObjectIdentifier(nil), in[i]...)
	}
	return out
}

func cloneIPs(in []net.IP) []net.IP {
	out := append([]net.IP(nil), in...)
	for i := range out {
		out[i] = append(net.IP(nil), in[i]...)
	}
	return out
}
