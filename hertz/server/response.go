package server

// ResponseMeta 返回元数据
type ResponseMeta struct {
	RequestId string    `json:"omitempty"`
	Action    string    `json:"omitempty"`
	Version   string    `json:"omitempty"`
	Error     *ApiError `json:"omitempty"`
}

type Response struct {
	Meta   ResponseMeta `json:"omitempty"`
	Result interface{}  `json:"omitempty"`
}
