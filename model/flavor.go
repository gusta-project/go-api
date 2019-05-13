package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"github.com/twinj/uuid"
)

// Flavor struct
// Only the Vendor requires the gorm hints for preloading
type Flavor struct {
	ID       int     `json:"-"`    // don't publish
	Slug     string  `json:"slug"` // used as primary key for the API calls
	UUID     string  `json:"uuid"`
	Name     string  `json:"name"`
	Gravity  string  `json:"gravity"`
	VendorID int     `json:"vendor_id" json:"-"`
	Vendor   *Vendor `gorm:"foreignkey:UUID;association_foreignkey:VendorID" json:"vendor"`
}

// FlavorManager to provide for short method names and a cache

type FlavorManager struct {
	cache *cache.Cache
	db    *Manager // link back to manager
}

// NewFlavorManager -
func (m *Manager) NewFlavorManager() *FlavorManager {
	return &FlavorManager{db: m, cache: newCache()}
}

func (f *Flavor) String() string {
	return f.Slug
}

// slug for this flavor.
// Will panic if the vendor is not loaded.
func (f *Flavor) slug() string {
	// FIXME: how to ensure that vendor is loaded?
	if f.Vendor.ID == 0 {
		panic("called Flavor.slug() with vendor not loaded")
	}
	return slug.Make(fmt.Sprintf("%s %s", f.Vendor.Slug, f.Name))
}

// uuid for this flavor based on the slug.
// Will panic if the vendor is not loaded.
func (f *Flavor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, f).String()
}

// Get -
func (m *FlavorManager) Get(f *Flavor) *Flavor {
	m.db.Where(f).Preload("Vendor").Find(f)
	return f
}

// GetAll -
func (m *FlavorManager) GetAll() *[]Flavor {
	flavors := make([]Flavor, 0)
	m.db.Preload("Vendor").Find(&flavors)
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
	f.Slug = f.slug() // set before calling uuid
	scope.SetColumn("UUID", f.uuid())
	scope.SetColumn("Slug", f.Slug)
	return nil
}

// Create -
func (m *FlavorManager) Create(f *Flavor) error {
	log.Printf("AddFlavor: %v", f.Vendor)
	v := m.db.Vendor.Get(f.Vendor)
	if v == nil {
		return fmt.Errorf("no vendor for %v", f.Vendor)
	}
	f.VendorID = v.ID
	db := m.db.Where(f).FirstOrCreate(f)
	if db.Error != nil {
		return db.Error
	}
	f.Vendor = v
	return nil
}

// UpdateFlavor -
func (m *FlavorManager) Update(flavor *Flavor) error {
	db := m.db.Save(flavor)
	return db.Error
}
