package interfaces

type DbConnectInterface interface {
	Connect()
	SetMigrate(instan []interface{}) error
}
