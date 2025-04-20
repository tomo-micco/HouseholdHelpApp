package entities

import "time"

// 収入・支出
type Transaction struct {
	Id         uint32    `db:"id" json:"id"`
	CategoryId uint32    `db:"category_id" json:"category_id" binding:"required"`
	UserId     uint32    `db:"user_id" json:"user_id" binding:"required"`
	Amount     uint32    `db:"amount" json:"amount" binding:"required"`
	InOutType  string    `db:"in_out_type" json:"in_out_type" binding:"required"`
	Date       time.Time `db:"date" json:"date" binding:"required"`
}
