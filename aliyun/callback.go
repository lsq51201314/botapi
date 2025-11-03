package aliyun

// 流式回调
type Callback func(id, text string)

func (a *Aliyun) SetThinkCallback(cfunc Callback) {
	a.thinkcallback = cfunc
}

func (a *Aliyun) SetMessagesCallback(cfunc Callback) {
	a.messagescallback = cfunc
}
