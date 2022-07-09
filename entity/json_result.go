package entity

type JSONResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Result  interface{} `json:"result"`
	From    string      `json:"from"`
}
