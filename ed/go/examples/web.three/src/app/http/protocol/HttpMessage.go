package protocol

type HttpMessage struct {
	Error HttpErrorMessage
	Success HttpSuccessMessage
}

func (r HttpMessage) GetError () HttpErrorMessage {
	return r.Error
}

func (r HttpMessage) GetSuccess () HttpSuccessMessage {
	return r.Success
}
