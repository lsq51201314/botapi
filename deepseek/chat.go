package deepseek

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

func (d *DeepSeek) Chat(msg []Messages, reasoner ...bool) (think, messages string, err error) {
	//构建聊天
	obj := chat{
		Model:    "deepseek-chat",
		Messages: msg,
		Stream:   true,
	}
	if len(reasoner) > 0 && reasoner[0] {
		obj.Model = "deepseek-reasoner"
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
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	req.Header.Set("Authorization", "Bearer "+d.authorization)
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		switch resp.StatusCode {
		case 400: // 原因：请求体格式错误 解决方法：请根据错误信息提示修改请求体
			err = errors.New("格式错误")
			return
		case 401: // 原因：API key 错误，认证失败 解决方法：请检查您的 API key 是否正确，如没有 API key，请先 创建 API key
			err = errors.New("认证失败")
			return
		case 402: // 原因：账号余额不足 解决方法：请确认账户余额，并前往 充值 页面进行充值
			err = errors.New("余额不足")
			return
		case 422: // 原因：请求体参数错误 解决方法：请根据错误信息提示修改相关参数
			err = errors.New("参数错误")
			return
		case 429: // 原因：请求速率（TPM 或 RPM）达到上限 解决方法：请合理规划您的请求速率。
			err = errors.New("请求速率达到上限")
			return
		case 500: // 原因：服务器内部故障 解决方法：请等待后重试。若问题一直存在，请联系我们解决
			err = errors.New("服务器故障")
			return
		case 503: // 原因：服务器负载过高 解决方法：请稍后重试您的请求
			err = errors.New("服务器繁忙")
			return
		default:
			err = fmt.Errorf("未知状态码:%d", resp.StatusCode)
			return
		}
	}
	//处理消息
	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/event-stream; charset=utf-8" {
		err = fmt.Errorf("格式错误:%s", contentType)
		return
	}
	scanner := bufio.NewScanner(resp.Body)
	think = ""
	messages = ""
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
			think += v.Delta.ReasoningContent
			messages += v.Delta.Content
			if d.thinkcallback != nil {
				d.thinkcallback(stream.ID, v.Delta.ReasoningContent)
			}
			if d.messagescallback != nil {
				d.messagescallback(stream.ID, v.Delta.Content)
			}
			if v.FinishReason != "" {
				break
			}
		}
	}
	return
}
