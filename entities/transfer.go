package entities

import "time"

type Transfer struct {
	id               int
	user_id_penerima int
	user_id_pengirim int
	transfer_balance int
	created_at       time.Time
}
