package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/tomo-micco/HouseholdHelpApp/internal/entities"
)

// ユーザー操作用リポジトリインターフェース
type UserRepo interface {
	GetAllUsers() ([]entities.User, error)
	GetById(id int) (entities.User, error)
	CreateUser(user entities.User, tx *sqlx.Tx) error
	UpdateUser(user entities.User, tx *sqlx.Tx) error
	DeleteUser(id int, tx *sqlx.Tx) error
}

type UsersRepository struct {
	db *sqlx.DB
}

// コンストラクタ
func NewUsersRepository(db *sqlx.DB) *UsersRepository {
	return &UsersRepository{db}
}

// 全ユーザー取得
func (u *UsersRepository) GetAllUsers() ([]entities.User, error) {
	users := []entities.User{}
	err := u.db.Select(&users, "SELECT * FROM users")
	return users, err
}

// ユーザー取得
func (u *UsersRepository) GetById(id int) (entities.User, error) {
	user := entities.User{}
	err := u.db.Get(&user, "SELECT * FROM users WHERE id = ?", id)
	return user, err
}

// ユーザー作成
func (u *UsersRepository) CreateUser(user entities.User, tx *sqlx.Tx) error {
	_, err := tx.NamedExec("INSERT INTO users (name, email, password) VALUES (:name, :email, :password)", user)
	return err
}

// ユーザー更新
func (u *UsersRepository) UpdateUser(user entities.User, tx *sqlx.Tx) error {
	_, err := tx.NamedExec("UPDATE users SET name = :name, email = :email, password = :password WHERE id = :id", user)
	return err
}

// ユーザー削除
func (u *UsersRepository) DeleteUser(id int, tx *sqlx.Tx) error {
	_, err := tx.Exec("DELETE FROM users WHERE id = ?", id)
	return err
}
