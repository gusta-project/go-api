package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/twinj/uuid"
)

// Flavor struct
type Flavor struct {
	UUID       string `gorm:"type:uuid;not null;unique_index;primary_key" json:"uuid"`
	Name       string `gorm:"not null;unique_index:idx_name_vendor_uuid" json:"name"`
	VendorUUID string `sql:"type:uuid REFERENCES vendors(uuid)" json:"vendor_uuid"`
	Vendor     Vendor
}

func (f *Flavor) String() string {
	return fmt.Sprintf("%s %s", &f.Vendor, f.Name)
}

func (f *Flavor) GenUUID() string {
	return uuid.NewV3(NameSpaceUUID, f).String()
}

// HasFlavor check if the given flavor exists
func (m *Manager) GetFlavor(flavor *Flavor) *Flavor {
	var result = &Flavor{}
	if flavor.UUID != "" {
		m.Where("uuid=?", flavor.UUID).Find(&result)
	}
	if flavor.Name != "" && flavor.VendorUUID != "" {
		m.Where("name=? AND vendor_uuid=?", flavor.Name, flavor.VendorUUID).Find(&result)
	}
	if flavor.Name != "" && flavor.Vendor.UUID != "" {
		m.Where("name=? AND vendor_uuid=?", flavor.Name, flavor.Vendor.UUID).Find(&result)
	}
	if result.UUID != "" {
		result.Vendor = *m.GetVendor(&Vendor{UUID: result.VendorUUID})
	}
	return result
}

func (f *Flavor) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("UUID", f.GenUUID())
	return nil
}

func (m *Manager) AddFlavor(f *Flavor) *Flavor {
	if f.VendorUUID == "" {
		f.VendorUUID = f.Vendor.UUID
	}
	m.Create(f) // FIXME: how to catch errors
	return f
}

func (m *Manager) UpdateFlavor(flavor *Flavor) *Flavor {
	// if we want to update by Name or by Code only, we need to get the UUID first
	if flavor.UUID == "" {
		current := m.GetFlavor(flavor)
		flavor.UUID = current.UUID
	}
	m.Save(flavor)
	return flavor
}
