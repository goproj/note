package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	Id       int       `orm:"auto;pk"`
	Name     string    `orm:"size(64);unique;index"`
	Email    string    `orm:"size(64);unique;index"`
	Password string    `orm:"size(64);null"`
	Portrait string    `orm:"null"`
	Blocked  int8      `orm:"type(tinyint)"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime);column(create_at)"`
}

type Note struct {
	Id       int       `orm:"auto;pk"`
	UserId   int       `orm:"column(user_id);index"`
	Content  string    `orm:"size(1024)"`
	Done     int8      `orm:"type(tinyint)"`
	DoneAt   time.Time `orm:"null;type(datetime);column(done_at)"`
	CreateAt time.Time `orm:"auto_now_add;type(datetime);column(create_at)"`
}

func (u *User) TableName() string {
	return "user"
}

func (u *Note) TableName() string {
	return "note"
}

func (u *User) TableEngine() string {
	return "INNODB"
}

func (u *Note) TableEngine() string {
	return "INNODB"
}

func init() {
	orm.RegisterModel(new(User), new(Note))
}

// func main() {
// 	orm.RegisterDataBase("default", "mysql", "root:@/todo_note?charset=utf8&loc=Asia%2FShanghai", 30, 200)
// 	orm.RunCommand()
// }
