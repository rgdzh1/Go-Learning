package tool

import (
	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"image/color"
)

type CaptchaResult struct {
	Id           string `json:"id"`
	Base64Blob   string `json:"base_64_blob"`
	VertifyValue string `json:"vertify_value"`
}

/**
生成图形化验证码
*/
func GenerateCaptcha(context *gin.Context) CaptchaResult {
	parameters := base64Captcha.ConfigCharacter{
		Height:             30,
		Width:              60,
		Mode:               3,
		ComplexOfNoiseDot:  0,
		ComplexOfNoiseText: 0,
		IsUseSimpleFont:    true,
		IsShowHollowLine:   false,
		IsShowNoiseDot:     false,
		IsShowNoiseText:    false,
		IsShowSlimeLine:    false,
		IsShowSineLine:     false,
		CaptchaLen:         4,
		BgColor: &color.RGBA{
			R: 3,
			G: 102,
			B: 214,
			A: 254,
		},
	}
	captchaId, instance := base64Captcha.GenerateCaptcha("", parameters)
	base64Blob := base64Captcha.CaptchaWriteToBase64Encoding(instance)
	return CaptchaResult{
		Id:         captchaId,
		Base64Blob: base64Blob,
	}
}

/**
这里验证的时候会自动调用刚刚实现的SET POST
*/
func VerifyCaptcha(id string, value string) bool {
	result := base64Captcha.VerifyCaptcha(id, value)
	return result
}
