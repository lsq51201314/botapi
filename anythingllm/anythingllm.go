package anythingllm

import "fmt"

// 流式回调
type CallBack func(id string, message string)

type AnythingLLM struct {
	url           string
	authorization string
	timeout       int64
	callback      CallBack
}

func New(url, slug string, authorization string, timeout ...int64) *AnythingLLM {
	obj := AnythingLLM{
		url:           fmt.Sprintf(url+"/api/v1/workspace/%s/stream-chat", slug),
		authorization: authorization,
	}
	if len(timeout) > 0 {
		obj.timeout = timeout[0]
	}
	return &obj
}

func (a *AnythingLLM) SetCallback(cfunc CallBack) {
	a.callback = cfunc
}
