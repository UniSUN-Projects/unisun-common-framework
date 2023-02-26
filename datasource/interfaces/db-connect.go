package interfaces

type DbConnectInterface interface {
	Connect()
	SetMigrate(instan any) error
}
