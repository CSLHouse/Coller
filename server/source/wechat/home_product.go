package wechat

import (
	"context"
	wechatModel "github.com/flipped-aurora/gin-vue-admin/server/model/wechat"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderHomeProduct = initOrderRecommendProduct + 1

type initHomeProduct struct{}

// auto run
func init() {
	system.RegisterInit(initOrderHomeProduct, &initHomeProduct{})
}

func (i *initHomeProduct) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(
		&wechatModel.Product{},
	)
}

func (i *initHomeProduct) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&wechatModel.Product{})
}

func (i initHomeProduct) InitializerName() string {
	return wechatModel.Product{}.TableName()
}

func (i *initHomeProduct) InitializeData(ctx context.Context) (next context.Context, err error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}

	entities := []wechatModel.Product{
		{
			BrandId:                    3,
			ProductCategoryId:          19,
			FeightTemplateId:           0,
			ProductAttributeCategoryId: 3,
			Name:                       "华为 HUAWEI P20",
			Pic:                        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf58Ndefaac16.jpg",
			ProductSN:                  "6946605",
			DeleteStatus:               0,
			PublishStatus:              1,
			NewStatus:                  1,
			RecommandStatus:            1,
			VerifyStatus:               0,
			Sort:                       100,
			Sale:                       100,
			Price:                      3788.00,
			PromotionPrice:             3659.00,
			GiftGrowth:                 3788,
			GiftPoint:                  3788,
			UsePointLimit:              0,
			SubTitle:                   "AI智慧全面屏 6GB +64GB 亮黑色 全网通版 移动联通电信4G手机 双卡双待手机 双卡双待",
			Description:                "",
			OriginalPrice:              4288.00,
			Stock:                      1000,
			LowStock:                   0,
			Unit:                       "件",
			Weight:                     0.00,
			PreviewStatus:              1,
			ServiceIds:                 "2,3,1",
			Keywords:                   "",
			Note:                       "",
			AlbumPics:                  "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ab46a3cN616bdc41.jpg,http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ac1bf5fN2522b9dc.jpg",
			DetailTitle:                "",
			DetailDesc:                 "",
			DetailHTML:                 "<p><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ad44f1cNf51f3bb0.jpg\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ad44fa8Nfcf71c10.jpg\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ad44fa9N40e78ee0.jpg\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ad457f4N1c94bdda.jpg\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5ad457f5Nd30de41d.jpg\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180607/5b10fb0eN0eb053fb.jpg\\\" /></p>",
			DetailMobileHTML:           "<p><img src=\\\"//img20.360buyimg.com/vc/jfs/t1/81293/35/5822/369414/5d3fe77cE619c5487/6e775a52850feea5.jpg!q70.dpg.webp\\\" alt=\\\"\\\" width=\\\"750\\\" height=\\\"776\\\" /></p>\\n<p><img src=\\\"//img20.360buyimg.com/vc/jfs/t1/45300/25/11592/3658871/5d85ef66E92a8a911/083e47d8f662c582.jpg!q70.dpg.webp\\\" alt=\\\"\\\" width=\\\"596\\\" height=\\\"16383\\\" /></p>",
			PromotionStartDate:         "2023-01-10 15:49:38",
			PromotionEndDate:           "2023-01-31 00:00:00",
			PromotionPerLimit:          0,
			PromotionType:              1,
			BrandName:                  "华为",
			ProductCategoryName:        "手机通讯",
		},
		{
			BrandId:                    6,
			ProductCategoryId:          19,
			FeightTemplateId:           0,
			ProductAttributeCategoryId: 3,
			Name:                       "小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待",
			Pic:                        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20180615/xiaomi.jpg",
			ProductSN:                  "7437788",
			DeleteStatus:               0,
			PublishStatus:              1,
			NewStatus:                  1,
			RecommandStatus:            1,
			VerifyStatus:               0,
			Sort:                       0,
			Sale:                       99,
			Price:                      2699.00,
			PromotionPrice:             0,
			GiftGrowth:                 2699,
			GiftPoint:                  2699,
			UsePointLimit:              0,
			SubTitle:                   "骁龙845处理器，红外人脸解锁，AI变焦双摄，AI语音助手小米6X低至1299，点击抢购",
			Description:                "小米8 全面屏游戏智能手机 6GB+64GB 黑色 全网通4G 双卡双待",
			OriginalPrice:              2699.00,
			Stock:                      100,
			LowStock:                   0,
			Unit:                       "",
			Weight:                     0.00,
			PreviewStatus:              0,
			ServiceIds:                 "1",
			Keywords:                   "",
			Note:                       "",
			AlbumPics:                  "",
			DetailTitle:                "",
			DetailDesc:                 "",
			DetailHTML:                 "<p style=\\\"text-align: center;\\\"><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/95935/9/19330/159477/5e9ecc13E5b8db8ae/8e954021a0835fb7.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/99224/22/19596/137593/5e9ecc13E34ef2113/2b362b90d8378ee1.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/93131/25/19321/107691/5e9ecc13E41e8addf/202a5f84d9129878.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/101843/19/19533/66831/5e9ecc13Ecb7f9d53/4fdd807266583c1e.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/99629/36/19016/59882/5e9ecc13E1f5beef5/1e5af3528f366e70.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/72343/29/8945/84548/5d6e5c67Ea07b1125/702791816b90eb3d.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/65403/35/9017/129532/5d6e5c68E3f2d0546/9ec771eb6e04a75a.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/58261/33/9380/106603/5d6e5c68E05a99177/2b5b9e29eed05b08.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/51968/40/9688/113552/5d6e5c68E5052b312/8969d83124cb78a4.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/75491/9/9101/146883/5d6e5c68E3db89775/c1aa57e78ded4e44.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/75063/11/9107/127874/5d6e5c68E0b1dfc61/10dd6000ce213375.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/47452/25/9579/108279/5d6e5c68Ea9002f3b/865b5d32ffd9da75.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/82146/26/9077/87075/5d6e5c68Ef63bccc8/556de5665a35a3f2.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/82804/21/9019/124404/5d6e5c69E06a8f575/0f861f8c4636c546.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/46044/10/9734/107686/5d6e5c69Edd5e66c7/a8c7b9324e271dbd.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/40729/32/13755/45997/5d6e5c69Eafee3664/6a3457a4efdb79c5.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/76254/34/9039/96195/5d6e5c69E3c71c809/49033c0b7024c208.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/79825/20/9065/90864/5d6e5c69E1e62ef89/a4d3ce383425a666.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/49939/21/9618/106207/5d6e5c6aEf7b1d4fd/0f5e963c66be3d0c.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/64035/7/9001/115159/5d6e5c6aE6919dfdf/39dfe4840157ad81.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/53389/3/9616/99637/5d6e5c6aEa33b9f35/b8f6aa26e72616a3.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/63219/6/9026/81576/5d6e5c6aEa9c74b49/b4fa364437531012.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/42146/27/13902/80437/5d6e5c6bE30c31ce9/475d4d54c7334313.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/45317/28/9596/78175/5d6e5c6bEce31e4b7/5675858b6933565c.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/60080/1/9112/138722/5d6e5c6bEefd9fc62/7ece7460a36d2fcc.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/51525/13/9549/36018/5d6e5c73Ebbccae6c/99bc2770dccc042b.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/61948/20/9088/72918/5d6e5c73Eab7aef5c/6f21e2f85cf478fa.jpg!q70.dpg.webp\\\" /></p>', '<p style=\\\"text-align: center;\\\"><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/95935/9/19330/159477/5e9ecc13E5b8db8ae/8e954021a0835fb7.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/99224/22/19596/137593/5e9ecc13E34ef2113/2b362b90d8378ee1.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/93131/25/19321/107691/5e9ecc13E41e8addf/202a5f84d9129878.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/101843/19/19533/66831/5e9ecc13Ecb7f9d53/4fdd807266583c1e.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWareDetail/jfs/t1/99629/36/19016/59882/5e9ecc13E1f5beef5/1e5af3528f366e70.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/72343/29/8945/84548/5d6e5c67Ea07b1125/702791816b90eb3d.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/65403/35/9017/129532/5d6e5c68E3f2d0546/9ec771eb6e04a75a.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/58261/33/9380/106603/5d6e5c68E05a99177/2b5b9e29eed05b08.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/51968/40/9688/113552/5d6e5c68E5052b312/8969d83124cb78a4.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/75491/9/9101/146883/5d6e5c68E3db89775/c1aa57e78ded4e44.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/75063/11/9107/127874/5d6e5c68E0b1dfc61/10dd6000ce213375.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/47452/25/9579/108279/5d6e5c68Ea9002f3b/865b5d32ffd9da75.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/82146/26/9077/87075/5d6e5c68Ef63bccc8/556de5665a35a3f2.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/82804/21/9019/124404/5d6e5c69E06a8f575/0f861f8c4636c546.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/46044/10/9734/107686/5d6e5c69Edd5e66c7/a8c7b9324e271dbd.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/40729/32/13755/45997/5d6e5c69Eafee3664/6a3457a4efdb79c5.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/76254/34/9039/96195/5d6e5c69E3c71c809/49033c0b7024c208.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/79825/20/9065/90864/5d6e5c69E1e62ef89/a4d3ce383425a666.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/49939/21/9618/106207/5d6e5c6aEf7b1d4fd/0f5e963c66be3d0c.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/64035/7/9001/115159/5d6e5c6aE6919dfdf/39dfe4840157ad81.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/53389/3/9616/99637/5d6e5c6aEa33b9f35/b8f6aa26e72616a3.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/63219/6/9026/81576/5d6e5c6aEa9c74b49/b4fa364437531012.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/42146/27/13902/80437/5d6e5c6bE30c31ce9/475d4d54c7334313.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/45317/28/9596/78175/5d6e5c6bEce31e4b7/5675858b6933565c.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/60080/1/9112/138722/5d6e5c6bEefd9fc62/7ece7460a36d2fcc.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/51525/13/9549/36018/5d6e5c73Ebbccae6c/99bc2770dccc042b.jpg!q70.dpg.webp\\\" /><img src=\\\"//img30.360buyimg.com/popWaterMark/jfs/t1/61948/20/9088/72918/5d6e5c73Eab7aef5c/6f21e2f85cf478fa.jpg!q70.dpg.webp\\\" /></p>",
			DetailMobileHTML:           "",
			PromotionStartDate:         "",
			PromotionEndDate:           "",
			PromotionPerLimit:          0,
			PromotionType:              3,
			BrandName:                  "小米",
			ProductCategoryName:        "手机通讯",
		},
		{
			BrandId:                    51,
			ProductCategoryId:          53,
			FeightTemplateId:           0,
			ProductAttributeCategoryId: 3,
			Name:                       "Apple iPad 10.9英寸平板电脑 2022年款",
			Pic:                        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/ipad_001.jpg",
			ProductSN:                  "100044025833",
			DeleteStatus:               0,
			PublishStatus:              1,
			NewStatus:                  0,
			RecommandStatus:            0,
			VerifyStatus:               0,
			Sort:                       0,
			Sale:                       0,
			Price:                      3599.00,
			PromotionPrice:             0.00,
			GiftGrowth:                 0,
			GiftPoint:                  0,
			UsePointLimit:              0,
			SubTitle:                   "【11.11大爱超大爱】iPad9代限量抢购，价格优惠，更享以旧换新至高补贴325元！！快来抢购吧！！",
			Description:                "",
			OriginalPrice:              3599.00,
			Stock:                      1000,
			LowStock:                   0,
			Unit:                       "",
			Weight:                     0.00,
			PreviewStatus:              0,
			ServiceIds:                 "1,2,3",
			Keywords:                   "",
			Note:                       "",
			AlbumPics:                  "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/ipad_002.jpg",
			DetailTitle:                "",
			DetailDesc:                 "",
			DetailHTML:                 "<div style=\\\"margin: 0 auto;\\\"><img style=\\\"max-width: 390px;\\\" src=\\\"//img12.360buyimg.com/cms/jfs/t1/75040/28/21026/295081/634ed154E981e4d10/2ceef91d6f2b2273.jpg!q70.dpg.webp\\\" /> <img style=\\\"max-width: 390px;\\\" src=\\\"//img13.360buyimg.com/cms/jfs/t1/191028/1/28802/401291/634ed15eEb234dc40/5ab89f83531e1023.jpg!q70.dpg.webp\\\" /> <img style=\\\"max-width: 390px;\\\" src=\\\"//img14.360buyimg.com/cms/jfs/t1/32758/24/18599/330590/634ed15eEc3db173c/c52953dc8008a576.jpg!q70.dpg.webp\\\" /></div>",
			DetailMobileHTML:           "",
			PromotionStartDate:         "",
			PromotionEndDate:           "",
			PromotionPerLimit:          0,
			PromotionType:              0,
			BrandName:                  "苹果",
			ProductCategoryName:        "平板电脑",
		},
		{
			BrandId:                    6,
			ProductCategoryId:          54,
			FeightTemplateId:           0,
			ProductAttributeCategoryId: 13,
			Name:                       "小米 Xiaomi Book Pro 14 2022 锐龙版 2.8K超清大师屏 高端轻薄笔记本电脑",
			Pic:                        "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_001.jpg",
			ProductSN:                  "100023207945",
			DeleteStatus:               0,
			PublishStatus:              1,
			NewStatus:                  0,
			RecommandStatus:            1,
			VerifyStatus:               0,
			Sort:                       0,
			Sale:                       0,
			Price:                      5599.00,
			PromotionPrice:             0.00,
			GiftGrowth:                 0,
			GiftPoint:                  0,
			UsePointLimit:              0,
			SubTitle:                   "【双十一大促来袭】指定型号至高优惠1000，以旧换新至高补贴1000元，晒单赢好礼",
			Description:                "",
			OriginalPrice:              5599.00,
			Stock:                      500,
			LowStock:                   0,
			Unit:                       "",
			Weight:                     0.00,
			PreviewStatus:              0,
			ServiceIds:                 "",
			Keywords:                   "",
			Note:                       "",
			AlbumPics:                  "http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_002.jpg",
			DetailTitle:                "",
			DetailDesc:                 "",
			DetailHTML:                 "",
			DetailMobileHTML:           "<div class=\\\"ssd-module-mobile-wrap\\\">\\n<div class=\\\"ssd-module M16667778180631\\\" data-id=\\\"M16667778180631\\\"><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_snipaste_01.png\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_snipaste_02.png\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_snipaste_03.png\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_snipaste_04.png\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_snipaste_05.png\\\" /><img class=\\\"wscnph\\\" src=\\\"http://macro-oss.oss-cn-shenzhen.aliyuncs.com/mall/images/20221028/xiaomi_computer_snipaste_06.png\\\" /></div>\\n<div class=\\\"ssd-module M16667778180631\\\" data-id=\\\"M16667778180631\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16534569204241\\\" data-id=\\\"M16534569204241\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16667778416884\\\" data-id=\\\"M16667778416884\\\">\\n<div class=\\\"ssd-widget-text W16667778440496\\\">&nbsp;</div>\\n</div>\\n<div class=\\\"ssd-module M165303211067557\\\" data-id=\\\"M165303211067557\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16530320459861\\\" data-id=\\\"M16530320459861\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16530320460062\\\" data-id=\\\"M16530320460062\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16530320460153\\\" data-id=\\\"M16530320460153\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16530320460366\\\" data-id=\\\"M16530320460366\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16530320460477\\\" data-id=\\\"M16530320460477\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16530320460578\\\" data-id=\\\"M16530320460578\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M16530320460699\\\" data-id=\\\"M16530320460699\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204608110\\\" data-id=\\\"M165303204608110\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204609511\\\" data-id=\\\"M165303204609511\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204611213\\\" data-id=\\\"M165303204611213\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204612714\\\" data-id=\\\"M165303204612714\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204614115\\\" data-id=\\\"M165303204614115\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204615516\\\" data-id=\\\"M165303204615516\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204617417\\\" data-id=\\\"M165303204617417\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204618818\\\" data-id=\\\"M165303204618818\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204620219\\\" data-id=\\\"M165303204620219\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204621620\\\" data-id=\\\"M165303204621620\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204622921\\\" data-id=\\\"M165303204622921\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204624522\\\" data-id=\\\"M165303204624522\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204626024\\\" data-id=\\\"M165303204626024\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204627525\\\" data-id=\\\"M165303204627525\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204629127\\\" data-id=\\\"M165303204629127\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204632330\\\" data-id=\\\"M165303204632330\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204634031\\\" data-id=\\\"M165303204634031\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204635832\\\" data-id=\\\"M165303204635832\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204637533\\\" data-id=\\\"M165303204637533\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204639334\\\" data-id=\\\"M165303204639334\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204642235\\\" data-id=\\\"M165303204642235\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204647936\\\" data-id=\\\"M165303204647936\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204651037\\\" data-id=\\\"M165303204651037\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204653838\\\" data-id=\\\"M165303204653838\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204656239\\\" data-id=\\\"M165303204656239\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204659141\\\" data-id=\\\"M165303204659141\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204661943\\\" data-id=\\\"M165303204661943\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204665844\\\" data-id=\\\"M165303204665844\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204668045\\\" data-id=\\\"M165303204668045\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204670146\\\" data-id=\\\"M165303204670146\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204672147\\\" data-id=\\\"M165303204672147\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204674348\\\" data-id=\\\"M165303204674348\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204676749\\\" data-id=\\\"M165303204676749\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204681352\\\" data-id=\\\"M165303204681352\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204683553\\\" data-id=\\\"M165303204683553\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204685855\\\" data-id=\\\"M165303204685855\\\">&nbsp;</div>\\n<div class=\\\"ssd-module M165303204688356\\\" data-id=\\\"M165303204688356\\\">&nbsp;</div>\\n</div>",
			PromotionStartDate:         "",
			PromotionEndDate:           "",
			PromotionPerLimit:          0,
			PromotionType:              0,
			BrandName:                  "小米",
			ProductCategoryName:        "笔记本",
		},
	}
	if err = db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrap(err, wechatModel.Product{}.TableName()+"表数据初始化失败!")
	}
	next = context.WithValue(ctx, i.InitializerName(), entities)

	return next, err
}

func (i *initHomeProduct) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("brand_id = ?", 6).First(&wechatModel.Product{}).Error, gorm.ErrRecordNotFound) { // 判断是否存在数据
		return false
	}
	return true
}
