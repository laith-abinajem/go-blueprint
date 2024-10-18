package db

import (
	"GoProject/config"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type MysqlDB struct {
	DB *gorm.DB
}

func NewMysqlDB(cfg *config.Config) (*MysqlDB, error) {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySQL.MysqlUser,
		cfg.MySQL.MysqlPassword,
		cfg.MySQL.MysqlHost,
		cfg.MySQL.MysqlPort,
		cfg.MySQL.MysqlDBName,
	)

	db, err := gorm.Open(mysql.Open(conn), &gorm.Config{
		PrepareStmt:                              true,
		DisableForeignKeyConstraintWhenMigrating: true,
		SkipDefaultTransaction:                   true, // Disable automatic transactions for read-only operations

		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "vfx_",
			SingularTable: true,
			NoLowerCase:   true,
		},
	})
	if err != nil {
		return nil, err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings
	sqlDB.SetConnMaxLifetime(5 * time.Minute) // Maximum connection lifetime
	sqlDB.SetMaxIdleConns(10)                 // Maximum idle connections
	sqlDB.SetMaxOpenConns(100)                // Maximum open connections

	dbc := &MysqlDB{
		DB: db,
	}

	return dbc, nil
}
func Migrate(cfg *config.Config) error {
	db, err := NewMysqlDB(cfg)
	if err != nil {
		return err
	}
	if err := db.DB.AutoMigrate(); err != nil {
		return err
	}
	return nil
}
