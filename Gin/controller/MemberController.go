package controller

import (
	"Go-Learning/Gin/model"
	"Go-Learning/Gin/param"
	"Go-Learning/Gin/service"
	"Go-Learning/Gin/tool"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type MemberController struct {
}

func (receiver *MemberController) Router(engine *gin.Engine) {
	engine.GET("/api/sencode", receiver.SendSmsCode)
	engine.POST("/api/login_sms", receiver.smsLogin)
	engine.GET("/api/captcha", receiver.captcha)
	engine.POST("/api/vertifycha", receiver.vertifyCaptcha)
	engine.POST("/api/login_pwd", receiver.nameLogin)
	engine.POST("/api/upload/avator", receiver.uploadAvator)
}

// 发送验证码
func (receiver *MemberController) SendSmsCode(context *gin.Context) {
	phone, exist := context.GetQuery("phone")
	if !exist {
		tool.Failed(context, "参数解析失败")
		return
	}
	isSend := new(service.MemberService).SendCode(phone)
	if isSend {
		tool.Success(context, "参数解析成功")
		return
	}
	tool.Failed(context, "短信发送失败")
}

/**
短信登录
*/
func (receiver *MemberController) smsLogin(context *gin.Context) {
	mSmsLoginParam := new(param.SmsLoginParam)
	err := tool.Decode(context.Request.Body, mSmsLoginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	// 完成登录
	member := new(service.MemberService).SmsLogin(mSmsLoginParam)
	if member != nil {
		// 格式化得到session
		sess, _ := json.Marshal(member)
		// 保存session
		err := tool.SetSeess(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登录失败,session没有保存下来")
			return
		}
		tool.Success(context, member)
		return
	}
	tool.Failed(context, "登录失败")
}

/**
生成验证码,并返回
*/
func (receiver *MemberController) captcha(context *gin.Context) {
	//生成验证码,并返回客户端
	generateCaptcha := tool.GenerateCaptcha(context)
	// 返回给客户端
	tool.Success(context, map[string]interface{}{
		// 返回数据
		"captcha_result": generateCaptcha,
	})
}

/**
验证验证码是否正确
*/
func (receiver *MemberController) vertifyCaptcha(context *gin.Context) {
	captcha := new(tool.CaptchaResult)
	err := tool.Decode(context.Request.Body, captcha)
	if err != nil {
		tool.Failed(context, "解析失败")
		return
	}
	result := tool.VerifyCaptcha(captcha.Id, captcha.VertifyValue)
	if result {
		fmt.Println("验证码成功")
	} else {
		fmt.Println("验证码失败")
	}
}

/**
用户登录
*/
func (receiver *MemberController) nameLogin(context *gin.Context) {
	// 解析参数
	loginParam := new(param.LoginParam)
	err := tool.Decode(context.Request.Body, loginParam)
	if err != nil {
		tool.Failed(context, "参数解析失败")
		return
	}
	// 验证验证码
	validate := tool.VerifyCaptcha(loginParam.Id, loginParam.Value)
	if !validate {
		tool.Failed(context, "验证码不正确,清重新验证")
		return
	}
	// 登录
	memberService := service.MemberService{}
	member := memberService.Login(loginParam.Name, loginParam.Password)
	if member.Id != 0 {
		// 格式化得到session
		sess, _ := json.Marshal(member)
		// 保存session
		err := tool.SetSeess(context, "user_"+string(member.Id), sess)
		if err != nil {
			tool.Failed(context, "登录失败,session没有保存下来")
			return
		}
		tool.Success(context, &member)
		return
	}
	tool.Failed(context, "登录失败")
}

/**
头像上传
*/
func (receiver *MemberController) uploadAvator(context *gin.Context) {
	// 解析上传参数
	userId := context.PostForm("user_id")
	file, err := context.FormFile("avatar")
	if err != nil || userId == "" {
		tool.Failed(context, "参数解析失败")
		return
	}
	// 判断user对应得用户是否登录
	seess := tool.GetSeess(context, "user_"+userId)
	if seess == nil {
		tool.Failed(context, "参数不合法")
		return
	}
	member := new(model.Member)
	json.Unmarshal(seess.([]byte), member)
	// file保存到本地

	fileName := "./Gin/uploadFile/" + strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	err = context.SaveUploadedFile(file, fileName)
	if err != nil {
		tool.Failed(context, "头像更新失败")
		return
	}
	// 将保存后得文件本地路径,保存到用户表中得头像字段
	memberService := service.MemberService{}
	path := memberService.UploadAvatar(member.Id, fileName[1:])
	if path != "" {
		tool.Success(context, "http://localhost:8090"+path)
		return
	}
	// 返回结果
	tool.Failed(context, "上传失败")
}
