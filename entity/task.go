package entity

type Task struct {
	GormModel
	Title       string `gorm:"not null" json:"title" valid:"required~Title is required"`
	Description string `gorm:"not null" json:"description" valid:"required~Description is required"`
	Status      bool   `gorm:"not null" json:"status" valid:"required~Status is required"`
	UserID      uint
	CategoryID  uint
	User        User
	Category    Category
}
