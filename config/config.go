package config

import (
	"log"
	"tansan/monitoring/metricsModel"
	"tansan/user/userModel"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadConfig() {
    viper.SetConfigName("config")
    viper.SetConfigType("json")
    viper.AddConfigPath(".")
    viper.AutomaticEnv()  // Add this line

    if err := viper.ReadInConfig(); err != nil {
        log.Fatalf("Error reading config file, %s", err)
    }

    initDatabase()
}

func initDatabase() {
    dsn := "host=localhost user=" + viper.GetString("POSTGRES_USER") +
           " password=" + viper.GetString("POSTGRES_PASSWORD") +
           " dbname=" + viper.GetString("POSTGRES_DB") +
           " port=5432 sslmode=disable TimeZone=Asia/Bangkok"
    var err error
    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database, %s", err)
    }

    // Auto migrate the User model
    DB.AutoMigrate(&userModel.User{})
    DB.AutoMigrate(&metricsModel.Metrics{})
}
