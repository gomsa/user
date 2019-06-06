package main

import (
	// 公共引入
	"github.com/micro/go-log"
	micro "github.com/micro/go-micro"
	k8s "github.com/micro/kubernetes/go/micro"

	"github.com/gomsa/tools/config"
	// 执行数据迁移
	_ "github.com/gomsa/user-srv/database/migrations"

	"github.com/gomsa/user-srv/hander"
	authPB "github.com/gomsa/user-srv/proto/auth"
	permissionPB "github.com/gomsa/user-srv/proto/permission"
	rolePB "github.com/gomsa/user-srv/proto/role"
	userPB "github.com/gomsa/user-srv/proto/user"
	db "github.com/gomsa/user-srv/providers/database"
	"github.com/gomsa/user-srv/service"
)

func main() {
	// 加载配置
	conf := &config.Config{}
	conf.LoadFile("config.yaml")

	srv := k8s.NewService(
		micro.Name(conf.GetApp()),
		micro.Version(conf.GetVersion()),
	)
	srv.Init()

	// 用户服务实现
	repo := &service.UserRepository{db.DB}
	userPB.RegisterUsersHandler(srv.Server(), &hander.User{repo})

	// token 服务实现
	token := &service.TokenService{}
	authPB.RegisterAuthHandler(srv.Server(), &hander.Auth{token, repo})

	// 权限服务实现
	prepo := &service.PermissionRepository{db.DB}
	permissionPB.RegisterPermissionsHandler(srv.Server(), &hander.Permission{prepo})

	// 角色服务实现
	rrepo := &service.RoleRepository{db.DB}
	rolePB.RegisterRolesHandler(srv.Server(), &hander.Role{rrepo})
	// Run the server
	if err := srv.Run(); err != nil {
		log.Log(err)
	}
	log.Log("serviser run ...")
}
