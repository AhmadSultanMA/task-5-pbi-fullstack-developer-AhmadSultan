package handler

//handler
import (
	"RakaminProject/models"
	"RakaminProject/repository"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PhotoHandler --> interface to Photo handler
type PhotoHandler interface {
	GetPhoto(*gin.Context)
	AddPhoto(*gin.Context)
	UpdatePhoto(*gin.Context)
	DeletePhoto(*gin.Context)
}

type photoHandler struct {
	repo repository.PhotoRepository
}

// NewPhotoHandler --> returns new handler for Photo entity
func NewPhotoHandler() PhotoHandler {
	return &photoHandler{
		repo: repository.NewPhotoRepository(),
	}
}

func (h *photoHandler) GetPhoto(ctx *gin.Context) {
	prodStr := ctx.Param("photo")
	prodID, err := strconv.Atoi(prodStr)
	userID := uint(ctx.MustGet("userID").(float64))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Photo, err := h.repo.GetPhoto(userID, prodID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, Photo)

}

func (h *photoHandler) AddPhoto(ctx *gin.Context) {
	var Photo models.Photo
	if err := ctx.ShouldBindJSON(&Photo); err != nil {

		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userID := uint(ctx.MustGet("userID").(float64))
	Photo.UserID = userID
	Photo, err := h.repo.AddPhoto(userID, Photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, "Photo has been created")

}
func (h *photoHandler) UpdatePhoto(ctx *gin.Context) {

	var updateData map[string]interface{}
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := ctx.Param("photo")
	intID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	userID := uint(ctx.MustGet("userID").(float64))
	repo := repository.NewPhotoRepository()
	photo, err := repo.GetPhoto(userID, intID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if photo.UserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you dont have permission"})
		return
	}

	photo, err = h.repo.UpdatePhoto(uint(intID), updateData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, "Photo has been updated")

}
func (h *photoHandler) DeletePhoto(ctx *gin.Context) {
	var Photo models.Photo

	id := ctx.Param("photo")
	intID, err := strconv.Atoi(id)
	Photo.ID = uint(intID)
	userID := uint(ctx.MustGet("userID").(float64))

	repo := repository.NewPhotoRepository()
	photo, err := repo.GetPhoto(userID, intID)
	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if photo.UserID != userID {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "you dont have permission"})
		return
	}

	Photos, err := h.repo.DeletePhoto(Photo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	ctx.JSON(http.StatusOK, Photos)

}
