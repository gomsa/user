package casbin

import (
	"fmt"

	"github.com/casbin/casbin"
	"github.com/casbin/gorm-adapter"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gomsa/tools/env"
)

// DB 管理包
var (
	// DB 连接
	Enforcer *casbin.Enforcer
)

func init() {
	Driver := env.Getenv("DB_DRIVER", "mysql")
	// Host 主机地址
	Host := env.Getenv("DB_HOST", "127.0.0.1")
	// Port 主机端口
	Port := env.Getenv("DB_PORT", "3306")
	// User 用户名
	User := env.Getenv("DB_USER", "root")
	// Password 密码
	Password := env.Getenv("DB_PASSWORD", "123456")
	// DbName 数据库名称
	DbName := env.Getenv("DB_NAME", "srv_user")
	// Charset 数据库编码
	Charset := env.Getenv("DB_CHARSET", "utf8")
	a := gormadapter.NewAdapter(Driver, fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
		User, Password, Host, Port, DbName, Charset,
	), true) // Your driver and data source.
	Enforcer = casbin.NewEnforcer("rbac_model.conf", a)
	Enforcer.LoadPolicy()
}
