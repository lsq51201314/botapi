package aliyun

//请求参数
type chat struct {
	Model          string     `json:"model"` //模型ID
	Messages       []Messages `json:"messages"`
	Stream         bool       `json:"stream"`
	EnableThinking bool       `json:"enable_thinking"`
}

/**
用户当前输入的期望模型执行指令。一个列表内多个字典，支持多轮对话。
对话列表，每个列表项为一个message object，message object中包含用户role和content两部分信息：
role可选值为user、assistant、system；
role为system时，不校验content空值，且message中system只能位于开头，即messages[0]位置；
role为user时说明是用户提问，role为assistant时说明是模型回答，而content为实际的对话内容；
单轮/多轮对话中，最后一个message中role必须为user，content为用户输入的最新问题，其余结果除system角色外都为历史信息拼接送入
messages中，assistant和user的role只能交替出现，assistant后只能跟user，user后只能跟assistant。
**/
type Messages struct {
	Role    string `json:"role"` //user、assistant、system
	Content string `json:"content"`
}

type ChatStream struct {
	ID      string        `json:"id"`
	Created int64         `json:"created"`
	Choices []ChatChoices `json:"choices"`
}

type ChatChoices struct {
	Delta struct {
		Content          string `json:"content"`
	} `json:"delta"`
	FinishReason string `json:"finish_reason"` //stop, length, content_filter, tool_calls, insufficient_system_resource
}
