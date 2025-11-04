package aliyun

// 流式回调
type CallBack func(id string, message string)

func (a *Aliyun) SetCallback(cfunc CallBack) {
	a.callback = cfunc
}
