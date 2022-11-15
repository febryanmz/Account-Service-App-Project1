package entities

import "database/sql"

type Users struct {
	Id         int
	Telp       string
	Pass       string
	Firstname  string
	Lastname   string
	Saldo      sql.NullInt64
	Created_at int
	Updated_at int
}
