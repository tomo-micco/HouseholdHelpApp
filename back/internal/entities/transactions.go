package entities

import "time"

// 収入・支出
type Transaction struct {
	Id         uint32    `db:"id"`
	CategoryId uint32    `db:"category_id"`
	UserId     uint32    `db:"user_id"`
	Amount     uint32    `db:"amount"`
	InOutType  string    `db:"in_out_type"`
	Date       time.Time `db:"date"`
}
