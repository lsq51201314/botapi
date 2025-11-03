package ctyun

type CTYun struct {
	model            string
	url              string
	authorization    string
	thinkcallback    Callback
	messagescallback Callback
	timeout          int64
}

// 新建实例
func New(model, authorization string, timeout ...int64) *CTYun {
	obj := CTYun{
		model:         model,
		url:           "https://wishub-x6.ctyun.cn/v1/chat/completions",
		authorization: authorization,
	}
	if len(timeout) > 0 {
		obj.timeout = timeout[0]
	}
	return &obj
}
