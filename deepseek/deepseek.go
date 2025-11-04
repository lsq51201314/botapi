package deepseek

type DeepSeek struct {
	url           string
	authorization string
	callback      CallBack
	timeout       int64
}

// 新建实例
func New(authorization string, timeout ...int64) *DeepSeek {
	obj := DeepSeek{
		url:           "https://api.deepseek.com/chat/completions",
		authorization: authorization,
	}
	if len(timeout) > 0 {
		obj.timeout = timeout[0]
	}
	return &obj
}
