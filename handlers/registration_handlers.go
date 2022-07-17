package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"it1shka.com/my-recipes/database"
	"it1shka.com/my-recipes/myutils"
)

const WEAK_PASSWORD_MESSAGE = `
	your password is weak.
	Password should contain uppercase letter,
	special character, number and be at least 8 characters
	long.
`

func getLoginHandler(ctx *gin.Context) {
	msg := ctx.GetString("error")
	ctx.HTML(http.StatusOK, "login.html", gin.H{
		"error": msg,
	})
}

func getRegisterHandler(ctx *gin.Context) {
	msg := ctx.GetString("error")
	ctx.HTML(http.StatusOK, "register.html", gin.H{
		"error": msg,
	})
}

type LoginFormData struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

func postLoginHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	defer session.Save()

	loginError := func(message string) {
		errmsg := fmt.Sprintf("Failed to login: %s", message)
		session.Set("error", errmsg)
		ctx.Redirect(http.StatusFound, "/login")
	}

	var loginData LoginFormData
	if err := ctx.ShouldBind(&loginData); err != nil {
		loginError("form validation failed")
		return
	}

	user, exists := database.FindUserByEmail(loginData.Email)
	if !exists {
		loginError("user not found")
		return
	}

	correctPwd := myutils.CheckPasswordHash(loginData.Password, user.Password)
	if !correctPwd {
		loginError("incorrect password")
		return
	}

	// log in logic
	session.Set("userid", user.ID)
	session.Set("username", user.Name)
	session.Set("useremail", user.Email)

	ctx.Redirect(http.StatusFound, "/")
}

type RegisterFormData struct {
	Name            string `form:"name" binding:"required,alpha"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required"`
	ConfirmPassword string `form:"confirm_password" binding:"required"`
}

func postRegisterHandler(ctx *gin.Context) {
	session := sessions.Default(ctx)
	defer session.Save()

	regError := func(message string) {
		errmsg := fmt.Sprintf("Failed to register: %s", message)
		session.Set("error", errmsg)
		ctx.Redirect(http.StatusFound, "/register")
	}

	var regData RegisterFormData
	if err := ctx.ShouldBind(&regData); err != nil {
		regError("form validation failed")
		return
	}

	_, alreadyRegistered := database.FindUserByEmail(regData.Email)
	if alreadyRegistered {
		regError("account already has been registered")
		return
	}

	if regData.Password != regData.ConfirmPassword {
		regError("confirm the password correctly")
		return
	}

	if !myutils.IsPasswordStrong(regData.Password) {
		regError(WEAK_PASSWORD_MESSAGE)
		return
	}

	password, err := myutils.HashPassword(regData.Password)
	if err != nil {
		regError("internal error occured")
		return
	}

	user, err := database.CreateUser(regData.Name, regData.Email, password)
	if err != nil {
		regError("failed to register your account due to internal error")
		return
	}

	// log in logic
	session.Set("userid", user.ID)
	session.Set("username", user.Name)
	session.Set("useremail", user.Email)

	ctx.Redirect(http.StatusFound, "/")
}

func getLogoutHander(ctx *gin.Context) {
	session := sessions.Default(ctx)
	defer session.Save()

	session.Delete("userid")
	session.Delete("username")
	session.Delete("useremail")

	ctx.Redirect(http.StatusFound, "/")
}
