package services

import (
	"github.com/jinzhu/gorm"
	"github.com/lmuench/skynet/internal/pkg/services/types"
)

// GetPrototypes returns all prototypes
func (s Prototypes) GetPrototypes() []types.Prototype {
	var prototypes []types.Prototype
	s.DB.Preload("Units").Find(&prototypes)
	return prototypes
}

// GetPrototype returns prototype with provided ID
func (s Prototypes) GetPrototype(id int) types.Prototype {
	var prototype types.Prototype
	s.DB.First(&prototype, id)
	return prototype
}

// Prototypes ...
type Prototypes struct {
	DB *gorm.DB
}
