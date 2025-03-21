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
	Id       uint32
	Name     string
	Email    string
	Password Password
}

func NewUser() *User {
	return &User{}
}
