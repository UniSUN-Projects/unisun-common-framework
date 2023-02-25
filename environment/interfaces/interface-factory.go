package interfaces

type FactoryInterface interface {
	Default(model *interface{})
	Custom(model *interface{})
}
