package config

import (
	"errors"
	"log"
	"os"
"time"
	"github.com/spf13/viper"
)
type CorsConfig struct {
	AllowedOrigin string
}

type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Redis    RedisConfig
}

type ServerConfig struct {
	Port    string
	runmode string
}

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DbName   string
	SSLMode string //disable ssl for local development
	Maxidleconnections int
	Maxopenconnections int
	Connmaxtimeout time.Duration
}
type RedisConfig struct {
	Host               string
	Port               string
	User               string
	Password           string
	DbName             string
	SSLMode            bool
	minidleconnections int
	poolsize           int
	pooltimeout        int
	dialTimeout time.Duration
    readTimeout  time.Duration
    writeTimeout time.Duration
	Idlecheckfrequency time.Duration
}
func GetConfig() *Config{
	cfgPath := GetConfigPath(os.Getenv("APP_ENV"))
	v,err :=LoadConfig(cfgPath,"yml")
	if err != nil{
		log.Fatalf("Error in load config %v",err)
	}
	cfg, err := ParseConfig(v)

	if err != nil{
		log.Fatalf("Error in parse config %v",err)
	}
	return cfg
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var cfg Config
	err := v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("cant parse the config:%v", err)
		return nil, err
	}
return &cfg, nil
}
func LoadConfig(filename string, fileType string) (*viper.Viper, error) {
    v := viper.New()
    v.SetConfigType(fileType)
    v.SetConfigFile(filename) 
    v.AutomaticEnv()

    err := v.ReadInConfig()
    if err != nil {
        if _, ok := err.(viper.ConfigFileNotFoundError); ok {
            return nil, errors.New("config file not found")
        }
        return nil, err
    }
    return v, nil
}

// get the file address
func GetConfigPath(env string) string {
	base := "/home/melika/projects/bookstore-project/src/api/config/"

	switch env {
	case "docker":
		return base + "config-docker.yml"
	case "production":
		return base + "config-production.yml"
	default:
		return base + "config-development.yml"
	}
}
