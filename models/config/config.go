package config

type AppConfig struct {
	Server Server
	DB     DB
}

type Server struct {
	Port       int  `yaml:"port"`
	InRotation bool `yaml:"in_rotation"`
}

type DB struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}
