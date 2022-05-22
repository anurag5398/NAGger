package entities

type Record struct {
	BaseEntity `gorm:"embedded"`
	User       string `gorm:"type:varchar(255)" json:"user"`
	Content    string `gorm:"type:varchar(10000)" json:"content"`
}
