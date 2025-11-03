package vocechat

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type upload struct {
	Path string `json:"path"`
}

func (b *Bot) Upload(data []byte, filename string) (path string, err error) {
	content_type := mime.TypeByExtension(filepath.Ext(filename))
	if content_type == "" {
		content_type = "application/octet-stream"
	}
	var fid string
	if fid, err = b.prepare(content_type, filename); err != nil {
		return
	}
	//准备表单
	var requestBody bytes.Buffer
	writer := multipart.NewWriter(&requestBody)
	if err = writer.WriteField("file_id", fid); err != nil {
		return
	}
	var file *os.File
	if file, err = os.Open(filename); err != nil {
		return
	}
	defer file.Close()
	var part io.Writer
	if part, err = writer.CreateFormFile("chunk_data", filepath.Base(filename)); err != nil {
		return
	}
	if _, err = io.Copy(part, file); err != nil {
		return
	}
	if err = writer.WriteField("chunk_is_last", "true"); err != nil {
		return
	}
	if err = writer.Close(); err != nil {
		return
	}
	//提交数据
	client := &http.Client{
		Timeout: 10 * time.Second,
	}
	var req *http.Request
	if req, err = http.NewRequest("POST", b.url+"/api/bot/file/upload", &requestBody); err != nil {
		return
	}
	req.Header.Set("accept", "application/json; charset=utf-8")
	req.Header.Set("X-API-Key", b.key)
	req.Header.Set("Content-Type", writer.FormDataContentType())
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
	//处理数据
	var body []byte
	if body, err = io.ReadAll(resp.Body); err != nil {
		return
	}
	var obj upload
	if err = json.Unmarshal(body, &obj); err != nil {
		return
	}
	path = obj.Path
	return
}
