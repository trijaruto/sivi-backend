package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type MySql struct {
	ListMysql map[string]*sql.DB
}

type MySqlParams struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
	Driver   string
}

func (r *MySql) MySqlMultipleConnection(options ...MySqlParams) error {
	m := make(map[string]*sql.DB)
	for _, opt := range options {
		db, err := sql.Open(fmt.Sprintf("%s", opt.Driver), fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", opt.User, opt.Password, opt.Host, opt.Port, opt.Schema))

		if err != nil {
			panic(err)
		}
		err = db.Ping()
		if err != nil {
			db.Close()
			panic(err)
		}
		m[opt.Name] = db
	}
	r.ListMysql = m
	return nil
}
