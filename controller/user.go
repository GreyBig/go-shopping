package controller

import (
	"shopping/common"
	"shopping/common/response"
	"shopping/models"
	"shopping/service"

	"github.com/gin-gonic/gin"
)

var user service.WebUserService

// WebGetCaptcha 后台管理前端，获取验证码
func WebGetCaptcha(c *gin.Context) {
	id, b64s, _ := common.GenerateCaptcha()
	data := map[string]interface{}{"captchaId": id, "captchaImg": b64s}
	response.Success("操作成功", data, c)
}

// WebUserLogin 后台管理前端，用户登录
func WebUserLogin(c *gin.Context) {
	var param models.AdminWebUserLoginVO
	if err := c.ShouldBind(&param); err != nil {
		response.Failed("请求参数无效", c)
		return
	}

	// 检查验证码
	if !common.VerifyCaptcha(param.CaptchaId, param.CaptchaValue) {
		response.Failed("验证码错误", c)
		return
	}

	// 生成token
	uid := user.Login(param)
	if uid > 0 {
		token, _ := common.GenerateToken(param.Username)
		userInfo := models.AdminWebUserInfo{
			Uid:   uid,
			Token: token,
		}
		response.Success("登录成功", userInfo, c)
		return
	}
	response.Failed("用户名或密码错误", c)
}
