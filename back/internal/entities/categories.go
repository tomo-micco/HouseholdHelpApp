package entities

// 項目
type Category struct {
	Id     uint32 `db:"id" json:"id"`
	UserId uint32 `db:"user_id" json:"user_id"`
	Name   string `db:"name" json:"name"`
}
