package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string `mapstructure:"DB_HOST"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBTimeZone string `mapstructure:"DB_TIME_ZONE"`
	ApiPort    string `mapstructure:"API_PORT"`
}

func InitConfig() (Config, error) {
	cfg := new(Config)

	err := godotenv.Load(os.Getenv("CONFIG_FILE"))
	if err != nil {
		return *cfg, err
	}

	v := viper.New()

	if err = v.BindEnv("DB_HOST"); err != nil {
		return Config{}, err
	}
	if err = v.BindEnv("DB_NAME"); err != nil {
		return Config{}, err
	}
	if err = v.BindEnv("DB_USER"); err != nil {
		return Config{}, err
	}
	if err = v.BindEnv("DB_PASSWORD"); err != nil {
		return Config{}, err
	}
	if err = v.BindEnv("DB_PORT"); err != nil {
		return Config{}, err
	}
	if err = v.BindEnv("DB_TIME_ZONE"); err != nil {
		return Config{}, err
	}
	if err = v.BindEnv("API_PORT"); err != nil {
		return Config{}, err
	}

	if err := v.Unmarshal(cfg); err != nil {
		return *cfg, err
	}

	return *cfg, nil
}

func (c *Config) DBConnectionString() string {
	return "host=" + c.DBHost + " port=" + c.DBPort + " user=" + c.DBUser + " dbname=" + c.DBName + " password=" + c.DBPassword + " timezone=" + c.DBTimeZone + " sslmode=disable connect_timeout=10"
}
