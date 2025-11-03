package anythingllm

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func (a *AnythingLLM) Chat(msg string, reset ...bool) (message string, err error) {
	//构建数据
	obj := reqBody{
		Message: msg,
		Mode:    "chat",
		Stream:  true,
		Reset:   false,
	}
	if len(reset) > 0 {
		obj.Reset = reset[0]
	}
	var buf []byte
	if buf, err = json.Marshal(&obj); err != nil {
		return
	}
	//提交数据
	client := &http.Client{}
	if a.timeout > 0 {
		//超时
		client.Timeout = time.Duration(a.timeout) * time.Second
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", a.url, bytes.NewBuffer(buf)); err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+a.authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "text/event-stream")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Connection", "keep-alive")
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		err = fmt.Errorf("未知状态码:%d", resp.StatusCode)
		return
	}
	//处理消息
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/event-stream" {
		err = fmt.Errorf("格式错误:%s", contentType)
		return
	}
	scanner := bufio.NewScanner(resp.Body)
	message = ""
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || len(line) <= 5 {
			continue
		}
		line = line[5:]

		var info repInfo
		if err = json.Unmarshal([]byte(line), &info); err != nil {
			return
		}

		if info.Type == "textResponseChunk" {
			message += info.TextResponse
			if a.callback != nil {
				a.callback(info.UUID, info.TextResponse)
			}
		}
	}
	return
}
