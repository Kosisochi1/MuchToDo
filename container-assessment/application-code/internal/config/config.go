package config

import (
	"os"
	"strconv"
	"github.com/spf13/viper"
)

// Config stores all configuration of the application.
type Config struct {
	ServerPort       string `mapstructure:"PORT"`
	MongoURI         string `mapstructure:"MONGO_URI"`
	DBName           string `mapstructure:"DB_NAME"`
	JWTSecretKey     string `mapstructure:"JWT_SECRET_KEY"`
	JWTExpirationHours int    `mapstructure:"JWT_EXPIRATION_HOURS"`
	EnableCache      bool   `mapstructure:"ENABLE_CACHE"`
	RedisAddr        string `mapstructure:"REDIS_ADDR"`
	RedisPassword    string `mapstructure:"REDIS_PASSWORD"`
	LogLevel      string `mapstructure:"LOG_LEVEL"`
	LogFormat     string `mapstructure:"LOG_FORMAT"`
}

// LoadConfig reads configuration from file or environment variables.
func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Set default values
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("ENABLE_CACHE", false)
	viper.SetDefault("JWT_EXPIRATION_HOURS", 72)

	// err = viper.ReadInConfig()
	// if err != nil {
	// 	if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
	// 		return
	// 	}
	// }
	err = viper.ReadInConfig()
    if err != nil {
	 if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
		// Fail ONLY if file exists but is invalid
		return config, err
	    }
	// If .env is missing, ignore and rely on environment variables
    }


	err = viper.Unmarshal(&config)
	if err != nil {
		return
	}

	// Force override from environment (Docker guarantees these exist)
	if mongoURI := os.Getenv("MONGO_URI"); mongoURI != "" {
		config.MongoURI = mongoURI
	}
	if dbName := os.Getenv("DB_NAME"); dbName != "" {
		config.DBName = dbName
	}
	if port := os.Getenv("PORT"); port != "" {
		config.ServerPort = port
	}
	if jwtKey := os.Getenv("JWT_SECRET_KEY"); jwtKey != "" {
		config.JWTSecretKey = jwtKey
	}
	if jwtExp := os.Getenv("JWT_EXPIRATION_HOURS"); jwtExp != "" {
		if v, err := strconv.Atoi(jwtExp); err == nil {
			config.JWTExpirationHours = v
		}
	}
	if enableCache := os.Getenv("ENABLE_CACHE"); enableCache != "" {
		config.EnableCache = enableCache == "true"
	}
	if logLevel := os.Getenv("LOG_LEVEL"); logLevel != "" {
		config.LogLevel = logLevel
	}
	if logFormat := os.Getenv("LOG_FORMAT"); logFormat != "" {
		config.LogFormat = logFormat
	}

	return




	
}

