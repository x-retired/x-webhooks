package initialize

import (
	_ "github.com/xiexianbin/webhooks/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// Init Database Connection
func InitDatabase() {
	// db type
	dbType := beego.AppConfig.String("db_type")
	// db connection name
	dbAlias := beego.AppConfig.String(dbType + "::db_alias")
	// db name
	dbName := beego.AppConfig.String(dbType + "::db_name")

	switch dbType {
	case "sqlite3":
		orm.RegisterDataBase(dbAlias, dbType, dbName)
	case "mysql":
		// db host or domain
		dbHost := beego.AppConfig.String(dbType + "::db_host")
		// db port
		dbPort := beego.AppConfig.String(dbType + "::db_port")
		// db username
		dbUser := beego.AppConfig.String(dbType + "::db_user")
		// db password
		dbPwd := beego.AppConfig.String(dbType + "::db_pwd")
		// db charset
		dbCharset := beego.AppConfig.String(dbType + "::db_charset")
		orm.RegisterDataBase(dbAlias, dbType, dbUser+":"+dbPwd+"@tcp("+dbHost+":"+
			dbPort+")/"+dbName+"?charset="+dbCharset, 30)
	}

	// if dev mode, show db sql info
	isDev := beego.AppConfig.String("runmode") == "dev"
	// sync db
	orm.RunSyncdb("default", false, isDev)
	if isDev {
		orm.Debug = isDev
	}
}
