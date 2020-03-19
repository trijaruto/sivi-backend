package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
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
	Database string
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
	Driver   string
	URI      string
}

func (r *Pgsql) PgsqlMultipleConnection(options ...PgsqlParams) error {
	m := make(map[string]*sql.DB)
	for _, opt := range options {
		db, err := sql.Open(fmt.Sprintf("%s", opt.Driver), fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
			opt.Host, opt.Port, opt.User, opt.Password, opt.Database))

		if err != nil {
			fmt.Println("err", err)
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
