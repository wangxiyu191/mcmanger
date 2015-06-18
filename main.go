package main

import (
	//"fmt"
	"github.com/astaxie/beego/config"
	//"github.com/go-martini/martini"
	"mcmanger/mc"
	"mcmanger/monitor"
	"mcmanger/qc"
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
	/*
		qcStatus, err := qc.Status()
		fmt.Println("QingCloud Instances:" + qcStatus)

		mcStatus, err := mc.Status()
		fmt.Printf("%d player online.\n", mcStatus.Players.Online)
		for _, player := range mcStatus.Players.Sample {
			fmt.Println(player.ID + "  " + player.Name)
		}

		err = qc.Boot()
		fmt.Println(err)
	*/
	monitor.Start(
		conf.String("Monitor::FirstDuration"),
		conf.String("Monitor::SecondDuration"),
	)
}
