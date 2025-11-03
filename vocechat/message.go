package vocechat

import (
	"github.com/gin-gonic/gin"
)

type msg struct {
	Detail detail `json:"detail"`
	Target target `json:"target"`
	Type   string `json:"type"`
}

type detail struct {
	Content     string `json:"content"`
	ContentType string `json:"content_type"`
}

type target struct {
	Gid int64 `json:"gid"`
}

func (b *Bot) message(c *gin.Context) {
	var p msg
	if bind(c, nil, &p) {
		if p.Type == "chat" && p.Detail.ContentType == "text/plain" {
			if b.call != nil {
				b.call(b, p.Target.Gid, p.Detail.Content)
			}
		}
		sendSuccess(c)
	}
}
