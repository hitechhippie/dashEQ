package eqdb

import (
	"dasheq/internal/config"
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

func Connect(c *config.ServerConfig) (*Connection, error) {
	// Build the MySQL connection configuration
	cfg := mysql.Config{
		User:                 c.DBuser,
		Passwd:               c.DBpass,
		Net:                  c.DBnet,
		Addr:                 c.DBaddr,
		DBName:               c.DBname,
		AllowNativePasswords: true,
	}

	var err error
	var mysqldb *sql.DB
	var database Connection

	mysqldb, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}
	err = mysqldb.Ping()
	if err != nil {
		return nil, err
	}

	database.Target = mysqldb

	if err != nil {
		return nil, err
	}

	return &database, nil
}

func Close(c *Connection) error {
	err := c.Target.Close()
	if err != nil {
		return err
	}

	return nil
}
