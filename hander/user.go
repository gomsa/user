package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/user-srv/proto/user"
	"github.com/gomsa/user-srv/service"

	"golang.org/x/crypto/bcrypt"
)

// User 用户结构
type User struct {
	Repo service.Repository
}

// Create 创建用户
func (srv *User) Create(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	// Generates a hashed version of our password
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPass)
	user, err := srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("创建用户失败:%v", err)
	}
	res.User = user
	res.Valid = true
	return err
}

// Exist 用户是否存在
func (srv *User) Exist(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	res.Valid = srv.Repo.Exist(req)
	return err
}

// Get 获取用户
func (srv *User) Get(ctx context.Context, req *pb.User, res *pb.Response) (err error) {
	user, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.User = user
	return err
}

// GetAll 获取所有用户
func (srv *User) GetAll(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	users, err := srv.Repo.GetAll()
	if err != nil {
		return err
	}
	res.Users = users
	return err
}
