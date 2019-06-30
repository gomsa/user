package main

import (
	"context"
	"fmt"
	"testing"

	_ "github.com/gomsa/user/database/migrations"
	db "github.com/gomsa/user/providers/database"

	"github.com/gomsa/user/hander"
	authPB "github.com/gomsa/user/proto/auth"
	permissionPB "github.com/gomsa/user/proto/permission"
	userPB "github.com/gomsa/user/proto/user"
	"github.com/gomsa/user/service"
)

func TestPermissionsUpdateOrCreate(t *testing.T) {
	req := &permissionPB.Permission{
		Service: "user-api", Method: "Auth.Auth1", Name: "用户授权3", Description: "用户登录授权返回 token 权限。",
	}
	repo := &service.PermissionRepository{db.DB}
	h := hander.Permission{repo}
	res := &permissionPB.Response{}
	err := h.UpdateOrCreate(context.TODO(), req, res)
	fmt.Println(req, res, err)
}
func TestUserCreate(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Username: `bvbv0115`,
		Password: `123456`,
		Mobile:   `13953186115`,
		Email:    `bvbv0a115@qq.com`,
		Name:     `bvbv0111`,
		Avatar:   `https://wpimg.wallstcn.com/f778738c-e4f8-4870-b634-56703b4acafe.gif`,
		Roles:    `admin,editor`,
		Origin:   `user`,
	}
	res := &userPB.Response{}
	err := h.Create(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestUserIsExist(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Username: `bvbv0111`,
		Mobile:   `13953186115`,
		Email:    `bvbv0a1@qq.com`,
	}
	res := &userPB.Response{}
	err := h.Exist(context.TODO(), req, res)
	// fmt.Println(req, res.Valid, err)
	t.Log(req, res, err)
}
func TestUserGet(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Username: `bvbv0111`,
	}
	res := &userPB.Response{}
	err := h.Get(context.TODO(), req, res)
	// fmt.Println("UserGet", res, err)
	t.Log(req, res, err)
}

func TestUserList(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.ListQuery{
		Limit: 20,
		Page:  1,
		Sort:  "created_at desc",
	}
	res := &userPB.Response{}
	err := h.List(context.TODO(), req, res)
	// fmt.Println("UserList", res, err)
	t.Log(req, res, err)
}
func TestUserUpdate(t *testing.T) {
	// repo := &service.UserRepository{db.DB}
	// h := hander.User{repo}
	// req := &userPB.User{
	// 	Id:       `66527a06-5a16-4bab-8ed8-6cd66ff867b3`,
	// 	Username: `bvbv0111`,
	// 	Name:     `newbvbv`,
	// }
	// res := &userPB.Response{}
	// err := h.Update(context.TODO(), req, res)
	// fmt.Println("UserUpdate", req, res, err)
	// t.Log(req, res, err)
}
func TestUserDelete(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	h := hander.User{repo}
	req := &userPB.User{
		Id: `66527a06-5a16-4bab-8ed8-6cd66ff867b3`,
	}
	res := &userPB.Response{}
	err := h.Delete(context.TODO(), req, res)
	// fmt.Println("UserDelete", req, res, err)
	t.Log(req, res, err)
}

// Auth
func TestAuth(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.User{
		Username: `bvbv0111`,
		Password: `123456`,
	}
	res := &authPB.Token{}
	err := h.Auth(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestAuthById(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.User{
		Id: `93267c95-ef44-4e49-b865-0ad6c84c646c`,
	}
	res := &authPB.Token{}
	err := h.AuthById(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}

func TestValidateToken(t *testing.T) {
	repo := &service.UserRepository{db.DB}
	token := &service.TokenService{}
	h := hander.Auth{token, repo}
	req := &authPB.Request{
		Token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVc2VyIjp7ImlkIjoiOTMyNjdjOTUtZWY0NC00ZTQ5LWI4NjUtMGFkNmM4NGM2NDZjIiwibmFtZSI6ImJ2YnYwMTEiLCJtb2JpbGUiOiIxMzk1NDM4NjExNCIsImVtYWlsIjoiYnZidjBhMUBxcS5jb20iLCJwYXNzd29yZCI6IiQyYSQxMCRFM3VDdzRpa3JNaXhaSjhWd21tYmRlMERyUlA3WDNaM0xOMEdMZ2Z5SWxNTWx0MGlZazZDRyIsIm9yaWdpbiI6InVzZXItc2VydmljZSIsImNyZWF0ZWRfYXQiOiIyMDE5LTA1LTA2VDE2OjI1OjMyKzA4OjAwIiwidXBkYXRlZF9hdCI6IjIwMTktMDUtMDZUMTY6MjU6MzIrMDg6MDAifSwiZXhwIjoxNTU3NDUyMDEzLCJpc3MiOiJ1c2VyLXNlcnZpY2UifQ.kQKnSW-qDGyVf3XahRFkkTx5wP_PyIyM_RGqiBK4pfI`,
	}
	res := &authPB.Token{}
	err := h.ValidateToken(context.TODO(), req, res)
	// fmt.Println(req, res, err)
	t.Log(req, res, err)
}
