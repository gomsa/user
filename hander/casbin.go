package hander

import (
	"context"
	"fmt"

	"github.com/casbin/casbin"
	pb "github.com/gomsa/user-srv/proto/casbin"
)

// Casbin 权限管理
type Casbin struct {
	Enforcer *casbin.Enforcer
}

// AddPermissions 添加权限
func (srv *Casbin) AddPermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	for _, permission := range req.Permissions {
		res.Valid = srv.Enforcer.AddPermissionForUser(req.User, []string{permission.Service, permission.Method}...)
		if !res.Valid {
			return fmt.Errorf("添加权限失败")
		}
	}
	return err
}

// DeletePermissions 根据角色名删除权限
func (srv *Casbin) DeletePermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Enforcer.DeletePermissionsForUser(req.User)
	return err
}

// GetPermissions 获取权限
func (srv *Casbin) GetPermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	permissions := srv.Enforcer.GetPermissionsForUser(req.User)
	for _, permission := range permissions {
		res.Permissions = append(res.Permissions, &pb.Permission{
			Service: permission[0],
			Method:  permission[1],
		})
	}
	return err
}
