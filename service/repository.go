package service

import (
	"fmt"

	"github.com/micro/go-log"

	pb "github.com/gomsa/user-srv/proto/user"

	"github.com/jinzhu/gorm"
)

//Repository 仓库接口
type Repository interface {
	Create(user *pb.User) (*pb.User, error)
	Exist(user *pb.User) bool
	Get(user *pb.User) (*pb.User, error)
	GetAll() ([]*pb.User, error)
}

// UserRepository 用户仓库
type UserRepository struct {
	DB *gorm.DB
}

// Create 创建用户
// bug 无用户名创建用户可能引起 bug
func (repo *UserRepository) Create(user *pb.User) (*pb.User, error) {
	if exist := repo.Exist(user); exist == true {
		return user, fmt.Errorf("注册用户已存在")
	}
	err := repo.DB.Create(user).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return user, fmt.Errorf("注册用户失败")
	}
	return user, nil
}

// Exist 检测用户是否已经存在
func (repo *UserRepository) Exist(user *pb.User) bool {
	var count int
	if user.Name != "" {
		repo.DB.Model(&user).Where("name = ?", user.Name).Count(&count)
		if count > 0 {
			return true
		}
	}
	if user.Mobile != "" {
		repo.DB.Model(&user).Where("mobile = ?", user.Mobile).Count(&count)
		if count > 0 {
			return true
		}
	}
	if user.Email != "" {
		repo.DB.Model(&user).Where("email = ?", user.Email).Count(&count)
		if count > 0 {
			return true
		}
	}
	return false
}

// Get 获取用户信息
func (repo *UserRepository) Get(user *pb.User) (*pb.User, error) {
	if user.Id != "" {
		if err := repo.DB.Model(&user).Where("id = ?", user.Id).Find(&user).Error; err != nil {
			return nil, err
		}
	}
	if user.Name != "" {
		if err := repo.DB.Model(&user).Where("name = ?", user.Name).Find(&user).Error; err != nil {
			return nil, err
		}
	}
	if user.Mobile != "" {
		if err := repo.DB.Model(&user).Where("mobile = ?", user.Mobile).Find(&user).Error; err != nil {
			return nil, err
		}
	}
	if user.Email != "" {
		if err := repo.DB.Model(&user).Where("email = ?", user.Email).Find(&user).Error; err != nil {
			return nil, err
		}
	}
	return user, nil
}

// GetAll 获取所有用户信息
func (repo *UserRepository) GetAll() (users []*pb.User, err error) {
	if err := repo.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
