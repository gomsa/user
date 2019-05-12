package auth

import (
	"context"
	"errors"
	"log"
	"strings"

	authClient "github.com/gomsa/user-srv/client"
	authPb "github.com/gomsa/user-srv/proto/auth"
	"github.com/gomsa/mpwechat-service/providers/config"

	"github.com/micro/go-micro/metadata"
	"github.com/micro/go-micro/server"
)


// AuthWrapper 是一个高阶函数，入参是 ”下一步“ 函数，出参是认证函数
// 在返回的函数内部处理完认证逻辑后，再手动调用 fn() 进行下一步处理
// token 是从 consignment-ci 上下文中取出的，再调用 user-service 将其做验证
// 认证通过则 fn() 继续执行，否则报错
func AuthWrapper(fn server.HandlerFunc) server.HandlerFunc {
	return func(ctx context.Context, req server.Request, resp interface{}) (err error) {
		meta, ok := metadata.FromContext(ctx)
		if !ok {
			return errors.New("no auth meta-data found in request")
		}
		// Note this is now uppercase (not entirely sure why this is...)
		token := strings.Split(meta["Authorization"], "Bearer ")[1]
		// Auth here
		authResp, err := authClient.Auth.ValidatePermission(context.Background(), &authPb.Request{
			Token:   token,
			Service: req.Service(),
			Method:  req.Method(),
		})
		if err != nil {
			return err
		}
		err = fn(ctx, req, resp)
		return err
	}
}
