package server

type TLSConfig struct {
	CertFile string
	KeyFile  string
}

type Config struct {
	Address string `default:"0.0.0.0"`
	Port    string `default:"3000"`
	TLS     TLSConfig
}
