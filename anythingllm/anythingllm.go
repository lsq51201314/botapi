package anythingllm

import "fmt"

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
