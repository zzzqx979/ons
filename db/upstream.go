package db

import (
	"github.com/sirupsen/logrus"
	"ons/util/errors"
)

type Upstream struct {
	Id   int64
	Ons1 string
	Ons2 string
	Ons3 string
	Ms1  string
	Ms2  string
	Ms3  string
}

func AddUpstream(upstream Upstream) error {
	affected, err := GetDB().Insert(&upstream)
	if affected != 1 {
		logrus.Errorf("AddUpstream() affected row incorrect.")
	}
	errors.ErrLine(err)
	return err
}

func UpdateUpstream(upstream Upstream) error {
	affected, err := GetDB().ID(upstream.Id).AllCols().Update(&upstream)
	if affected != 1 {
		logrus.Errorf("UpdateUpstream() affected row incorrect.")
	}
	errors.ErrLine(err)
	return err
}

func DeleteUpstream(id int) error {
	affected, err := GetDB().ID(id).Delete(&Upstream{})
	if affected != 1 {
		logrus.Errorf("DeleteUpstream() affected row incorrect.")
	}
	errors.ErrLine(err)
	return err
}

func GetUpstream() (upstream Upstream, err error) {
	var has bool
	has, err = GetDB().Desc("id").Get(&upstream)
	if !has {
		logrus.Errorf("GetUpstream() could not found record.")
	}
	errors.ErrLine(err)
	return
}

func GetOnsUpstream() (onsUpStream []string, err error) {
	var has bool
	upstream := &Upstream{}
	has, err = GetDB().Desc("id").Get(upstream)
	if !has {
		logrus.Errorf("GetUpstream() could not found record.")
	}
	errors.ErrLine(err)
	onsUpStream = append(onsUpStream, upstream.Ons1, upstream.Ons2, upstream.Ons3)
	return
}

func GetMsUpstream() (onsUpStream []string, err error) {
	var has bool
	upstream := &Upstream{}
	has, err = GetDB().Desc("id").Get(upstream)
	if !has {
		logrus.Errorf("GetUpstream() could not found record.")
	}
	errors.ErrLine(err)
	onsUpStream = append(onsUpStream, upstream.Ms1, upstream.Ms2, upstream.Ms3)
	return
}
