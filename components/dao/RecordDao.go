package dao

import (
	"NAGger/models/entities"
	"context"
)

type RecordDaoInterface interface {
	CreateRecord(c context.Context, record *entities.Record) (err error)
}

type RecordDao struct {
	BaseDao `inject:"inline"`
}

func (self *RecordDao) CreateRecord(c context.Context, record *entities.Record) (err error) {
	err = self.Create(c, record)
	return
}
