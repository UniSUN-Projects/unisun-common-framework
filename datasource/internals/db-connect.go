package internals

import (
	"fmt"
	"time"

	"github.com/UniSUN-Projects/unisun-common-framework/datasource/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type OptionHandle struct{}

func NewOptionHandle() *OptionHandle {
	return &OptionHandle{}
}

func (*OptionHandle) Connect() {
	var err error
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", models.Config.AppConfig.Datasource.Host, models.Config.AppConfig.Datasource.Username,
		models.Config.AppConfig.Datasource.Password, models.Config.AppConfig.Datasource.DBName, models.Config.AppConfig.Datasource.Port,
		models.Config.AppConfig.Datasource.Properties.SSLMode, models.Config.AppConfig.Datasource.Properties.TimeZone)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbInstance, _ := DB.DB()
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	dbInstance.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	dbInstance.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	dbInstance.SetConnMaxLifetime(time.Hour)
	defer dbInstance.Close()
}

func (*OptionHandle) SetMigrate(instan any) error {
	return DB.AutoMigrate(instan)
}

func (*OptionHandle) Close() {
	dbInstance, _ := DB.DB()
	_ = dbInstance.Close()
}
