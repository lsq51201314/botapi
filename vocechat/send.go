package vocechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cast"
)

func (b *Bot) SendText(gid int64, text string) (err error) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", b.url+"/api/bot/send_to_group/"+cast.ToString(gid), bytes.NewBuffer([]byte(text))); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-Key", b.key)
	req.Header.Set("Content-Type", "text/plain")
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		err = fmt.Errorf("错误状态:%d", resp.StatusCode)
		return
	}
	return
}

type file struct {
	Path string `json:"path"`
}

func (b *Bot) SendFile(gid int64, path string) (err error) {
	//构建数据
	obj := file{
		Path: path,
	}
	var buf []byte
	if buf, err = json.Marshal(&obj); err != nil {
		return
	}
	//提交数据
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", b.url+"/api/bot/send_to_group/"+cast.ToString(gid), bytes.NewBuffer(buf)); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-Key", b.key)
	req.Header.Set("Content-Type", "vocechat/file")
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		err = fmt.Errorf("错误状态:%d", resp.StatusCode)
		return
	}
	return
}
