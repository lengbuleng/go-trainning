package v1

import (
	"app/internal/models"
	"app/internal/service/user_service"
	"app/pkg/e"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		e.Err("bind", err)
	}

	if err := user.Save(); err != nil {
		e.Err("注册失败", err)
	}

	tokenSvc := user_service.TokenSvc{}
	token, err := tokenSvc.MakeToken(user)
	if err != nil {
		e.Err("token生成失败", err)
	}

	e.NewsOk(c, "注册成功", "user", e.Map{"userid": user.Id, "username": user.Name, "token": token})

}

func Login(c *gin.Context) {

	user := &models.User{}
	if err := c.ShouldBind(user); err != nil {
		e.Err("bind", err)
	}

	if err := user.Login(); err != nil {
		e.Err("登录失败", err)
	}

	tokenSvc := user_service.TokenSvc{}
	token, err := tokenSvc.MakeToken(user)
	if err != nil {
		e.Err("token生成失败", err)
	}

	e.NewsOk(c, "登录成功", "user", e.Map{"userid": user.Id, "username": user.Name, "token": token})

}
