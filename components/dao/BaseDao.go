package dao

import (
	"NAGger/commons/helper"
	"context"
)

type BaseDaoInterface interface {
	Create(c context.Context, entity interface{}) (err error)
}

type BaseDao struct {
}

func (self *BaseDao) Create(c context.Context, entity interface{}) (err error) {
	err = helper.Tx(c).Create(entity).Error
	return
}
