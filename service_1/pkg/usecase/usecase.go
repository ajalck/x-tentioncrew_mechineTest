package usecase

import (
	"github.com/ajalck/service_1/pkg/models"
	repository "github.com/ajalck/service_1/pkg/repository/interfaces"
	usecase "github.com/ajalck/service_1/pkg/usecase/interfaces"
)

type userUseCase struct {
	repo repository.UserRepo
}

func NewUserUseCase(repo repository.UserRepo) usecase.UserUseCase {
	return &userUseCase{repo}
}

func (u userUseCase) CreateUser(body *models.User) error {

	err := u.repo.CreateUser(body)
	if err != nil {
		return err
	}
	return nil

}
func (u userUseCase) GetUserByID(id uint) (*models.User, error) {

	user, err := u.repo.GetUserByID(id)
	if err != nil {
		return &models.User{}, err
	}
	return user, nil

}
func (u userUseCase) UpdateUser(id uint, body *models.User) error {

	err := u.repo.UpdateUser(id, body)
	if err != nil {
		return err
	}
	return nil

}
func (u userUseCase) DeleteUser(id uint) error {

	err := u.repo.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil

}
