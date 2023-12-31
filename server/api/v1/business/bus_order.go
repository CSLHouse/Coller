package business

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	businessRes "github.com/flipped-aurora/gin-vue-admin/server/model/business/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type OrderApi struct{}

//func (e *OrderApi) CreateVIPOrder(c *gin.Context) {
//	var order business.VIPOrder
//	err := c.ShouldBindJSON(&order)
//	if err != nil {
//		response.FailWithMessage(err.Error(), c)
//		return
//	}
//
//	order.SysUserAuthorityID = utils.GetUserAuthorityId(c)
//	order.ComboType = comboData.ComboType
//	order.State = 1 // 默认Vip会员
//	order.IsNew = true
//	err = orderService.CreateVIPOrder(order)
//	if err != nil {
//		global.GVA_LOG.Error("创建失败!", zap.Error(err))
//		response.FailWithMessage("创建失败", c)
//		return
//	}
//	response.OkWithMessage("创建成功", c)
//}

// GetVIPOrderList
// @Tags      ExaCustomer
// @Summary   分页获取权限客户列表
// @Security  ApiKeyAuth
// @accept    application/json
// @Produce   application/json
// @Param     data  query     request.PageInfo                                        true  "页码, 每页大小"
// @Success   200   {object}  response.Response{data=response.PageResult,msg=string}  "分页获取权限客户列表,返回包括列表,总数,页码,每页数量"
// @Router    /customer/customerList [get]
func (e *OrderApi) GetVIPOrderList(c *gin.Context) {
	var pageInfo request.OrderSearchInfo
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
	orderList, total, err := orderService.GetVIPOrderInfoList(utils.GetUserAuthorityId(c), pageInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	orderData := make([]businessRes.VipOrderResModel, 0)
	allOrder := orderList.([]business.VIPOrder)
	for _, order := range allOrder {
		orderRes := businessRes.VipOrderResModel{}
		orderRes.ID = order.ID
		orderRes.OrderID = order.OrderID
		orderRes.Telephone = order.Telephone
		orderRes.MemberName = order.Member.MemberName
		orderRes.ComboId = order.Member.ComboId
		orderRes.ComboType = order.Member.Combo.ComboName
		orderRes.ComboPrice = order.Member.Combo.ComboPrice
		orderRes.BuyDate = order.BuyDate
		orderRes.State = order.State
		orderRes.IsNew = order.IsNew
		orderRes.Type = order.Type
		orderRes.Collection = order.Member.Collection
		orderData = append(orderData, orderRes)
	}
	response.OkWithDetailed(response.PageResult{
		List:     orderData,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}, "获取成功", c)
}

// GetVIPStatementList 获取流水数据
func (e *OrderApi) GetVIPStatementList(c *gin.Context) {
	var statisticsInfo request.StatisticsSearchInfo
	err := c.ShouldBindQuery(&statisticsInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	orderList, err := orderService.GetVIPStatementInfoList(utils.GetUserAuthorityId(c), statisticsInfo)
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.PageResult{
		List: orderList,
	}, "获取成功", c)
}

// GetVIPStatisticsList 获取统计数据
func (e *OrderApi) GetVIPStatisticsList(c *gin.Context) {
	statistics, err := orderService.GetVIPStatisticsInfoList(utils.GetUserAuthorityId(c))
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败"+err.Error(), c)
		return
	}
	response.OkWithDetailed(response.AllResult{
		List: statistics,
	}, "获取成功", c)
}
