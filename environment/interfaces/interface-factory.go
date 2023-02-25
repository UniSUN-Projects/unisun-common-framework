package interfaces

type FactoryInterface interface {
	Default(model *any)
	Custom(model *any)
}
