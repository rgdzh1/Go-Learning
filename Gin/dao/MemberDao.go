package dao

import (
	"Go-Learning/Gin/model"
	"Go-Learning/Gin/tool"
	"log"
)

type MemberDao struct {
	*tool.OrmTool
}

func (receiver *MemberDao) InsertCode(sms *model.SmsCode) int64 {
	result, err := receiver.InsertOne(sms)
	if err != nil {
		log.Fatal(err.Error())
	}
	return result
}
