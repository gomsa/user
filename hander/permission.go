package hander

import (
	"context"
	"fmt"

	pb "github.com/gomsa/user-srv/proto/permission"
	"github.com/gomsa/user-srv/service"
)

// Permission 权限结构
type Permission struct {
	Repo service.PRepository
}

// List 获取所有权限
func (srv *Permission) List(ctx context.Context, req *pb.ListQuery, res *pb.Response) (err error) {
	permissions, err := srv.Repo.List(req)
	if err != nil {
		return err
	}
	res.Permissions = permissions
	return err
}

// Get 获取权限
func (srv *Permission) Get(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	permission, err := srv.Repo.Get(req)
	if err != nil {
		return err
	}
	res.Permission = permission
	return err
}

// Create 创建权限
func (srv *Permission) Create(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	_, err = srv.Repo.Create(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("添加权限失败")
	}
	res.Valid = true
	return err
}

// Update 更新权限
func (srv *Permission) Update(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	valid, err := srv.Repo.Update(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("更新权限失败")
	}
	res.Valid = valid
	return err
}

// Delete 删除权限
func (srv *Permission) Delete(ctx context.Context, req *pb.Permission, res *pb.Response) (err error) {
	valid, err := srv.Repo.Delete(req)
	if err != nil {
		res.Valid = false
		return fmt.Errorf("删除权限失败")
	}
	res.Valid = valid
	return err
}
