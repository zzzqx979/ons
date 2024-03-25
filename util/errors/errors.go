package errors

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"runtime"
)

var errMap map[int]string

const (
	InternalErr = 1000 + iota
	ValueZeroErr
	DBErr
)

func init() {
	errMap = map[int]string{
		InternalErr:  "内部服务异常",
		ValueZeroErr: "系统约定保留0值",
		DBErr:        "数据库异常:%+v",
	}
}

func GetCode(status int) string {
	return errMap[status]
}

// New
func New(status int) error {
	return errors.New(GetCode(status))
}

// Newf
func Newf(status int, args ...interface{}) error {

	errMsg := GetCode(status)
	if len(args) != 0 {
		errMsg = fmt.Sprintf(errMsg, args...)
	}
	return errors.New(errMsg)
}

func ErrLine(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(0)
		logrus.Errorln("file:", file, ":", line, Newf(DBErr, err.Error()))
	}
}
