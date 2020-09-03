package config

type Config struct {
	AccessKey string `toml:"accessKey"`
	SecretKey string `toml:"secretKey"`
	UseSSL    bool   `toml:"useSSL"`
	Pwd       string `toml:"pwd"`
	EndPoint  string `toml:"endPoint"`
}

func New() *Config {
	return &Config{}
}
