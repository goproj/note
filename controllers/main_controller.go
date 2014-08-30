package controllers

import (
	"github.com/goproj/note/g"
	"github.com/goproj/note/models"
	"strconv"
)

type MainController struct {
	BaseController
}

func (this *MainController) CheckLogin() {
	if this.RetrieveUserFromCookie() {
		return
	}

	this.Abort("401")
}

func (this *MainController) RetrieveUserFromCookie() bool {
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

func (this *MainController) Me() {
	// just a test
	this.ServeDataJson(this.CurrentUser)
}
