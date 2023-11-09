package bootstrap

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
	"todo_list/utils"
)

//定制 gin 框架 Validator属性

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		//	注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidatorMobile)

		//	注册自定义json tag 函数
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
