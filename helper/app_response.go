package helper

type ObjectKosongResponse struct {
}

type ListKosongResponse struct {
}

type DefaultResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type DefaultListResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type DefaultErrorResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type DefaultLoginResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Token  string      `json:"token"`
}
