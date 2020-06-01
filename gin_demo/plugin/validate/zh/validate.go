package zh

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

var (
	validate = validator.New()
	uni      = ut.New(zh.New())
	trans, _ = uni.GetTranslator("zh")
)

func init() {
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		panic(err)
	}
}

/************************************************************************
* 参数校验, 如果有错误, 将错误翻译成中文
* @Param obj 校验的对象， tag 字段默认是 validate
* @Ret   infos 错误信息
* @Ret   ok  如果参数校验通过， 返回 [], true, nil
* @Ret   err 如果失败，  1. 参数字段错误， 返回 [...], false, nil
*                       2. 其他的错误 返回 [], false , err
*************************************************************************/
func ValidateAndTrans2Zh(obj interface{}) (infos []string, ok bool, err error) {
	err = validate.Struct(obj)

	if err == nil {
		return infos, true, nil
	}

	if CanTrans, ok := err.(validator.ValidationErrors); ok {
		translate := CanTrans.Translate(trans)

		for _, v := range translate {
			infos = append(infos, v)
		}
		return infos, false, nil
	}

	return infos, false, err
}
