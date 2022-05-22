package managers

import (
	"NAGger/components/dao"
	"NAGger/models/entities"
	"context"
	"time"
)

type RecordManagerInterface interface {
	CreateRecord(c context.Context, content string) (err error)
	// GetRecords : startDate and endDate can be nil -> One month default in that case
	// If keyword is provided and nil dates, then search in entire records
	GetRecords(c context.Context, startDate time.Time, endDate time.Time, keywords []string) (output map[string]string, err error)
}

type RecordManager struct {
	BaseManager `inject:"inline"`
	Dao         dao.RecordDaoInterface `inject:""`
}

func (self *RecordManager) CreateRecord(c context.Context, content string) (err error) {
	err = self.UseOrCreateTx(c, func(c context.Context) (err error) {
		record := entities.Record{
			Content: content,
		}
		err = self.Dao.CreateRecord(c, &record)
		return
	})
	return
}
