package request

//用来存放所有用户相关的请求结构体，并实现Validator接口

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required,mobile"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 自定义错误消息

//注：由于在 InitializeValidator() 方法中，使用 RegisterTagNameFunc() 注册了自定义 json tag， 所以在 GetMessages() 中自定义错误信息 key 值时，需使用 json tag 名称

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"name.required":     "用户名称不能为空",
		"mobile.required":   "用户手机号不能为空",
		"mobile.mobile":     "用户手机号格式不正确",
		"password.required": "用户密码不能为空",
	}
}

type Login struct {
	Mobile   string `json:"mobile" form:"mobile" binding:"required,mobile"`
	Password string `json:"password" form:"password" binding:"required"`
}

func (login Login) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"mobile.required":   "手机号码不能为空",
		"mobile.mobile":     "手机号码格式不正确",
		"password.required": "密码不能为空",
	}
}
