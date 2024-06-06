package bnrnethttp

type Config struct {
	EnableTLS        bool     `json:"enable_tls" mapstructure:"enable_tls"`
	CertificatePaths []string `json:"certificate_paths" mapstructure:"certificate_paths"`
	KeyFile          string   `json:"key_file" mapstructure:"key_file"`
	CertificateFile  string   `json:"certificate_file" mapstructure:"certificate_file"`
}
