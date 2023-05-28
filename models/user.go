package models

type User struct {
	Model
	Email     string     `json:"email,omitempty" gorm:"not null;unique"`
	Name      string     `json:"name" gorm:"not null"`
	Bio       string     `json:"bio" gorm:"type:text"`
	Password  string     `json:"-" gorm:"not null"`
	Image     string     `json:"image"`
	ImageName string     `json:"imageName"`
	Provider  string     `json:"provider" gorm:"type:enum('GOOGLE','EMAIL');default:EMAIL;not null" binding:"required"`
	Projects  []*Project `json:"projects,omitempty"`
}
