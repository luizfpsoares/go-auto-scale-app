package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {
	route := gin.Default()
	route.GET("/info", InfoHandler)
	route.GET("/fibonacci/:qtd", fibonacciHandler)

	route.Run()
}

type Info struct {
	Hostname string `json:"hostname"`
	OS       string `json:"os"`
}

func OsInfo() Info {
	var data Info
	data.Hostname, _ = os.Hostname()
	data.OS = runtime.GOOS
	return data
}

func InfoHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": OsInfo()})
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func fibonacciHandler(c *gin.Context) {
	qtd := c.Param("qtd")
	num, err := strconv.Atoi(qtd)
	if err != nil {
		log.Println("Erro na conversão:", err)
		c.JSON(400, gin.H{"parametro invalido": err})
		return
	}
	res := fibonacci(num)
	c.JSON(http.StatusOK, gin.H{"result": res})
}
