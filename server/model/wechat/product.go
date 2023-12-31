package wechat

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// HomeAdvertise 首页轮播广告表
type HomeAdvertise struct {
	global.GVA_MODEL
	Name       string `json:"name" gorm:"not null;comment:名称;size:100"`                  // 套餐ID
	Type       int    `json:"type" form:"type" gorm:"comment:轮播位置：0->PC首页轮播；1->app首页轮播"` // 套餐名称
	Pic        string `json:"pic" form:"pic" gorm:"comment:图片"`                          // 套餐类型
	StartTime  string `json:"startTime" form:"startTime" gorm:"comment:开始时间"`            // 套餐价格
	EndTime    string `json:"endTime" form:"endTime" gorm:"comment:结束时间"`                // 天数/次数/金额
	State      int    `json:"state" form:"state" gorm:"comment:上下线状态：0->下线；1->上线"`       // 状态
	ClickCount int    `json:"clickCount" form:"clickCount" gorm:"comment:点击数"`
	OrderCount int    `json:"orderCount" form:"orderCount" gorm:"comment:下单数"`
	Url        string `json:"url" form:"url" gorm:"comment:链接地址;size:500"`
	Note       string `json:"note" form:"note" gorm:"comment:备注;size:500"`
	Sort       int    `json:"sort" form:"sort" gorm:"comment:排序"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeAdvertise) TableName() string {
	return "sms_home_advertise"
}

// HomeBrand 品牌表
type HomeBrand struct {
	global.GVA_MODEL
	Name                string `json:"name" gorm:"not null;comment:名称;size:100"`
	FirstLetter         string `json:"firstLetter" form:"firstLetter" `
	Sort                int    `json:"sort" form:"sort" gorm:"comment:排序"`
	FactoryStatus       int    `json:"factoryStatus" form:"factoryStatus" gorm:"comment:是否是品牌制造商：0->不是；1->是"`
	ShowStatus          int    `json:"showStatus" form:"showStatus" gorm:"comment:是否显示"`
	ProductCount        int    `json:"productCount" form:"productCount" gorm:"comment:产品数量"`
	ProductCommentCount int    `json:"productCommentCount" form:"productCommentCount" gorm:"comment:产品评论数量"`
	Logo                string `json:"logo" gorm:"not null;comment:品牌logo"`
	BigPic              string `json:"bigPic" gorm:"not null;comment:专区大图"`
	BrandStory          string `json:"brandStory" gorm:"not null;comment:品牌故事"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeBrand) TableName() string {
	return "sms_home_brand"
}

// HomeFlashPromotion 限时购表
type HomeFlashPromotion struct {
	global.GVA_MODEL
	Title     string `json:"title" gorm:"not null;comment:名称;size:100"`
	StartDate string `json:"startDate" form:"startDate" gorm:"comment:开始日期"`     // 套餐价格
	EndDate   string `json:"endDate" form:"endDate" gorm:"comment:结束日期"`         // 天数/次数/金额
	Status    int    `json:"status" form:"status" gorm:"comment:状态：0->下线；1->上线"` // 状态
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeFlashPromotion) TableName() string {
	return "sms_home_new_product"
}

// HomeNewProduct 新鲜好物表
type HomeNewProduct struct {
	global.GVA_MODEL
	ProductId       int    `json:"productId" gorm:"not null;comment:物品序号"`
	ProductName     string `json:"productName" form:"productName" gorm:"comment:物品名称"`                     // 套餐价格
	RecommendStatus int    `json:"recommendStatus" form:"recommendStatus" gorm:"comment:推荐状态：0->下线；1->上线"` // 状态
	Sort            int    `json:"sort" form:"sort" gorm:"comment:排序"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeNewProduct) TableName() string {
	return "home_new_product"
}

// HomeHotProduct 首页推荐专题表 猜你喜欢
type HomeHotProduct struct {
	global.GVA_MODEL
	ProductId       int    `json:"productId" gorm:"not null;comment:物品序号;size:100"`
	ProductName     string `json:"productName" form:"productName" gorm:"comment:物品名称"`                     // 套餐价格
	RecommendStatus int    `json:"recommendStatus" form:"recommendStatus" gorm:"comment:推荐状态：0->下线；1->上线"` // 状态
	Sort            int    `json:"sort" form:"sort" gorm:"comment:排序"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeHotProduct) TableName() string {
	return "sms_home_recommend_subject"
}

// HomeRecommendProduct 人气推荐商品表
type HomeRecommendProduct struct {
	global.GVA_MODEL
	ProductName     string `json:"productName" form:"productName" gorm:"comment:物品名称"`                     // 套餐价格
	RecommendStatus int    `json:"recommendStatus" form:"recommendStatus" gorm:"comment:推荐状态：0->下线；1->上线"` // 状态
	Sort            int    `json:"sort" form:"sort" gorm:"comment:排序"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
	ProductId int         `json:"productId" gorm:"not null;comment:物品序号;size:100"`
	Product   HomeProduct `json:"product" gorm:"foreignKey:ProductId;references:ID;comment:商品"`
}

func (HomeRecommendProduct) TableName() string {
	return "sms_home_recommend_product"
}

// HomeProduct 商品信息
type HomeProduct struct {
	global.GVA_MODEL
	BrandId                    int     `json:"brandId" gorm:"comment:品牌序号"`
	ProductCategoryId          int     `json:"productCategoryId" `
	FeightTemplateId           int     `json:"feightTemplateId" `
	ProductAttributeCategoryId int     `json:"productAttributeCategoryId"`
	Name                       string  `json:"name" gorm:"not null;comment:名称;size:100"`
	Pic                        string  `json:"pic" form:"pic" gorm:"comment:图片"`
	ProductSN                  string  `json:"productSN" form:"productSN" gorm:"comment:货号"`
	DeleteStatus               int     `json:"deleteStatus" gorm:"comment:删除状态：0->未删除；1->已删除"`
	PublishStatus              int     `json:"publishStatus" gorm:"comment:上架状态：0->下架；1->上架"`
	NewStatus                  int     `json:"newStatus" gorm:"comment:新品状态:0->不是新品；1->新品"`
	RecommandStatus            int     `json:"recommandStatus" gorm:"comment:推荐状态；0->不推荐；1->推荐"`
	VerifyStatus               int     `json:"verifyStatus" gorm:"comment:审核状态：0->未审核；1->审核通过"`
	Sort                       int     `json:"sort" form:"sort" gorm:"comment:排序"`
	Sale                       int     `json:"sale" form:"sale" gorm:"comment:销量"`
	Price                      float32 `json:"price" form:"price" gorm:"comment:价格"`
	PromotionPrice             int     `json:"promotionPrice" form:"promotionPrice" gorm:"comment:促销价格"`
	GiftGrowth                 int     `json:"giftGrowth" form:"giftGrowth" gorm:"comment:赠送的成长值"`
	GiftPoint                  int     `json:"giftPoint" form:"giftPoint" gorm:"comment:赠送的积分"`
	UsePointLimit              int     `json:"usePointLimit" form:"usePointLimit" gorm:"comment:限制使用的积分数"`
	SubTitle                   string  `json:"subTitle" form:"subTitle" gorm:"comment:副标题"`
	Description                string  `json:"description" form:"description" gorm:"comment:商品描述"`
	OriginalPrice              float32 `json:"originalPrice" form:"originalPrice" gorm:"comment:市场价"`
	Stock                      int     `json:"stock" form:"stock" gorm:"comment:库存"`
	LowStock                   int     `json:"lowStock" form:"lowStock" gorm:"comment:库存预警值"`
	Unit                       string  `json:"unit" form:"unit" gorm:"comment:单位"`
	Weight                     float32 `json:"weight" form:"weight" gorm:"comment:商品重量，默认为克"`
	PreviewStatus              int     `json:"previewStatus" gorm:"comment:是否为预告商品：0->不是；1->是"`
	ServiceIds                 string  `json:"serviceIds" form:"serviceIds" gorm:"comment:以逗号分割的产品服务：1->无忧退货；2->快速退款；3->免费包邮"`
	Keywords                   string  `json:"keywords" form:"keywords" gorm:"comment:关键字"`
	Note                       string  `json:"note" form:"note" gorm:"comment:备注"`
	AlbumPics                  string  `json:"albumPics" form:"albumPics" gorm:"comment:画册图片(头图)，连产品图片限制为5张，以逗号分割"`
	DetailTitle                string  `json:"detailTitle" form:"detailTitle" gorm:"comment:详情标题"`
	DetailDesc                 string  `json:"detailDesc" form:"detailDesc" gorm:"comment:详情描述"`
	DetailHTML                 string  `json:"detailHTML" form:"detailHTML" gorm:"type:text;comment:产品详情网页内容"`
	DetailMobileHTML           string  `json:"detailMobileHTML" form:"detailMobileHTML" gorm:"type:text;comment:移动端网页详情"`
	PromotionStartDate         string  `json:"promotionStartDate" form:"promotionStartDate" gorm:"type:text;comment:促销开始日期"`
	PromotionEndDate           string  `json:"promotionEndDate" form:"promotionEndDate" gorm:"comment:促销结束日期"`
	PromotionPerLimit          int     `json:"promotionPerLimit" form:"promotionPerLimit" gorm:"comment:活动限购数量"`
	PromotionType              int     `json:"promotionType" form:"promotionType" gorm:"comment:促销类型：0->没有促销使用原价;1->使用促销价；2->使用会员价；3->使用阶梯价格；4->使用满减价格；5->限时购"`
	BrandName                  string  `json:"brandName" form:"brandName" gorm:"comment:品牌名称"`
	ProductCategoryName        string  `json:"productCategoryName" form:"productCategoryName" gorm:"comment:商品分类名称"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeProduct) TableName() string {
	return "pms_product"
}

// HomeProductAttributeCategory 商品分类 产品属性分类表
type HomeProductAttributeCategory struct {
	global.GVA_MODEL
	Name           string `json:"name" gorm:"not null;comment:名称;size:100"`
	AttributeCount int    `json:"attributeCount" gorm:"not null；comment:属性数量"`
	ParamCount     int    `json:"paramCount" gorm:"not null；comment:参数数量"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeProductAttributeCategory) TableName() string {
	return "pms_product_attribute_category"
}

// HomeProductCategory 产品分类
type HomeProductCategory struct {
	global.GVA_MODEL
	ParentId     int    `json:"parentId" gorm:"not null；comment:上级分类的编号：0表示一级分类"`
	Name         string `json:"name" gorm:"not null;comment:名称;size:100"`
	Level        int    `json:"level" gorm:"not null；comment:分类级别：0->1级；1->2级;size:1"`
	ProductCount int    `json:"productCount"`
	ProductUnit  string `json:"productUnit" gorm:"comment:单位"`
	NavStatus    int    `json:"navStatus" gorm:"not null；comment:是否显示在导航栏：0->不显示；1->显示;size:1"`
	ShowStatus   int    `json:"showStatus" gorm:"not null；comment:显示状态：0->不显示；1->显示;size:1"`
	Sort         int    `json:"sort" gorm:"size:11"`
	Icon         string `json:"icon" gorm:"comment:图标"`
	Keywords     string `json:"keywords"`
	Description  string `json:"description" gorm:"not null；comment:描述"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeProductCategory) TableName() string {
	return "pms_product_category"
}

// HomeProductAttribute 商品属性参数表
type HomeProductAttribute struct {
	global.GVA_MODEL
	ProductAttributeCategoryId int    `json:"productAttributeCategoryId" `
	Name                       string `json:"name" gorm:"not null;comment:名称;size:100"`
	SelectType                 int    `json:"selectType" gorm:"not null；comment:属性选择类型：0->唯一；1->单选；2->多选;size:1"`
	InputType                  int    `json:"inputType" gorm:"comment:属性录入方式：0->手工录入；1->从列表中选取"`
	InputList                  string `json:"inputList" gorm:"comment:可选值列表，以逗号隔开"`
	Sort                       int    `json:"sort" gorm:"size:11"`
	FilterType                 int    `json:"filterType" gorm:"not null；comment:分类筛选样式：1->普通；1->颜色;size:1"`
	SearchType                 int    `json:"searchType" gorm:"not null；comment:检索类型；0->不需要进行检索；1->关键字检索；2->范围检索;size:1"`
	RelatedStatus              int    `json:"relatedStatus" gorm:"comment:相同属性产品是否关联；0->不关联；1->关联;size:1"`
	HandAddStatus              int    `json:"handAddStatus" gorm:"comment:是否支持手动新增；0->不支持；1->支持;size:1"`
	Type                       int    `json:"type" gorm:"not null；comment:属性的类型；0->规格；1->参数;size:1"`
	//SysUserAuthorityID int   `json:"sys_user_authority_id" form:"sys_user_authority_id" gorm:"comment:管理角色ID"`
}

func (HomeProductAttribute) TableName() string {
	return "pms_product_attribute"
}

// HomeProductAttributeValue 存储产品参数信息的表
type HomeProductAttributeValue struct {
	global.GVA_MODEL
	ProductId          int    `json:"productId" gorm:"null;default null"`
	ProductAttributeId int    `json:"productAttributeId" gorm:"null;default null"`
	Value              string `json:"value" gorm:"not null;comment:手动添加规格或参数的值，参数单值，规格有多个时以逗号隔开;size:64"`
}

func (HomeProductAttributeValue) TableName() string {
	return "pms_product_attribute_value"
}

// ProductFullReduction 产品满减表(只针对同商品)
type ProductFullReduction struct {
	global.GVA_MODEL
	ProductId   int     `json:"productId" gorm:"null;default null"`
	FullPrice   float32 `json:"fullPrice" gorm:"null;default null"`
	ReducePrice float32 `json:"reducePrice" gorm:"null;default null"`
}

func (ProductFullReduction) TableName() string {
	return "pms_product_full_reduction"
}

// ProductLadder 产品阶梯价格表(只针对同商品)
type ProductLadder struct {
	global.GVA_MODEL
	ProductId int     `json:"productId" gorm:"null;default null"`
	Count     int     `json:"count" gorm:"null;default null;comment:满足的商品数量;"`
	Discount  float32 `json:"discount" gorm:"null;default null;comment:折扣;"`
	Price     float32 `json:"price" gorm:"null;default null;comment:折后价格;"`
}

func (ProductLadder) TableName() string {
	return "pms_product_ladder"
}

// SKUStock sku的库存
type SKUStock struct {
	global.GVA_MODEL
	ProductId      int     `json:"productId" gorm:"null;default null"`
	SkuCode        string  `json:"skuCode" gorm:"null;default null;comment:sku编码;"`
	Price          float32 `json:"price" gorm:"null;default null;comment:价格;"`
	Stock          int     `json:"stock" gorm:"null;default null;comment:库存;"`
	LowStock       int     `json:"lowStock" gorm:"null;default null;comment:预警库存;"`
	Pic            string  `json:"pic" gorm:"null;default null;comment:展示图片;"`
	Sale           int     `json:"sale" gorm:"null;default null;comment:销量;"`
	PromotionPrice float32 `json:"promotionPrice" gorm:"null;default null;comment:单品促销价格;"`
	LockStock      int     `json:"lockStock" gorm:"null;default null;comment:锁定库存;"`
	SpData         string  `json:"spData" gorm:"null;default null;comment:商品销售属性，json格式;"`
}

func (SKUStock) TableName() string {
	return "pms_sku_stock"
}

// CartItem 购物车表
type CartItem struct {
	global.GVA_MODEL
	ProductId         int     `json:"productId" gorm:"null;default null"`
	ProductSkuId      int     `json:"productSkuId" gorm:"null;default null;"`
	UserId            int     `json:"user_id" gorm:" not null;"`
	Quantity          int     `json:"quantity" gorm:"null;default null;comment:购买数量;"`
	Price             float32 `json:"price" gorm:"null;default null;comment:添加到购物车的价格;"`
	ProductPic        string  `json:"productPic" gorm:"null;default null;comment:商品主图;"`
	ProductName       string  `json:"productName" gorm:"null;default null;comment:商品名称;"`
	ProductSubTitle   string  `json:"productSubTitle" gorm:"null;default null;comment:商品副标题（卖点）;"`
	ProductSkuCode    string  `json:"productSkuCode" gorm:"null;default null;comment:商品sku条码;"`
	MemberNickname    string  `json:"memberNickname" gorm:"null;default null;comment:会员昵称;"`
	DeleteStatus      int     `json:"deleteStatus" gorm:"null;default null;comment:是否删除;"`
	ProductCategoryId int     `json:"productCategoryId" gorm:"null;default null;comment:商品分类;"`
	ProductBrand      string  `json:"productBrand" gorm:"null;default null;comment:品牌;"`
	ProductSn         string  `json:"productSn" gorm:"null;default null;"`
	ProductAttr       string  `json:"productAttr" gorm:"null;default null;comment:商品销售属性:[{\"key\":\"颜色\",\"value\":\"颜色\"},{\"key\":\"容量\",\"value\":\"4G\"}];"`
	// TODO: AuthorityId可删除
	AuthorityId int `json:"authorityId" gorm:"default:888;comment:用户角色ID"` // 用户角色ID
}

func (CartItem) TableName() string {
	return "oms_cart_item"
}
