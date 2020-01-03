package protocol

// Canonical HTTP protocol level message.
// Any HTTP response must be provided using this message.
type HttpMessage struct {
	Error HttpErrorMessage `json:"error"`
	Success HttpSuccessMessage `json:"success"`
}

func (r HttpMessage) GetError () HttpErrorMessage {
	return r.Error
}

func (r HttpMessage) GetSuccess () HttpSuccessMessage {
	return r.Success
}
