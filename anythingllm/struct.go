package anythingllm

type reqBody struct {
	Message string `json:"message"`
	Mode    string `json:"mode"`
	Stream  bool   `json:"stream"`
	Reset   bool   `json:"reset"`
}

type repInfo struct {
	UUID         string `json:"uuid"`
	Type         string `json:"type"`
	TextResponse string `json:"textResponse"`
	Close        bool   `json:"close"`
	Error        bool   `json:"error"`
}
