package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	ClashControlApi = "http://127.0.0.1:9090"
)

var (
	client = &http.Client{
		Timeout: time.Second * 5,
	}

	originModel *Mode = nil
)

type ClashConfig struct {
	Port           int           `json:"port"`
	SocksPort      int           `json:"socks-port"`
	RedirPort      int           `json:"redir-port"`
	TproxyPort     int           `json:"tproxy-port"`
	MixedPort      int           `json:"mixed-port"`
	Authentication []interface{} `json:"authentication"`
	AllowLan       bool          `json:"allow-lan"`
	BindAddress    string        `json:"bind-address"`
	Mode           Mode          `json:"mode"`
	LogLevel       string        `json:"log-level"`
	Ipv6           bool          `json:"ipv6"`
}

type Mode string

const (
	ModeRule   Mode = "rule"
	ModeGlobal Mode = "global"
	ModeDirect Mode = "direct"
	ModeScript Mode = "script"
)

func check() bool {
	// 测试api是否可用
	_, err := client.Get(ClashControlApi)
	if err != nil {
		return false
	}
	return true
}

func DisableClashProxy() {
	if !check() {
		return
	}

	resp, err := client.Get(fmt.Sprintf("%s/conigs", ClashControlApi))
	if err != nil {
		log.Println(err)
		return
	}

	defer resp.Body.Close()

	// 读取resp body 至 ClashConfig
	var config ClashConfig
	err = json.NewDecoder(resp.Body).Decode(&config)
	if err != nil {
		log.Println(err)
		return
	}
	originModel = &config.Mode

	// 将代理规则设置为直连
	UpdateMode(ModeDirect)
}

func UpdateMode(newMode Mode) {
	request, err := http.NewRequest(http.MethodPut, fmt.Sprintf("%s/configs", ClashControlApi), strings.NewReader(fmt.Sprintf(`{"mode":"%s"}`, newMode)))
	if err != nil {
		log.Println(err)
		return
	}
	_, err = client.Do(request)
	if err != nil {
		log.Println(err)
		return
	}
}

func ResetMode() {
	if originModel != nil {
		UpdateMode(*originModel)
	}
}
