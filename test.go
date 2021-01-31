package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
)

var UserData map[string]string

func init(){
	UserData = map[string]string{
		"test": "test",
	}
}

func CheckUserIsExist(usertest string) bool{
	_, IsExist := UserData[usertest]
	return IsExist
}

func CheckPassword(InputPassword string,RealPassword string) error{
	if InputPassword == RealPassword {
		return nil
	} else {
		return errors.New("wrong password")
	}
}

func Auth(InputUsername string, InputPassword string) error{
	IsExist := CheckUserIsExist(InputUsername);
	if IsExist == false {
		return errors.New("User not exist")
	} else {
		return CheckPassword(InputPassword,UserData[InputUsername])
	}
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func LoginAuth(c *gin.Context){
	var username string
	var password string
	if in,IsExist := c.GetPostForm("username"); in!="" && IsExist {
		username = in
	} else {
		c.HTML(http.StatusBadRequest,"index.html",gin.H{"error" : errors.New("必須輸入使用者名稱")})
		return
	}
	if in,IsExist := c.GetPostForm("password"); in!="" && IsExist {
		password = in
	} else {
		c.HTML(http.StatusBadRequest,"index.html",gin.H{"error" : errors.New("必須輸入使用者密碼")})
		return
	}

	if err:=Auth(username,password); err==nil{
		//c.HTML(http.StatusOK,"index.html",gin.H{"success" : "登入成功"})
		c.Redirect(302,"http://127.0.0.1:8888/profile")
		return 
	} else {
		c.HTML(http.StatusBadRequest,"index.html",gin.H{"error" : err})
		return
	}
}

func Profile(c *gin.Context) {
	c.HTML(http.StatusOK, "profile.html",nil)
}

func main() {
	//gin.SetMode(gin.ReleaseMode)
	server := gin.Default()
	server.LoadHTMLGlob("template/*")
	server.GET("/login", LoginPage)
	server.POST("/login",LoginAuth)
	server.GET("/profile",Profile)
	server.Run(":8888")
}

