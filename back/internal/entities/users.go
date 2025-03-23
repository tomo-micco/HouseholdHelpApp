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
	Id       uint32   `db:"id"`
	Name     string   `db:"name"`
	Email    string   `db:"email"`
	Password Password `db:"password"`
}

func NewUser() *User {
	return &User{}
}
