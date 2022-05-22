package dao

import (
	"NAGger/components/commons"
	"context"
	"github.com/sirupsen/logrus"
)

type BaseDaoInterface interface {
	// Create TODO: Verify if this creates same entity again and again
	//Create entity
	Create(c context.Context, entity interface{}) (err error)
}

type BaseDao struct {
	//TODO: Implement Client Error Mapper
}

func (self *BaseDao) Create(c context.Context, entity interface{}) (err error) {
	var txErr error
	if txErr := commons.Tx(c).Create(entity).Error; txErr != nil {
		logrus.WithContext(c).WithFields(logrus.Fields{"Entity": entity}).WithField("Err", txErr).Error("Error while creating entity")
	}
	//TODO: err conversion from txErr to err
	err = txErr
	return
}
