package model

import (
	"errors"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

// Vendor struct
type Vendor struct {
	UUID string `gorm:"type:uuid;not null;unique_index;primary_key" json:"uuid"`
	Name string `gorm:"not null;unique_index" json:"name"`
	Code string `gorm:"not null;unique_index" json:"code"`
	URL  string `gorm:"" json:"url"`
}

func (v *Vendor) String() string {
	return fmt.Sprintf("%s %s", v.Name, v.Code)
}

func (v *Vendor) GenUUID() string {
	return uuid.NewV3(NameSpaceUUID, v).String()
}

// HasVendor check if the given vendor exists
func (m *Manager) GetVendor(vendor *Vendor) *Vendor {
	var result = &Vendor{}
	if vendor.UUID != "" {
		m.Where("uuid=?", vendor.UUID).Find(&result)
	}
	if vendor.Name != "" {
		m.Where("name=?", vendor.Name).Find(&result)
	}
	if vendor.Code != "" {
		m.Where("code=?", vendor.Code).Find(&result)
	}
	return result
}

// BeforeCreate check if name & code is set and generate a UUID
func (v *Vendor) BeforeCreate(scope *gorm.Scope) error {
	if v.Name == "" {
		return errors.New("name must be set")
	}
	if v.Code == "" {
		return errors.New("code must be set")
	}
	scope.SetColumn("UUID", v.GenUUID())
	return nil
}

// AddVendor -
func (m *Manager) AddVendor(v *Vendor) error {
	db := m.Create(v)
	return db.Error
}

// UpdateVendor -
func (m *Manager) UpdateVendor(vendor *Vendor) error {
	db := m.Save(vendor)
	return db.Error
}
