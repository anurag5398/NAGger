package config

type AppConfig struct {
	Server Server
}

type Server struct {
	Port       int  `yaml:"port"`
	InRotation bool `yaml:"in_rotation"`
}
