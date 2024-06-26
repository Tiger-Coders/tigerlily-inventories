package inventory

import (
	"github.com/Tiger-Coders/tigerlily-inventories/internal/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Create an interface to prevent unwanted use of these methods

type InventoryRepo struct {
	db *gorm.DB
}

// Receives the db instance as argument and sets it in the struct before returning the struct itself
func NewInventoryRepo(db *gorm.DB) *InventoryRepo {
	return &InventoryRepo{
		db: db,
	}
}

func (m InventoryRepo) GetAllInventories(limit, offset int32) (items []*models.Skus, err error) {
	m.db.Debug().Find(&items)
	return
}
