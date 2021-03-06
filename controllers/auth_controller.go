package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/goproj/note/g"
	"github.com/goproj/note/models"
)

type AuthController struct {
	BaseController
}

func (this *AuthController) Index() {
	this.TplNames = "index.html"
}

func (this *AuthController) Register() {
	type JsonInfo struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var jsonInfo JsonInfo
	if json.Unmarshal(this.Ctx.Input.RequestBody, &jsonInfo) != nil {
		this.ServeErrMsg("input format error")
		return
	}

	email := jsonInfo.Email
	name := jsonInfo.Username
	password := jsonInfo.Password

	if email == "" {
		this.ServeErrMsg("邮箱必须填写")
		return
	}

	if name == "" {
		name = email
	}

	userInDB := models.FindUserByEmail(email)
	if userInDB != nil {
		this.ServeErrMsg("该邮箱已经注册过啦")
		return
	}

	userInDB = models.FindUserByName(name)
	if userInDB != nil {
		this.ServeErrMsg("该昵称已经注册过啦")
		return
	}

	u := &models.User{Name: name, Email: email, Password: password}
	err := u.Insert()
	if err != nil {
		this.ServeErrMsg(err.Error())
		return
	}

	// write cookie: max-age, path, domain, secure, httponly
	this.Ctx.SetSecureCookie(g.SecretKey, g.CookieName, fmt.Sprint("%d", u.Id), 2592000, "/", "", false, true)

	this.ServeDataJson(map[string]interface{}{
		"id":       u.Id,
		"name":     u.Name,
		"email":    u.Email,
		"portrait": models.Avatar90(u.Email),
	})
}

func (this *AuthController) Logout() {
	// write cookie: max-age, path, domain, secure, httponly
	this.Ctx.SetSecureCookie(g.SecretKey, g.CookieName, "-1", 0, "/", "", false, true)
	this.ServeOk()
}

func (this *AuthController) Login() {

	type JsonInfo struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	var jsonInfo JsonInfo
	if json.Unmarshal(this.Ctx.Input.RequestBody, &jsonInfo) != nil {
		this.ServeErrMsg("input format error")
		return
	}

	email := jsonInfo.Email
	password := jsonInfo.Password
	if email == "" {
		this.ServeErrMsg("邮箱必须填写")
		return
	}
	u := models.FindUserByEmail(email)
	if u == nil {
		this.ServeErrMsg("该邮箱不存在")
		return
	}

	if u.Password != password {
		this.ServeErrMsg("密码错误")
		return
	}

	if u.Blocked == 1 {
		this.ServeErrMsg("u'r blocked")
		return
	}

	// write cookie: max-age, path, domain, secure, httponly
	this.Ctx.SetSecureCookie(g.SecretKey, g.CookieName, fmt.Sprintf("%d", u.Id), 2592000, "/", "", false, true)
	if u.Portrait == "" {
		u.Portrait = models.Avatar90(u.Email)
	}

	this.ServeDataJson(map[string]interface{}{
		"id":       u.Id,
		"name":     u.Name,
		"email":    u.Email,
		"portrait": u.Portrait,
	})

}
