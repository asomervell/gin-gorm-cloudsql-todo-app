package main

import (
	"fmt"

	"github.com/ektagarg/gin-gorm-todo-app/Config"
	"github.com/ektagarg/gin-gorm-todo-app/Models"
	"github.com/ektagarg/gin-gorm-todo-app/Routes"
	"github.com/jinzhu/gorm"
)

var err error

func main() {

	DSN := Config.DbURL(Config.BuildDBConfig())
	fmt.Println("DSN: ", DSN)
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
