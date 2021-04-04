package main

import (
	"fmt"
	"os"

	"github.com/ektagarg/gin-gorm-todo-app/Config"
	"github.com/ektagarg/gin-gorm-todo-app/Models"
	"github.com/ektagarg/gin-gorm-todo-app/Routes"
	"github.com/jinzhu/gorm"
)

var err error

var (
	dbUser                 = os.Getenv("DB_USER")                  // e.g. 'my-db-user'
	dbPwd                  = os.Getenv("DB_PASS")                  // e.g. 'my-db-password'
	instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
	dbName                 = os.Getenv("DB_NAME")                  // e.g. 'my-database'
	socketDir              = "/cloudsql"
	dbURI                  string
)

func main() {

	dbURI = fmt.Sprintf("%s:%s@unix(/%s/%s)/%s?parseTime=true", dbUser, dbPwd, socketDir, instanceConnectionName, dbName)
	Config.DB, err = gorm.Open("mysql", dbURI)

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
