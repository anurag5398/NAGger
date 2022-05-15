package managers

import (
	"NAGger/commons/helper"
	"NAGger/models"
	"context"
	"errors"
	"gorm.io/gorm"
)

type BaseManager struct {
	Db *gorm.DB `inject:""`
}

func (self *BaseManager) UseOrCreateTx(c context.Context, run func(c context.Context) (err error), readOnly ...bool) (err error) {
	var txErr error

	//Check if Context has Tx
	if helper.Tx(c) != nil {
		//First Preference to use existing tx if supplied
		txErr = run(c)
	} else if len(readOnly) > 0 && readOnly[0] {
		// If Read only use DB as Transaction
		readOnlyError := run(context.WithValue(c, models.CONTEXT_TX, self.Db.WithContext(c)))
		return readOnlyError
	} else {
		//Else create New Transaction
		txErr = self.Db.WithContext(c).Transaction(func(tx *gorm.DB) error {
			return run(context.WithValue(c, models.CONTEXT_TX, tx))

		})
	}

	/* Convert Transaction error to Client Error */
	if ok := errors.As(txErr, &err); txErr != nil && !ok {
		//If Transaction function returns non Client Error
		//raise a New Server Error
		err = errors.New(txErr.Error())
	}
	return
}
