package entity

type Category struct {
	GormModel
	Type string `gorm:"not null" json:"type" valid:"required~Type is required"`
	Task []Task `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"task"`
}
