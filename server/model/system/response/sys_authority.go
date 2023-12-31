package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/system"

type SysAuthorityResponse struct {
	Authority system.SysAuthority `json:"authority"`
}

type SysAuthorityCopyResponse struct {
	Authority      system.SysAuthority `json:"authority"`
	OldAuthorityId int                 `json:"oldAuthorityId"` // 旧角色ID
}
