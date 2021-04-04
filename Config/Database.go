package Config

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

// DBConfig represents db configuration
type DBConfig struct {
	User                   string
	Password               string
	InstanceConnectionName string
	DBName                 string
	SocketDir              string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		User:                   os.Getenv("DB_USER"),
		Password:               os.Getenv("DB_PASS"),
		InstanceConnectionName: os.Getenv("INSTANCE_CONNECTION_NAME"),
		DBName:                 os.Getenv("DB_NAME"),
		SocketDir:              "cloudsql",
	}
	return &dbConfig
}

func DbURL(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@unix(/%s/%s)/%s?parseTime=true",
		dbConfig.User,
		dbConfig.Password,
		dbConfig.SocketDir,
		dbConfig.InstanceConnectionName,
		dbConfig.DBName,
	)
}
