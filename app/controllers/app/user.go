package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"todo_list/app/common/request"
	"todo_list/app/common/response"
	"todo_list/app/services"
)

//校验入参 调用UserService注册逻辑

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	fmt.Println("form", form, &form)
	if err := c.ShouldBindJSON(&form); err != nil {
		fmt.Println("err", err, &form)
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
