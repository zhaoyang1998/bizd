package utils

import (
	"bizd/metion/model"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type RequestParam struct {
	Msgtype  string   `json:"msgtype"`
	Markdown MarkDown `json:"markdown"`
}

type MarkDown struct {
	Content string `json:"content"`
}

func SendWxMessage(contentInfo string, url string) {
	mark := MarkDown{
		Content: contentInfo,
	}
	reqParam := RequestParam{
		Msgtype:  "markdown",
		Markdown: mark,
	}
	data, _ := json.Marshal(reqParam)
	res, err := http.Post(url, "application/json", bytes.NewBuffer(data))
	if err != nil {
		log.Print("信息发送失败，内容：" + contentInfo + "机器人地址：" + url + "错误信息：" + url)
	}
	log.Print("信息发送成功，内容：" + contentInfo + "机器人地址：" + url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
}

func SendConductorMessage(conductorInfo model.Conductor) {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=91a2ca6f-3dd3-4b15-bd15-939d263979ed"
	info := "开周会啦！ 这周轮到你主持" + "<font color='#dd0000'>" + conductorInfo.Username + "</font>\n" + "<@" + conductorInfo.WxName + ">"
	SendWxMessage(info, url)
}

func SendTimeNotice(MsgFromCron model.MsgFromCron) {
	url := MsgFromCron.Receive
	info := "test"
	SendWxMessage(info, url)
}
