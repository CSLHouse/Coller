package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderHomeProductAttributeCategory = initOrderHomeProduct + 1

type initHomeProductAttributeCategory struct{}

// auto run
func init() {
	system.RegisterInit(initOrderHomeProductAttributeCategory, &initHomeProductAttributeCategory{})
}

func (i *initHomeProductAttributeCategory) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.HomeProductAttributeCategory{},
	)
}

func (i *initHomeProductAttributeCategory) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.HomeProductAttributeCategory{})
}

func (i initHomeProductAttributeCategory) InitializerName() string {
	return wechatModel.HomeProductAttributeCategory{}.TableName()
}

func (i *initHomeProductAttributeCategory) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.HomeProductAttributeCategory{
		{
			Name:           "服装-T恤",
			AttributeCount: 2,
			ParamCount:     5,
		},
		{
			Name:           "服装-裤装",
			AttributeCount: 2,
			ParamCount:     4,
		},
		{
			Name:           "手机数码-手机通讯",
			AttributeCount: 2,
			ParamCount:     4,
		},
		{
			Name:           "配件",
			AttributeCount: 0,
			ParamCount:     0,
		},
		{
			Name:           "居家",
			AttributeCount: 0,
			ParamCount:     0,
		},
		{
			Name:           "洗护",
			AttributeCount: 0,
			ParamCount:     0,
		},
		{
			Name:           "测试分类",
			AttributeCount: 0,
			ParamCount:     0,
		},
		{
			Name:           "服装-鞋帽",
			AttributeCount: 4,
			ParamCount:     0,
		},
		{
			Name:           "家用电器-电视",
			AttributeCount: 2,
			ParamCount:     4,
		},
		{
			Name:           "电脑办公-笔记本",
			AttributeCount: 2,
			ParamCount:     3,
		},
		{
			Name:           "家用电器-厨卫大电",
			AttributeCount: 1,
			ParamCount:     3,
		},
		{
			Name:           "电脑办公-硬盘",
			AttributeCount: 2,
			ParamCount:     5,
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.HomeProductAttributeCategory{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initHomeProductAttributeCategory) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("name = ?", "配件").First(&wechatModel.HomeProductAttributeCategory{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
