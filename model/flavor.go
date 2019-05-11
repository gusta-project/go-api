package model

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

// Flavor struct
// Only the Vendor requires the gorm hints for preloading
type Flavor struct {
	ID       int     `json:"-"` // don't publish
	UUID     string  `json:"uuid"`
	Slug     string  `json:"slug"`
	Name     string  `json:"name"`
	Gravity  string  `json:"gravity"`
	VendorID int     `json:"vendor_id" json:"-"`
	Vendor   *Vendor `gorm:"foreignkey:UUID;association_foreignkey:VendorID" json:"vendor"`
}

func (f *Flavor) String() string {
	// FIXME: how to ensure that vendor is loaded?
	if f.Vendor.ID == 0 {
		return fmt.Sprintf("[NotLoaded!] %s", f.Name)
	}
	return fmt.Sprintf("%s %s", f.Vendor.Slug, f.Name)
}

// slug for this flavor.
// Will panic if the vendor is not loaded.
func (f *Flavor) slug() string {
	// FIXME: how to ensure that vendor is loaded?
	if f.Vendor.ID == 0 {
		panic("called Flavor.String() with vendor not loaded")
	}
	return slug.Make(fmt.Sprintf("%s %s", f.Vendor.Slug, f.Name))
}

// uuid for this flavor based on the slug.
// Will panic if the vendor is not loaded.
func (f *Flavor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, f.slug()).String()
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
	scope.SetColumn("Slug", f.slug())
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
