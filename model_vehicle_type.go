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

//VehicleType repusents a type of vehicle
type VehicleType struct {
	ID uint64 `json:"id,omitempty" gorm:"primary_key"`

	Make string `json:"make,omitempty"`

	Model string `json:"model,omitempty"`

	Color string `json:"color,omitempty"`

	YearOfManufacture uint8 `json:"yearOfManufacture,omitempty"`

	PassengerCount uint8 `json:"passengerCount,omitempty"`

	// CO2Emissions per gram per mile
	CO2Emissions uint8 `json:"CO2Emissions,omitempty"`

	CreatedAt time.Time

	UpdatedAt time.Time

	DeletedAt *time.Time `sql:"index"`
}

//BeforeCreate sets the ID to a id generated by snowflake also sets creation and updated at times
func (vehicle *VehicleType) BeforeCreate(scope *gorm.Scope) error {
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