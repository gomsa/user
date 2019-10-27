package gorm

import (
	"fmt"

	_ "github.com/alexbrainman/odbc"
	"github.com/jinzhu/gorm"
)

// odbcConnection 创建数据库连接
func odbcConnection(conf *Config) (*gorm.DB, error) {
	return gorm.Open(
		"odbc",
		fmt.Sprintf(
			"driver=freetds;server=%s;port=%s;database=%s;uid=%s;pwd=%s;TDS_Version=8.0;clientcharset=%s",
			conf.Host, conf.Port, conf.DbName, conf.User, conf.Password, conf.Charset,
		),
	)
}
