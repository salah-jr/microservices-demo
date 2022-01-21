package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salah-jr/microservices/DB"
	"net/http"
)

var authorizationToken string = "asdasdasdasdasd"

func main() {

	DB.Migrate()

	g := gin.Default()
	g.GET("/", Welcome)
	g.POST("RemoveStoreToken", RemoveStoreToken)
	g.POST("logout", Logout)
	g.Use(Services())

	g.Run("127.0.0.1:8080")
}

func Welcome(g *gin.Context)  {
	g.JSON(http.StatusOK, gin.H{
		"message": "ok",
		"status": true,
		"data": "",
	})
}

func RemoveStoreToken(g *gin.Context)  {
	var removeToken DB.RemoveToken
	if err := g.ShouldBindJSON(&removeToken); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Token not valid.",
			"status":  false,
			"data":    "",
		})
	}
	if removeToken.AuthToken != authorizationToken{
		g.JSON(http.StatusBadRequest, gin.H{
			"message": "Token not valid.",
			"status":  false,
			"data":    "",
		})
	}
	DB.Con.Where("user_id", removeToken.UserId).Unscoped().Delete(&DB.Token{})
	token := DB.Token{
		UserId: removeToken.UserId,
		Token: removeToken.UserToken,
	}
	DB.Con.Create(&token)
	g.JSON(http.StatusOK, gin.H{
		"message": "Token Updated.",
		"status":  true,
		"data":    "",
	})
}
func Logout(g *gin.Context) {
	token := g.GetHeader("Authorization")
	if token == "" {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
			"status":  false,
			"data":    "",
		})
		return
	}
	var fToken DB.Token
	DB.Con.Where("token", token).First(&fToken)
	if fToken.ID == 0 {
		g.JSON(http.StatusUnauthorized, gin.H{
			"message": "Not Authorized",
			"status":  false,
			"data":    "",
		})
		return
	}
	DB.Con.Where("token", token).Unscoped().Delete(&DB.Token{})
	g.JSON(http.StatusOK, gin.H{
		"message": "Done logout.",
		"status":  true,
		"data":    "",
	})
}

