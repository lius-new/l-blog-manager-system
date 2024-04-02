package utils

import (
	"crypto/sha256"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Hash: 计算文件的Hash
func Hash(file io.Reader) (string, error) {
	hash := sha256.New()
	if _, err := io.Copy(hash, file); err != nil {
		log.Fatal(err)
		return "", err
	}
	sum := hash.Sum(nil)
	return fmt.Sprintf("%x", sum), nil
}

// GetFileSuffix: 获取文件后缀
func GetFileSuffix(filename string) string {
	if len(filename) == 0 {
		return ""
	}
	tempStr := strings.Split(filename, ".")
	if len(tempStr) < 1 {
		return ""
	}
	return tempStr[len(tempStr)-1]
}

// FileExist: 判断文件夹是否存在指定文件
func FileExist(filename, currentPath string) (string, error) {
	dirs, err := os.ReadDir(currentPath)

	if err != nil {
		return "", err
	}

	for _, v := range dirs {
		if strings.Contains(v.Name(), filename) {
			return v.Name(), nil
		}
	}
	return "", nil
}

// CreateDir: 创建文件夹
func CreateDir(dir string) error {
	_, err := os.Stat(dir)
	if err == nil {
		return nil
	}
	if os.IsNotExist(err) {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
