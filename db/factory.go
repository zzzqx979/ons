package db

import (
	"github.com/sirupsen/logrus"
	"ons/util/errors"
	"runtime"
	"time"
)

type Factory struct {
	Id        int64     `json:"id"`
	Name      string    `json:"name"`
	Code      string    `json:"code" xorm:"notnull unique"`
	Status    int       `json:"status"`
	Note      string    `json:"note"`
	UpdatedAt time.Time `json:"-"`
	DeletedAt time.Time `json:"-"`
}

func AddFactory(f Factory) error {
	f.UpdatedAt = time.Now()
	affected, err := GetDB().Insert(&f)
	if affected != 1 {
		logrus.Errorln("AddFactory() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func DelFactory(id int) error {
	affected, err := GetDB().ID(id).Delete(&Factory{})
	if affected != 1 {
		logrus.Errorln("DelFactory() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func UpdateFactory(f Factory) error {
	f.UpdatedAt = time.Now()
	affected, err := GetDB().ID(f.Id).AllCols().Update(&f)
	if affected != 1 {
		logrus.Errorln("UpdateFactory() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func GetFactories(page, pageSize int, name, code string) (count int64, Factories []Factory, err error) {
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
	count, err = dbCount.Count(&Factory{})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.Errorln("file:", file, ":", line, errors.Newf(errors.DBErr, err.Error()))
		return
	}
	err = dbFind.Limit(pageSize, pageSize*(page-1)).Find(&Factories)
	errors.ErrLine(err)
	return
}
