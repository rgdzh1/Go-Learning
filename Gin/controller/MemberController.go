package controller

import (
	"Go-Learning/Gin/param"
	"Go-Learning/Gin/service"
	"Go-Learning/Gin/tool"
	"github.com/gin-gonic/gin"
)

type MemberController struct {
}

func (receiver *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sencode", receiver.SendSmsCode)
	engine.POST("/api/login_sms", receiver.smsLogin)
}

// 发送验证码
func (receiver *MemberController) SendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		//context.JSON(200, map[string]interface{}{
		//	"code": 0,
		//	"msg":  "参数解析失败",
		//	"data": "",
		//})
		tool.Failed(context, "参数解析失败")
		return
	}
	isSend := new(service.MemberService).SendCode(phone)
	if isSend {
		//context.JSON(200, map[string]interface{}{
		//	"code": 1,
		//	"msg":  "发送成功",
		//	"data": "",
		//})
		tool.Success(context, "参数解析成功")
		return
	}
	//context.JSON(200, map[string]interface{}{
	//	"code": 0,
	//	"msg":  "发送失败",
	//	"data": "",
	//})
	tool.Failed(context, "短信发送失败")
}

func (receiver *MemberController) smsLogin(context *gin.Context) {
	mSmsLoginParam := new(param.SmsLoginParam)
	err := tool.Decode(context.Request.Body, mSmsLoginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	// 完成登录

}
