package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

// Flavor struct
// Only the Vendor requires the gorm hints for preloading
type Flavor struct {
	UUID       string  `json:"uuid"`
	Name       string  `json:"name"`
	VendorUUID string  `json:"vendor_uuid"`
	Vendor     *Vendor `gorm:"foreignkey:UUID;association_foreignkey:VendorUUID" json:"vendor"`
}

func (f *Flavor) String() string {
	return fmt.Sprintf("%s %s", f.VendorUUID, f.Name)
}

func (f *Flavor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, f).String()
}

// GetFlavor -
func (m *Manager) GetFlavor(uuid string) *Flavor {
	f := &Flavor{}
	m.Where("uuid=?", uuid).Preload("Vendor").Find(&f)
	return f
}

// GetFlavors -
func (m *Manager) GetFlavors() *[]Flavor {
	flavors := make([]Flavor, 0)
	m.Preload("Vendor").Find(&flavors)
	return &flavors
}

// BeforeCreate checks
func (f *Flavor) BeforeCreate(scope *gorm.Scope) error {
	if f.Name == "" {
		return errors.New("name must be set")
	}
	if f.VendorUUID == "" {
		return errors.New("vendor_uuid must be set")
	}
	scope.SetColumn("UUID", f.uuid())
	return nil
}

// AddFlavor -
func (m *Manager) AddFlavor(f *Flavor) error {
	db := m.Create(f)
	if db.Error != nil {
		return db.Error
	}
	f.Vendor = m.GetVendor(f.VendorUUID)
	return nil
}

// UpdateFlavor -
func (m *Manager) UpdateFlavor(flavor *Flavor) error {
	db := m.Save(flavor)
	return db.Error
}
