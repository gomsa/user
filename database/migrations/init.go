package migrations

import (
	"context"

	"github.com/gomsa/tools/env"
	"github.com/gomsa/user/hander"
	permissionPB "github.com/gomsa/user/proto/permission"
	rolePB "github.com/gomsa/user/proto/role"
	userPB "github.com/gomsa/user/proto/user"
	casbinPB "github.com/gomsa/user/proto/casbin"
	db "github.com/gomsa/user/providers/database"
	"github.com/gomsa/user/service"
	"github.com/micro/go-micro/util/log"

	"github.com/gomsa/user/providers/casbin"
)

func init() {
	user()
	permission()
	role()

	seeds()
}

// user 用户数据迁移
func user() {
	user := &userPB.User{}
	if !db.DB.HasTable(&user) {
		db.DB.Exec(`
			CREATE TABLE users (
			id varchar(36) NOT NULL,
			username varchar(64) DEFAULT NULL,
			mobile varchar(11) DEFAULT NULL,
			email varchar(64) DEFAULT NULL,
			password varchar(128) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			avatar varchar(128) DEFAULT NULL,
			origin varchar(32) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY username (username),
			UNIQUE KEY mobile (mobile),
			UNIQUE KEY email (email)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}
// permission 权限数据迁移
func permission() {
	permission := &permissionPB.Permission{}
	if !db.DB.HasTable(&permission) {
		db.DB.Exec(`
			CREATE TABLE permissions (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			service varchar(64) DEFAULT NULL,
			method varchar(64) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			description varchar(128) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY service_OR_method (service,method)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

// role 角色数据迁移
func role() {
	role := &rolePB.Role{}
	if !db.DB.HasTable(&role) {
		db.DB.Exec(`
			CREATE TABLE roles (
			id int(11) unsigned NOT NULL AUTO_INCREMENT,
			label varchar(64) DEFAULT NULL,
			name varchar(64) DEFAULT NULL,
			description varchar(128) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id),
			UNIQUE KEY label (label)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}

func seeds() {
	CreateRole()
	CreateUser()
}

func CreateRole() {
	// 角色服务实现
	repo := &service.RoleRepository{db.DB}
	h := hander.Role{repo}
	req := &rolePB.Role{
		Label: `root`,
		Name: `超级管理员`,
		Description: `超级管理员拥有全部权限`,
	}
	res := &rolePB.Response{}
	err := h.Create(context.TODO(), req, res)
	// AddRole
	log.Log(err)
}

// CreateUser 填充文件
func CreateUser() {
	password := env.Getenv("ADMIN_PASSWORD", "123456")
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Username: `admin`,
		Password: password,
		Origin:   `user`,
	}
	res := &userPB.Response{}
	err := h.Create(context.TODO(), req, res)
	// 增加用户 root 权限
	addRole(res.User.Id,`root`)
	// AddRole
	log.Log(err)
}
// AddRole 增加用户角色
func addRole(userID string, role string) {
	h := hander.Casbin{casbin.Enforcer}
	req := &casbinPB.Request{
		UserID: userID,
		Role: role,
	}
	res := &casbinPB.Response{}
	err := h.AddRole(context.TODO(), req, res)
	log.Log(err)
}