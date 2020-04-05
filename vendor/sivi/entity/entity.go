package entity

type ResponseHttp struct {
	Code    int8        `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
