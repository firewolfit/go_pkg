package file

import (
	"os"
	"path/filepath"
)

// NewFile 创建文件
// fileName 文件名
// override 是否覆盖已存在的文件
func NewFile(fileName string, override bool) error {
	existed := IsExisted(fileName, false)
	if existed {
		if !override { // 覆盖则删除
			return nil
		}
	}
	dir := filepath.Dir(fileName)
	folderExisted := IsExisted(dir, true)
	if !folderExisted {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return err
		}
	}
	if _, err := os.Create(fileName); err != nil {
		return err

	}
	return nil
}

func NewFolder(folder string) error {
	if existed := IsExisted(folder, true); existed {
		return nil
	}
	if err := os.MkdirAll(folder, os.ModePerm); err != nil {
		return err
	}
	return nil
}

func IsExisted(filename string, folder bool) bool {
	fileInfo, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return fileInfo.IsDir() == folder
}
