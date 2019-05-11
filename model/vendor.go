package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

// Vendor struct
type Vendor struct {
	ID   int    `json:"-"`
	UUID string `json:"uuid"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Code string `json:"code"`
	URL  string `json:"url"`
}

func (v *Vendor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Code)
}

func (v *Vendor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, v).String()
}

// GetVendor -
func (m *Manager) GetVendor(v *Vendor) *Vendor {
	m.Where(v).Find(v)
	return v
}

// GetVendors -
func (m *Manager) GetVendors() *[]Vendor {
	vendors := make([]Vendor, 0)
	m.Find(&vendors)
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
	scope.SetColumn("UUID", v.uuid())
	return nil
}

// AddVendor -
func (m *Manager) AddVendor(v *Vendor) error {
	db := m.Where(v).FirstOrCreate(v)
	return db.Error
}

// UpdateVendor -
func (m *Manager) UpdateVendor(vendor *Vendor) error {
	db := m.Save(vendor)
	return db.Error
}
