package deepseek

// 流式回调
type Callback func(id, text string)

func (d *DeepSeek) SetThinkCallback(cfunc Callback) {
	d.thinkcallback = cfunc
}

func (d *DeepSeek) SetMessagesCallback(cfunc Callback) {
	d.messagescallback = cfunc
}
