package web

import (
	"fmt"
	"log"
	"mcmanger/mc"
	"mcmanger/qc"
)

func startHandler() (int, string) {
	qcStatus, err := qc.Status()
	if err != nil {
		log.Println(err)
		return 500, "获取青云实例状态发生错误。"
	}
	if qcStatus == "running" {
		return 500, "服务器正在运行。"
	}
	err = qc.Start()
	if err != nil {
		log.Println(err)
		return 500, "启动青云实例发生错误"
	}
	return 200, "服务器正在启动……Enjoy!"
}
func statusHandler() (int, string) {
	var response string

	qcStatus, err := qc.Status()
	if err != nil {
		log.Println(err)
		return 500, "获取青云实例状态发生错误。"
	}
	response += fmt.Sprintln("青云实例状态：" + qcStatus)
	if qcStatus == "running" {
		mcStatus, err := mc.Status()
		if err != nil {
			log.Println(err)
			return 500, "获取Minecraft服务器状态发生错误。"
		}
		response += fmt.Sprintf("%d人正在进行游戏:\n", mcStatus.Players.Online)
		for _, player := range mcStatus.Players.Sample {
			response += fmt.Sprintln(player.Name)
		}
	}
	return 200, response

}
