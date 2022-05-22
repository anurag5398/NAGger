package commons

import (
	"NAGger/models"
	"context"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// Tx Extracts Transaction from Model and Returns */
func Tx(c context.Context) (tx *gorm.DB) {
	//Context Must be Passed
	if c != nil {
		//Check If Context Has Tx
		if value := c.Value(models.CONTEXT_TX); value != nil {
			//Extract and Return
			tx = value.(*gorm.DB)
		} else {
			log.Trace("Context Passed Without Transaction")
		}
	} else {
		log.Trace("Nil Context Passed")
	}
	return
}
