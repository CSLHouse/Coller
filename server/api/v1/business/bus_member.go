package business

import (
	"github.com/bwmarrin/snowflake"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessReq "github.com/flipped-aurora/gin-vue-admin/server/model/business/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	date_conversion "github.com/flipped-aurora/gin-vue-admin/server/utils/timer"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type MemberApi struct{}

// CreateVIPMember
// @Tags      ExaCustomer
// @Summary   创建客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户用户名, 客户手机号码"
// @Success   200   {object}  response.Response{msg=string}  "创建客户"
// @Router    /customer/customer [post]
func (e *MemberApi) CreateVIPMember(c *gin.Context) {
	var member businessReq.CreateCardRequest
	err := c.ShouldBindJSON(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	userId := utils.GetUserID(c)
	isTmp := false
	if len(member.UserName) < 1 && member.ComboId < 1 && member.RemainTimes < 1 {
		isTmp = true
	}
	var customer business.Customer
	customer.Telephone = member.Telephone
	customer.UserName = member.UserName
	err = memberService.CreateCustomerFormWeb(&customer)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}

	comboData, err := comboService.GetVIPComboById(member.ComboId)
	var card business.VIPCard
	card.CardId = member.CardID
	card.UserName = member.UserName
	card.Telephone = member.Telephone
	card.StartDate = member.StartDate
	card.Deadline = date_conversion.DateYearLater(member.StartDate, 1)
	card.State = 1
	card.RemainTimes += comboData.Times
	card.IsNew = true
	card.ComboId = member.ComboId
	card.Collection = member.Collection
	card.SysUserId = userId
	card.CustomerId = customer.ID
	card.StoreName = utils.GetNickName(c)

	var certificate business.VIPCertificate
	certificate.Telephone = member.Telephone
	certificate.StoreName = utils.GetNickName(c)
	certificate.SysUserId = utils.GetUserID(c)
	err = memberService.CreateVIPCertificate(&certificate)
	if err != nil {
		global.GVA_LOG.Error("创建会员消费凭证失败!", zap.Error(err))
		response.FailWithMessage("创建会员消费凭证失败", c)
		return
	}

	// 临时会员
	if isTmp {
		err = memberService.CreateVIPCard(&card)
		if err != nil {
			global.GVA_LOG.Error("创建会员卡失败!", zap.Error(err))
			response.FailWithMessage("创建会员卡失败", c)
			return
		}
		response.OkWithMessage("创建成功", c)
		return
	}

	var order business.VIPOrder
	n, err := snowflake.NewNode(1)
	if err != nil {
		global.GVA_LOG.Error("创建id失败!", zap.Error(err))
	}
	id := n.Generate()
	order.OrderID = int64(id)
	order.Telephone = member.Telephone
	//order.MemberName = member.MemberName
	//order.ComboId = member.ComboId
	//order.BuyDate = member.StartDate
	order.State = 1
	order.IsNew = true
	order.Type = 1
	//order.Collection = member.Collection
	order.SysUserId = userId
	//err = memberService.CreateVIPMember(member)
	//if err != nil {
	//	global.GVA_LOG.Error("创建失败!", zap.Error(err))
	//	response.FailWithMessage("创建失败", c)
	//	return
	//}

	//err = orderService.CreateVIPOrder(order)
	//if err != nil {
	//	global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
	//	response.FailWithMessage("创建订单失败", c)
	//	return
	//}

	var statement business.VIPStatement
	statement.Recharge = member.Collection
	statement.NewMember = 1
	statement.SysUserId = userId
	//err = orderService.CreateVIPStatement(statement)
	//if err != nil {
	//	global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
	//	response.FailWithMessage("创建订单失败", c)
	//	return
	//}
	// TODO: 事务处理，错误回退
	var statistics business.VIPStatistics
	statistics.TotalStream = float64(member.Collection)
	statistics.TotalOrder = 1
	statistics.TotalMember = 1
	statistics.SysUserId = userId
	//err = orderService.CreateVIPStatistics(statistics)
	//if err != nil {
	//	global.GVA_LOG.Error("统计失败!", zap.Error(err))
	//	response.FailWithMessage("统计失败", c)
	//	return
	//}

	err = memberService.CreateVIPMemberSynchronous(&card, &order, &statement, &statistics)
	if err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
		return
	}
	response.OkWithMessage("创建成功", c)
}

// DeleteExaCustomer
// @Tags      ExaCustomer
// @Summary   删除客户
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID"
// @Success   200   {object}  response.Response{msg=string}  "删除客户"
// @Router    /customer/customer [delete]
func (e *MemberApi) DeleteVIPMemberById(c *gin.Context) {
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

	err = memberService.DeleteVIPMemberById(reqId.ID)
	if err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
		return
	}
	response.OkWithMessage("删除成功", c)
}

// UpdateExaCustomer
// @Tags      ExaCustomer
// @Summary   更新客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  body      example.ExaCustomer            true  "客户ID, 客户信息"
// @Success   200   {object}  response.Response{msg=string}  "更新客户信息"
// @Router    /customer/customer [put]
func (e *MemberApi) UpdateVIPMember(c *gin.Context) {
	var member businessReq.CreateCardRequest
	err := c.ShouldBindJSON(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(member, utils.MemberVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	var customer business.Customer
	customer.Telephone = member.Telephone
	customer.UserName = member.UserName
	err = memberService.CreateCustomerFormWeb(&customer)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}

	comboData, err := comboService.GetVIPComboById(member.ComboId)
	var card business.VIPCard
	card.CardId = member.CardID
	card.StartDate = member.StartDate
	card.Deadline = date_conversion.DateYearLater(member.StartDate, 1)
	card.RemainTimes += comboData.Times
	card.ComboId = member.ComboId
	card.Collection = member.Collection
	card.CustomerId = customer.ID

	err = memberService.UpdateVIPCard(&card)
	if err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

// GetExaCustomer
// @Tags      ExaCustomer
// @Summary   获取单一客户信息
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     example.ExaCustomer                                                true  "客户ID"
// @Success   200   {object}  response.Response{data=exampleRes.ExaCustomerResponse,msg=string}  "获取单一客户信息,返回包括客户详情"
// @Router    /customer/customer [get]
//func (e *MemberApi) GetVIPMember(c *gin.Context) {
//	var member business.Customer
//	err := c.ShouldBindQuery(&member)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	err = utils.Verify(member.GVA_MODEL, utils.IdVerify)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//	data, err := memberService.GetVIPMemberById(member.ID)
//	if err != nil {
//		global.GVA_LOG.Error("获取失败!", zap.Error(err))
//		response.FailWithMessage("获取失败", c)
//		return
//	}
//	response.OkWithDetailed(businessRes.VIPMemberResponse{Member: data}, "获取成功", c)
//}

// GetVIPMemberList
// @Tags      ExaCustomer
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *MemberApi) GetVIPMemberList(c *gin.Context) {
	var pageInfo request.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(pageInfo, utils.PageInfoVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	memberList, total, err := memberService.GetVIPMemberInfoList(utils.GetUserID(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithDetailed(response.PageResult{
		List:     memberList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

func (e *MemberApi) SearchVIPMember(c *gin.Context) {
	var searchInfo request.MemberSearchInfo
	err := c.ShouldBindQuery(&searchInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, total, err := memberService.SearchVIPMember(utils.GetUserID(c), searchInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List:     data,
		Total:    total,
		Page:     searchInfo.Page,
		PageSize: searchInfo.PageSize,
	}, "获取成功", c)
}

func (e *MemberApi) SearchVIPCard(c *gin.Context) {
	var cardInfo request.CardInfo
	err := c.ShouldBindQuery(&cardInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	data, err := memberService.SearchVipCard(utils.GetUserID(c), cardInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	response.OkWithDetailed(response.AllResult{
		List: data,
	}, "获取成功", c)
}

// 续费
func (e *MemberApi) RenewVIPCard(c *gin.Context) {
	var member businessReq.RenewCardRequest
	err := c.ShouldBindQuery(&member)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	oldData, err := memberService.GetVIPCardById(member.ID)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}
	comboData, err := comboService.GetVIPComboById(member.ComboId)
	remainTimes := oldData.RemainTimes
	oldData.RemainTimes = remainTimes + comboData.Times + member.Times
	oldData.ComboId = member.ComboId
	oldData.StartDate = date_conversion.BuildTheDayStr()
	oldData.Deadline = date_conversion.DateYearLater(oldData.StartDate, 1)
	oldData.IsNew = false

	//err = memberService.UpdateVIPMember(&member)
	//if err != nil {
	//	global.GVA_LOG.Error("更新失败!", zap.Error(err))
	//	response.FailWithMessage("更新失败", c)
	//	return
	//}
	var statement business.VIPStatement
	statement.Recharge = member.Collection
	statement.SysUserId = oldData.SysUserId
	//err = orderService.CreateVIPStatement(statement)
	//if err != nil {
	//	global.GVA_LOG.Error("创建订单失败!", zap.Error(err))
	//	response.FailWithMessage("创建订单失败", c)
	//	return
	//}

	var statistics business.VIPStatistics
	statistics.TotalStream = float64(member.Collection)
	statistics.TotalOrder = 1
	statistics.SysUserId = oldData.SysUserId
	//err = orderService.CreateVIPStatistics(statistics)
	//if err != nil {
	//	global.GVA_LOG.Error("统计失败!", zap.Error(err))
	//	response.FailWithMessage("统计失败", c)
	//	return
	//}
	err = memberService.UpdateVIPMemberSynchronous(&oldData, &statement, &statistics)
	if err != nil {
		global.GVA_LOG.Error("统计失败!", zap.Error(err))
		response.FailWithMessage("统计失败", c)
		return
	}
	response.OkWithMessage("更新成功", c)
}

func (e *MemberApi) GetVipCardList(c *gin.Context) {
	var cardInfo request.CardInfo
	err := c.ShouldBindQuery(&cardInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(cardInfo, utils.CardVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	memberList, err := memberService.GetVIPCardByTelephone(cardInfo.OnlyId)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}

	response.OkWithData(memberList, c)
}

func (e *MemberApi) GetCertificateList(c *gin.Context) {
	var cardInfo request.CardInfo
	err := c.ShouldBindQuery(&cardInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	err = utils.Verify(cardInfo, utils.CardVerify)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	certificateList, err := memberService.GetVIPCertificateByTelephone(cardInfo.OnlyId)
	if err != nil {
		global.GVA_LOG.Error("获取消费凭证失败!", zap.Error(err))
		response.FailWithMessage("获取消费凭证失败"+err.Error(), c)
		return
	}

	response.OkWithData(certificateList, c)
}
