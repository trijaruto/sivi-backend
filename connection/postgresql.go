package connection

import (
	"database/sql"
	"fmt"
)

func NewConnectionPgsql(hostname string, port int, username string, password string, schema string) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		hostname, port, username, password, schema))
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, err
}

type Pgsql struct {
	ListPgsql map[string]*sql.DB
}

type PgsqlParams struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
}

func (r *Pgsql) NewPgsqlMultipleConnection(options ...PgsqlParams) error {
	m := make(map[string]*sql.DB)
	for _, opt := range options {
		db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
			opt.Host, opt.Port, opt.User, opt.Password, opt.Schema))

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
	r.ListPgsql = m
	return nil
}
