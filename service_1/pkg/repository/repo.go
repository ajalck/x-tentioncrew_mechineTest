package repository

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/ajalck/service_1/pkg/models"

	repoInt "github.com/ajalck/service_1/pkg/repository/interfaces"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type userRepo struct {
	sDB *gorm.DB
	rDB *redis.Client
}

func NewUserRepo(sdb *gorm.DB, rdb *redis.Client) repoInt.UserRepo {
	return &userRepo{sdb, rdb}
}
func (r *userRepo) CreateUser(body *models.User) error {
	user := models.User{
		FirstName: body.FirstName,
		LastName:  body.LastName,
		Email:     body.Email,
		Phone:     body.Phone,
	}
	result := r.sDB.Create(&user)

	if result.Error != nil {
		return result.Error
	}
	jsonbytes, _ := json.Marshal(&user)
	fmt.Println(string(jsonbytes))
	key := strconv.Itoa(int(user.ID))
	err := r.rDB.Set(context.Background(), key, string(jsonbytes), 30*24*time.Hour).Err()
	if err != nil {
		return err
	}
	return nil
}
func (r *userRepo) GetUserByID(id uint) (*models.User, error) {
	user := &models.User{}
	key := strconv.Itoa(int(id))
	result, _ := r.rDB.Get(context.Background(), key).Result()
	if result == "" {
		r.sDB.Model(models.User{}).Where("id", id).First(&user)
		return user, nil
	}
	if len(result) > 0 {
		if err := json.Unmarshal([]byte(result), &user); err != nil {
			return user, err
		}
		if user.ID == 0 {
			return user, errors.New("user not found")
		}
	}
	return user, nil
}
func (r *userRepo) UpdateUser(id uint, body *models.User) error {

	result := r.sDB.Model(&models.User{}).Where("id", id).
		Updates(&models.User{FirstName: body.FirstName, LastName: body.LastName, Email: body.Email, Phone: body.Phone})
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected != 1 {
		return errors.New("no user found")
	}
	jsonbytes, _ := json.Marshal(&body)
	key := strconv.Itoa(int(id))
	res := r.rDB.GetSet(context.Background(), key, string(jsonbytes))
	if res.Err() != redis.Nil {
		return res.Err()
	}
	return nil

}
func (r *userRepo) DeleteUser(id uint) error {
	result := r.sDB.Where("id", id).Delete(&models.User{})
	if result.Error != nil {
		return result.Error
	}
	key := strconv.Itoa(int(id))
	res := r.rDB.Del(context.Background(), key)
	if res.Err() != redis.Nil {
		return res.Err()
	}
	return nil
}
