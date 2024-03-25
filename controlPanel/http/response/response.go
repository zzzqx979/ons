package response

type TraceResponse struct {
	Code    uint32      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	Trace   Trace       `json:"trace"`
}

type Response struct {
	Code    uint32      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type Trace struct {
	Ons []TraceItem `json:"ons"`
	Ms  []TraceItem `json:"ms"`
}

type TraceItem struct {
	Upstream string `json:"upstream"`
	Time     string `json:"time"`
}
