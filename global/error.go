package global

//自定义错误码

// 将项目中可能存在的错误都统一存放到这里，为每一种类型错误都定义一个错误码，便于在开发中快速定位错误

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError  CustomError
	ValidatorError CustomError
	TokenError     CustomError
}

var Errors = CustomErrors{
	BusinessError:  CustomError{40000, "业务错误"},
	ValidatorError: CustomError{42000, "请求参数错误"},
	TokenError:     CustomError{40100, "登录授权失败"},
}
