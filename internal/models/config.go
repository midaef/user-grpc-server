package models

// Config ...
type Config struct {
	Server 	 Server `yaml:"server"`
	Database Database `yaml:"database"`
	Token    Token `yaml:"token"`
	Logger   Logger `yaml:"logger"`
}

// Server ...
type Server struct {
	Port string `yaml:"port"`
}

// Database ...
type Database struct {
	Uri string `yaml:"uri"`
}

// Token ...
type Token struct {
	SecretKey string `yaml:"secret_key"`
	PublicKey string `yaml:"public_key"`
	Salt      string `yaml:"salt"`
}

// Logger ...
type Logger struct {
	LogLevel  string `yaml:"log_level"`
}