package fileU

import (
	"errors"
	"os"
)

// 目录是否存在
func IsExistDir(dirPath string) (bool, error) {
	f, err := os.Stat(dirPath)
	if err == nil {
		if f.IsDir() {
			return true, nil
		}
		return false, errors.New(dirPath + " 不是目录")
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 创建文件夹
func MkDir(dirPath string) error {
	exist, err := IsExistDir(dirPath)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	err = os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return errors.New("目录创建失败：" + err.Error())
	}

	return nil
}
