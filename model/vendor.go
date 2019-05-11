package model

import (
	"errors"
	"fmt"
	"log"

	"github.com/gosimple/slug"
	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

// Vendor struct
type Vendor struct {
	ID   int    `json:"-"` // don't publish
	UUID string `json:"uuid"`
	Slug string `json:"slug"`
	Name string `json:"name"`
	Code string `json:"code"`
	URL  string `json:"url"`
}

func (v *Vendor) String() string {
	return fmt.Sprintf("%s %s", v.Code, v.Name)
}

func (v *Vendor) slug() string {
	return slug.Make(fmt.Sprintf("%s %s", v.Code, v.Name))
}

func (v *Vendor) uuid() string {
	return uuid.NewV3(NameSpaceUUID, v).String()
}

// GetVendor -
func (m *Manager) GetVendor(v *Vendor) *Vendor {
	err := m.HandleError(m.Where(v).Find(v))
	if err != nil {
		log.Printf("error in GetVendor: %v", err)
	}
	return v
}

// GetVendors -
func (m *Manager) GetVendors() *[]Vendor {
	vendors := make([]Vendor, 0)
	err := m.HandleError(m.Find(&vendors))
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
	scope.SetColumn("UUID", v.uuid())
	scope.SetColumn("Slug", v.slug())
	return nil
}

// AddVendor -
func (m *Manager) AddVendor(v *Vendor) error {
	return m.HandleError(m.Where(v).FirstOrCreate(v))
}

// DeleteVendor -
func (m *Manager) DeleteVendor(v *Vendor) error {
	return m.HandleError(m.Where(v).Delete(v))
}

// UpdateVendor -
func (m *Manager) UpdateVendor(vendor *Vendor) error {
	return m.HandleError(m.Save(vendor))
}
