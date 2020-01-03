package protocol

// Canonical HTTP error message.
// This structure must be used to handle any error.
type HttpErrorMessage struct {
    Code int `json:"code"`
    Description interface{} `json:"description"`
    Data interface{} `json:"data"`
}

// Wrapper for HTTP error message,
// which helps to handle errors in more elegant way.
func HttpError(code int, desc string) HttpMessage {
    e := HttpErrorMessage{Code: code, Description: desc}

    return HttpMessage{Error: e, Success: HttpSuccessMessage{}}
}

// Wrapper for errors gained from "panic".
// In addition to regular error information,
// this one provides additional custom payload.
func HttpException(data interface{}) HttpMessage {
    e := HttpErrorMessage{
        Code: 500,
        Description: "Internal Server Error.",
        Data: data,
    }

    return HttpMessage{Error: e, Success: HttpSuccessMessage{}}
}
