package protocol

type HttpErrorMessage struct {
	Code int
	Description PayloadInterface
}

func HttpError(code int, desc string) HttpMessage {
	e := HttpErrorMessage{Code: code, Description: desc}

	return HttpMessage{Error: e, Success: HttpSuccessMessage{}}
}
