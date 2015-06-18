package monitor

import (
	"log"
	"mcmanger/mc"
	"mcmanger/qc"
	"time"
)

func getPlayerNum() (playerNum int, err error) {
	response, err := mc.Status()
	if err != nil {
		return 0, err
	}
	return response.Players.Online, nil
}

func parseDuration(durstr string) (duration time.Duration) {
	duration, err := time.ParseDuration(durstr)
	if err != nil {
		panic(err)
	}
	return duration
}

func Start(first string, second string) {
	firstDuration := parseDuration(first)
	secondDuration := parseDuration(first)

	timer := time.NewTimer(firstDuration)
	nonePlayerTimes := 0 //已经连续几次检测到服务器没有人
	for {
		<-timer.C
		timer.Reset(firstDuration) //循环监测

		//服务器是否开机
		qcStatus, err := qc.Status()
		if err != nil {
			log.Println("[Monitor]Error:" + err.Error())
			continue
		}
		if qcStatus != "running" {
			//log.Println("[Monitor]Server is not Running.")
			continue
		}

		playerNum, err := getPlayerNum()
		if err != nil {
			log.Println("[Monitor]Error:" + err.Error())
			continue
		}

		if playerNum == 0 {
			if nonePlayerTimes == 1 {
				log.Println("[Monitor]第二次发现无人在线，停机")
				nonePlayerTimes = 0
				qc.Stop()
			} else {
				log.Println("[Monitor]第一次发现无人在线，进入第二阶段监测")
				nonePlayerTimes++
				timer.Reset(secondDuration)
			}
		} else {
			log.Printf("[Monitor]有%d人在线", playerNum)
			nonePlayerTimes = 0
		}

	}

}
