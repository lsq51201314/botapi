package aliyun

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func (d *Aliyun) Chat(msg []Messages, reasoner ...bool) (messages string, err error) {
	//构建聊天
	obj := chat{
		Model:          d.model,
		Messages:       msg,
		Stream:         true,
		EnableThinking: false,
	}
	if len(reasoner) > 0 && reasoner[0] {
		obj.EnableThinking = true
	}
	var buf []byte
	if buf, err = json.Marshal(&obj); err != nil {
		return
	}
	//提交数据
	client := &http.Client{}
	if d.timeout > 0 {
		//超时
		client.Timeout = time.Duration(d.timeout) * time.Second
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", d.url, bytes.NewBuffer(buf)); err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+d.authorization)
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		var data []byte
		if data, err = io.ReadAll(resp.Body); err != nil {
			return
		}
		err = fmt.Errorf("错误状态(%d):%s", resp.StatusCode, string(data))
		return
	}
	//处理消息
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/event-stream;charset=UTF-8" {
		err = fmt.Errorf("格式错误:%s", contentType)
		return
	}
	scanner := bufio.NewScanner(resp.Body)
	messages = ""
	finish := false
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || len(line) <= 5 {
			continue
		}
		if line == "data: [DONE]" {
			break
		}
		line = line[5:]

		var stream ChatStream
		if err = json.Unmarshal([]byte(line), &stream); err != nil {
			return
		}

		for _, v := range stream.Choices {
			if v.Delta.Role == "assistant" { //肯定在第一个
				messages = "<think>"
				if d.callback != nil {
					d.callback(stream.ID, "<think>")
				}
			}
			if !finish && v.Delta.Content != "" && v.Delta.ReasoningContent == "" {
				finish = true
				messages += "</think>"
				if d.callback != nil {
					d.callback(stream.ID, "</think>")
				}
			}

			str := v.Delta.Content + v.Delta.ReasoningContent //必定有一个为空，直接相加
			messages += str
			if d.callback != nil {
				d.callback(stream.ID, str)
			}

			if v.FinishReason != "" {
				break
			}
		}
	}
	return
}
