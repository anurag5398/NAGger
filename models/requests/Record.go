package requests

type RecordCreate struct {
	Content string `json:"content" binding:"required"`
}
