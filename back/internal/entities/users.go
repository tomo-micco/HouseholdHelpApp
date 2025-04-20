package entities

// パスワード型
type Password string

// ログ出力時にパスワードを伏せ字にする
func (p Password) String() string {
	return "****"
}

// ログ出力時にパスワードを伏せ字にする
func (p Password) GoString() string {
	return "****"
}

// ユーザー
type User struct {
	Id       uint32   `db:"id" json:"id"`
	Name     string   `db:"name" json:"name" binding:"required"`
	Email    string   `db:"email" json:"email" binding:"required"`
	Password Password `db:"password" json:"password" binding:"required"`
}

func NewUser() *User {
	return &User{}
}
