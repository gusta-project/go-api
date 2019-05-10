package flavor

import (
	"github.com/jinzhu/gorm"
	"github.com/pscn/flavor2go/model"
)

// Flavor struct
type Flavor struct {
	gorm.Model `json:"-"`
	UUID       string `gorm:"not null;unique_index;primary_key" json:"uuid"`
	Name       string `gorm:"not null;unique_index:idx_name_vendor_uuid" json:"name"`
	VendorUUID string `gorm:"not null;unique_index:idx_name_vendor_uuid" json:"vendor_uuuid"`
}

// Manager struct
type Manager struct {
	db *model.DB
}

// New create a new *Manager for flavors
func New(db *model.DB) (*Manager, error) {
	db.AutoMigrate(&Flavor{})
	return &Manager{db: db}, nil
}

// Has check if the given flavor exists
func (m *Manager) Has(name, VendorUUID string) bool {
	if err := m.db.Where("name=? AND vendor_uuid=?", name, VendorUUID).Find(&Flavor{}).Error; err != nil {
		return false
	}
	return true
}
