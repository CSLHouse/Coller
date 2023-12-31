package business

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type MemberRouter struct{}

func (e *MemberRouter) InitMemberRouter(Router *gin.RouterGroup) {
	businessRouter := Router.Group("business").Use(middleware.OperationRecord())
	businessRouterWithoutRecord := Router.Group("business")
	vipMemberApi := v1.ApiGroupApp.BusinessApiGroup.MemberApi
	{
		businessRouter.POST("member", vipMemberApi.CreateVIPMember)   // 创建客户
		businessRouter.PUT("member", vipMemberApi.UpdateVIPMember)    // 更新客户
		businessRouter.DELETE("member", vipMemberApi.DeleteVIPMember) // 删除客户
		businessRouter.POST("renew", vipMemberApi.RenewVIPCard)       // VIP客户续费
	}
	{
		businessRouterWithoutRecord.GET("member", vipMemberApi.GetVIPMember)         // 获取单一客户信息
		businessRouterWithoutRecord.GET("memberList", vipMemberApi.GetVIPMemberList) // 获取客户列表
		businessRouterWithoutRecord.POST("memberSearch", vipMemberApi.SearchVIPMember)
		businessRouterWithoutRecord.POST("searchCard", vipMemberApi.SearchVIPCard)
	}
}
