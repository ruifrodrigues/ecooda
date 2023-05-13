package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	gorm.Dialector
	*gorm.DB
}

type IDatabase interface {
	OpenConnection() (Database, error)
	CloseConnection() error
}

func NewDatabase(dialer IDialer) Database {
	return Database{
		Dialector: dialer.Dial(),
	}
}

func (db Database) OpenConnection() (Database, error) {
	database := Database{}

	sqlDB, err := gorm.Open(db.Dialector, &gorm.Config{})
	if err != nil {
		return database, err
	}

	database.DB = sqlDB

	return database, nil
}

func (db Database) CloseConnection() error {
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

type Dialer struct {
	DSN string
}

type IDialer interface {
	Dial() gorm.Dialector
}

func (d Dialer) Dial() gorm.Dialector {
	return mysql.Open(d.DSN)
}

func NewMySQLDialer(conf Config) IDialer {
	return Dialer{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
			conf.Values.DBUsername,
			conf.Values.DBPassword,
			conf.Values.DBHost,
			conf.Values.DBPort,
			conf.Values.DBName,
		),
	}
}

func NewSQLiteDialer(conf Config) IDialer {
	return Dialer{
		DSN: conf.Values.DBTest + ".db",
	}
}
