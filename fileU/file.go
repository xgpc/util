package fileU

import (
	"errors"
	"io/ioutil"
	"os"
)

// 读取文件的所有内容
func FileReadAllContent(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
}

// 读取文件的所有内容，如有异常自动抛出
func FileReadAllContentThrow(path string) []byte {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		// todo: 写错误日志
		panic(err)
	}
	return bs
}

// 文件是否存在
func FileExist(filePath string) (bool, error) {
	f, err := os.Stat(filePath)
	if err == nil {
		if f.IsDir() {
			return false, errors.New(filePath + " 不是文件")
		}
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

// 创建文件
func MkFile(filePath string) error {
	exist, err := FileExist(filePath)
	if err != nil {
		return err
	}
	if exist {
		return nil
	}

	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		return errors.New("文件创建失败：" + err.Error())
	}
	_ = file.Close()

	return nil
}

// 读取文件大小
func GetFileSize(filePath string) (int64, bool) {
	f, err := os.Stat(filePath)
	if err == nil {
		if f.IsDir() {
			return 0, false
		}
		return f.Size(), true
	}
	return 0, false
}
