package dao

import (
	"Go-Learning/Gin/model"
	"Go-Learning/Gin/tool"
	"fmt"
	"log"
)

type MemberDao struct {
	*tool.OrmTool
}

// 验证手机号和验证码是否存在
func (receiver MemberDao) ValidateSmsCode(phone string, code string) *model.SmsCode {
	smsCode := new(model.SmsCode)
	if _, err := receiver.Where("phone=? and code = ?", phone, code).Get(smsCode); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return smsCode
}

// 查询手机号
func (receiver MemberDao) QueryByPhone(phone string) *model.Member {
	member := new(model.Member)
	if _, err := receiver.Where("mobile = ?", phone).Get(member); err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return member
}

// 插入用户信息
func (receiver MemberDao) InsertMember(user *model.Member) {
	result, err := receiver.InsertOne(user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	user.Id = result
}

func (receiver *MemberDao) InsertCode(sms *model.SmsCode) int64 {
	result, err := receiver.InsertOne(sms)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
func (receiver *MemberDao) Query(name string, password string) *model.Member {
	user := new(model.Member)
	_, err := receiver.Where("user_name = ? and password = ?", name, password).Get(user)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	return user
}

func (receiver MemberDao) UpdateMemberAvatar(id int64, fileName string) int64 {
	member := model.Member{Avatar: fileName}
	result, err := receiver.Where("id = ?", id).Update(&member)
	if err != nil {
		fmt.Println(err.Error())
		return 0
	}
	return result
}
