package db

import (
	"github.com/sirupsen/logrus"
	"ons/util/errors"
	"runtime"
	"time"
)

const (
	StatusEnabled  = 1
	StatusDisabled = 0
)

type Product struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code" xorm:"notnull unique"`
	Status    int       `json:"status"`
	Note      string    `json:"note"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func AddProduct(p Product) error {
	affected, err := GetDB().Insert(&p)
	if affected != 1 {
		logrus.Errorln("AddProduct() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func DelProduct(id int) error {
	affected, err := GetDB().ID(id).Delete(&Product{})
	if affected != 1 {
		logrus.Errorln("DelProduct() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func UpdateProduct(p Product) error {
	affected, err := GetDB().ID(p.Id).AllCols().Update(&p)
	if affected != 1 {
		logrus.Errorln("UpdateProduct() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func GetProducts(page, pageSize int, name, code string) (count int64, products []Product, err error) {
	dbCount := GetDB().NewSession()
	dbFind := GetDB().NewSession()
	if name != "" {
		dbCount = dbCount.Where("name = ?", name)
		dbFind = dbFind.Where("name = ?", name)
	}
	if code != "" {
		dbCount = dbCount.Where("code = ?", code)
		dbFind = dbFind.Where("code = ?", code)
	}
	count, err = dbCount.Count(&Product{})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.Errorln("file:", file, ":", line, errors.Newf(errors.DBErr, err.Error()))
		return
	}
	err = dbFind.Limit(pageSize, pageSize*(page-1)).Find(&products)
	errors.ErrLine(err)
	return
}
