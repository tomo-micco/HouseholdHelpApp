package usecases

import (
	"github.com/jmoiron/sqlx"
	"github.com/tomo-micco/HouseholdHelpApp/internal/entities"
	"github.com/tomo-micco/HouseholdHelpApp/internal/repositories"
)

type UsersUseCase struct {
	userRepository *repositories.UsersRepository
	db             *sqlx.DB
}

func NewUsersUseCase(userRepository *repositories.UsersRepository, db *sqlx.DB) *UsersUseCase {
	return &UsersUseCase{
		userRepository: userRepository,
		db:             db,
	}
}

func (u *UsersUseCase) GetAllUsers() ([]entities.User, error) {
	return u.userRepository.GetAllUsers()
}

func (u *UsersUseCase) GetById(id int) (entities.User, error) {
	return u.userRepository.GetById(id)
}

func (u *UsersUseCase) CreateUser(user entities.User) error {
	tx := u.db.MustBegin()

	err := u.userRepository.CreateUser(user, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *UsersUseCase) UpdateUser(user entities.User) error {
	tx := u.db.MustBegin()

	err := u.userRepository.UpdateUser(user, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u *UsersUseCase) DeleteUser(id int) error {
	tx := u.db.MustBegin()

	err := u.userRepository.DeleteUser(id, tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
