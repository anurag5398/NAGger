package handlers

import (
	"NAGger/components/managers"
	"NAGger/models/requests"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type RecordHandler struct {
	RecordManager managers.RecordManagerInterface `inject:""`
}

// @Summary Create Record
// @Description Create Record
// @Tags Record
// @Param request body requests.RecordCreate true "Request body"
// @Success 200 {object} string
// @Router /record/create [post]
func (self *RecordHandler) CreateRecord(c *gin.Context) {
	var content requests.RecordCreate
	var err error
	if err = c.ShouldBindBodyWith(content, binding.JSON); err == nil {
		if err = self.RecordManager.CreateRecord(c, content.Content); err == nil {
			c.JSON(http.StatusAccepted, gin.H{
				"message": "Record created successfully",
			})
		}
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})

	}
}
