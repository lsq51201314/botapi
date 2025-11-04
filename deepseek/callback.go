package deepseek

// 流式回调
type CallBack func(id string, message string)

func (a *DeepSeek) SetCallback(cfunc CallBack) {
	a.callback = cfunc
}
