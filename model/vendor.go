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

// Vendor struct
type Vendor struct {
	ID   int    `json:"-"`    // don't publish
	Slug string `json:"slug"` // used as primary key for the API calls
	UUID string `json:"uuid"`
	Name string `json:"name"`
	Code string `json:"code"`
	URL  string `json:"url"`
}

// VendorManager to provide for short method names and a cache
type VendorManager struct {
	cache *cache.Cache
	db    *Manager // link back to manager
}

// NewVendorManager -
func (m *Manager) NewVendorManager() *VendorManager {
	return &VendorManager{db: m, cache: newCache()}
}

func (v *Vendor) String() string {
	return v.Slug
}

// slug transparently populates vendor.Slug
func (v *Vendor) slug() string {
	return slug.Make(fmt.Sprintf("%s %s", v.Code, v.Name))
}

func (v *Vendor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, v).String()
}

func (m *VendorManager) getFromCache(slug string) *Vendor {
	if slug == "" {
		return nil
	}
	if c, found := m.cache.Get(slug); found {
		v := c.(*Vendor)
		log.Printf("from cache: %s", v)
		return v
	}
	return nil
}

func (m *VendorManager) storeToCache(v *Vendor) {
	if v.Slug != "" {
		m.cache.SetDefault(v.Slug, v)
	}
}

func (m *VendorManager) findOne(v *Vendor) *gorm.DB {
	return m.db.Model(v).Where("slug=?", v.Slug)
}

// Get -
func (m *VendorManager) Get(v *Vendor) *Vendor {
	result := m.getFromCache(v.Slug)
	if result != nil {
		return result
	}
	result = &Vendor{}
	err := m.db.HandleError(m.findOne(v).Find(result))
	if err != nil {
		log.Printf("error in GetVendor: %v", err)
		return nil
	}
	m.storeToCache(result)
	return result
}

// GetAll -
func (m *VendorManager) GetAll() *[]Vendor {
	vendors := make([]Vendor, 0)
	err := m.db.HandleError(m.db.Find(&vendors))
	if err != nil {
		log.Printf("error in GetVendors: %v", err)
	}
	return &vendors
}

// BeforeCreate check if name & code is set and generate a UUID
func (v *Vendor) BeforeCreate(scope *gorm.Scope) error {
	if v.Name == "" {
		return errors.New("name must be set")
	}
	if v.Code == "" {
		return errors.New("code must be set")
	}
	v.Slug = v.slug() // set before calling uuid
	scope.SetColumn("UUID", v.uuid())
	scope.SetColumn("Slug", v.Slug)
	return nil
}

// Create -
func (m *VendorManager) Create(v *Vendor) error {
	err := m.db.HandleError(m.findOne(v).FirstOrCreate(v))
	m.storeToCache(v)
	return err
}

// Delete -
func (m *VendorManager) Delete(v *Vendor) error {
	return m.db.HandleError(m.findOne(v).Delete(v))
}

// Update vendor identified by s with values from v
// returns the modified vendor or nil if the vendor is not found or an error occured
func (m *VendorManager) Update(s *Vendor, v *Vendor) *Vendor {
	db := m.findOne(s).Update(v)
	err := m.db.HandleError(db)
	if err != nil {
		log.Printf("Error: Update: %v", err)
		return nil
	}
	if db.RowsAffected == 0 {
		log.Printf("Info: Update: no rows affected")
		return nil
	}

	// Verify / update slug
	v = m.Get(s)
	if v == nil { // should not happen as we checked RowsAffected above
		// panic?
		log.Printf("Error: Update affected rows, but could not fetch the result")
		return nil
	}
	newSlug := v.slug()
	if newSlug != v.Slug { // changed => update
		db = m.findOne(s).Update(&Vendor{Slug: newSlug})
		err = m.db.HandleError(db)
	}
	return v
}
