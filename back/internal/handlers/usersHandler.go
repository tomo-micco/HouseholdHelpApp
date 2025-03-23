package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/tomo-micco/HouseholdHelpApp/internal/entities"
	"github.com/tomo-micco/HouseholdHelpApp/internal/repositories"
)

type UsersHandler struct {
	db       *sqlx.DB
	userRepo repositories.UserRepo
}

// コンストラクタ
func NewUsersHandler(db *sqlx.DB, userRepo repositories.UserRepo) *UsersHandler {
	return &UsersHandler{db: db, userRepo: userRepo}
}

// ユーザー取得
func (u UsersHandler) GetUsers(c *gin.Context) {
	users, err := u.userRepo.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, users)
}

// IDに該当するユーザー取得
func (u UsersHandler) GetById(c gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		c.Abort()
		return
	}

	user, err := u.userRepo.GetById(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, user)
}

// ユーザー作成
func (u UsersHandler) CreateUser(c *gin.Context) {

	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	tx := u.db.MustBegin()

	if err := u.userRepo.CreateUser(user, tx); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": rollbackErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	if err := tx.Commit(); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": rollbackErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, user)
}

// ユーザー更新
func (u UsersHandler) UpdateUser(c *gin.Context) {
	var user entities.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	tx := u.db.MustBegin()

	if err := u.userRepo.UpdateUser(user, tx); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": rollbackErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	if err := tx.Commit(); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": rollbackErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, nil)
}

// ユーザー削除
func (u UsersHandler) DeleteUser(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		c.Abort()
		return
	}

	tx := u.db.MustBegin()
	defer tx.Rollback()

	if err := u.userRepo.DeleteUser(id, tx); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": rollbackErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	if err := tx.Commit(); err != nil {
		rollbackErr := tx.Rollback()
		if rollbackErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": rollbackErr.Error()})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, nil)
}
