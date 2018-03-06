package protocol

type HttpSuccessMessage struct {
	Code int
	Data PayloadInterface
}

func HttpSuccess(code int, data PayloadInterface) HttpMessage {
	s := HttpSuccessMessage{Code: code, Data: data}

	return HttpMessage{Error: HttpErrorMessage{}, Success: s}
}
