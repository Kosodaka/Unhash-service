package entity

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	HTTP
	HashConf
}

type HTTP struct {
	Host string `env:"HTTP_HOST" envDefault:"localhost"`
	Port string `env:"HTTP_PORT" envDefault:"8080"`
}

func (c *Config) GetHTTPPort() string {
	return c.HTTP.Port
}

func (c *Config) GetHTTPHost() string {
	return c.HTTP.Host
}

type HashConf struct {
	HMACSecret string `env:"HMAC_SECRET" envDefault:"-"`
}

func (c *Config) GetHMACSecret() string {
	return c.HashConf.HMACSecret
}

func NewConfig(cgfPath string) (*Config, error) {
	cfg := &Config{}

	err := cleanenv.ReadConfig(cgfPath, cfg)
	if err != nil {
		return nil, err
	}

	err = cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
