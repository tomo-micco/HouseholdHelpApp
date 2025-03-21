package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/tomo-micco/HouseholdHelpApp/internal/entities"
)

type CategoryRepo interface {
	GetAllCategories() ([]entities.Category, error)
	CreateCategory(category entities.Category) error
	UpdateCategory(category entities.Category) error
	DeleteCategory(id uint32) error
}

type CategoriesRepository struct {
	db sqlx.DB
}

// コンストラクタ
func NewCategoriesRepository(db sqlx.DB) *CategoriesRepository {
	return &CategoriesRepository{
		db: db,
	}
}

// 全件取得
func (c *CategoriesRepository) GetAllCategories() ([]entities.Category, error) {
	categories := []entities.Category{}
	err := c.db.Select(&categories, "SELECT * FROM categories")
	return categories, err
}

// 登録
func (c *CategoriesRepository) CreateCategory(category entities.Category) error {
	_, err := c.db.NamedExec("INSERT INTO categories (name) VALUES (:name)", category)
	return err
}

// 更新
func (c *CategoriesRepository) UpdateCategory(category entities.Category) error {
	_, err := c.db.NamedExec("UPDATE categories SET name = :name WHERE id = :id AND user_id = :user_id", category)
	return err
}

// 削除
func (c *CategoriesRepository) DeleteCategory(id uint32) error {
	_, err := c.db.Exec("DELETE FROM categories WHERE id = $1", id)
	return err
}
