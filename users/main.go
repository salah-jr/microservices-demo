package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salah-jr/microservices/DB"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"github.com/kirinlabs/HttpRequest"
)

var authorizationToken string = "asdasdasdasdasd"
var authUrl string = "http://127.0.0.1:8080"

func main() {

	g := gin.Default()
	users := g.Group("users")
	{
		users.GET("/", Welcome)
		users.POST("/login", Login)
	}
	g.Run("127.0.0.1:8070")
}


func Login(g *gin.Context) {
	var login DB.Login
	if err := g.ShouldBindJSON(&login); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Please enter your email and password",
			"status":  false,
			"data":    "",
		})
	}
	var user DB.User
	DB.Con.Where("email", login.Email).Where("password", login.Password).First(&user)

	if user.ID == 0 {
		g.JSON(http.StatusNotFound, gin.H{
			"message": "Your email or your password not valid.",
			"status":  false,
			"data":    "",
		})
	}
	token, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 3)
	user.Token = string(token)
	DB.Con.Save(&user)
	req := HttpRequest.NewRequest()
	req.JSON().Post(authUrl+"/RemoveStoreToken", map[string]interface{}{
		"auth_token": authorizationToken,
		"user_token": user.Token,
		"user_id": user.ID,
	})
	g.JSON(http.StatusOK, gin.H{
		"message": "You are logged in..",
		"status":  true,
		"data":    user,
	})
}



func Welcome(g *gin.Context) {
	g.JSON(http.StatusOK, gin.H{
		"message": "Welcome in user microservice",
		"status":  true,
		"data":    "",
	})
}
