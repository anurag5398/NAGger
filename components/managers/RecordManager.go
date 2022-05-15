package managers

import (
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
}

func (self *RecordManager) CreateRecord(c context.Context, content string) (err error) {
	return
}