package handlers

import (
	"NAGger/models/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RecordHandler struct {
}

func (self *RecordHandler) CreateRecord(c *gin.Context) {
	var content requests.RecordCreate
	if err := c.ShouldBindBodyWith(content, binding.JSON); err == nil {

	}
}
