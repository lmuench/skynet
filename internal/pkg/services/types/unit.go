package types

import "github.com/jinzhu/gorm"

// Unit type
type Unit struct {
	gorm.Model
	PrototypeID uint
	Health      int
}
