package pkg

import (
	"errors"
	"fmt"
	"foods-spider-server/com"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/zh"
	"github.com/go-playground/locales/zh_Hant_TW"
	ut "github.com/go-playground/universal-translator"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	zhTranslations "github.com/go-playground/validator/v10/translations/zh"
)

/*
Trans 一个全局翻译器
*/
var Trans ut.Translator

/*
InitLocaleTrans 初始化翻译器
*/
func InitLocaleTrans(locale string) (err error) {
	// 修改gin框架中的Validator引擎属性，实现自定制 （使用断言转化）
	if v, assertOk := binding.Validator.Engine().(*validator.Validate); assertOk {

		//使用为结构的 JSON 表示指定的名称，而不是普通的 Go 字段名称
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		// SignUpNewUser 注册自定义校验方法
		//v.RegisterStructValidation(SignUpNewUserParamStructLevelValidation, parameters.SignUpNewUser{})
		zhT := zh.New()          // 中文简体翻译器
		zhTw := zh_Hant_TW.New() // 台湾省繁体字
		enT := en.New()          // 英文翻译器

		// 第一个参数是备用（fallback）的语言环境
		// 后面的参数是应该支持的语言环境（支持多个）
		// uni := ut.New(zhT, zhT) 也是可以的
		uni := ut.New(enT, zhT, enT, zhTw) //将来就从这些注册中间去找 Translator ！

		// locale 通常也取决于 http 请求头的 'Accept-Language'
		var ok bool
		//根据需要的语言，通过 GetTranslator 初始化一个翻译器 也可以使用 uni.FindTranslator(...) 传入多个locale进行查找

		Trans, ok = uni.GetTranslator(locale) //这个传入的要和 github.com/go-playground/locales/(xxx) xxx对应，所以下面switch也要对应上！
		if !ok {
			return fmt.Errorf("uni.GetTranslator(%s) failed", locale)
		}

		// 注册翻译器
		switch locale {
		case "en":
			_ = enTranslations.RegisterDefaultTranslations(v, Trans)
		case "zh":
			_ = zhTranslations.RegisterDefaultTranslations(v, Trans)
		default:
			_ = enTranslations.RegisterDefaultTranslations(v, Trans)
		}
		return nil
	}
	return nil
}

// RemoveStructName var Trans *validator.TranslationFunc
func RemoveStructName(errors validator.ValidationErrorsTranslations) validator.ValidationErrorsTranslations {
	newTransErrors := make(map[string]string)
	for key, value := range errors {
		newTransErrors[strings.Split(key, ".")[1]] = value
	}
	return newTransErrors
}

/*
ErrorValidator 处理json映射到参数校验的错误处理！
判断是validator异常还是其他异常
*/
func ErrorValidator(context *gin.Context, err error) {
	var errs validator.ValidationErrors
	assertOk := errors.As(err, &errs)
	if !assertOk {
		/*
			! assert_ok 表示断言失败！非validator.ValidationErrors类型错误直接返回
		*/
		zap.L().Error(err.Error())
		com.ResponseData(context, com.CodeFailed, err.Error())
		return
	}
	/*
		断言成功则错误为 validator.ValidationErrors类型错误，则进行翻译
	*/
	zap.L().Error(fmt.Sprintf("参数异常%v", RemoveStructName(errs.Translate(Trans))))
	com.ResponseData(context, com.CodeParameterError, RemoveStructName(errs.Translate(Trans)))
}
