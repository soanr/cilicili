package main

import (
	"cilicili-go/api/email"
	"cilicili-go/client"
	"cilicili-go/conf"
	_ "cilicili-go/conf"
	"github.com/kataras/iris/v12"
	"os"
	"strconv"
)

func main() {

	//获取启动端口
	port := os.Args[len(os.Args)-1]

	i, err := strconv.Atoi(port)
	if err != nil {
		panic(err)
	}

	//注册到eureka注册中心
	client.RegisterToEureka(conf.Config.Eureka.Address, i, "email-service")

	app := iris.New()
	email.RegisterAllRouter(app)
	err = app.Run(iris.Addr(":" + port))
	if err != nil {
		panic(err)
	}

}
