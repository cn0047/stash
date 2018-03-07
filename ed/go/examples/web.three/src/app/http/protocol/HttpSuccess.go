package protocol

// Canonical HTTP success message.
// This structure must be used to reply with success message.
// This structure is - container for payload shipping.
type HttpSuccessMessage struct {
    Code int `json:"code"`
    Data interface{} `json:"data"`
}

// Wrapper for HTTP success message,
// which helps to provide payload in more elegant way.
func HttpSuccess(code int, data interface{}) HttpMessage {
    s := HttpSuccessMessage{Code: code, Data: data}

    return HttpMessage{Error: HttpErrorMessage{}, Success: s}
}
