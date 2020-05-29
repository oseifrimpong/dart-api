package delivery

import (
	"net/http"

	"dart-api/api/model"
	"dart-api/api/service"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

type userController struct {
	// put logs here
	svc service.Service
}

func NewUserController(svc service.Service) *userController {
	return &userController{svc}
}

func (u *userController) GetAll(ctx *gin.Context) {
	users, err := u.svc.GetAll()
	if len(users) == 0 || err != nil {
		ctx.Status(http.StatusNoContent)
		return
	}
	ctx.SecureJSON(http.StatusOK, users)
}

func (u *userController) GetByID(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	user, err := u.svc.GetByID(id)
	if user == nil || err != nil {
		ctx.Status(http.StatusNotFound)
		return
	}
	ctx.SecureJSON(http.StatusOK, user)
}

func (u *userController) Create(ctx *gin.Context) {
	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.SecureJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u.svc.Create(&user)
	ctx.Status(http.StatusCreated)
}

// update func goes here
func (u *userController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	var user model.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = id
	u.svc.Update(&user)
	ctx.Status(http.StatusOK)
}

func (u *userController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if _, err := uuid.FromString(id); err != nil {
		ctx.Status(http.StatusBadRequest)
		return
	}

	u.svc.Delete(id)
	ctx.Status(http.StatusNoContent)
}
