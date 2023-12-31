package business

import (
	"bytes"
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/business"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"go.uber.org/zap"
	"strconv"
	"strings"
)

type VIPOrderService struct{}

var VIPOrderServiceApp = new(VIPOrderService)

func (exa *VIPOrderService) CreateVIPOrder(e *business.VIPOrder) (err error) {
	err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVIPOrderInfoList
//@description: 获取套餐列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *VIPOrderService) GetVIPOrderInfoList(sysUserAuthorityID int, searchInfo request.OrderSearchInfo) (list interface{}, total int64, err error) {
	limit := searchInfo.PageSize
	offset := searchInfo.PageSize * (searchInfo.Page - 1)

	var orderList []business.VIPOrder
	cmd := fmt.Sprintf("sys_user_authority_id = %d", sysUserAuthorityID)
	if searchInfo.Telephone >= 1000 {
		cmd += fmt.Sprintf(" and telephone like '%%%d%%'", searchInfo.Telephone)
	}
	if len(searchInfo.OrderId) > 1 {
		cmd += fmt.Sprintf(" and order_id like '%%%s%%'", searchInfo.OrderId)
	}
	if searchInfo.Type > 0 {
		cmd += fmt.Sprintf(" and type = %d", searchInfo.Type)
	}
	if len(searchInfo.BuyDate) > 1 {
		cmd += fmt.Sprintf(" and buy_date = %s", searchInfo.BuyDate)
	}

	if limit > 0 && offset > 0 {
		cmd += fmt.Sprintf(" limit %d offset %d", limit, offset)
	}
	db := global.GVA_DB.Model(&business.VIPOrder{})
	err = db.Where(cmd).Count(&total).Error
	if err != nil {
		return orderList, total, err
	} else {
		err = db.Limit(limit).Offset(offset).Preload("Member").Preload("Member.Combo").Where(cmd).Find(&orderList).Error
	}

	//err = global.GVA_DB.Where("sysUserAuthorityID = ? and telephone like ?", sysUserAuthorityID, telephone+"%").First(&member).Error
	return orderList, total, err
}

func (exa *VIPOrderService) CreateVIPStatement(e *business.VIPStatement) (err error) {
	var sql bytes.Buffer
	sql.WriteString("insert into bus_statement(date, recharge, card_number,new_member,consume_number,sys_user_authority_id) values (")
	sql.WriteString("\"")
	sql.WriteString(strings.TrimSpace(e.Date))
	sql.WriteString("\",")
	sql.WriteString(strconv.Itoa(int(e.Recharge)))
	sql.WriteString(",")
	sql.WriteString(strconv.Itoa(int(e.CardNumber)))
	sql.WriteString(",")
	sql.WriteString(strconv.Itoa(int(e.NewMember)))
	sql.WriteString(",")
	sql.WriteString(strconv.Itoa(int(e.ConsumeNumber)))
	sql.WriteString(",")
	sql.WriteString(strconv.Itoa(int(e.SysUserAuthorityID)))
	sql.WriteString(") ON DUPLICATE KEY UPDATE ")
	sql.WriteString("recharge=recharge+")
	sql.WriteString(strconv.Itoa(int(e.Recharge)))
	sql.WriteString(",card_number=card_number+")
	sql.WriteString(strconv.Itoa(int(e.CardNumber)))
	sql.WriteString(",new_member=new_member+")
	sql.WriteString(strconv.Itoa(int(e.NewMember)))
	sql.WriteString(",consume_number=consume_number+")
	sql.WriteString(strconv.Itoa(int(e.ConsumeNumber)))
	sql.WriteString(";")
	fmt.Println("------sql:", sql.String())
	err = global.GVA_DB.Exec(sql.String()).Error

	if err != nil {
		global.GVA_LOG.Error("创建统计失败!", zap.Error(err))
		return err
	}
	//err = global.GVA_DB.Create(&e).Error
	return err
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetVIPStatementInfoList
//@description: 获取流水列表
//@param: sysUserAuthorityID string, info request.PageInfo
//@return: list interface{}, total int64, err error

func (exa *VIPOrderService) GetVIPStatementInfoList(sysUserAuthorityID int, searchInfo request.StatisticsSearchInfo) (list interface{}, err error) {
	var sql bytes.Buffer
	sql.WriteString("Select date,recharge,card_number,new_member,consume_number from bus_statement where sys_user_authority_id = ")
	sql.WriteString(strconv.Itoa(int(sysUserAuthorityID)))
	if len(searchInfo.StartDate) >= 10 {
		sql.WriteString(" and date > ")
		sql.WriteString(strings.TrimSpace(searchInfo.StartDate))
	}
	if len(searchInfo.EndDate) >= 10 {
		sql.WriteString(" and date < ")
		sql.WriteString(strings.TrimSpace(searchInfo.EndDate))
	}

	var orderList []business.VIPStatement
	//fmt.Println(sql.String())
	rows, err := global.GVA_DB.Raw(sql.String()).Rows()
	if err != nil {
		return orderList, err
	}

	for rows.Next() {
		mould := business.VIPStatement{}
		err = rows.Scan(&mould.Date, &mould.Recharge, &mould.CardNumber, &mould.NewMember, &mould.ConsumeNumber)
		orderList = append(orderList, mould)
	}
	rows.Close()

	//err = global.GVA_DB.Where("sysUserAuthorityID = ? and telephone like ?", sysUserAuthorityID, telephone+"%").First(&member).Error
	return orderList, err
}

func (exa *VIPOrderService) BuildVIPStatistics(e *business.VIPStatistics) (cmd string) {
	var sql bytes.Buffer
	sql.WriteString("insert into bus_statistics(total_stream, total_order, total_member,total_consumer,sys_user_authority_id) values (")
	sql.WriteString("\"")
	sql.WriteString(strconv.FormatFloat(e.TotalStream, 'f', 2, 64))
	sql.WriteString("\",")
	sql.WriteString(strconv.Itoa(int(e.TotalOrder)))
	sql.WriteString(",")
	sql.WriteString(strconv.Itoa(int(e.TotalMember)))
	sql.WriteString(",")
	sql.WriteString(strconv.Itoa(int(e.TotalConsumer)))
	sql.WriteString(",")
	sql.WriteString(strconv.Itoa(int(e.SysUserAuthorityID)))
	sql.WriteString(") ON DUPLICATE KEY UPDATE ")
	sql.WriteString("total_stream=total_stream+")
	sql.WriteString(strconv.Itoa(int(e.TotalStream)))
	sql.WriteString(",total_order=total_order+")
	sql.WriteString(strconv.Itoa(int(e.TotalOrder)))
	sql.WriteString(",total_member=total_member+")
	sql.WriteString(strconv.Itoa(int(e.TotalMember)))
	sql.WriteString(",total_consumer=total_consumer+")
	sql.WriteString(strconv.Itoa(int(e.TotalConsumer)))

	sql.WriteString(";")
	return sql.String()
}

func (exa *VIPOrderService) CreateVIPStatistics(e *business.VIPStatistics) (err error) {
	sql := VIPOrderServiceApp.BuildVIPStatistics(e)
	//fmt.Println("------sql:", sql.String())
	err = global.GVA_DB.Exec(sql).Error

	if err != nil {
		global.GVA_LOG.Error("创建统计失败!", zap.Error(err))
		return err
	}
	//err = global.GVA_DB.Create(&e).Error
	return err
}

func (exa *VIPOrderService) GetVIPStatisticsInfoList(sysUserAuthorityID int) (list interface{}, err error) {
	var statistics business.VIPStatistics
	err = global.GVA_DB.Where("sys_user_authority_id = ? ", sysUserAuthorityID).First(&statistics).Error
	return statistics, err
}
