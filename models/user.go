package models

import (
	"github.com/astaxie/beego/orm"
)

func (this *User) FillAttrs() bool {
	o := orm.NewOrm()
	err := o.Read(this)
	if err == orm.ErrNoRows {
		return false
	} else if err == orm.ErrMissPK {
		return false
	} else {
		return true
	}
}

func (this *User) Insert() error {
	o := orm.NewOrm()
	id, err := o.Insert(this)
	if err != nil {
		return err
	}
	this.Id = int(id)
	return nil
}

func (this *User) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func FindUserByEmail(email string) *User {
	user := User{Email: email}
	err := orm.NewOrm().Read(&user, "Email")
	if err != nil {
		return nil
	}

	return &user
}

func FindUserByName(name string) *User {
	user := User{Name: name}
	err := orm.NewOrm().Read(&user, "Name")
	if err != nil {
		return nil
	}

	return &user
}
