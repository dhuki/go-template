package configloader

import "time"

var (
	Conf Config

	searchPath = []string{
		"$HOME/.go-rest-template",
		"$GOPATH/src/github/go-rest-template",
		".",
		"/app",
	}

	configName = map[string]string{
		"LOCAL": "config/config.local.yaml",
		"DEV":   "config/config.dev.yaml",
		"UAT":   "config/config.uat.yaml",
		"PROD":  "config/config.prod.yaml",
	}
)

type Config struct {
	App          Application    `mapstructure:"app"`
	ConnDatabase DatabaseConfig `mapstructure:"postgres"`
	Redis        RedisConfig    `mapstructure:"redis"`
}

type Application struct {
	Env       string        `mapstructure:"env"`
	Name      string        `mapstructure:"name"`
	Port      int           `mapstructure:"port"`
	LogFormat string        `mapstructure:"logFormat"`
	Timeout   time.Duration `mapstructure:"timeout"`
}

type DatabaseConfig struct {
	DbConnectionInfo
	Slave  DBInfo `mapstructure:"slave"`
	Master DBInfo `mapstructure:"master"`
}

type DBInfo struct {
	DBName   string `mapstructure:"dbName"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
	Port     int    `mapstructure:"port"`
	Schema   string `mapstructure:"schema"`
	User     string `mapstructure:"user"`
	Debug    bool   `mapstructure:"debug"`
}

type DbConnectionInfo struct {
	SetMaxIdleCons    int `mapstructure:"maxIdleConnections"`
	SetMaxOpenCons    int `mapstructure:"maxOpenConnections"`
	SetConMaxIdleTime int `mapstructure:"setConMaxIdleTime"`
	SetConMaxLifetime int `mapstructure:"connectTimeout"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	DB       int    `mapstructure:"db"`
	Password string `mapstructure:"password"`
}
