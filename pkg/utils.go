package pkg

import (
	"go.uber.org/zap"
	"net/url"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// MkAllPathDir 定义一个创建文件目录的方法
func MkAllPathDir(basePath string) (string, error) {
	//	1.获取当前时间,并且格式化时间
	folderName := time.Now().Format("2006/01/02")
	folderPath := filepath.Join(basePath, folderName)
	//使用MkdirAll会创建多层级目录
	err := os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return "", err
	}
	return folderPath + "/", nil
}

// String2Int32SliceCLD 字符串 ， 隔开的转为 int32切片
func String2Int32SliceCLD(str string) (r []int32) {
	split := strings.Split(strings.TrimSpace(str), ",")
	r = make([]int32, 0, len(split))
	for _, v := range split {
		i, err := strconv.Atoi(v)
		if err != nil {
			zap.L().Error(err.Error())
			continue
		}
		r = append(r, int32(i))
	}
	return
}

// StructToMap 结构体转map
func StructToMap(obj interface{}) map[string]interface{} {
	values := reflect.ValueOf(obj)
	fields := reflect.TypeOf(obj)
	result := make(map[string]interface{})
	for i := 0; i < fields.NumField(); i++ {
		result[fields.Field(i).Name] = values.Field(i).Interface()
	}

	return result
}

// ChineseTransToUnicode 中文转unicode
func ChineseTransToUnicode(chinese string) string {
	textQuoted := strconv.QuoteToASCII(chinese)
	return textQuoted[1 : len(textQuoted)-1]
}
func URIEncoder(params []string) string {
	values := url.Values{}
	for i := 0; i < len(params); i += 2 {
		values.Add(params[i], params[i+1])
	}
	return values.Encode()
}

func SpaceTextSpiteToPureTextArray(str string) []string {
	ret := make([]string, 0)
	spaceStr := strings.TrimSpace(str)
	splitStrArr := strings.Split(spaceStr, "\n")
	for i := 0; i < len(splitStrArr); i++ {
		if len(strings.TrimSpace(splitStrArr[i])) > 0 {
			ret = append(ret, strings.TrimSpace(splitStrArr[i]))
		}
	}
	return ret
}
