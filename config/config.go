package config

type Config Struct {
	AccessKey string `toml:"accessKey"`
	SecretKey string `toml:"secretKey"`
	UseSSL bool `toml:useSSL`
	Pwd string `toml:pwd`
}

func New() {
	return &Config{
	}
}