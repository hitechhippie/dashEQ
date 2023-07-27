package db

import (
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

type ConnectionConfig struct {
	User   string
	Pass   string
	Net    string
	Addr   string
	Dbname string
}

type Connection struct {
	Target *sql.DB
}

func Connect(c *ConnectionConfig) (*Connection, error) {
	// Build the MySQL connection configuration
	cfg := mysql.Config{
		User:                 c.User,
		Passwd:               c.Pass,
		Net:                  c.Net,
		Addr:                 c.Addr,
		DBName:               c.Dbname,
		AllowNativePasswords: true,
	}

	var err error
	var mysqldb *sql.DB
	var database Connection

	mysqldb, err = sql.Open("mysql", cfg.FormatDSN())

	database.Target = mysqldb

	if err != nil {
		return nil, err
	}

	return &database, nil
}

func Close(d *Connection) error {
	err := d.Target.Close()
	return err
}
