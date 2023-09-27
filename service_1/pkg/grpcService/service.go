package service

import (
	"context"

	"github.com/ajalck/service_1/pkg/pb"
	"gorm.io/gorm"
)

type UserServer struct {
	DB *gorm.DB
	pb.UnimplementedUsersServer
}

func (u *UserServer) ListUsers(c context.Context, req *pb.RequestParams) (*pb.ResponseData, error) {
	users := &pb.ResponseData{}
	var countOfUsers int64
	result := u.DB.Table("users").Select("id", "CONCAT(first_name, ' ', last_name) as full_name").Find(&users).Count(&countOfUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
