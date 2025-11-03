package vocechat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	"strings"
	"time"
)

type prepare struct {
	ContentType string `json:"content_type"`
	FileName    string `json:"filename"`
}

func (b *Bot) prepare(content_type, filename string) (fid string, err error) {
	//构建数据
	obj := prepare{
		ContentType: content_type,
		FileName:    filepath.Base(filename),
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
	if req, err = http.NewRequest("POST", b.url+"/api/bot/file/prepare", bytes.NewBuffer(buf)); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-Key", b.key)
	req.Header.Set("Content-Type", "application/json; charset=utf-8")
	var resp *http.Response
	if resp, err = client.Do(req); err != nil {
		return
	}
	defer resp.Body.Close()
	//处理状态
	if resp.StatusCode != 200 {
		if resp.StatusCode == 401 {
			err = errors.New("无效的凭证")
			return
		}
		err = fmt.Errorf("未知状态码:%d", resp.StatusCode)
		return
	}
	//处理数据
	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	fid = strings.ReplaceAll(string(body), `"`, "") //必须
	return
}
