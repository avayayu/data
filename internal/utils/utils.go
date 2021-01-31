package utils

import (
	"fmt"
	"os"
	"unicode"
)

//LcFirst 首字母小写
func LcFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

//EnsureDir 确保路径存在
func EnsureDir(dirName string) error {

	err := os.Mkdir(dirName, os.ModeDir)

	if err == nil || os.IsExist(err) {
		return nil
	}
	return err

}


func ParseToStr(mp map[string]interface{}) string {
	values := ""
	for key, val := range mp {
		values += "&" + key + "=" + fmt.Sprintf("%v",val)
	}
	temp := values[1:]
	values = "?" + temp
	return values
}
