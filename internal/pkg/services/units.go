package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lmuench/skynet/internal/pkg/services/types"
)

// GetUnits returns all units
func (s Units) GetUnits() []types.Unit {
	var units []types.Unit
	s.DB.Find(&units)
	return units
}

// GetUnit returns unit with provided ID
func (s Units) GetUnit(id int) types.Unit {
	var unit types.Unit
	s.DB.First(&unit, id)
	return unit
}

// Units ...
type Units struct {
	DB *gorm.DB
}
