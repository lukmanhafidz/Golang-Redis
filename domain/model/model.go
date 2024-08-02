package model

type RedisConfig struct {
	Host string `env:"host"`
	Port string `env:"port"`
}
