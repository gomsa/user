package middleware

import (
	"context"
	"fmt"
	"testing"

	"github.com/micro/go-micro/server"
)

func TestConfig(t *testing.T) {
	Pata := []map[string]interface{}{
		{"service": "Users", "method": "Create", "auth": false, "policy": false, "display_name": "创建用户", "description": "创建新用户权限。"},
		{"service": "Users", "method": "Exist", "auth": false, "policy": false, "display_name": "检测用户", "description": "检测用户是否存在权限。"},
		{"service": "Users", "method": "Get", "auth": true, "policy": false, "display_name": "用户查询", "description": "查询用户信息权限。"},
		{"service": "Auth", "method": "Auth", "auth": false, "policy": false, "display_name": "用户授权", "description": "用户登录授权返回 token 权限。"},
		{"service": "Auth", "method": "ValidateToken", "auth": false, "policy": false, "display_name": "权限认证", "description": "权限相关认证权限。"},
	}
	h := Handler{Permission{
		Data: Pata,
	}}
	hf := h.Wrapper(HandlerFunc)
	err := hf(context.Background(), nil, nil)
	fmt.Println(err)
}
func HandlerFunc(ctx context.Context, req server.Request, rsp interface{}) error {
	fmt.Println(ctx, req, rsp)
	return nil
}
