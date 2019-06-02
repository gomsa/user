package service

import (
	"fmt"
	// 公共引入
	"github.com/micro/go-log"

	pb "github.com/gomsa/user-srv/proto/role"

	"github.com/jinzhu/gorm"
)

//RRepository 仓库接口
type RRepository interface {
	Create(role *pb.Role) (*pb.Role, error)
	Delete(role *pb.Role) (bool, error)
	Update(role *pb.Role) (bool, error)
	Get(role *pb.Role) (*pb.Role, error)
	List(req *pb.ListQuery) ([]*pb.Role, error)
}

// RoleRepository 角色仓库
type RoleRepository struct {
	DB *gorm.DB
}

// List 获取所有角色信息
func (repo *RoleRepository) List(req *pb.ListQuery) (roles []*pb.Role, err error) {
	db := repo.DB
	// 分页
	var limit, offset int64
	if req.Limit > 0 {
		limit = req.Limit
	} else {
		limit = 10
	}
	if req.Page > 1 {
		offset = (req.Page - 1) * limit
	} else {
		offset = -1
	}

	// 排序
	var sort string
	if req.Sort != "" {
		sort = req.Sort
	} else {
		sort = "created_at desc"
	}
	// 查询条件
	if req.Name != "" {
		db = db.Where("rolename like ?", "%"+req.Name+"%")
	}
	if err := db.Order(sort).Limit(limit).Offset(offset).Find(&roles).Error; err != nil {
		log.Log(err)
		return nil, err
	}
	return roles, nil
}

// Get 获取角色信息
func (repo *RoleRepository) Get(p *pb.Role) (*pb.Role, error) {
	if p.Id > 0 {
		if err := repo.DB.Model(&p).Where("id = ?", p.Id).Find(&p).Error; err != nil {
			return nil, err
		}
	}
	if p.Name != "" {
		if err := repo.DB.Model(&p).Where("name = ?", p.Name).Find(&p).Error; err != nil {
			return nil, err
		}
	}
	if p.DisplayName != "" {
		if err := repo.DB.Model(&p).Where("display_name = ?", p.DisplayName).Find(&p).Error; err != nil {
			return nil, err
		}
	}
	return p, nil
}

// Create 创建角色
// bug 无角色名创建角色可能引起 bug
func (repo *RoleRepository) Create(p *pb.Role) (*pb.Role, error) {
	err := repo.DB.Create(p).Error
	if err != nil {
		// 写入数据库未知失败记录
		log.Log(err)
		return p, fmt.Errorf("添加角色失败")
	}
	return p, nil
}

// Update 更新角色
func (repo *RoleRepository) Update(p *pb.Role) (bool, error) {
	if p.Id > 0 {
		return false, fmt.Errorf("请传入更新id")
	}
	id := &pb.Role{
		Id: p.Id,
	}
	err := repo.DB.Model(id).Updates(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}

// Delete 删除角色
func (repo *RoleRepository) Delete(p *pb.Role) (bool, error) {
	err := repo.DB.Delete(p).Error
	if err != nil {
		log.Log(err)
		return false, err
	}
	return true, nil
}
