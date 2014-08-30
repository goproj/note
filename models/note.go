package models

import (
	"github.com/astaxie/beego/orm"
)

func (this *Note) FillAttrs() bool {
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

func (this *Note) Insert() error {
	o := orm.NewOrm()
	id, err := o.Insert(this)
	if err != nil {
		return err
	}
	this.Id = int(id)
	return nil
}

func (this *Note) Delete() error {
	if _, err := orm.NewOrm().Delete(this); err != nil {
		return err
	}
	return nil
}

func (this *Note) Update(fields ...string) error {
	if _, err := orm.NewOrm().Update(this, fields...); err != nil {
		return err
	}
	return nil
}

func TodoNotes(userId int) []*Note {
	var notes []*Note
	Notes().Filter("UserId", userId).Filter("Done", 0).All(&notes)
	return notes
}

func DoneNotes(userId int) []*Note {
	var notes []*Note
	Notes().Filter("UserId", userId).Filter("Done", 1).All(&notes)
	return notes
}

func Notes() orm.QuerySeter {
	return orm.NewOrm().QueryTable("note").OrderBy("-Id")
}
