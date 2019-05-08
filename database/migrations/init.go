package migrations

import (
	pb "github.com/gomsa/user-srv/proto/user"
	db "github.com/gomsa/user-srv/providers/database"
)

func init() {
	user()
}

// user 用户数据迁移
func user() {
	user := &pb.User{}
	if !db.DB.HasTable(&user) {
		db.DB.Exec(`
			CREATE TABLE users (
			id varchar(36) NOT NULL,
			name varchar(64) DEFAULT NULL,
			mobile varchar(11) DEFAULT NULL,
			email varchar(64) DEFAULT NULL,
			password varchar(128) DEFAULT NULL,
			origin varchar(32) DEFAULT NULL,
			created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			xxx_unrecognized varbinary(255) DEFAULT NULL,
			xxx_sizecache int(11) DEFAULT NULL,
			PRIMARY KEY (id)
			) ENGINE=InnoDB DEFAULT CHARSET=utf8;
		`)
	}
}
