package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"strconv"
	"time"
)

// CustomClaims 自定义声明类型 并内嵌jwt.RegisteredClaims
// jwt包自带的jwt.RegisteredClaims只包含了官方字段
// 假设我们这里需要额外记录一个username字段，所以要自定义结构体
// 如果想要保存更多信息，都可以添加到这个结构体中
type CustomClaims struct {
	// 可根据需要自行添加字段
	Username             string `json:"username"`
	UserId               string `json:"user_id"`
	OpenId               string `json:"openid"`
	jwt.RegisteredClaims        // 内嵌的一些声明
}

// GenToken 生成JWT
func GenToken(username string, openid string, userId int64) (string, error) {
	// 创建一个我们自己的声明
	claims := CustomClaims{
		Username: username, // 自定义字段
		UserId:   strconv.FormatInt(userId, 10),
		OpenId:   openid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 72)), // 过期时间
			Issuer:    Config.JwtConfig.Issuer,                            // 签发人
		},
	}
	// 使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// 使用指定的secret签名并获得完整的编码后的字符串token
	// 这里不能直接传字符串！会报错！ key is invalid!
	return token.SignedString([]byte(Config.JwtConfig.Secret))
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (payload *CustomClaims, parseErr error) {
	payload = new(CustomClaims)
	// 这里 第三个参数fun ，考虑到可以通过函数做一个组合，其实本质还是通过token,载体 密钥 校验的！
	token, err := jwt.ParseWithClaims(tokenString, payload, func(token *jwt.Token) (i interface{}, err error) {
		return []byte(Config.JwtConfig.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid { // 校验token
		return payload, nil
	}
	zap.L().Info("无效的token")
	return nil, errors.New("无效的token")
}
