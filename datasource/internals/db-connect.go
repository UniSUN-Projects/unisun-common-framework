package internals

import (
	"fmt"

	"github.com/UniSUN-Projects/unisun-common-framework/datasource/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type OptionHandle struct{}

func NewOptionHandle() *OptionHandle {
	return &OptionHandle{}
}

func (*OptionHandle) Connect() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", models.Config.AppConfig.Datasource.Host, models.Config.AppConfig.Datasource.Username,
		models.Config.AppConfig.Datasource.Password, models.Config.AppConfig.Datasource.DBName, models.Config.AppConfig.Datasource.Port,
		models.Config.AppConfig.Datasource.Properties.SSLMode, models.Config.AppConfig.Datasource.Properties.TimeZone)
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbInstance, _ := db.DB()
	defer func() {
		dbInstance.Close()
	}()
}

func (*OptionHandle) SetMigrate(instan any) error {
	return db.AutoMigrate(instan)
}

func (*OptionHandle) Get() *gorm.DB {
	return db
}
