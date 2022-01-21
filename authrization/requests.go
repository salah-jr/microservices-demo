package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Post(g *gin.Context, url string) {

	body, err := g.GetRawData()
	if err != nil {
		fmt.Println("error in read request body!")
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		fmt.Println("error")
	}

	for key, _ := range g.Request.Header {
		req.Header.Set(key, g.Request.Header.Get(key))
	}
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("error")
	}
	defer response.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	if len(result) == 0 {
		g.JSON(response.StatusCode, gin.H{
			"message": "No response from this ms",
			"status":  false,
			"data":    "",
		})
		return
	}

	g.JSON(response.StatusCode, result)
	return
}

func Get(g *gin.Context, url string) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Println("error")
	}

	for key, _ := range g.Request.Header {
		req.Header.Set(key, g.Request.Header.Get(key))
	}
	client := &http.Client{Timeout: time.Second * 10}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println("error")
	}
	defer response.Body.Close()

	var result map[string]interface{}

	json.NewDecoder(response.Body).Decode(&result)

	if len(result) == 0 {
		g.JSON(response.StatusCode, gin.H{
			"message": "No response from this ms",
			"status":  false,
			"data":    "",
		})
		return
	}

	g.JSON(response.StatusCode, result)
	return
}
