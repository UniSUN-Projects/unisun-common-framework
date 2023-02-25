package interfaces

type RequestInterface interface {
	Request(model *any) error
}
