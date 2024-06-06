package bnrnethttp

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"net/http"
	"os"

	"github.com/go-resty/resty/v2"
)

type Client interface {
	Delete(url string, resp interface{}) (*ResponseMetadata, error)

	GetJSON(url string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PostJSON(url string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	DeleteJSON(url string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PutJSON(url string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PatchJSON(url string, req interface{}, resp interface{}) (*ResponseMetadata, error)

	GetJSONWithBearerAndHeaders(url string, bearerToken string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PostJSONWithBearerAndHeaders(url string, bearerToken string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	DeleteJSONWithBearerAndHeaders(url string, bearerToken string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PatchJSONWithBearerAndHeaders(url string, bearerToken string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PutJSONWithBearerAndHeaders(url string, bearerToken string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)

	PostJSONWithBasicAuthAndHeaders(url string, username, password string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)

	PostFormWithBasicAuthAndHeaders(url string, username, password string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)

	PostFormWithBearerAndHeaders(url string, bearerToken string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PostFormWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)

	PostUrlEncodedWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PostUrlEncodedWithHeadersXMLResponse(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)

	GetJSONWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PostJSONWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PutJSONWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	PatchJSONWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	DeleteJSONWithHeaders(url string, headers *map[string]string, req interface{}, resp interface{}) (*ResponseMetadata, error)
	DeleteWithHeaders(url string, headers *map[string]string, resp interface{}) (*ResponseMetadata, error)

	PostXMLStringWithBasicAuthKeyAndHeaders(url string, basicAuthenKey string, headers *map[string]string, xmlBodyReq string, resp interface{}) (*ResponseMetadata, error)

	PostMultipartFormDataWithAuthToken(url string, accessToken string, forms interface{}, files interface{}, resp interface{}) error

	CreateQueryParamString(model interface{}) string
}

type defaultClient struct {
	client *resty.Client
}

func NewClient() Client {

	cfg := &Config{
		EnableTLS: false,
	}

	c, _ := NewClientWithConfig(cfg)
	return c
}

func NewClientWithConfig(cfg *Config) (Client, error) {

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}

	if cfg.EnableTLS {
		caCertPool := x509.NewCertPool()
		for _, p := range cfg.CertificatePaths {
			caCert, err := os.ReadFile(p)
			if err != nil {
				return nil, err
			}
			block, _ := pem.Decode(caCert)
			if block == nil {
				return nil, fmt.Errorf("Decode ca file fail")
			}
			if block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
				return nil, fmt.Errorf("Decode ca block file fail")
			}

			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				return nil, fmt.Errorf("ParseCertificate ca block file fail")
			}

			caCertPool.AddCert(cert)
		}

		tr.TLSClientConfig = &tls.Config{
			RootCAs: caCertPool,
		}
	}

	if cfg.KeyFile != "" && cfg.CertificateFile != "" {
		cert, err := tls.LoadX509KeyPair(cfg.CertificateFile, cfg.KeyFile)
		if err != nil {
			return nil, err
		}
		tr.TLSClientConfig = &tls.Config{
			Certificates: []tls.Certificate{cert},
			// InsecureSkipVerify false means validate server certificate
			InsecureSkipVerify: false,
		}
	}

	hc := http.Client{
		Transport: tr,
	}

	client := resty.NewWithClient(&hc)

	return &defaultClient{
		client: client,
	}, nil
}
