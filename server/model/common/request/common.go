package request

// PageInfo Paging common input parameter structure
type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

// GetById Find by id structure
type GetById struct {
	ID int `json:"id" form:"id"` // 主键ID
}

func (r *GetById) Uint() int {
	return r.ID
}

type IdsReq struct {
	Ids []int `json:"ids" form:"ids"`
}

// GetAuthorityId Get role by id structure
type GetAuthorityId struct {
	AuthorityId int `json:"authorityId" form:"authorityId"` // 角色ID
}

type Empty struct{}

type MemberSearchInfo struct {
	Telephone  int    `json:"telephone" form:"telephone"`   // 联系电话
	MemberName string `json:"memberName" form:"memberName"` // 姓名
	Deadline   string `json:"deadline" form:"deadline"`     // 截止时间
	State      int    `json:"state" form:"state"`           // 状态
	Page       int    `json:"page" form:"page"`             // 页码
	PageSize   int    `json:"pageSize" form:"pageSize"`     // 每页大小
}

type CardInfo struct {
	OnlyId int `json:"onlyId" form:"onlyId"` // 唯一标识
}

type ConsumeSearchInfo struct {
	Telephone int `json:"telephone" form:"telephone"` // 联系电话
	State     int `json:"state" form:"state"`         // 状态
	Page      int `json:"page" form:"page"`           // 页码
	PageSize  int `json:"pageSize" form:"pageSize"`   // 每页大小
}

type OrderSearchInfo struct {
	OrderId   string `json:"orderId" form:"orderId" gorm:"comment:订单编号"`
	Telephone int    `json:"telephone" form:"telephone"` // 联系电话
	Type      int    `json:"type" form:"type" gorm:"comment:订单类型"`
	BuyDate   string `json:"buyDate" form:"buyDate" gorm:"comment:购买日期"`
	Page      int    `json:"page" form:"page"`         // 页码
	PageSize  int    `json:"pageSize" form:"pageSize"` // 每页大小
}

type StatisticsSearchInfo struct {
	StartDate string `json:"startDate" form:"startDate" gorm:"comment:开始日期"`
	EndDate   string `json:"endDate" form:"endDate" gorm:"comment:结束日期"`
}

type ProductSearchInfo struct {
	Name                string `json:"name" form:"name" gorm:"comment:商品编号"`
	BrandId             int    `json:"brandId" form:"brandId" gorm:"comment:物品序号"`
	ProductSN           string `json:"productSN" form:"productSN" gorm:"comment:货号"`
	ProductCategoryName string `json:"productCategoryName" form:"productCategoryName" gorm:"comment:商品分类"`
	BrandName           string `json:"brandName" form:"brandName" gorm:"comment:品牌"`
	PublishStatus       int    `json:"publishStatus" form:"publishStatus" gorm:"comment:上架状态"`
	VerifyStatus        int    `json:"verifyStatus" form:"verifyStatus" gorm:"comment:审核状态"`
	Page                int    `json:"page" form:"page"`         // 页码
	PageSize            int    `json:"pageSize" form:"pageSize"` // 每页大小
}

type TagSearchInfo struct {
	Tag      int `json:"tag" form:"tag"`           // 商品分类id
	State    int `json:"state" form:"state"`       // 状态
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// KeySearchInfo 根据数据库列名查找
type KeySearchInfo struct {
	Key      string `json:"key" form:"key"`           // 数据库列名
	ID       int    `json:"id" form:"id"`             // ID
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
}

type SortSearchInfo struct {
	Key      string `json:"key" form:"key"`           // 数据库列名
	ID       int    `json:"id" form:"id"`             // ID
	Sort     int    `json:"sort" form:"sort"`         // ID
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
}

type KeyWordInfo struct {
	ID      int    `json:"id" form:"id"`           // 主键ID
	Keyword string `json:"keyword" form:"keyword"` //关键字
}

type QuantityInfo struct {
	ID       int `json:"id" form:"id"`             // 主键ID
	Quantity int `json:"quantity" form:"quantity"` // 数量
}
