package ctyun

// 流式回调
type Callback func(id, text string)

func (c *CTYun) SetThinkCallback(cfunc Callback) {
	c.thinkcallback = cfunc
}

func (c *CTYun) SetMessagesCallback(cfunc Callback) {
	c.messagescallback = cfunc
}
