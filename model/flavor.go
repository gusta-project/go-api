package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

// Flavor struct
type Flavor struct {
	UUID       string `gorm:"type:uuid;not null;unique_index;primary_key" json:"uuid"`
	Name       string `gorm:"not null;unique_index:idx_name_vendor_uuid" json:"name"`
	VendorUUID string `gorm:"not null;unique_index:idx_name_vendor_uuid" json:"vendor_uuid"`
	Vendor     Vendor `gorm:"foreignkey:vendor_uuid;association_foreignkey:uuid" json:"vendor"`
}

func (f *Flavor) String() string {
	return fmt.Sprintf("%s %s", &f.Vendor, f.Name)
}

func (f *Flavor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, f).String()
}

// GetFlavor -
func (m *Manager) GetFlavor(uuid string) *Flavor {
	flavor := &Flavor{}
	m.Where("uuid=?", uuid).Find(&flavor)
	if flavor.UUID != "" {
		flavor.Vendor = *m.GetVendor(flavor.VendorUUID)
	}
	return flavor
}

// GetFlavors -
func (m *Manager) GetFlavors() *[]Flavor {
	flavors := make([]Flavor, 0)
	m.Find(&flavors)
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
	db := m.Create(f) // FIXME: how to catch errors
	return db.Error
}

// UpdateFlavor -
func (m *Manager) UpdateFlavor(flavor *Flavor) error {
	db := m.Save(flavor)
	return db.Error
}
