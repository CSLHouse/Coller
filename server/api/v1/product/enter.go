package product

import "github.com/flipped-aurora/gin-vue-admin/server/service"

type ApiGroup struct {
	FlashApi
}

var (
	wechatService                = service.ServiceGroupApp.WechatServiceGroup.HomeService
	accountService               = service.ServiceGroupApp.WechatServiceGroup.AccountService
	orderService                 = service.ServiceGroupApp.WechatServiceGroup.OrderService
	jspaymentService             = service.ServiceGroupApp.PaymentServiceGroup.PayMentService
	fileUploadAndDownloadService = service.ServiceGroupApp.ExampleServiceGroup.FileUploadAndDownloadService
)
