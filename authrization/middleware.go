package main

import (
	"github.com/gin-gonic/gin"
	"github.com/salah-jr/microservices/DB"
	"net/http"
	"strconv"
	"strings"
)

func Services() gin.HandlerFunc {
	return func(g *gin.Context) {
		url := g.Request.RequestURI
		segments := strings.Split(url, "/")
		segmentOne := strings.Split(segments[1], "?")
		services := ServicesList()
		if ok := services[segmentOne[0]]; ok == "" {
			g.JSON(http.StatusNotFound, gin.H{
				"message": "Not found",
				"status":  false,
				"data":    "",
			})
			return
			g.Abort()
		}

		serviceurl := services[segmentOne[0]]
		newUrl := serviceurl + url
		var token DB.Token
		authList := AuthRoutesList()
		for route, _ := range authList {
			if strings.Contains(url, route) {
				authorized := g.GetHeader("Authorization")
				if authorized == "" {
					g.JSON(http.StatusUnauthorized, gin.H{
						"message": "Not Authorized",
						"status":  false,
						"data":    "",
					})
					return
					g.Abort()
				}
				DB.Con.Where("token", authorized).First(&token)
				if token.ID == 0 {
					g.JSON(http.StatusUnauthorized, gin.H{
						"message": "Not Authorized",
						"status":  false,
						"data":    "",
					})
					return
					g.Abort()
				}
				g.Request.Header.Set("USER_ID", strconv.Itoa(token.UserId))
			}
		}

		method := strings.ToLower(g.Request.Method)
		switch method {
		case "get":
			Get(g, newUrl)
			break
		case "post":
			Post(g, newUrl)
			break
		}
	}
}

func ServicesList() map[string]string {
	m := make(map[string]string)
	m["users"] = "http://127.0.0.1:8070"
	m["posts"] = "http://127.0.0.1:8090"

	return m
}

func AuthRoutesList() map[string]bool {
	m := make(map[string]bool)
	m["my-posts"] = true

	return m
}
