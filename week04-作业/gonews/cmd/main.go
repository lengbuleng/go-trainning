package main

import (
	"app/api"
	"app/internal/dao"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {

	dao.Connect()

	g := gin.Default()
	api.InitRouter(g)
	g.Static("/static", "./static")

	err := g.Run(":1688")
	if err != nil {
		fmt.Println("服务启动失败" + err.Error())
	}

}
