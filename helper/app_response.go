package helper

type ObjectKosongResponse struct {
}

type ListKosongResponse struct {
}

type DefaultPaginationResponse struct {
	Code     int         `json:"code"`
	Status   bool        `json:"status"`
	Message  string      `json:"message"`
	Page     int         `json:"page"`
	Total    int64       `json:"total"`
	LastPage int         `json:"last_page"`
	Data     interface{} `json:"data"`
}
type DefaultResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DefaultListResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type DefaultErrorResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

type DefaultLoginResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Token   string      `json:"token"`
}
