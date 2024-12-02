package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func DownloadFile(c *gin.Context) {
	fileID := c.Param("id")

	// 查找文件信息
	var filePath string
	var fileName string
	for _, file := range fileList {
		if file.ID == fileID {
			filePath = file.Path
			fileName = file.Name
			break
		}
	}

	if filePath == "" {
		c.String(http.StatusNotFound, "File not found")
		return
	}

	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+fileName)
	c.Header("Content-Type", "application/octet-stream")
	c.File(filePath) // 直接发送文件
}
