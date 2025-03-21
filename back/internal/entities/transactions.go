package entities

import "time"

// 収入・支出
type Transaction struct {
	Id         uint32
	CategoryId uint32
	UserId     uint32
	Amount     uint32
	InOutType  string
	Date       time.Time
}
