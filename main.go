package main

import (
	//"fmt"
	"github.com/astaxie/beego/config"
	//"github.com/go-martini/martini"
	"mcmanger/mc"
	"mcmanger/monitor"
	"mcmanger/qc"
	"mcmanger/web"
)

func main() {
	//ini配置信息
	conf, err := config.NewConfig("ini", "config.ini")
	if err != nil {
		panic(err)
	}
	qc.Init(
		conf.String("QingCloud::Zone"),
		conf.String("QingCloud::AccessKeyId"),
		conf.String("QingCloud::SecretAccessKey"),
		conf.String("QingCloud::ServerId"),
	)
	mc.Init(conf.String("Minecraft::Address"))

	go monitor.Start(
		conf.String("Monitor::FirstDuration"),
		conf.String("Monitor::SecondDuration"),
	)
	web.Start(conf.DefaultInt("Web::Port", 8123))
}
