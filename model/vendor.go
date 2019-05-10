package model

import (
	"fmt"

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

func (m *Manager) AddVendor(vendor *Vendor) *Vendor {
	if vendor.UUID == "" {
		vendor.UUID = vendor.GenUUID()
	}
	m.Create(vendor) // FIXME: how to catch errors
	return m.GetVendor(vendor)
}

func (m *Manager) UpdateVendor(vendor *Vendor) *Vendor {
	// if we want to update by Name or by Code only, we need to get the UUID first
	if vendor.UUID == "" {
		current := m.GetVendor(vendor)
		vendor.UUID = current.UUID
	}
	m.Save(vendor)
	return vendor
}
