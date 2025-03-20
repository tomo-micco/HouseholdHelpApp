package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/tomo-micco/HouseholdHelpApp/internal/entities"
)

/*
 * ユーザー操作用リポジトリインターフェース
 */
type UserRepo interface {
	GetAllUsers() ([]entities.User, error)
	GetById(id uint32) (entities.User, error)
	CreateUser(user entities.User) error
	UpdateUser(user entities.User) error
	DeleteUser(id uint32) error
}

type UsersRepository struct {
	db sqlx.DB
}

/*
 * コンストラクタ
 */
func NewUsersRepository(db sqlx.DB) *UsersRepository {
	return &UsersRepository{db}
}

/*
 * 全ユーザー取得
 */
func (u *UsersRepository) GetAllUsers() ([]entities.User, error) {
	users := []entities.User{}
	err := u.db.Select(&users, "SELECT * FROM users")
	return users, err
}

/*
 * ユーザー取得
 */
func (u *UsersRepository) GetById(id uint32) (entities.User, error) {
	user := entities.User{}
	err := u.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	return user, err
}

/*
* ユーザー作成
 */
func (u *UsersRepository) CreateUser(user entities.User) error {
	_, err := u.db.NamedExec("INSERT INTO users (name, email, password) VALUES (:name, :email, :password)", user)
	return err
}

/*
 * ユーザー更新
 */
func (u *UsersRepository) UpdateUser(user entities.User) error {
	_, err := u.db.NamedExec("UPDATE users SET name = :name, email = :email, password = :password WHERE id = :id", user)
	return err
}

/*
 * ユーザー削除
 */
func (u *UsersRepository) DeleteUser(id uint32) error {
	_, err := u.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}
