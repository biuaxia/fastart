package core

import "gorm.io/gorm/schema"

const (
	// authentication key of cookie
	COOKIE_AUTH_KEY = "_ak"

	USERNAME_KEY = "_username"
	PASSWORD_KEY = "_password"

	DEFAULT_SERVER_PORT = 6179

	// db table's prefix. fart10_ means current version is fart:1.0.x
	TABLE_PREFIX = "fart10_"

	VERSION = "1.0.0"
)

type Config interface {
	Installed() bool
	ServerPort() int
	// get the db type
	DbType() string
	// get the mysql url. eg. fart:fart123@tcp(127.0.0.1:3306)/fart?charset=utf8&parseTime=True&loc=Local
	MysqlUrl() string
	// get the sqlite path
	SqliteFolder() string
	// files storage location.
	MatterPath() string
	// table name strategy
	NamingStrategy() schema.NamingStrategy
	// when installed by user. Write configs to fart.json
	FinishInstall(dbType string, mysqlPort int, mysqlHost string, mysqlSchema string, mysqlUsername string, mysqlPassword string, mysqlCharset string)
}
