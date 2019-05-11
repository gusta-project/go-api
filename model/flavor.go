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
	ID       int     `json:"-"`
	UUID     string  `json:"uuid"`
	Slug     string  `json:"slug"`
	Name     string  `json:"name"`
	VendorID int     `json:"vendor_id" json:"-"`
	Vendor   *Vendor `gorm:"foreignkey:UUID;association_foreignkey:VendorID" json:"vendor"`
}

func (f *Flavor) String() string {
	return fmt.Sprintf("%s %s", f.Vendor.Slug, f.Name)
}

func (f *Flavor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, f).String()
}

// GetFlavor -
func (m *Manager) GetFlavor(f *Flavor) *Flavor {
	m.Where(f).Preload("Vendor").Find(f)
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
	if f.VendorID == 0 {
		return errors.New("vendor id must be set")
	}
	scope.SetColumn("UUID", f.uuid())
	return nil
}

// AddFlavor -
func (m *Manager) AddFlavor(f *Flavor) error {
	vendor := m.GetVendor(&Vendor{ID: f.VendorID})
	if vendor.UUID == "" {
		return fmt.Errorf("no vendor with id=%d", f.VendorID)
	}
	db := m.Where(f).FirstOrCreate(f)
	if db.Error != nil {
		return db.Error
	}
	f.Vendor = vendor
	return nil
}

// UpdateFlavor -
func (m *Manager) UpdateFlavor(flavor *Flavor) error {
	db := m.Save(flavor)
	return db.Error
}
