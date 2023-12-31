package wechat

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	wechatReq "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/request"
	wechatRes "github.com/flipped-aurora/gin-vue-admin/server/model/wechat/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type WXAccountApi struct{}

func (b *WXAccountApi) WXLogin(c *gin.Context) {
	var login wechatReq.WXLogin
	err := c.ShouldBindJSON(&login)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(login.Code) < 1 {
		fmt.Println("----login:", login)
		response.FailWithMessage(err.Error(), c)
		return
	}
	//发送jscode，获得用户的open_id
	wechatClient := middleware.NewWechatClient(nil)
	wxMap, err := wechatClient.WXLogin(login.Code)
	if err != nil {
		global.GVA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败", c)
		return
	}
	var session wechatRes.WXLoginRes
	session.OpenID = wxMap["openid"]
	if len(session.OpenID) < 1 {
		global.GVA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败", c)
		return
	}
	var wxUser wechat.WXUser
	wxUser.OpenId = wxMap["openid"]
	wxUser.SessionKey = wxMap["session_key"]
	//wxUser.Token = b.CreateToken(wxMap["openid"], userInfo.NickName)
	wxUser.Count = 1
	fmt.Println("---wxUser:", wxUser)
	err = accountService.CreateWXAccount(wxUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(session, "获取成功", c)
}

func (b *WXAccountApi) UpdateUserInfo(c *gin.Context) {
	var userInfo wechatReq.WXUserInfo
	err := c.ShouldBindJSON(&userInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("---userInfo:", userInfo)
	err = utils.Verify(userInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var wxUser wechat.WXUser
	wxUser.OpenId = userInfo.OpenID
	wxUser.NickName = userInfo.NickName
	wxUser.Gender = userInfo.Gender
	wxUser.AvatarUrl = userInfo.AvatarUrl
	wxUser.Token = b.CreateToken(userInfo.OpenID, userInfo.NickName)
	fmt.Println("---wxUser:", wxUser)
	err = accountService.UpdateWXAccountInfo(wxUser)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.OkWithDetailed(wxUser, "更新成功", c)
}

// CreateToken 登录以后签发jwt
func (b *WXAccountApi) CreateToken(openId string, nickName string) string {
	j := &utils.JWT{SigningKey: []byte(global.GVA_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateWXClaims(systemReq.WXBaseClaims{
		OpenId:   openId,
		NickName: nickName,
	})
	token, err := j.CreateWXToken(claims)
	if err != nil {
		global.GVA_LOG.Error("获取token失败!", zap.Error(err))
		return ""
	}
	return token
}

// GetUserInfo 获取用户详情
func (b *WXAccountApi) GetUserInfo(c *gin.Context) {
	var user wechatReq.UserTag
	err := c.ShouldBindQuery(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wxUser, err := accountService.GetWXAccountByOpenID(user.OpenID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithData(wxUser, c)
}

func (b *WXAccountApi) ParsePhoneNumber(c *gin.Context) {
	var loginInfo wechatReq.WXPhoneNumber
	err := c.ShouldBindJSON(&loginInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	wxUser, err := accountService.GetWXAccountByOpenID(loginInfo.OpenID)
	if len(wxUser.SessionKey) < 1 {
		fmt.Println("----login:", loginInfo)
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("----SessionKey:", wxUser.SessionKey)
	//发送jscode，获得用户的open_id
	wechatClient, err := middleware.DecryptWXOpenData(wxUser.SessionKey, loginInfo.EncryptedData, loginInfo.Iv)

	if err != nil {
		global.GVA_LOG.Error("解密失败!", zap.Error(err))
		response.FailWithMessage("解密失败", c)
		return
	}
	fmt.Println("----wechatClient:", wechatClient)
	var phoneNumber = wechatClient["phoneNumber"].(string)
	if len(phoneNumber) < 1 {
		global.GVA_LOG.Error("登录失败!", zap.Error(err))
		response.FailWithMessage("登录失败", c)
		return
	}
	err = accountService.UpdateWXAccountPhone(loginInfo.OpenID, phoneNumber)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	response.OkWithDetailed(wechatRes.WXPhoneNum{
		PhoneNumber: phoneNumber,
	}, "获取成功", c)
}

// CheckPhoneNumber 查询是否有手机号 true 有 false 无
func (b *WXAccountApi) CheckPhoneNumber(c *gin.Context) {
	var loginInfo wechatReq.UserTag
	err := c.ShouldBindQuery(&loginInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(loginInfo.OpenID) < 1 {
		fmt.Println("----login:", loginInfo)
		response.FailWithMessage(err.Error(), c)
		return
	}

	userInfo, err := accountService.CheckWXAccountPhone(loginInfo.OpenID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if len(userInfo.PhoneNumber) < 11 {
		response.OkWithData(false, c)
		return
	}
	response.OkWithData(true, c)
}

func (b *WXAccountApi) CreateMemberReceiveAddress(c *gin.Context) {
	var address wechat.MemberReceiveAddress
	err := c.ShouldBindJSON(&address)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if userId < 1 {
		response.FailWithMessage("获取UserID失败", c)
		return
	}
	address.UserId = userId
	err = accountService.CreateMemberReceiveAddress(&address)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// GetMemberReceiveAddressList 获取会员地址列表
func (b *WXAccountApi) GetMemberReceiveAddressList(c *gin.Context) {
	var user wechatReq.UserTag
	err := c.ShouldBindQuery(&user)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	if userId < 1 {
		response.FailWithMessage("获取UserID失败", c)
		return
	}
	addressList, err := accountService.GetMemberReceiveAddressList(userId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithData(addressList, c)
}

// GetMemberReceiveAddressById 获取会员地址列表
func (b *WXAccountApi) GetMemberReceiveAddressById(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindQuery(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	address, err := accountService.GetMemberReceiveAddressById(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	response.OkWithData(address, c)
}

// UpdateMemberReceiveAddress 更新会员地址列表
func (b *WXAccountApi) UpdateMemberReceiveAddress(c *gin.Context) {
	var address wechat.MemberReceiveAddress
	err := c.ShouldBindJSON(&address)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = accountService.UpdateMemberReceiveAddress(&address)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// DeleteMemberReceiveAddress 删除会员地址列表
func (b *WXAccountApi) DeleteMemberReceiveAddress(c *gin.Context) {
	var reqId request.GetById
	err := c.ShouldBindJSON(&reqId)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(reqId, utils.IdVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	err = accountService.DeleteMemberReceiveAddress(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}
