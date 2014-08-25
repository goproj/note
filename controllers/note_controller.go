package controllers

import (
	"fmt"
	"strconv"
)

type NoteController struct {
	BaseController
}

func (this *NoteController) Read() {
	idStr := this.Ctx.Input.Params[":id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	this.ServeErrMsg(fmt.Sprintf("read id:%d not implemented", id))
}

func (this *NoteController) Delete() {
	idStr := this.Ctx.Input.Params[":id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	this.ServeErrMsg(fmt.Sprintf("delete id:%d not implemented", id))
}

func (this *NoteController) Update() {
	idStr := this.Ctx.Input.Params[":id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	this.ServeErrMsg(fmt.Sprintf("update id:%d not implemented", id))
}

func (this *NoteController) Add() {
	this.ServeErrMsg(fmt.Sprintf("add method not implemented"))
}

func (this *NoteController) Todo() {
	this.ServeErrMsg("todo not implemented")
}

func (this *NoteController) History() {
	this.ServeErrMsg("history not implemented")
}
