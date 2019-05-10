package model

import (
	"fmt"

	"github.com/twinj/uuid"
)

// Flavor struct
type Flavor struct {
	UUID       string `gorm:"type:uuid;not null;unique_index;primary_key" json:"uuid"`
	Name       string `gorm:"not null;unique_index:idx_name_vendor_uuid" json:"name"`
	VendorUUID string `sql:"type:uuid REFERENCES vendors(uuid)" json:"vendor_uuid"`
	Vendor     Vendor
}

func (v *Flavor) String() string {
	return fmt.Sprintf("%s %s", &v.Vendor, v.Name)
}

func (v *Flavor) GenUUID() string {
	return uuid.NewV3(NameSpaceUUID, v).String()
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

func (m *Manager) AddFlavor(flavor *Flavor) *Flavor {
	if flavor.UUID == "" {
		flavor.UUID = flavor.GenUUID()
	}
	if flavor.VendorUUID == "" {
		flavor.VendorUUID = flavor.Vendor.UUID
	}
	m.Create(flavor) // FIXME: how to catch errors
	return m.GetFlavor(flavor)
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
