package service

import (
	"Go-Learning/Gin/dao"
	"Go-Learning/Gin/model"
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
