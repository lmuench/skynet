package types

import "github.com/jinzhu/gorm"

// Prototype type
type Prototype struct {
	gorm.Model
	Name  string
	Units []Unit `gorm:"foreignkey:PrototypeRefer"`
}
