package handler

//handler
import (
	"RakaminProject/models"
	"RakaminProject/repository"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// UserHandler -> interface to User entity
type UserHandler interface {
	AddUser(*gin.Context)
	GetUser(*gin.Context)
	SignInUser(*gin.Context)
	UpdateUser(*gin.Context)
	DeleteUser(*gin.Context)
	GetPhoto(*gin.Context)
}

type userHandler struct {
	repo repository.UserRepository
}

// NewUserHandler --> returns new user handler
func NewUserHandler() UserHandler {

	return &userHandler{
		repo: repository.NewUserRepository(),
	}
}

func hashPassword(pass *string) {
	bytePass := []byte(*pass)
	hPass, _ := bcrypt.GenerateFromPassword(bytePass, bcrypt.DefaultCost)
	*pass = string(hPass)
}

func comparePassword(dbPass, pass string) bool {
	return bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(pass)) == nil
}

func (h *userHandler) GetUser(ctx *gin.Context) {
	id := uint(ctx.MustGet("userID").(float64))
	user, err := h.repo.GetUser(uint(id))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)

}

func (h *userHandler) SignInUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	dbUser, err := h.repo.GetByEmail(user.Email)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "No Such User Found"})
		return

	}
	if isTrue := comparePassword(dbUser.Password, user.Password); isTrue {
		fmt.Println("user before", dbUser.ID)
		token := GenerateToken(dbUser.ID)
		ctx.JSON(http.StatusOK, gin.H{"msg": "Successfully SignIN", "token": token})
		return
	}
	ctx.JSON(http.StatusInternalServerError, gin.H{"msg": "Password not matched"})
	return

}

func (h *userHandler) AddUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	hashPassword(&user.Password)
	user, err := h.repo.AddUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	user.Password = ""
	ctx.JSON(http.StatusOK, user)

}

func (h *userHandler) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id := ctx.Param("user")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	user.ID = uint(intID)
	userID := uint(ctx.MustGet("userID").(float64))
	repo := repository.NewUserRepository()
	users, err := repo.GetUser(uint(intID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if users.ID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you dont have permission"})
		return
	}
	user, err = h.repo.UpdateUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, user)
}

func (h *userHandler) DeleteUser(ctx *gin.Context) {
	var user models.User
	id := ctx.Param("user")
	intID, _ := strconv.Atoi(id)
	user.ID = uint(intID)
	userID := uint(ctx.MustGet("userID").(float64))

	repo := repository.NewUserRepository()
	users, err := repo.GetUser(uint(intID))
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if users.ID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you dont have permission"})
		return
	}

	User, err := h.repo.DeleteUser(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, User)

}

func (h *userHandler) GetPhoto(ctx *gin.Context) {
	userID := uint(ctx.MustGet("userID").(float64))
	if photos, err := h.repo.GetAllPhoto(userID); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, photos)
	}
}
