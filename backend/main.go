package main

import (
	"filesys/database"
	"filesys/files"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func main() {
	database.InitDB()
	// 加载 savefile 目录中的文件
	err := files.LoadFilesFromDir("./savefile")
	if err != nil {
		panic(err)
	}
	r := gin.Default()
	// 配置CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // 允许所有源，也可以指定特定的源
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.POST("/upload", files.UploadFile)
	r.GET("/list", files.UpdateFileList)
	r.GET("/download/:id", files.DownloadFile) // 文件下载

	r.Run(":8222")
}
