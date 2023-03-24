package main

import (
	"CloudflareSpeedTest/task"
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

func main() {
	// 提示
	fmt.Println("使用之前建议关闭代理软件，否则测速可能不准确")
	task.Routines = 600
	task.InitRandSeed() // 置随机数种子
	//task.Disable = true
	task.IPFile = "./lib/CloudflareSpeedTest/ip.txt"
	task.TestCount = 35

	//utils.DisableClashProxy()

	// 开始延迟测速
	pingData := task.NewPing().Run().FilterDelay()
	// 开始下载测速
	speedData := task.TestDownloadSpeed(pingData)
	//speedData.Print()          // 打印结果
	//utils.ExportCsv(speedData) // 输出文件

	ipList := make([]string, speedData.Len())

	for key, data := range speedData {
		ipList[key] = data.IP.String()
	}

	// 读取yaml
	clashFile, err := ioutil.ReadFile("./clash.yaml")
	if err != nil {
		panic(err)
	}
	// 解析 yaml 文件
	var node yaml.Node
	err = yaml.Unmarshal(clashFile, &node)

	if err != nil {
		panic(err)
	}
	var proxiesKey int
	for key, childNode := range node.Content[0].Content {
		if childNode.Value == "proxies" {
			proxiesKey = key
		}
	}

	// 将clash.yaml 中的ip替换为测速后的ip
	if proxiesKey != 0 {
		proxiesNode := node.Content[0].Content[proxiesKey+1]
		for proxyKey, proxyNode := range proxiesNode.Content {
			for key, proxyChildNode := range proxyNode.Content {
				if proxyChildNode.Value == "server" {
					var serverKey = key + 1
					proxyNode.Content[serverKey].Value = ipList[proxyKey]
					break
				}
			}
		}

	}

	// 将更新后的配置写入 yaml 文件
	updatedYaml, err := yaml.Marshal(&node)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("clash.yaml", updatedYaml, 0644)
	if err != nil {
		panic(err)
	}

	//utils.ResetMode()
}
