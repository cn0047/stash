package protocol

type ResponseInterface interface {
	GetError() ErrorInterface
	GetSuccess() SuccessInterface
}
