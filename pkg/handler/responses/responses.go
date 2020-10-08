package responses

type Response struct {
	Result   interface{} `json:"result"`
	Messages string      `json:"message"`
	Success  bool        `json:"success"`
}
