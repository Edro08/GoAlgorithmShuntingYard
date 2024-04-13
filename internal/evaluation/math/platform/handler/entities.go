package handler

type Request struct {
	Infix string `json:"infix"`
}

type Response struct {
	Infix   string  `json:"infix"`
	Postfix string  `json:"postfix"`
	Result  float64 `json:"result"`
}

type ErrResponse struct {
	Message string `json:"message"`
}
