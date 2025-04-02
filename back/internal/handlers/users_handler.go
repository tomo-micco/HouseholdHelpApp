package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/tomo-micco/HouseholdHelpApp/internal/entities"
	"github.com/tomo-micco/HouseholdHelpApp/internal/repositories"
	"github.com/tomo-micco/HouseholdHelpApp/internal/usecases"
)

type UsersHandler struct {
	db *sqlx.DB
}

// コンストラクタ
func NewUsersHandler(db *sqlx.DB) UsersHandler {
	return UsersHandler{db: db}
}

// ユーザー取得
func (u UsersHandler) GetUsers(c *gin.Context) {
	usersRepository := repositories.NewUsersRepository(u.db)
	usersUseCase := usecases.NewUsersUseCase(usersRepository, u.db)

	users, err := usersUseCase.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, users)
}

// IDに該当するユーザー取得
func (u UsersHandler) GetById(c *gin.Context) {

	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		c.Abort()
		return
	}

	userRepository := repositories.NewUsersRepository(u.db)
	user, err := userRepository.GetById(id)
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

	userRepository := repositories.NewUsersRepository(u.db)
	userUseCase := usecases.NewUsersUseCase(userRepository, u.db)
	if err := userUseCase.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	userRepository := repositories.NewUsersRepository(u.db)
	userUseCase := usecases.NewUsersUseCase(userRepository, u.db)
	if err := userUseCase.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	userRepository := repositories.NewUsersRepository(u.db)
	usersUseCase := usecases.NewUsersUseCase(userRepository, u.db)
	if err := usersUseCase.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, nil)
}
