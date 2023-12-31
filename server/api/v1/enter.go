package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/business"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/wechat"
)

type ApiGroup struct {
	SystemApiGroup   system.ApiGroup
	ExampleApiGroup  example.ApiGroup
	BusinessApiGroup business.ApiGroup
	WechatApiGroup   wechat.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
