package entities

import "time"

type TopUp struct {
	Id         int
	User_id    int
	Balance    int
	Created_at time.Time
}
