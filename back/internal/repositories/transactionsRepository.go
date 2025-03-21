package repositories

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/tomo-micco/HouseholdHelpApp/internal/entities"
)

type TransactionRepo interface {
	GetTransactionsByDate(date time.Time) ([]entities.Transaction, error)
	CreateTransaction(transaction entities.Transaction) error
	UpdateTransaction(transaction entities.Transaction) error
	DeleteTransaction(id uint32) error
}

type TransactionsRepository struct {
	db sqlx.DB
}

// コンストラクタ
func NewTransactionRepository(db *sqlx.DB) *TransactionsRepository {
	return &TransactionsRepository{
		db: *db,
	}
}

// 日付に該当する取引情報取得
func (t TransactionsRepository) GetTransactionsByDate(date time.Time) ([]entities.Transaction, error) {
	var transactions []entities.Transaction
	err := t.db.Select(&transactions, "SELECT * FROM transactions WHERE date = $1", date)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

// 取引情報登録
func (t TransactionsRepository) CreateTransaction(transaction entities.Transaction) error {
	_, err := t.db.NamedExec("INSERT INTO transactions (category_id, user_id, amount, in_out_type, date) VALUES (:category_id, :user_id, :amount, :in_out_type, :date)", transaction)
	return err
}

// 取引情報更新
func (t TransactionsRepository) UpdateTransaction(transaction entities.Transaction) error {
	_, err := t.db.NamedExec("UPDATE transactions SET category_id = :category_id, user_id = :user_id, amount = :amount, in_out_type = :in_out_type, date = :date WHERE id = :id", transaction)
	return err
}

// 取引情報削除
func (t TransactionsRepository) DeleteTransaction(id uint32) error {
	_, err := t.db.Exec("DELETE FROM transactions WHERE id = $1", id)
	return err
}
