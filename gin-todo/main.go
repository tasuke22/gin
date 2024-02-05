package main

import (
	"fmt"
	"github.com/tasuke/gin-todo/models"
	"github.com/tasuke/gin-todo/pkg/setting"
	"github.com/tasuke/gin-todo/routers"
)

func init() {
	setting.Setup("conf/development.ini")
	models.Setup()
}

func main() {
	fmt.Println(setting.DatabaseSetting.Host)
	r := routers.SetupRouter()
	r.Run(":8080")
}
