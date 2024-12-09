package files

import (
	"crypto/md5"
	"encoding/hex"
	"filesys/database"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

type File struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Path string `json:"path"`
	MD5  string `json:"md5"`
}

var fileList []File // 全局文件列表

func calculateMD5(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}

func generateID() string {
	return uuid.New().String() // 使用UUID生成唯一ID
}

func findFileByMD5(md5 string) *File {
	for _, file := range fileList {
		if file.MD5 == md5 {
			return &file
		}
	}
	return nil
}

func findFileByName(name string) *File {
	for _, file := range fileList {
		if file.Name == name {
			return &file
		}
	}
	return nil
}

func generateUniqueFilePath(fileName string) string {
	ext := filepath.Ext(fileName)
	baseName := fileName[0 : len(fileName)-len(ext)]
	newFileName := baseName + "_" + generateID() + ext
	return "./uploads/" + newFileName
}

func saveFileToDB(file File) {
	fileList = append(fileList, file)
}

func UploadFile(c *gin.Context) {
	file, _ := c.FormFile("file")
	savePath := "./savefile/" + file.Filename
	tmpsavePath := "./tmpfile/" + file.Filename

	if err := c.SaveUploadedFile(file, tmpsavePath); err != nil {
		c.String(http.StatusBadRequest, "Upload failed: %v", err)
		return
	}

	md5Value, err := calculateMD5(tmpsavePath)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to calculate MD5: %v", err)
		return
	}

	// 检查文件是否已经存在
	var existingFile File
	err = database.DB_Conn.QueryRow(`SELECT id, name, path, md5 FROM files WHERE md5 = $1`, md5Value).Scan(&existingFile.ID, &existingFile.Name, &existingFile.Path, &existingFile.MD5)
	if err == nil {
		os.Remove(tmpsavePath) // 删除刚刚上传的重复文件
		c.JSON(http.StatusConflict, gin.H{"message": "File already exists", "file": existingFile})
		return
	}

	// 检查是否有同名但不同MD5的文件
	var sameNameFile File
	err = db.QueryRow(`SELECT id, name, path, md5 FROM files WHERE name = $1`, file.Filename).Scan(&sameNameFile.ID, &sameNameFile.Name, &sameNameFile.Path, &sameNameFile.MD5)
	if err == nil && sameNameFile.MD5 != md5Value {
		newPath := generateUniqueFilePath(file.Filename)
		os.Rename(tmpsavePath, newPath)
		savePath = newPath
	} else {
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.String(http.StatusBadRequest, "Upload failed: %v", err)
			return
		}
	}

	// 插入文件信息到数据库
	var fileID int
	err = db.QueryRow(`INSERT INTO files (name, path, md5) VALUES ($1, $2, $3) RETURNING id`, file.Filename, savePath, md5Value).Scan(&fileID)
	if err != nil {
		c.String(http.StatusInternalServerError, "Failed to save file info: %v", err)
		return
	}

	newFile := File{ID: fileID, Name: file.Filename, Path: savePath, MD5: md5Value}
	os.Remove(tmpsavePath) // 删除临时文件

	c.JSON(http.StatusOK, gin.H{"message": "Upload successful", "file": newFile})
}
