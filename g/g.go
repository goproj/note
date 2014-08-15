package g

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	CookieName = "note"
)

var (
	Static    string
	SecretKey string
	RunMode   string
)

func ReadConfig() {
	Static = beego.AppConfig.String("static")
	SecretKey = beego.AppConfig.String("secret_key")
	RunMode = beego.AppConfig.String("runmode")
}

func InitDB() {
	dbuser := beego.AppConfig.String("dbuser")
	dbpass := beego.AppConfig.String("dbpass")
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbname := beego.AppConfig.String("dbname")
	dblink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbuser, dbpass, dbhost, dbport, dbname)
	// dblink = "root:1234@/uic?charset=utf8&loc=Asia%2FChongqing"

	orm.RegisterDriver("mysql", orm.DR_MySQL)
	orm.RegisterDataBase("default", "mysql", dblink+"&loc=Asia%2FChongqing", 30, 200)
	// orm.DefaultTimeLoc = time.UTC

	if RunMode == "dev" {
		orm.Debug = true
	}
}
