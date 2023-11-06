package request

//用来存放所有用户相关的请求结构体，并实现Validator接口

type Register struct {
	Name     string `form:"name" json:"name" binding:"required"`
	Mobile   string `form:"mobile" json:"mobile" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 自定义错误消息

func (register Register) GetMessages() ValidatorMessages {
	return ValidatorMessages{
		"Name.required":     "用户名称不能为空",
		"Mobile.required":   "用户手机号不能为空",
		"Password.required": "用户密码不能为空",
	}
}
