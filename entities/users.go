package entities

import "time"

type Users struct {
	Id         int
	Telp       string
	Pass       string
	Firstname  string
	Lastname   string
	Saldo      int
	Created_at time.Time
	Updated_at time.Time
}
