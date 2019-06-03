package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/user-srv/proto/role"
	"github.com/gomsa/user-srv/service"
)

// Role 角色结构
type Role struct {
	Repo service.RRepository
}

// List 获取所有角色
func (srv *Role) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	roles, err := srv.Repo.List(req)
	total, err := srv.Repo.Total(req)
	if err != nil {
		return err
	}
	res.Roles = roles
	res.Total = total
	return err
}

// Get 获取角色
func (srv *Role) Get(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	role, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Role = role
	return err
}

// Create 创建角色
func (srv *Role) Create(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加角色失败")
	}
	res.Valid = true
	return err
}

// Update 更新角色
func (srv *Role) Update(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新角色失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除角色
func (srv *Role) Delete(ctx context.Context, req *pb.Role, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除角色失败")
	}
	res.Valid = valid
	return err
}
