package db

import (
	"ons/util/errors"
	"runtime"
	"time"

	"github.com/sirupsen/logrus"
)

type OnsInfo struct {
	Id          int64     `json:"id"`
	Tid         string    `json:"tid" xorm:"notnull unique"`
	ProductCode string    `json:"product_code" xorm:"notnull index(codes)"`
	FactoryCode string    `json:"factory_code" xorm:"notnull index(codes)"`
	Epc         string    `json:"epc"`
	Note        string    `json:"note"`
	UpdatedAt   time.Time `json:"-" xorm:"updated"`
}

// 添加Tid和Epc的映射关系
func AddOnsInfo(info OnsInfo) error {
	affected, err := GetDB().Insert(&info)
	if affected != 1 {
		logrus.Errorf("AddOnsInfo() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func GetOnsInfoByTid(tid string) (info OnsInfo, has bool, err error) {
	info.Tid = tid
	has, err = GetDB().Get(&info)
	errors.ErrLine(err)
	return
}

func GetOnsInfoByEpc(epc string) (info OnsInfo, has bool, err error) {
	info.Epc = epc
	has, err = GetDB().Get(&info)
	errors.ErrLine(err)
	return
}

func UpdateOnsInfo(info OnsInfo) error {
	affected, err := GetDB().ID(info.Id).AllCols().Update(&info)
	if affected != 1 {
		logrus.Errorln("UpdateOnsInfo() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func DeleteOnsInfo(id int) error {
	affected, err := GetDB().ID(id).Delete(&OnsInfo{})
	if affected != 1 {
		logrus.Errorln("UpdateOnsInfo() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

type GetOnsInfosQuery struct {
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
	Tid         string `form:"tid"`
	ProductCode string `form:"product_code"`
	FactoryCode string `form:"Factory_code"`
}

func GetOnsInfos(q GetOnsInfosQuery) (count int64, infos []OnsInfo, err error) {
	dbCount := GetDB().NewSession()
	dbFind := GetDB().NewSession()
	if q.Tid != "" {
		dbCount = dbCount.Where("tid = ?", q.Tid)
		dbFind = dbFind.Where("tid = ?", q.Tid)
	}
	if q.FactoryCode != "" {
		dbCount = dbCount.Where("factory_code = ?", q.FactoryCode)
		dbFind = dbFind.Where("factory_code = ?", q.FactoryCode)
	}
	if q.ProductCode != "" {
		dbCount = dbCount.Where("product_code = ?", q.ProductCode)
		dbFind = dbFind.Where("product_code = ?", q.ProductCode)
	}
	count, err = dbCount.Count(&OnsInfo{})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.Errorln("file:", file, ":", line, errors.Newf(errors.DBErr, err.Error()))
		return
	}

	err = dbFind.Limit(q.PageSize, q.PageSize*(q.Page-1)).Find(&infos)
	errors.ErrLine(err)
	return
}
