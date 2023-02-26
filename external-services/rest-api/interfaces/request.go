package interfaces

type RequestInterface interface {
	Call(model any) error
}
