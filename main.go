package main

import (
	"fmt"

	"github.com/asomervell/gin-gorm-cloudsql-todo-app/Config"
	"github.com/asomervell/gin-gorm-cloudsql-todo-app/Models"
	"github.com/asomervell/gin-gorm-cloudsql-todo-app/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	DSN := Config.DbURL(Config.BuildDBConfig())
	Config.DB, err = gorm.Open("mysql", DSN)

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Todo{})

	r := Routes.SetupRouter()
	// running
	r.Run()
}
