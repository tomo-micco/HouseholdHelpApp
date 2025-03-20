package entities

/*
 * ユーザー
 */
type User struct {
	Id       uint32
	Name     string
	Email    string
	Password string
}

func NewUser() *User {
	return &User{}
}
