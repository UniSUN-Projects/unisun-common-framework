package environment

type FactoryInterface interface {
	Default(model *any)
	Custom(model *any)
}
