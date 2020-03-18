package connection

import (
	"database/sql"
	"fmt"

	_ "gopkg.in/goracle.v2"
)

type Oracle struct {
	ListOracle map[string]*sql.DB
}

type OracleParams struct {
	Name     string
	Host     string
	Port     int
	User     string
	Password string
	Schema   string
	Driver   string
	Sid      string
}

func (r *Oracle) OracleMultipleConnection(options ...OracleParams) error {
	m := make(map[string]*sql.DB)
	for _, opt := range options {
		//db, err := sql.Open("ora", "ISISALL/ISISALL@10.1.35.3:1521/APPDB")
		fmt.Println("   driver : ", fmt.Sprintf("%s", opt.Driver))
		db, err := sql.Open(fmt.Sprintf("%s", opt.Driver), fmt.Sprintf("%s/%s@%s:%d/%s", opt.User, opt.Password, opt.Host, opt.Port, opt.Sid))

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
	r.ListOracle = m
	return nil
}
