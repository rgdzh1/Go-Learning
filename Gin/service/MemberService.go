package service

import (
	"Go-Learning/Gin/dao"
	"Go-Learning/Gin/model"
	"Go-Learning/Gin/param"
	"Go-Learning/Gin/tool"
	"encoding/json"
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
	"log"
	"math/rand"
	"time"
)

type MemberService struct {
}

/**
用户登录
*/
func (receiver MemberService) SmsLogin(param *param.SmsLoginParam) *model.Member {
	// 1.获取手机号和验证码
	mDao := dao.MemberDao{
		OrmTool: tool.DBEngine,
	}
	// 2.验证手机号+验证码是否正确
	smsCode := mDao.ValidateSmsCode(param.Phone, param.Code)
	if smsCode.Id == 0 {
		return nil
	}
	// 3.根据手机号member表中查询记录
	member := mDao.QueryByPhone(smsCode.Phone)
	if member.Id != 0 {
		// 用户已经存在,返回登录成功
		return member
	}
	// 4.新创建一个member记录,并保存
	user := new(model.Member)
	user.UserName = param.Phone
	user.Mobile = param.Phone
	user.RegisterTime = time.Now()
	// 插入用户
	mDao.InsertMember(user)
	return user
}

/**
发送密码
*/
func (receiver MemberService) SendCode(phone string) bool {
	// 1.产生验证码 %04v 表示产生的四位数的验证字符串,后面4个0  10000
	// %06v 产生四位数验证字符串,后面6个0 1000000
	code := fmt.Sprintf("%04v", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(10000))
	// 调用阿里云SDK
	config := tool.GetConfig()
	client, err := dysmsapi.NewClientWithAccessKey(config.Sms.RegionId, config.Sms.AppKey, config.Sms.AppSecret)
	if err != nil {
		log.Fatal(err.Error())
		return false
	}
	request := dysmsapi.CreateSendSmsRequest()
	request.Scheme = "https"
	request.SignName = config.Sms.SignName
	request.TemplateCode = config.Sms.TemplateCode
	request.PhoneNumbers = phone
	param, err := json.Marshal(map[string]interface{}{
		"code": code,
	})
	request.TemplateParam = string(param)
	response, err := client.SendSms(request)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	// log.Fatal(response)
	if response.Code == "OK" {
		// 将用户获取的验证码保存到数据库中
		smsCode := model.SmsCode{
			Phone:      phone,
			Code:       code,
			BizId:      response.BizId,
			CreateTime: time.Now(),
		}
		// 创建MemberDao对象,调用xorm插入数据的方法
		memberDao := dao.MemberDao{
			OrmTool: tool.DBEngine,
		}
		result := memberDao.InsertCode(&smsCode)
		return result > 0
	} else {

		smsCode := model.SmsCode{
			Phone:      phone,
			Code:       code,
			BizId:      response.BizId,
			CreateTime: time.Now(),
		}
		memberDao := dao.MemberDao{
			tool.DBEngine,
		}
		result := memberDao.InsertCode(&smsCode)
		log.Println("短信信息插入code", result, "打印错误日志", response)
	}
	return false
}

/**
登录
*/
func (receiver MemberService) Login(name string, password string) *model.Member {
	// 1 使用用户名+密码查询用户是否存在,如果存在就直接返回
	memberDao := dao.MemberDao{
		tool.DBEngine,
	}
	member := memberDao.Query(name, password)
	if member.Id != 0 {
		return member
	}
	// 2 用户信息不存在,作为新用户保存到数据库中
	user := new(model.Member)
	user.UserName = name
	user.Password = password
	user.RegisterTime = time.Now()
	memberDao.InsertMember(user)
	return user
}

/**
更新图像
*/
func (receiver MemberService) UploadAvatar(id int64, fileName string) string {
	memberDao := dao.MemberDao{tool.DBEngine}
	result := memberDao.UpdateMemberAvatar(id, fileName)
	if result == 0 {
		return ""
	}
	return fileName
}
