package qc

import (
	"encoding/json"
	"github.com/magicshui/qingcloud-go"
	"github.com/magicshui/qingcloud-go/instance"
	//"log"
)

//青云操作

var (
	accessKeyId     string
	secretAccessKey string
	serverId        string
	zone            string
	client          *instance.Client
)

type descriM struct {
	Status string `json:"status"`
}

func Init(z string, id string, secret string, server string) {
	accessKeyId = id
	secretAccessKey = secret
	serverId = server
	zone = z
	client = instance.NewClient(zone, id, secret)
}

func Start() (err error) {
	_, err = client.StartInstances(qingcloud.Params{
		{"instances.1", serverId},
	})
	return err
}

func Stop() (err error) {
	_, err = client.StopInstances(qingcloud.Params{
		{"instances.1", serverId},
	})
	return err
}

func Status() (status string, err error) {
	result, err := client.DescribeInstances(qingcloud.Params{
		{"instances.1", serverId},
	})

	var response struct {
		Instances []descriM `json:"instance_set"`
	}

	err = json.Unmarshal(result, &response)
	if err != nil {
		return "", err
	}
	//log.Println(response.Count)
	//log.Println(len(response.Instances))

	return response.Instances[0].Status, err
}
