package controllers

import (
	"encoding/json"
	"github.com/goproj/note/g"
	"github.com/goproj/note/models"
	"strconv"
	"fmt"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) CheckLogin() {
	if this.RetrieveUserFromCookie() {
		return
	}

	this.Abort("401")
}


func (this *NoteController) RetrieveUserFromCookie() bool {
	uid, succ := this.Ctx.GetSecureCookie(g.SecretKey, g.CookieName)
	if !succ {
		return false
	}

	id, err := strconv.ParseInt(uid, 10, 64)
	if err != nil {
		return false
	}

	if id <= 0 {
		return false
	}

	u := &models.User{Id: int(id)}
	if u.FillAttrs() {
		this.CurrentUser = u
	} else {
		return false
	}

	if u.Blocked == 1 {
		return false
	}

	return true
}

func (this *NoteController) Read() {
	idStr := this.Ctx.Input.Params[":id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	n := models.Note{Id: id}
	if !n.FillAttrs() {
		this.ServeErrMsg("no such note")
		return
	}

	if n.UserId != this.CurrentUser.Id {
		this.ServeErrMsg("no privilege")
		return
	}

	this.ServeDataJson(n)

}

func (this *NoteController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	n := models.Note{Id: id}
	if !n.FillAttrs() {
		this.ServeErrMsg("no such note")
		return
	}

	if n.UserId != this.CurrentUser.Id {
		this.ServeErrMsg("no privilege")
		return
	}

	if err = n.Delete(); err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	this.ServeOk()
}

func (this *NoteController) Update() {
	idStr := this.Ctx.Input.Params[":id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	n := models.Note{Id: id}
	if !n.FillAttrs() {
		this.ServeErrMsg("no such note")
		return
	}

	if n.UserId != this.CurrentUser.Id {
		this.ServeErrMsg("no privilege")
		return
	}

	type JsonInfo struct {
		Content string `json:"content"`
	}

	var jsonInfo JsonInfo
	if json.Unmarshal(this.Ctx.Input.RequestBody, &jsonInfo) != nil {
		this.ServeErrMsg("input format error")
		return
	}

	n.Content = jsonInfo.Content
	if err = n.Update("Content"); err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	this.ServeOk()
}

func (this *NoteController) Add() {
	type JsonInfo struct {
		Content string `json:"content"`
	}

	var jsonInfo JsonInfo
	if json.Unmarshal(this.Ctx.Input.RequestBody, &jsonInfo) != nil {
		this.ServeErrMsg("input format error")
		return
	}

	if jsonInfo.Content == "" {
		this.ServeErrMsg("content is blank")
		return
	}

	fmt.Println(this.CurrentUser)

	n := models.Note{Content: jsonInfo.Content, UserId: this.CurrentUser.Id}
	if e := n.Insert(); e != nil {
		this.ServeErrMsg(e.Error())
	} else {
		this.ServeDataJson(n.Id)
	}
}

func (this *NoteController) Mark() {
	idStr := this.Ctx.Input.Params[":id"]
	doneStr := this.Ctx.Input.Params[":done"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.ServeErrMsg("id format error")
		return
	}

	var done int
	done, err = strconv.Atoi(doneStr)
	if err != nil {
		this.ServeErrMsg("done format error")
		return
	}

	n := models.Note{Id: id}
	if !n.FillAttrs() {
		this.ServeErrMsg("no such note")
		return
	}

	if n.UserId != this.CurrentUser.Id {
		this.ServeErrMsg("no privilege")
		return
	}

	n.Done = int8(done)
	if err = n.Update("Done"); err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	this.ServeOk()
}

func (this *NoteController) Todo() {
	this.ServeDataJson(models.TodoNotes(this.CurrentUser.Id))
}

func (this *NoteController) Done() {
	this.ServeDataJson(models.DoneNotes(this.CurrentUser.Id))
}
