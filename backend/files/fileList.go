package files

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func LoadFilesFromDir(directory string) error {
	entries, err := os.ReadDir(directory)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			filePath := filepath.Join(directory, entry.Name())
			md5Value, err := calculateMD5(filePath)
			if err != nil {
				return err
			}
			newFile := File{ID: generateID(), Name: entry.Name(), Path: filePath, MD5: md5Value}
			saveFileToDB(newFile)
		}
	}
	return nil
}

func UpdateFileList(c *gin.Context) {
	c.JSON(http.StatusOK, fileList)
	return
} // 返回全局文件列表 }
