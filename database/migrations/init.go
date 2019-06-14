package migrations

import (
	permissionPB "github.com/gomsa/user-srv/proto/permission"
	rolePB "github.com/gomsa/user-srv/proto/role"
	userPD "github.com/gomsa/user-srv/proto/user"
	db "github.com/gomsa/user-srv/providers/database"
)

func init() {
	user()
	permission()
	role()
}

// user 用户数据迁移
func user() {
	user := &userPD.User{}
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
