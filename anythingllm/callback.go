package anythingllm

// 流式回调
type CallBack func(id string, message string)

func (a *AnythingLLM) SetCallback(cfunc CallBack) {
	a.callback = cfunc
}
