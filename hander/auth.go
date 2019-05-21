package hander

import (
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/gomsa/tools/uitl"
	pb "github.com/gomsa/user-srv/proto/auth"
	userPb "github.com/gomsa/user-srv/proto/user"
	"github.com/gomsa/user-srv/service"
)

// Auth 授权服务处理
type Auth struct {
	TokenService service.Authable
	Repo         service.Repository
}

// AuthById 只通过 id 获取 jwt token
// bug 只能服务之间调用如果前端调用会不验证获取权限
func (srv *Auth) AuthById(ctx context.Context, req *pb.User, res *pb.Token) (err error) {
	user, err := srv.Repo.Get(&userPb.User{
		Id: req.Id,
	})
	if err != nil {
		return err
	}
	err = uitl.Data2Data(user, req)
	if err != nil {
		return err
	}
	t, err := srv.TokenService.Encode(req)
	if err != nil {
		return err
	}
	res.Token = t
	return nil
}

// Auth 授权认证
// 返回token
func (srv *Auth) Auth(ctx context.Context, req *pb.User, res *pb.Token) (err error) {
	// 在 part3 中直接传参 &pb.User 去查找用户
	// 会导致 req 的值完全是数据库中的记录值
	// 即 req.Password 与 u.Password 都是加密后的密码
	// 将无法通过验证
	user, err := srv.Repo.Get(&userPb.User{
		Id:       req.Id,
		Username: req.Username,
		Email:    req.Email,
		Mobile:   req.Mobile,
	})
	if err != nil {
		return errors.New("用户不存在")
	}
	// 进行密码验证
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return errors.New("密码错误")
	}
	err = uitl.Data2Data(user, req)
	if err != nil {
		return err
	}
	t, err := srv.TokenService.Encode(req)
	if err != nil {
		return err
	}
	res.Token = t
	return nil
}

// ValidateToken 验证 token
// 并且验证 token 所属用户相关权限
func (srv *Auth) ValidateToken(ctx context.Context, req *pb.Request, res *pb.Token) (err error) {
	// Decode token
	claims, err := srv.TokenService.Decode(req.Token)
	if err != nil {
		return err
	}
	if claims.User.Id == "" {
		return errors.New("invalid user")
	}
	res.User = claims.User
	res.Valid = true
	return err
}
