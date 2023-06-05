package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB interface {
	OpenConnection() error
	CloseConnection() error
	Ctx() Database
}

type Database struct {
	gorm.Dialector
	*gorm.DB
}

func NewDatabase(dialer Dialector) DB {
	return &Database{
		Dialector: dialer.Dial(),
	}
}

func (db *Database) OpenConnection() error {
	config := new(gorm.Config)
	//config.Logger = logger.Default.LogMode(logger.Info)

	var err error

	if db.DB, err = gorm.Open(db.Dialector, config); err != nil {
		return err
	}

	return nil
}

func (db *Database) CloseConnection() error {
	sqlDB, err := db.DB.DB()
	if err != nil {
		return err
	}

	err = sqlDB.Close()
	if err != nil {
		return err
	}

	return nil
}

func (db *Database) Ctx() Database {
	if db.DB == nil {
		panic("Database Connection not open!")
	}

	return *db
}

type Dialector interface {
	Dial() gorm.Dialector
}

type Dialer struct {
	DSN string
}

func NewMySQLDialer(conf Config) Dialector {
	return &Dialer{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			conf.Values.DBUsername,
			conf.Values.DBPassword,
			conf.Values.DBHost,
			conf.Values.DBPort,
			conf.Values.DBName,
		),
	}
}

func (d *Dialer) Dial() gorm.Dialector {
	return mysql.Open(d.DSN)
}

func NewSQLiteDialer(conf Config) Dialector {
	return &Dialer{
		DSN: conf.Values.DBTest + ".db",
	}
}
