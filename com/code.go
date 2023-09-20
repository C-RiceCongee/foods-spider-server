package com

type ResCode uint

const (
	CodeSuccess = 1000 + iota
	CodeFailed
	CodeParameterError
	CodeUserExists
	CodeUserNotExists
	CodeExists
	CodeUsernamePasswordIncorrect
	CodeNoAuthError
	CodeServerError
	CodeForbidden
)

var codeMessageMap = map[ResCode]string{
	CodeSuccess:                   "success",
	CodeFailed:                    "failed",
	CodeParameterError:            "请求参数错误",
	CodeUserExists:                "用户已经存在",
	CodeUserNotExists:             "用户不存在",
	CodeUsernamePasswordIncorrect: "用户名或密码错误",
	CodeExists:                    "重复,已经存在",
	CodeNoAuthError:               "登录失效",
	CodeServerError:               "服务器异常",
	CodeForbidden:                 "无权限请求，拒绝访问",
}

func GetResponseMessage(code ResCode, customMsg string) string {
	if customMsg == "" {
		s, ok := codeMessageMap[code]
		if ok {
			return s
		}
		return "服务器繁忙"
	}
	return customMsg
}
