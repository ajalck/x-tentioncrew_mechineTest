package interfaces

import(
	"github.com/ajalck/service_1/pkg/models"
)

type UserUseCase interface {
	CreateUser(body *models.User) error
	GetUserByID(id uint)(*models.User, error)
	UpdateUser(id uint, body *models.User) error
	DeleteUser(id uint) error
}
