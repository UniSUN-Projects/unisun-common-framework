package restapi

type RequestInterface interface {
	Request(model *any) error
}
