package config

type Config struct {
	DataBase Database `yaml:"database"`
	JWT      JWT      `yaml:"jwt"`
	Qiniu    Qiniu    `yaml:"qiniu"`
}

type Database struct {
	Type     string `yaml:"type"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Name     string `yaml:"name"`
}

type JWT struct {
	SigningKey  string `yaml:"signingKey"`
	ExpiresTime int64  `yaml:"expiresTime"`
	BufferTime  int64  `yaml:"bufferTime"`
}

type Qiniu struct {
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}
