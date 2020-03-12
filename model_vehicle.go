/*
 * VehicleService
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0
 * Contact: johnfg2610\"gmail.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Vehicle - Represents a vehicle
type Vehicle struct {
	ID uint64 `json:"id,omitempty" gorm:"primary_key" binding:"isdefault"`

	VehicleTypeID uint64 `json:"vehicleTypeId" binding:"required"`
	//VehicleType *VehicleType `json:"vehicleType,omitempty" gorm:"foreignkey:ID"`

	PlateNumber uint16 `json:"plateNumber" gorm:"UNIQUE_INDEX" binding:"required"`

	Millage uint64 `json:"millage" binding:"required"`

	MOTExpiry *time.Time `json:"MOTExpiry,omitempty"`

	PlateExpiry *time.Time `json:"plateExpiry,omitempty"`

	InsuranceExpiry *time.Time `json:"insuranceExpiry,omitempty"`

	RoadTaxExpiry *time.Time `json:"roadTaxExpiry,omitempty"`

	BusinessID uint64 `json:"businessID" gorm:"INDEX" binding:"isdefault"`

	DriverID uint64 `json:"driverID" binding:"required" binding:"isdefault"`

	IsActive bool `json:"isActive" gorm:"default:'true'" binding:"isdefault"`

	Registration string `json:"registration,omitempty" gorm:"UNIQUE_INDEX" binding:"required"`

	CreatedAt time.Time `json:"CreatedAt,omitempty" binding:"isdefault"`

	UpdatedAt time.Time `json:"UpdatedAt,omitempty" binding:"isdefault"`

	DeletedAt *time.Time `json:"DeletedAt,omitempty" binding:"isdefault" sql:"index"`
}

//BeforeCreate sets the ID to a id generated by snowflake also sets creation and updated at times
func (vehicle *Vehicle) BeforeCreate(scope *gorm.Scope) error {
	id, err := sf.NextID()
	if err != nil {
		return err
	}
	scope.SetColumn("ID", id)
	scope.SetColumn("CreatedAt", time.Now())
	scope.SetColumn("UpdatedAt", time.Now())
	scope.SetColumn("DeletedAt", nil)
	return nil
}
