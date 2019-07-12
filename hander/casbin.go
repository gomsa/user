package hander

import (
	"context"
	"errors"

	"github.com/casbin/casbin"
	"github.com/micro/go-micro/metadata"

	pb "github.com/gomsa/user/proto/casbin"
)

// Casbin 权限管理
type Casbin struct {
	Enforcer *casbin.Enforcer
}

// AddPermission 添加权限
func (srv *Casbin) AddPermission(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Enforcer.AddPermissionForUser(req.Role, []string{req.Permission}...)
	if !res.Valid {
		return errors.New("添加权限失败")
	}
	return err
}

// DeletePermissions 根据角色名删除权限
func (srv *Casbin) DeletePermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Enforcer.DeletePermissionsForUser(req.Role)
	return err
}

// UpdatePermissions 重新设置角色所有权限
func (srv *Casbin) UpdatePermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Enforcer.DeletePermissionsForUser(req.Role)
	for _, permission := range req.Permissions {
		res.Valid = srv.Enforcer.AddPermissionForUser(req.Role, []string{permission}...)
		if !res.Valid {
			return errors.New("添加权限失败")
		}
	}
	return err
}

// GetPermissions 获取权限
func (srv *Casbin) GetPermissions(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	permissions := srv.Enforcer.GetPermissionsForUser(req.Role)
	for _, permission := range permissions {
		res.Permissions = append(res.Permissions, permission[1])
	}
	return err
}

////////////

// AddRole 添加权限
func (srv *Casbin) AddRole(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Enforcer.AddRoleForUser(req.UserID, req.Role)
	if !res.Valid {
		return errors.New("添加角色失败")
	}
	return err
}

// DeleteRoles 根据角色名删除权限
func (srv *Casbin) DeleteRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Enforcer.DeleteRolesForUser(req.UserID)
	return err
}

// UpdateRoles 重新设置角色所有权限
func (srv *Casbin) UpdateRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Valid = srv.Enforcer.DeleteRolesForUser(req.UserID)
	for _, role := range req.Roles {
		res.Valid = srv.Enforcer.AddRoleForUser(req.UserID, role)
		if !res.Valid {
			return errors.New("添加角色失败")
		}
	}
	return err
}

// GetRoles 获取权限
func (srv *Casbin) GetRoles(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	res.Roles, err = srv.Enforcer.GetRolesForUser(req.UserID)
	return err
}

// Validate 验证权限
func (srv *Casbin) Validate(ctx context.Context, req *pb.Request, res *pb.Response) (err error) {
	meta, ok := metadata.FromContext(ctx)
	if !ok {
		return errors.New("no auth meta-data found in request")
	}
	if userID, ok := meta["user_id"]; ok {
		res.Valid = srv.Enforcer.Enforce(userID, meta["service"], meta["method"])
		if !res.Valid {
			return errors.New("您没有权限访问")
		}
	} else {
		return errors.New("Empty userID")
	}
	return err
}
