package main

import (
	"github.com/gin-gonic/gin"
	"github.com/namduong/project-layout/internal/database"
)

func main() {
	database.Connect()
	r := gin.Default()

	r.Run()
}
