package pddsdk

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var (
	CachePath = GetCurrentDirectory() + "/cache"
)

//args[0] = filename string
//args[1] = cacheLife int64
func FileGetCache(args ...interface{}) string {
	if len(args) < 1 || len(args) > 2 {
		return ""
	}
	filename := args[0].(string)
	cacheLife := args[1].(int64)
	filename = CachePath + "/" + filename + ".cache"
	fileInfo, exist := FileExists(filename)
	if exist {
		if cacheLife > 0 && (GetNow().Unix()-fileInfo.ModTime().Unix()) > cacheLife {
			return ""
		}
		fileData, err := ioutil.ReadFile(filename)
		if err != nil {
			return ""
		} else {
			return string(fileData)
		}
	} else {
		return ""
	}
}

//args[0] = filename string
//args[1] = fileData string
//args[2] = cacheLife int64
func FileSetCache(args ...interface{}) bool {
	if len(args) < 2 || len(args) > 3 {
		return false
	}
	filename := args[0].(string)
	fileData := args[1].(string)
	//cacheLife := args[2].(int64)
	var err error
	err = IsNotExistMkDir(CachePath)
	if err != nil {
		return false
	}
	filename = CachePath + "/" + filename + ".cache"
	err = ioutil.WriteFile(filename, []byte(fileData), 0777)
	if err != nil {
		return false
	} else {
		return true
	}
}

func GetCurrentDirectory() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径  filepath.Dir(os.Args[0])去除最后一个元素的路径
	if err != nil {
		panic(err.Error())
	}
	return strings.Replace(dir, "\\", "/", -1) //将\替换成/
}

func CheckNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func IsNotExistMkDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

func MkDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

func FileExists(filename string) (os.FileInfo, bool) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false
		}
	}
	return fileInfo, true
}
