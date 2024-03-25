package db

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"ons/util/errors"
	"runtime"
	"time"
)

type FunctionType string

type ObjectModel struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	Status      int       `json:"status"`
	ProductCode string    `json:"product_code" xorm:"notnull index(codes)"`
	FactoryCode string    `json:"factory_code" xorm:"notnull index(codes)"`
	Properties  string    `json:"properties"`
	Events      string    `json:"events"`
	Services    string    `json:"services"`
	Note        string    `json:"note"`
	UpdatedAt   time.Time `json:"-" xorm:"updated"`
	DeletedAt   time.Time `json:"-" xorm:"deleted"`
}

type TSL struct {
	Version string `json:"version"`
	Profile
	Properties []Property `json:"properties"`
	Events     []Event    `json:"events"`
	Services   []Service  `json:"services"`
}

type Profile struct {
	ProductCode string `json:"product_code"`
	FactoryCode string `json:"factory_code"`
}

func AddObjectModel(model ObjectModel) error {
	affected, err := GetDB().Insert(&model)
	if affected != 1 {
		logrus.Errorln("AddObjectModel() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func GetObjectModelByCodes(factoryCode, productCode string) (tsl TSL, has bool, err error) {
	model := ObjectModel{
		FactoryCode: factoryCode,
		ProductCode: productCode,
	}

	has, err = GetDB().Get(&model)
	errors.ErrLine(err)
	tsl, err = ParseTSL(model)
	errors.ErrLine(err)
	return
}

func ParseTSL(model ObjectModel) (tsl TSL, err error) {
	tsl.FactoryCode = model.FactoryCode
	tsl.ProductCode = model.ProductCode
	err = json.Unmarshal([]byte(model.Properties), &tsl.Properties)
	if err != nil {
		logrus.Errorf("ParseTSL() unmarshal properties err:%+v.\n", err)
		return
	}
	err = json.Unmarshal([]byte(model.Events), &tsl.Events)
	if err != nil {
		logrus.Errorf("ParseTSL() unmarshal events err:%+v.\n", err)
		return
	}
	err = json.Unmarshal([]byte(model.Services), &tsl.Services)
	if err != nil {
		logrus.Errorf("ParseTSL() unmarshal services err:%+v.\n", err)
	}
	return
}

type GetObjectModelsQuery struct {
	Page        int    `form:"page"`
	PageSize    int    `form:"page_size"`
	Name        string `form:"name"`
	ProductCode string `form:"product_code"`
	FactoryCode string `form:"factory_code"`
}

func GetObjectModels(q GetObjectModelsQuery) (count int64, models []ObjectModel, err error) {
	dbCount := GetDB().NewSession()
	dbFind := GetDB().NewSession()
	if q.Name != "" {
		dbCount = dbCount.Where("name = ?", q.Name)
		dbFind = dbFind.Where("name = ?", q.Name)
	}
	if q.FactoryCode != "" {
		dbCount = dbCount.Where("factory_code = ?", q.FactoryCode)
		dbFind = dbFind.Where("factory_code = ?", q.FactoryCode)
	}
	if q.ProductCode != "" {
		dbCount = dbCount.Where("product_code = ?", q.ProductCode)
		dbFind = dbFind.Where("product_code = ?", q.ProductCode)
	}
	count, err = dbCount.Count(&ObjectModel{})
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.Errorln("file:", file, ":", line, errors.Newf(errors.DBErr, err.Error()))
		return
	}

	err = dbFind.Limit(q.PageSize, q.PageSize*(q.Page-1)).Find(&models)
	errors.ErrLine(err)
	return
}

func UpdateObjectModel(model ObjectModel) error {
	affected, err := GetDB().ID(model.Id).AllCols().Update(&model)
	if affected != 1 {
		logrus.Errorln("UpdateObjectModel() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}

func DelObjectModel(id int) error {
	affected, err := GetDB().ID(id).Delete(&ObjectModel{})
	if affected != 1 {
		logrus.Errorln("DelObjectModel() affected row incorrect")
	}
	errors.ErrLine(err)
	return err
}
