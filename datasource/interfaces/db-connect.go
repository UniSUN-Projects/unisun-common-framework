package interfaces

import "gorm.io/gorm"

type DbConnectInterface interface {
	Connect()
	SetMigrate(instan any) error
	Get() *gorm.DB
	Close() error
}
