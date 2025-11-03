package aliyun

type Aliyun struct {
	model            string
	url              string
	authorization    string
	thinkcallback    Callback
	messagescallback Callback
	timeout          int64
}

// 新建实例
func New(model, authorization string, timeout ...int64) *Aliyun {
	obj := Aliyun{
		model:         model,
		url:           "https://dashscope.aliyuncs.com/compatible-mode/v1/chat/completions",
		authorization: authorization,
	}
	if len(timeout) > 0 {
		obj.timeout = timeout[0]
	}
	return &obj
}
