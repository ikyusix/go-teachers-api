package driver

import (
	"fmt"
	"net/url"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("driver/config.json")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Error(err)
	}
}

func ConnectionDB() *gorm.DB {
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.user")
	dbPass := viper.GetString("database.pass")
	dbName := viper.GetString("database.name")
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser,dbPass,dbHost,dbPort,dbName)
	val := url.Values{}
	val.Add("sslmode", "disable")
	connStr := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := gorm.Open("postgres", connStr)
	if err != nil {
		logrus.Fatal(err)
	}
	return dbConn
}
