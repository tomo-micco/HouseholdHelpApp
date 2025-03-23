package entities

// 項目
type Category struct {
	Id     uint32 `db:"id"`
	UserId uint32 `db:"user_id"`
	Name   string `db:"name"`
}
