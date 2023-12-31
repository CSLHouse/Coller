package business

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ComboRouter struct{}

func (e *ComboRouter) InitComboRouter(Router *gin.RouterGroup) {
	businessRouter := Router.Group("business").Use(middleware.OperationRecord())
	businessRouterWithoutRecord := Router.Group("business")
	exaComboApi := v1.ApiGroupApp.BusinessApiGroup.ComboApi
	{
		businessRouter.POST("combo", exaComboApi.CreateVIPCombo)   // 创建客户
		businessRouter.PUT("combo", exaComboApi.UpdateVIPCombo)    // 更新客户
		businessRouter.DELETE("combo", exaComboApi.DeleteVIPCombo) // 删除客户
	}
	{
		businessRouterWithoutRecord.GET("combo", exaComboApi.GetVIPCombo)         // 获取单一客户信息
		businessRouterWithoutRecord.GET("comboList", exaComboApi.GetVIPComboList) // 获取客户列表
		businessRouterWithoutRecord.GET("allCombo", exaComboApi.GetAllVIPCombos)  // 获取客户列表
	}
}
