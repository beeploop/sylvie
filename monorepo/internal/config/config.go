package config

type Config interface {
	Load() Config
}
