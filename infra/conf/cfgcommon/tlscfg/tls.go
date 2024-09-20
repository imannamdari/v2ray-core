package tlscfg

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"

	"github.com/imannamdari/v2ray-core/v5/common/platform/filesystem"
	"github.com/imannamdari/v2ray-core/v5/infra/conf/cfgcommon"
	"github.com/imannamdari/v2ray-core/v5/transport/internet/tls"
)

//go:generate go run github.com/imannamdari/v2ray-core/v5/common/errors/errorgen

type TLSConfig struct {
	Insecure                         bool                  `json:"allowInsecure"`
	Certs                            []*TLSCertConfig      `json:"certificates"`
	ServerName                       string                `json:"serverName"`
	ALPN                             *cfgcommon.StringList `json:"alpn"`
	EnableSessionResumption          bool                  `json:"enableSessionResumption"`
	DisableSystemRoot                bool                  `json:"disableSystemRoot"`
	PinnedPeerCertificateChainSha256 *[]string             `json:"pinnedPeerCertificateChainSha256"`
	VerifyClientCertificate          bool                  `json:"verifyClientCertificate"`
	EnableEch                        bool                  `json:"enableEch"`
	EchSetting                       *TLSEchSetting        `json:"echSetting"`
}

// Build implements Buildable.
func (c *TLSConfig) Build() (proto.Message, error) {
	config := new(tls.Config)
	config.Certificate = make([]*tls.Certificate, len(c.Certs))
	for idx, certConf := range c.Certs {
		cert, err := certConf.Build()
		if err != nil {
			return nil, err
		}
		config.Certificate[idx] = cert
	}
	serverName := c.ServerName
	config.AllowInsecure = c.Insecure
	config.VerifyClientCertificate = c.VerifyClientCertificate
	if len(c.ServerName) > 0 {
		config.ServerName = serverName
	}
	if c.ALPN != nil && len(*c.ALPN) > 0 {
		config.NextProtocol = []string(*c.ALPN)
	}
	config.EnableSessionResumption = c.EnableSessionResumption
	config.DisableSystemRoot = c.DisableSystemRoot

	if c.PinnedPeerCertificateChainSha256 != nil {
		config.PinnedPeerCertificateChainSha256 = [][]byte{}
		for _, v := range *c.PinnedPeerCertificateChainSha256 {
			hashValue, err := base64.StdEncoding.DecodeString(v)
			if err != nil {
				return nil, err
			}
			config.PinnedPeerCertificateChainSha256 = append(config.PinnedPeerCertificateChainSha256, hashValue)
		}
	}

	config.EnableEch = c.EnableEch
	if c.EchSetting != nil {
		config.EchSetting, _ = c.EchSetting.Build()
	}

	// handle ech
	if config.EnableEch {
		ech, err := config.FetchECH()
		if err != nil {
			fmt.Println("failed to get first ech")
		} else {
			tls.ECH = ech
		}
		fmt.Printf("new ech = %s\n", ech)
		go config.UpdateECHEvery(15 * time.Minute)
	}

	return config, nil
}

type TLSEchSetting struct {
	DnsAddr    string `json:"dnsAddr"`
	InitEchKey string `json:"initEchKey"`
}

func (c *TLSEchSetting) Build() (*tls.ECHSetting, error) {
	setting := new(tls.ECHSetting)

	setting.DnsAddr = c.DnsAddr
	setting.InitEchKey = c.InitEchKey

	return setting, nil
}

type TLSCertConfig struct {
	CertFile string   `json:"certificateFile"`
	CertStr  []string `json:"certificate"`
	KeyFile  string   `json:"keyFile"`
	KeyStr   []string `json:"key"`
	Usage    string   `json:"usage"`
}

// Build implements Buildable.
func (c *TLSCertConfig) Build() (*tls.Certificate, error) {
	certificate := new(tls.Certificate)

	cert, err := readFileOrString(c.CertFile, c.CertStr)
	if err != nil {
		return nil, newError("failed to parse certificate").Base(err)
	}
	certificate.Certificate = cert

	if len(c.KeyFile) > 0 || len(c.KeyStr) > 0 {
		key, err := readFileOrString(c.KeyFile, c.KeyStr)
		if err != nil {
			return nil, newError("failed to parse key").Base(err)
		}
		certificate.Key = key
	}

	switch strings.ToLower(c.Usage) {
	case "encipherment":
		certificate.Usage = tls.Certificate_ENCIPHERMENT
	case "verify":
		certificate.Usage = tls.Certificate_AUTHORITY_VERIFY
	case "verifyclient":
		certificate.Usage = tls.Certificate_AUTHORITY_VERIFY_CLIENT
	case "issue":
		certificate.Usage = tls.Certificate_AUTHORITY_ISSUE
	default:
		certificate.Usage = tls.Certificate_ENCIPHERMENT
	}

	return certificate, nil
}

func readFileOrString(f string, s []string) ([]byte, error) {
	if len(f) > 0 {
		return filesystem.ReadFile(f)
	}
	if len(s) > 0 {
		return []byte(strings.Join(s, "\n")), nil
	}
	return nil, newError("both file and bytes are empty.")
}
